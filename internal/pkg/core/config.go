package core

type (
	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}
	Http struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
)
