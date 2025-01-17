package main

import (
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/dzhisl/jupiter-airdrop-checker/api"
	"github.com/dzhisl/jupiter-airdrop-checker/types"
	"github.com/dzhisl/jupiter-airdrop-checker/utils"
	"github.com/jedib0t/go-pretty/v6/table"
)

var (
	emptyWallets   []mapStruct
	eligbleWallets []mapStruct
	sybilWallets   []mapStruct
	errorWallets   []utils.WalletMap
	mu             sync.Mutex // Mutex for thread-safe operations
)

type mapStruct struct {
	wallet   utils.WalletMap
	response types.GetAllocationResponse
}

func main() {
	cfg, err := utils.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}
	logger := utils.NewLogger()
	proxies := utils.SetupProxy()
	wallets := utils.SetupWallets(cfg.UsePrivateKeys)

	var wg sync.WaitGroup

	// Initial check for all wallets
	for _, wallet := range wallets {
		wg.Add(1)
		go func(w utils.WalletMap) {
			defer wg.Done()
			checkEligible(w, *logger, proxies)
		}(wallet)
		time.Sleep(time.Duration(cfg.Delay) * time.Millisecond)
	}

	wg.Wait() // Wait for all goroutines to finish

	// Retry failed wallets until none are left
	retries := 0
	const maxRetries = 5

	for len(errorWallets) > 0 && retries < maxRetries {
		logger.Info(fmt.Sprintf("Retrying failed wallets. Attempt %d", retries+1))
		failedWallets := errorWallets
		errorWallets = nil // Reset error wallet list for this iteration

		for _, wallet := range failedWallets {
			wg.Add(1)
			go func(w utils.WalletMap) {
				defer wg.Done()
				checkEligible(w, *logger, proxies)
			}(wallet)
			time.Sleep(time.Duration(cfg.Delay) * time.Millisecond)
		}

		wg.Wait()
		retries++
	}

	if len(errorWallets) > 0 {
		logger.Error(fmt.Sprintf("Some wallets still failed after %d retries: %d", maxRetries, len(errorWallets)))
	}

	// Log the results
	logger.Info(fmt.Sprintf("Total wallets: %d", len(wallets)))
	logger.Info(fmt.Sprintf("Total eligible wallets: %d", len(eligbleWallets)))
	logger.Info(fmt.Sprintf("Total sybil wallets: %d", len(sybilWallets)))
	logger.Info(fmt.Sprintf("Total empty wallets: %d", len(emptyWallets)))
	logger.Info(fmt.Sprintf("Total error wallets: %d", len(errorWallets)))

	// Render the table
	renderTable()
	logger.Info("All wallets checked.")
}

func checkEligible(wallet utils.WalletMap, logger utils.Logger, proxies []*url.URL) {
	logger.Info(fmt.Sprintf("Checking wallet: %s | %s", wallet.PrivateKey.String(), wallet.PublicKey.String()))

	result, err := api.GetJupAllocation(wallet.PublicKey.String(), proxies)

	mu.Lock() // Lock the mutex before modifying shared resources
	defer mu.Unlock()

	if err != nil {
		logger.Error(fmt.Sprintf("Error checking allocation for wallet: %s, err: %s", wallet.PublicKey.String(), err.Error()))
		errorWallets = append(errorWallets, wallet)
		return
	}

	if result.Data.TotalAllocated > 0 {
		eligbleWallets = append(eligbleWallets, mapStruct{wallet: wallet, response: result})
	} else if result.Data.LikelySybil == true {
		sybilWallets = append(sybilWallets, mapStruct{wallet: wallet, response: result})
	} else {
		emptyWallets = append(emptyWallets, mapStruct{wallet: wallet, response: result})
	}
}

func renderTable() {
	t := utils.SetupTable()
	var sum int
	for i, item := range eligbleWallets {
		sum += int(item.response.Data.TotalAllocated)
		t.AppendRow(table.Row{
			i, item.wallet.PrivateKey.String(), item.wallet.PublicKey.Short(5), item.response.Data.TotalAllocated,
		})
	}
	t.AppendFooter(table.Row{"", "", "Total", sum})

	t.Render()
}
