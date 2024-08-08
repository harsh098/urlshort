package internal

type Urls struct {
	Path   string `yaml:"path"`
	Target string `yaml:"target"`
}

type Config struct {
	Host  string `yaml:"host"`
	Port  string `yaml:"port"`
	Paths []Urls `yaml:"urls"`
}
