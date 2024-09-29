# Samples

This directory contains sample files used by the program.

## `ips.txt`

The [`ips.txt`](ips.txt) file serves as an example for the input provided via the `--ip-list` flag in the XrayPing command-line tool.

Currently, the `ips.txt` list includes multiple GCore IPv4 addresses.

## `config.json`

The [`config.json`](config.json) file is an example configuration to be used with the `--config` flag in the XrayPing command-line tool.

---

## `cloudflare_ips.txt`

The [`cloudflare_ips.txt`](cloudflare_ips.txt) file contains Cloudflare IP ranges, which can be used with the `random-test` sub-command to randomly select IPs for latency testing.

### Example Usage

```bash
xrayping random-test --subnet-list ./samples/cloudflare_ips.txt --count 25
```

This will randomly select 25 IP addresses from the Cloudflare IP ranges and test their latency using XrayPing.

---

### NOTE

1. If you want to modify the `config.json` file, only the `outbounds` parameter needs to be changed; other parameters should remain unchanged.
2. If necessary, you can modify other JSON keys (e.g., `inbound`, `dns`, etc.).
3. It is not mandatory to change the `outbounds.settings.vnext[0].address` key, but for better readability, it is recommended to leave it _empty_.
