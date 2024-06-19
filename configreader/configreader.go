package configreader

import (
	"bufio"
	"os"
	"strings"
	"unicode"

	loger "github.com/dredfort42/tools/logprinter"
)

// ConfigMap is a map containing configuration properties.
type ConfigMap map[string]string

// Global variable to store configuration
var Config ConfigMap = make(ConfigMap)

// Get the configuration from the /app/config.ini file if path parameter is nil or the .ini file located along the path
func GetConfig(path *string) (err error) {
	var configPath string

	if path == nil {
		configPath = "/app/config.ini"
	} else {
		configPath = *path
	}

	var file *os.File
	file, err = os.Open(configPath)

	if err != nil {
		loger.Error("Failed to open config file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) == 0 || !unicode.IsLetter(rune(line[0])) {
			continue
		}

		parameter, _, found := strings.Cut(line, "#")
		if found {
			line = parameter
		}

		before, after, found := strings.Cut(line, "=")
		if found {
			parameter := strings.TrimSpace(before)
			value := strings.TrimSpace(after)
			Config[parameter] = value
		}
	}

	err = scanner.Err()
	if err != nil {
		loger.Error("Failed to read configuration from file", err)
		return
	}

	loger.Success("Successfully read configuration from file", configPath)

	return
}

// PrintConfig prints a ConfigMap to stdout.
func PrintConfig(config ConfigMap) {
	loger.Info("Configuration")
	for key, value := range config {
		loger.Info(key, value)
	}
}
