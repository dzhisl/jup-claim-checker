package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
	"github.com/dzhisl/jupiter-airdrop-checker/types"
	"github.com/dzhisl/jupiter-airdrop-checker/utils"
)

// create TLS client with proxy to bypass cloudflare anti-bot
func createClient(proxies []*url.URL) (tls_client.HttpClient, error) {
	jar := tls_client.NewCookieJar()
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(10),
		tls_client.WithClientProfile(profiles.Firefox_133),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar),
	}

	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP client: %w", err)
	}

	if len(proxies) > 0 {
		randProxy := utils.GetRandomProxy(proxies)
		client.SetProxy(randProxy.String())
	}

	return client, nil
}

func GetJupAllocation(walletPubKey string, proxies []*url.URL) (types.GetAllocationResponse, error) {
	url := fmt.Sprintf("https://jupuary.jup.ag/api/allocation?wallet=%s", walletPubKey)
	client, err := createClient(proxies)
	if err != nil {
		return types.GetAllocationResponse{}, err
	}

	req, err := http.NewRequest("GET", url, nil)
	setRequestHeaders(req)
	resp, err := client.Do(req)
	if err != nil {
		return types.GetAllocationResponse{}, fmt.Errorf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return types.GetAllocationResponse{}, fmt.Errorf("Received non-200 response: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.GetAllocationResponse{}, fmt.Errorf("Failed to read response body: %v", err)
	}
	// Define the struct to unmarshal into
	var allocationResponse types.GetAllocationResponse

	// Unmarshal the JSON into the struct
	err = json.Unmarshal(body, &allocationResponse)
	if err != nil {
		return types.GetAllocationResponse{}, fmt.Errorf("Failed to unmarshal JSON response: %v", err)
	}

	// Print the unmarshaled data
	return allocationResponse, nil
}

func setRequestHeaders(req *http.Request) {
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("priority", "u=1, i")
	req.Header.Set("referer", fmt.Sprintf("https://jupuary.jup.ag/allocation/%s", req.URL.Query().Get("wallet")))
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="131", "Chromium";v="131", "Not_A Brand";v="24"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")
}
