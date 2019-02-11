package v1

// Config provides definition of the config at .born.yml
type Config struct {
	Pipeline  string                 `yaml:"pipeline"`
	Name      string                 `yaml:"name"`
	PreScript []string               `yaml:"pre_script"`
	Steps     map[string]interface{} `yaml:"steps"`
}

// StepConfig defines step at the .born.yml
type StepConfig struct {
	Image    string
	Commands []string `yaml:"commands"`
}
