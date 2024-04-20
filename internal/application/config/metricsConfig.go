package config

import "raffalda-api/pkg/mysync"

type MetricsConfig struct {
	IpAddress string
	Port      string
}

var (
	metricsConfigInst     = &MetricsConfig{}
	loadMetricsConfigOnce = mysync.NewOnce()
)

func Metrics() MetricsConfig {
	loadMetricsConfigOnce.Do(func() {
		env := Env()
		metricsConfigInst.IpAddress = env.IpAddress
		metricsConfigInst.Port = env.MetricPort
	})
	return *metricsConfigInst
}
