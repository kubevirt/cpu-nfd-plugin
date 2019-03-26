package util

type Config struct {
	ObsoleteCPUs []string `yaml:"obsoleteCPUs"`
	MinCPU       MinCPU   `yaml:"minCPU"`
}
type MinCPU struct {
	Intel string `yaml:"intel"`
	AMD   string `yaml:"amd"`
}
