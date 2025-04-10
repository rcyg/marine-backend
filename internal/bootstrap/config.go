package bootstrap

import (
	"marine-backend/cmd/flags"
	"marine-backend/internal/conf"
	"marine-backend/pkg/utils"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func InitConfig() {
	// TODO:
	configPath := filepath.Join(flags.DataDir, "config.json")
	log.Infof("reading config file: %s", configPath)
	if !utils.Exists(configPath) {
		log.Infof("config file not exists, creating default config file")
		_, err := utils.CreateNestedFile(configPath)
		if err != nil {
			log.Fatalf("failed to create config file: %+v", err)
		}
		conf.Conf = conf.DefaultConfig()
		if !utils.WriteJsonToFile(configPath, conf.Conf) {
			log.Fatalf("failed to create default config file")
		}
	} else {
		configBytes, err := os.ReadFile(configPath)
		if err != nil {
			log.Fatalf("reading config file error: %+v", err)
		}
		conf.Conf = conf.DefaultConfig()
		err = utils.Json.Unmarshal(configBytes, conf.Conf)
		if err != nil {
			log.Fatalf("load config error: %+v", err)
		}
		// update config.json struct
		confBody, err := utils.Json.MarshalIndent(conf.Conf, "", "  ")
		if err != nil {
			log.Fatalf("marshal config error: %+v", err)
		}
		err = os.WriteFile(configPath, confBody, 0o777)
		if err != nil {
			log.Fatalf("update config struct error: %+v", err)
		}
	}
	log.Debugf("config: %+v", conf.Conf)
}
