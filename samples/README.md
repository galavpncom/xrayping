# Samples

This directory contains sample files used by the program.

## `ips.txt`

The [`ips.txt`](ips.txt) file serves as an example for the input provided via the `--ip-list` flag in the XrayPing command-line tool.

Currently, the `ips.txt` list includes 28 GCore IPv4 addresses. We are working on adding new listings to this repository.

## `config.json`

The [`config.json`](config.json) file is an example configuration to be used with the `--config` flag in the XrayPing command-line tool.

---

### NOTE

1. If you want to modify the `config.json` file, only the `outbounds` parameter needs to be changed; other parameters should remain unchanged.
2. If necessary, you can modify other JSON keys (e.g., `inbound`, `dns`, etc.).
3. It is not mandatory to change the `outbounds.settings.vnext[0].address` key, but for better readability, it is recommended to leave it _empty_.
