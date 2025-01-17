package utils

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"os"
	"strings"

	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/jedib0t/go-pretty/v6/table"
	"gopkg.in/ini.v1"
)

type WalletMap struct {
	PrivateKey solana.PrivateKey
	PublicKey  solana.PublicKey
}

type config struct {
	UsePrivateKeys bool
	Delay          int
}

// SetupWallets reads wallets from a file and returns a slice of solana.PrivateKey
func SetupWallets(usePrivate bool) []WalletMap {
	// Define the file name
	fmt.Println("Setting up wallets")
	var m []WalletMap
	fileName := "data/private.txt"
	if usePrivate == false {
		fileName = "data/public.txt"
	}

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Use a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if usePrivate == true {
			privateKey := solana.MustPrivateKeyFromBase58(line)
			publicKey := privateKey.PublicKey()

			s := WalletMap{
				privateKey,
				publicKey,
			}
			m = append(m, s)
		} else {
			privateKey := solana.PrivateKey{}
			publicKey := solana.MustPublicKeyFromBase58(line)
			s := WalletMap{
				privateKey,
				publicKey,
			}
			m = append(m, s)
		}

	}

	// Check for errors while scanning
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return m
}

// SetupWallets reads wall
func SetupProxy() []*url.URL {
	// Define the file name
	fileName := "data/proxy.txt"

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Create a slice to store the proxies
	var proxies []*url.URL

	// Use a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")

		if len(parts) != 4 {
			log.Printf("Invalid proxy format: %s", line)
			continue
		}

		// Construct the proxy URL
		proxyURL := &url.URL{
			Scheme: "http", // You can change this to "https" if needed
			User:   url.UserPassword(parts[2], parts[3]),
			Host:   fmt.Sprintf("%s:%s", parts[0], parts[1]),
		}

		// Append the proxy URL to the slice
		proxies = append(proxies, proxyURL)
	}

	// Check for errors while scanning
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return proxies
}

func GetRandomProxy(proxies []*url.URL) *url.URL {
	if len(proxies) == 0 {
		return nil
	}
	rand.Seed(time.Now().UnixNano())
	randomProxy := proxies[rand.Intn(len(proxies))]
	return randomProxy
}

func SetupTable() table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	// Use a compact style

	// Optionally, set column widths
	t.SetStyle(table.StyleLight)

	// Set header and separator
	t.AppendHeader(table.Row{"#", "Private", "Public", "Total Allocation"})
	t.AppendSeparator()

	return t
}

func ParseConfig() (config, error) {
	cfg, err := ini.Load("cfg.ini")

	if err != nil {
		return config{}, fmt.Errorf("failed to load config file: %v", err)
	}

	usePrivateKeys, err := cfg.Section("settings").Key("usePrivateKeys").Bool()
	delay, err := cfg.Section("settings").Key("delay_in_ms").Int()

	if err != nil {
		return config{}, fmt.Errorf("invalid value for usePrivateKeys: %v", err)
	}

	fmt.Printf("Configuration loaded: usePrivateKeys=%v\n", usePrivateKeys)

	return config{
		UsePrivateKeys: usePrivateKeys,
		Delay:          delay,
	}, nil
}
