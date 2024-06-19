package configreader

import (
	"bufio"
	"errors"
	"flag"
	"os"
	"strings"
	"unicode"
)

// ConfigMap is a map containing configuration properties.
type ConfigMap map[string]string

// Global variable to store configuration
var Config ConfigMap = make(ConfigMap)

// Get the configuration from the /app/config.ini file if path parameter is nil or the .ini file located along the path
func GetConfig() (err error) {
	var path *string = flag.String("config", "/app/config.ini", "Path to the configuration file")

	flag.Parse()

	var file *os.File
	file, err = os.Open(*path)

	if err != nil {
		return errors.New("use the --config flag to specify the path to the .ini configuration file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) == 0 || !unicode.IsLetter(rune(line[0])) {
			continue
		}

		line = strings.Split(line, "#")[0]
		line = strings.Split(line, ";")[0]

		before, after, found := strings.Cut(line, "=")
		if found {
			parameter := strings.TrimSpace(before)
			value := strings.TrimSpace(after)
			Config[parameter] = value
		}
	}

	err = scanner.Err()
	if err != nil {
		return errors.New("failed to read configuration from .ini file")
	}

	return
}
