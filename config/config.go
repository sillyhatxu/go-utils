package config

import (
	"bytes"
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func ParseConfig(configFile string, input interface{}) {
	if fileInfo, err := os.Stat(configFile); err != nil {
		if os.IsNotExist(err) {
			log.Panicf("configuration file [%v] does not exist.", configFile)
		} else {
			log.Panicf("configuration file [%v] can not be stated. %v", configFile, err)
		}
	} else {
		if fileInfo.IsDir() {
			log.Panicf("%v is a directory name", configFile)
		}
	}

	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Panicf("read configuration file error. %v", err)
	}
	content = bytes.TrimSpace(content)

	err = toml.Unmarshal(content, &input)
	if err != nil {
		log.Panicf("unmarshal toml object error. %v", err)
	}
}
