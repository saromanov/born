package build

import (
	"fmt"

	structs "github.com/saromanov/born/structs/v1"
	"gopkg.in/yaml.v2"
)

// parseConfig provides parsing of the .born.yml file
func parseConfig(data []byte) (*structs.Config, error) {
	var c *structs.Config
	err := yaml.Unmarshal(data, &c)
	if err != nil {
		return nil, fmt.Errorf("unable to parse .born.yml: %v", err)
	}
	if len(c.Steps) > 0 {
		c.Steps = updateMap(c.Steps)
	}
	return c, nil
}

func updateMap(m map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range m {
		result[k] = v
	}
	return result
}
