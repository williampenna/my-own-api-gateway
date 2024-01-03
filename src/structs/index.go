package kong

type Service struct {
	Name    string   `yaml:"name"`
	URL     string   `yaml:"url"`
	Plugins []Plugin `yaml:"plugins"`
	Routes  []Route  `yaml:"routes"`
}

type Plugin struct {
	Name  string         `yaml:"name"`
	Input map[string]any `yaml:"input,omitempty"`
}

type Route struct {
	Name    string   `yaml:"name"`
	Paths   []string `yaml:"paths"`
	Methods []string `yaml:"methods"`
}
