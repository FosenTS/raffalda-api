package config

import (
	"path"
	"raffalda-api/pkg/mysync"
)

const merchandiseParserConfigFilename = "merchandise.parser.config.yaml"

type MerchandiseParserConfig struct {
	DataFilePath string `yaml:"dataFilePath" env-required:"true"`

	DataFileAbsPath string
}

var (
	merchandiseParserConfigInst     = &MerchandiseParserConfig{}
	loadMerchandiseParserConfigOnce = mysync.NewOnce()
)

func MerchandiseParser() MerchandiseParserConfig {
	loadMetricsConfigOnce.Do(func() {
		env := Env()
		readConfig(path.Join(env.ConfigAbsPath, merchandiseParserConfigFilename), merchandiseParserConfigInst)
		merchandiseParserConfigInst.DataFileAbsPath = path.Join(env.ProjectAbsPath, merchandiseParserConfigInst.DataFilePath)
	})

	return *merchandiseParserConfigInst
}
