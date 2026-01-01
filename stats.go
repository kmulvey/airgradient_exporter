package main

import "github.com/prometheus/client_golang/prometheus"

type AirGradientMeasures struct {
	PM01            float64 `json:"pm01"`
	PM02            float64 `json:"pm02"`
	PM10            float64 `json:"pm10"`
	PM01Standard    float64 `json:"pm01Standard"`
	PM02Standard    float64 `json:"pm02Standard"`
	PM10Standard    float64 `json:"pm10Standard"`
	PM003Count      float64 `json:"pm003Count"`
	PM005Count      float64 `json:"pm005Count"`
	PM01Count       float64 `json:"pm01Count"`
	PM02Count       float64 `json:"pm02Count"`
	PM50Count       float64 `json:"pm50Count"`
	PM10Count       float64 `json:"pm10Count"`
	PM02Compensated float64 `json:"pm02Compensated"`
	Atmp            float64 `json:"atmp"`
	AtmpCompensated float64 `json:"atmpCompensated"`
	Rhum            float64 `json:"rhum"`
	RhumCompensated float64 `json:"rhumCompensated"`
	Rco2            float64 `json:"rco2"`
	TvocIndex       float64 `json:"tvocIndex"`
	TvocRaw         float64 `json:"tvocRaw"`
	NoxIndex        int     `json:"noxIndex"`
	NoxRaw          float64 `json:"noxRaw"`
	Boot            int     `json:"boot"`
	BootCount       int     `json:"bootCount"`
	Wifi            int     `json:"wifi"`
	LedMode         string  `json:"ledMode"`
	SerialNo        string  `json:"serialno"`
	Firmware        string  `json:"firmware"`
	Model           string  `json:"model"`
}

var (
	// Device info
	airgradientInfo = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "airgradient_info",
			Help: "AirGradient device information",
		},
		[]string{"serial_number", "firmware", "model"},
	)

	// WiFi
	wifiRSSI = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "airgradient_wifi_rssi_dbm",
			Help: "WiFi signal strength in dBm",
		},
	)

	// PM measurements (atmospheric environment)
	pm01 = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "airgradient_pm01_ugm3",
			Help: "PM1.0 in ug/m3 (atmospheric environment)",
		},
	)

	pm02 = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "airgradient_pm02_ugm3",
			Help: "PM2.5 in ug/m3 (atmospheric environment)",
		},
	)

	pm10 = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "airgradient_pm10_ugm3",
			Help: "PM10 in ug/m3 (atmospheric environment)",
		},
	)

	pm02Compensated = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "airgradient_pm02_compensated_ugm3",
			Help: "PM2.5 in ug/m3 with correction applied (from fw version 3.1.4 onwards)",
		},
	)

	// PM measurements (standard particle)
	pm01Standard = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "airgradient_pm01_standard_ugm3",
			Help: "PM1.0 in ug/m3 (standard particle)",
		},
	)

	pm02Standard = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "airgradient_pm02_standard_ugm3",
			Help: "PM2.5 in ug/m3 (standard particle)",
		},
	)

	pm10Standard = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "airgradient_pm10_standard_ugm3",
			Help: "PM10 in ug/m3 (standard particle)",
		},
	)

	// Particle counts
	pm003Count = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "airgradient_pm003_count_pdl",
			Help: "Particle count 0.3um per dL",
		},
	)

	pm005Count = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "airgradient_pm005_count_pdl",
			Help: "Particle count 0.5um per dL",
		},
	)

	pm01Count = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "airgradient_pm01_count_pdl",
			Help: "Particle count 1.0um per dL",
		},
	)

	pm02Count = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "airgradient_pm02_count_pdl",
			Help: "Particle count 2.5um per dL",
		},
	)

	pm50Count = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "airgradient_pm50_count_pdl",
			Help: "Particle count 5.0um per dL (only for indoor monitor)",
		},
	)

	pm10Count = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "airgradient_pm10_count_pdl",
			Help: "Particle count 10um per dL (only for indoor monitor)",
		},
	)

	// CO2
	rco2 = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "airgradient_co2_ppm",
			Help: "CO2 in ppm",
		},
	)

	// Temperature
	atmp = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "airgradient_temperature_celsius",
			Help: "Temperature in Degrees Celsius",
		},
	)

	atmpCompensated = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "airgradient_temperature_compensated_celsius",
			Help: "Temperature in Degrees Celsius with correction applied",
		},
	)

	// Humidity
	rhum = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "airgradient_humidity_percent",
			Help: "Relative Humidity",
		},
	)

	rhumCompensated = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "airgradient_humidity_compensated_percent",
			Help: "Relative Humidity with correction applied",
		},
	)

	// VOC
	tvocIndex = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "airgradient_tvoc_index",
			Help: "Senisiron VOC Index",
		},
	)

	tvocRaw = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "airgradient_tvoc_raw",
			Help: "VOC raw value",
		},
	)

	// NOx
	noxIndex = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "airgradient_nox_index",
			Help: "Senisirion NOx Index",
		},
	)

	noxRaw = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "airgradient_nox_raw",
			Help: "NOx raw value",
		},
	)

	// Boot counter
	bootCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "airgradient_boots",
			Help: "Counts every measurement cycle. Low boot counts indicate restarts.",
		},
	)

	// LED mode
	ledModeInfo = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "airgradient_led_mode_info",
			Help: "Current configuration of the LED mode",
		},
		[]string{"led_mode"},
	)
)

func init() {
	prometheus.MustRegister(
		airgradientInfo,
		wifiRSSI,
		pm01,
		pm02,
		pm10,
		pm02Compensated,
		pm01Standard,
		pm02Standard,
		pm10Standard,
		pm003Count,
		pm005Count,
		pm01Count,
		pm02Count,
		pm50Count,
		pm10Count,
		rco2,
		atmp,
		atmpCompensated,
		rhum,
		rhumCompensated,
		tvocIndex,
		tvocRaw,
		noxIndex,
		noxRaw,
		bootCount,
		ledModeInfo,
	)
}

// updateMetrics updates all Prometheus metrics from AirGradient measures
func updateMetrics(m AirGradientMeasures) {
	airgradientInfo.WithLabelValues(m.SerialNo, m.Firmware, m.Model).Set(1)
	wifiRSSI.Set(float64(m.Wifi))
	pm01.Set(m.PM01)
	pm02.Set(m.PM02)
	pm10.Set(m.PM10)
	pm02Compensated.Set(m.PM02Compensated)
	pm01Standard.Set(m.PM01Standard)
	pm02Standard.Set(m.PM02Standard)
	pm10Standard.Set(m.PM10Standard)
	pm003Count.Set(m.PM003Count)
	pm005Count.Set(m.PM005Count)
	pm01Count.Set(m.PM01Count)
	pm02Count.Set(m.PM02Count)
	pm50Count.Set(m.PM50Count)
	pm10Count.Set(m.PM10Count)
	rco2.Set(float64(m.Rco2))
	atmp.Set(m.Atmp)
	atmpCompensated.Set(m.AtmpCompensated)
	rhum.Set(m.Rhum)
	rhumCompensated.Set(m.RhumCompensated)
	tvocIndex.Set(m.TvocIndex)
	tvocRaw.Set(m.TvocRaw)
	noxIndex.Set(float64(m.NoxIndex))
	noxRaw.Set(m.NoxRaw)
	bootCount.Set(float64(m.BootCount))
	ledModeInfo.WithLabelValues(m.LedMode).Set(1)
}
