package configutils

import (
	"strings"

	"github.com/panprogramadorgh/quickgoserver/pkg/fileutils"
)

func ReadConf(confKey string, confPath interface{}) (string, error) {
	var defaultConfPath string = "./pkg/confutils/go-server.conf"
	if str, ok := confPath.(string); ok {
		defaultConfPath = str
	}
	conf, err := fileutils.ReadFile(defaultConfPath)
	if err != nil {
		return "", nil
	}
	var confLines []string = strings.Split(conf, ";")

	if confKey == "*" {
		return strings.Join(confLines, ", "), nil
	}

	var confValue string
	for _, confLine := range confLines {
		keyValue := strings.Split(confLine, "=")
		if len(keyValue) < 2 {
			panic("malformed config")
		}
		key := strings.Trim(keyValue[0], " ")
		value := strings.Trim(keyValue[1], " ")
		if key == confKey {
			confValue = value
		}
	}

	return confValue, nil
}
