package config

type mode struct{}

func (mode) Local() string {
	return "local"
}

func (mode) Deploy() string {
	return "deploy"
}

var Mode = mode{}
