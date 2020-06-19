package api

import (
	"encoding/json"
	// "fmt"
	// "github.com/MuhtasimTanmoy/messaging_server/internal/package/driver"
	// "github.com/MuhtasimTanmoy/messaging_server/internal/package/logger"
)

const namespace string = "config"

// ConfigResult struct
type ConfigResult struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// LoadFromJSON load object from json
func (c *ConfigResult) LoadFromJSON(data []byte) (bool, error) {
	err := json.Unmarshal(data, &c)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ConvertToJSON converts object to json
func (c *ConfigResult) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&c)
	if err != nil {
		return "", err
	}
	return string(data), nil
}