# XrayPing

**XrayPing** is a command-line tool for testing the latency of multiple IP addresses using the Xray proxy and a SOCKS5 proxy. The tool measures latency and allows for concurrency, retries, and verbose output. It is designed to be modular, providing flexibility for future expansion and improvements.

## Features

- **SOCKS5 Proxy**: Latency testing is routed through a specified SOCKS5 proxy.
- **Xray Proxy Integration**: Leverages the Xray proxy to test IP addresses in sequence.
- **Concurrency**: Multiple IP addresses are tested concurrently, improving speed.
- **Retries**: Ability to retry latency tests for each IP multiple times to get the best result.
- **Verbose Mode**: Optional detailed logs during the execution of tests.
- **Colorized Output**: Provides color-coded feedback for success, warnings, and errors.

## Installation

1. **Run the `install.sh`**:

   ```bash
   chmod +x install.sh
   sudo ./install.sh
   ```

---

## Usage

Once the application is installed, you can run the tool to test latency for a list of IP addresses.

```bash
xrayping --config /path/to/config.json --ip-list /path/to/ips.txt --socks5 127.0.0.1:10808
```

The app will print out the results of the latency tests for each IP address, and you can control various parameters using the command-line flags.

---

## Flags

The following flags are available for the application:

| Flag          | Description                                                              | Default                                             |
| ------------- | ------------------------------------------------------------------------ | --------------------------------------------------- |
| `--config`    | Path to the Xray configuration file (required).                          |                                                     |
| `--ip-list`   | Path to the file containing the list of IP addresses to test (required). |                                                     |
| `--xray-path` | Path to the Xray binary.                                                 | `./core/xray`                                       |
| `--socks5`    | Address of the SOCKS5 proxy (e.g., `127.0.0.1:10808`).                   | `127.0.0.1:10808`                                   |
| `--url`       | URL to test latency against.                                             | `http://connectivitycheck.gstatic.com/generate_204` |
| `--retry`     | Number of retry attempts for latency tests.                              | `3`                                                 |
| `--verbose`   | Enable verbose output for more detailed logs.                            | `false`                                             |

---

## Examples

### 1. Basic Usage:

Test a list of IP addresses using default SOCKS5 settings:

```bash
xrayping --config ./config.json --ip-list ./ips.txt
```

### 2. Custom SOCKS5 Proxy:

```bash
xrayping --config ./config.json --ip-list ./ips.txt --socks5 127.0.0.1:8089
```

### 3. Custom URL and Retry Count:

```bash
xrayping --config ./config.json --ip-list ./ips.txt --socks5 127.0.0.1:8089 --url http://example.com --retry 5
```

### 4. Verbose Mode:

```bash
xrayping --config ./config.json --ip-list ./ips.txt --verbose
```

## Build (`Makefile`)

### Explanation of Targets:

1. **`build`**: Builds the application for the current platform.
2. **`build-linux-64`**: Cross-compiles the app for Linux on an amd64 architecture (64-bit).
3. **`build-linux-32`**: Cross-compiles the app for Linux on a 32-bit architecture.
4. **`build-linux-arm32-v7a`**: Cross-compiles the app for Linux on ARM 32-bit (ARMv7-a).
5. **`build-linux-arm64-v8a`**: Cross-compiles the app for Linux on ARM 64-bit (ARMv8-a).
6. **`build-linux`**: Builds the application for all the supported Linux platforms (64-bit, 32-bit, ARM32, ARM64).
7. **`format`**: Formats all Go source code using `go fmt`.
8. **`clean`**: Removes all build artifacts and binaries.
9. **`help`**: Provides a simple help message that lists all available Makefile targets.

### Usage:

- To build for the current platform: `make build`
- To build for specific Linux platforms (e.g., `amd64`): `make build-linux-64`
- To build for all Linux platforms: `make build-linux`
- To format your code: `make format`
- To clean up generated files: `make clean`
- To see available commands: `make help`

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

---

## Contact

For any inquiries or issues, feel free to contact the repository maintainer.

- GalaVPN <galavpn.com@gmail.com>
