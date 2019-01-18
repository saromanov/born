package v1

// Config provides definition of the config at .born.yml
type Config struct {
	Pipeline string       `yaml:"pipeline"`
	Name     string       `yaml:"name"`
	Steps    []StepConfig `yaml:"steps"`
}

// StepConfig defines step at the .born.yml
type StepConfig struct {
	Image    string
	Commands []string `yaml:"commands"`
}
