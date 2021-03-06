package gnb

import (
	"fmt"
	"free5gc-cli/logger"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// GNBConfig is the GNB configuration
var GNBConfig *Config

func checkErr(err error) {
	if err != nil {
		err = fmt.Errorf("[Configuration] %s", err.Error())
		logger.GNBLog.Fatal(err)
	}
}

// InitConfigFactory initialize the module configuration
func InitConfigFactory(f string, force bool) {

	if !force && GNBConfig != nil {
		return
	}

	content, err := ioutil.ReadFile(f)
	checkErr(err)

	GNBConfig = &Config{}

	err = yaml.Unmarshal([]byte(content), &GNBConfig)
	checkErr(err)
	logger.GNBLog.Infof("Successfully load gNB module configuration %s", f)
}
