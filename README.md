# Solana Jupiter Airdrop Eligibility Checker

This is a Go-based program that checks the eligibility of Solana wallets for a Jupiter airdrop. It categorizes wallets into the following groups:

- **Empty Wallets**: Wallets with no allocation.
- **Sybil Wallets**: Wallets flagged for suspicious activity by Jupiter and don't have allocation.
- **Eligible Wallets**: Wallets meeting the eligibility criteria for the airdrop.

## Features

- Reads wallets from either `data/public.txt` or `data/private.txt`.
- Supports proxies listed in `data/proxy.txt`.
- Configurable settings through `cfg.ini`.
- Adjustable delay between wallet checks to avoid rate limiting.

## Installation

### Prerequisites

Ensure you have Go installed on your system. If not, you can download it from the [official website](https://golang.org/dl/).

### Steps

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd <repository-folder>
   ```
2. Install dependencies:
   ```bash
   go mod install
   ```
3. Run the program:
   ```bash
   go run main.go
   ```

## Setup

### 1. Configure `cfg.ini`

Create or edit the `cfg.ini` file with the following structure:

```ini
[settings]
usePrivateKeys = false  # Set to true to scan private keys in data/private.txt
                        # Set to false to scan public keys in data/public.txt
delay_in_ms = 30        # Delay in milliseconds between wallet checks
```

### 2. Add Wallets

- If `usePrivateKeys` is set to `true`, populate `data/private.txt` with private keys (one per line).
- If `usePrivateKeys` is set to `false`, populate `data/public.txt` with public keys (one per line).

### 3. Add Proxies (optional)

Add proxies to `data/proxy.txt` in the following format:

```
<proxy-host>:<proxy-port>:<username>:<password>
```

## Usage

1. Ensure the `cfg.ini`, wallet files (`data/public.txt` or `data/private.txt`), and `data/proxy.txt` are properly set up.
2. Run the program:
   ```bash
   go run main.go
   ```
3. The program will:
   - Check wallet eligibility using the Jupiter API.
   - Categorize wallets as empty, Sybil, or eligible.
   - Print the results to the console.

## Example Configuration

### cfg.ini

```ini
[settings]
usePrivateKeys = false
delay_in_ms = 30
```

### data/public.txt

```
2eCrYZv5Z5yijaJP4RYS3pvWXzQXtyfg3eDHDKmfzj4L
8RyVcKQt8pDXbCGHSpeKksaJEAJhmxMNzQYoF9A4Z3pL
EyJPN5ywLYm635NFMrVVqtwpb1aL9AV93Bjgey4ySsQ8
```

### data/proxy.txt

```
proxy1.example.com:8080:username:password
proxy2.example.com:8080:username:password
```

## Notes

- Ensure proxies are valid and active to avoid connection issues.
- Delays can be adjusted in `cfg.ini` to prevent being rate-limited.
- Make sure the wallet file (`public.txt` or `private.txt`) corresponds to the `usePrivateKeys` setting in `cfg.ini`.

## Contributions

Feel free to fork the repository and submit pull requests for any enhancements or bug fixes.

---

Happy scanning!
