package configreader

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	loger "github.com/dredfort42/tools/logprinter"
)

// ConfigMap is a map containing configuration properties.
type ConfigMap map[string]string

// Global variable to store configuration
var config ConfigMap = make(ConfigMap)

// ReadConfig reads a configuration file to a ConfigMap and returns an error if it fails.
func ReadConfig(path string) error {
	file, err := os.Open(path)

	if err != nil {
		loger.PrintWarning("Failed to open file", path)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "#") && len(line) != 0 {
			before, after, found := strings.Cut(line, "=")
			if found {
				parameter := strings.TrimSpace(before)
				value := strings.TrimSpace(after)
				config[parameter] = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	loger.Success("Successfully read configuration from file", path)

	return nil
}

// Get configuration from global and local .cfg files and returns a ConfigMap and an error if it fails.
func GetConfig() (ConfigMap, error) {
	success := false

	// Read global config file
	if err := ReadConfig("/app/global.cfg"); err == nil {
		success = true
	} else if err := ReadConfig("./global.cfg"); err == nil {
		success = true
	}

	// Read local config file
	if err := ReadConfig("/app/local.cfg"); err == nil {
		success = true
	} else if err := ReadConfig("./local.cfg"); err == nil {
		success = true
	}

	if !success {
		loger.Error("Failed to read configuration", nil)
		return nil, fmt.Errorf("Failed to read configuration")
	} else {
		loger.Success("Successfully read configuration")
		return config, nil
	}
}

// PrintConfig prints a ConfigMap to stdout.
func PrintConfig(config ConfigMap) {
	loger.Info("Configuration")
	for key, value := range config {
		loger.Info(key, value)
	}
}
