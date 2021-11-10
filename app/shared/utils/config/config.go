package config

import (
	"errors"
	"fmt"
	consulApi "github.com/hashicorp/consul/api"
	"github.com/mubaiedj/go-clean-sample/app/shared/utils/log"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"os"
	"strings"
	"time"
)

const DURATION = time.Second * 5

func LoadSettings(squadName string, appName string, file string) {
	environment := os.Getenv("ENVIRONMENT")
	if "local" == environment || "" == environment {
		err := getLocalConfig()
		if err != nil {
			log.WithError(err).Fatal("no config found")
		}
		return
	}
	if len(environment) == 0 {
		log.Fatal("not environment to run")
	}
	if len(appName) == 0 {
		log.Fatal("not appName to run")
	}
	if len(file) == 0 {
		log.Fatal("not file to run")
	}
	getRemoteConfig(squadName, environment, appName, file)
}

func getRemoteConfig(squadName string, environment string, appName string, file string) {
	path := fmt.Sprintf("%s/%s/%s/%s", squadName, appName, environment, file)
	err := getConsulConfig(path)
	if err != nil {
		log.Error("error trying to get consul config: %s", err.Error())
	}
}

func WatchRemoteConfig(squad string, appName string, folder string) {
	go func(squadName string, appName string, folder string) {
		for {
			LoadSettings(squadName, appName, folder)
			time.Sleep(DURATION) // delay after each request
		}
	}(squad, appName, folder)
}

func Get(key string) interface{} {
	return viper.Get(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetInt64(key string) int64 {
	return viper.GetInt64(key)
}

func GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

func GetDuration(key string) time.Duration {
	return viper.GetDuration(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func AllSettings() map[string]interface{} {
	return viper.AllSettings()
}

func getViperRemoteConfigConsul(path string) error {
	endPoint := os.Getenv("CONSUL_HTTP_ADDR")
	secret := os.Getenv("CONSUL_HTTP_TOKEN")
	viper.AddSecureRemoteProvider("consul", endPoint, path, secret)
	viper.SetConfigType("yaml") // Need to explicitly set this to json
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	return err
}

func getConsulConfig(consulFolder string) error {
	client, err := consulApi.NewClient(consulApi.DefaultConfig())
	if err != nil {
		log.Error(fmt.Sprintf("error creating consul api client: %s", err.Error()))
	}
	log.Info("consul folder: %s", consulFolder)
	consulConfigs, _, err := client.KV().Get(consulFolder, nil)
	if err != nil {
		log.WithError(err).Error("error trying to get config from consul.")
	}
	if consulConfigs == nil {
		err = errors.New("no config found")
		log.WithError(err).Error("error trying to get config from consul")
	}
	consulRAW := string(consulConfigs.Value)
	viper.SetConfigType("yaml")
	err = viper.MergeConfig(strings.NewReader(consulRAW))
	if err != nil {
		log.WithError(err).Error("error trying to read config from viper")
	}
	return err
}

func getLocalConfig() error {
	viper.SetConfigName("dev.config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return errors.New(fmt.Sprintf("Fatal error config file: %s \n", err))
	}
	return nil
}
