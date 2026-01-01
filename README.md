# AirGradient Exporter

A Prometheus exporter for [AirGradient](https://www.airgradient.com/) air quality monitors. This exporter fetches data from the AirGradient device's local API and exposes it as Prometheus metrics.

## Features

- Scrapes the `/measures/current` endpoint of AirGradient devices.
- Exposes a wide range of metrics including PM1, PM2.5, PM10, CO2, Temperature, Humidity, TVOC, and NOx.
- Includes a Grafana dashboard for visualization.
- Supports systemd for running as a service.

## Installation

### From Source

```bash
go install github.com/kmulvey/airgradient_exporter@latest
```

### Building Locally

```bash
git clone https://github.com/kmulvey/airgradient_exporter.git
cd airgradient_exporter
go build
```

## Usage

```bash
./airgradient_exporter [flags]
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-addr` | `:9500` | Address for the Prometheus server to listen on. |
| `-url` | `http://airgradient.local/measures/current` | URL to fetch AirGradient measurements from. Replace `airgradient.local` with your device's IP or hostname. |
| `-timeout` | `30s` | Timeout duration between measurements. |
| `-version`, `-v` | `false` | Print version information. |

### Example

```bash
./airgradient_exporter -url http://192.168.1.50/measures/current -addr :9500
```

## Systemd Service

A systemd service file is provided in `airgradient_exporter.service`.

1. Copy the binary to `/usr/bin/`:
   ```bash
   sudo cp airgradient_exporter /usr/bin/
   ```

2. Copy the service file to `/etc/systemd/system/`:
   ```bash
   sudo cp airgradient_exporter.service /etc/systemd/system/
   ```

3. Edit the service file to configure the `-url` flag if your device is not at `http://airgradient.local/measures/current`.

4. Enable and start the service:
   ```bash
   sudo systemctl enable airgradient_exporter
   sudo systemctl start airgradient_exporter
   ```

## Metrics

The exporter exposes the following metrics (prefixed with `airgradient_`):

- **PM (Atmospheric & Standard):** `pm01`, `pm02`, `pm10`
- **Particle Counts:** `pm003_count`, `pm005_count`, `pm01_count`, `pm02_count`, `pm50_count`, `pm10_count`
- **CO2:** `co2_ppm`
- **Temperature:** `temperature_celsius`, `temperature_compensated_celsius`
- **Humidity:** `humidity_percent`, `humidity_compensated_percent`
- **TVOC:** `tvoc_index`, `tvoc_raw`
- **NOx:** `nox_index`, `nox_raw`
- **WiFi:** `wifi_rssi_dbm`
- **Device Info:** `info` (labels: serial_number, firmware, model), `boots`

## Grafana Dashboard

A Grafana dashboard is included in `grafana-dashboard.json`. You can import this JSON file directly into Grafana to visualize your AirGradient data.

## License

[MIT](LICENSE)
