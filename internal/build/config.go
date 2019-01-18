package build

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// parseConfig provides parsing of the .born.yml file
func parseConfig(data []byte) (*Config, error) {
	var c *Config
	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		return nil, fmt.Errorf("unable to parse .born.yml: %v", err)
	}

	return c, nil
}
