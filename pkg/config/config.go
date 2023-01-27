package config

import (
	"fmt"
	"runtime"
	"strings"
	"testing/internal/entity"

	"github.com/spf13/viper"
)

type Config struct {
	entity.ConfigInterface
	app            entity.App
	viper          *viper.Viper
	configuration  interface{}
	configFileName string
}

var TomlConfig []byte

const defaultConfigName = "config"

// создание нового экземпляра конфига реентерабельно возвращается интерфейс
// dbg ...string для отладки передаем строки первая печатается в консоль.. это видимо был тест параметров, пусть так пока
func NewInstance(app entity.App, cfgName string, configuration interface{}, dbg ...string) (*Config, error) {
	defer RecoverFmt("config:NewInstance")
	var cfg *Config
	var err error
	if len(dbg) > 0 {
		fmt.Printf("package:Config file:configuration.go func:GetInstance() msg:doOnce invoked by %v\n", dbg[0])
	}
	if cfg, err = initConfiguration(app, cfgName, configuration); err != nil {
		return nil, fmt.Errorf("config:NewInstance() %w", err)
	}
	if len(dbg) > 0 {
		fmt.Printf("Dbg=%v\n", dbg[0])
		pc, file, no, ok := runtime.Caller(1)
		if ok {
			details := runtime.FuncForPC(pc)
			fmt.Printf("package:Config file:configuration.go func:GetInstance() msg:called from %s#%d\n", file, no)
			if details != nil {
				fmt.Printf("package:Config file:configuration.go func:GetInstance() msg:called from %s\n", details.Name())
			}
		}
	}
	return cfg, nil
}

func initConfiguration(app entity.App, cfgName string, configuration interface{}) (*Config, error) {
	defer RecoverFmt("InitConfiguration")

	configName := defaultConfigName
	if cfgName != "" {
		configName = cfgName
	}
	viperOrigin := viper.GetViper()
	configFileName := configName + ".toml"

	viperOrigin.SetConfigName(configName)
	viperOrigin.SetConfigType("toml")
	viperOrigin.AddConfigPath(".")

	if err := viperOrigin.MergeConfig(strings.NewReader(string(TomlConfig))); err != nil {
		return nil, fmt.Errorf("viperOrigin.MergeConfig() %w", err)
	}

	if err := viper.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("viper.ConfigFileNotFoundError %w", err)
		} else {
			fmt.Printf("Config file ('%s') not found\n\r", configFileName)
		}
	}

	if err := viper.Unmarshal(configuration); err != nil {
		return nil, fmt.Errorf("viper.Unmarshal(configuration) %w", err)
	}

	cfg := &Config{
		app:            app,
		configuration:  configuration,
		configFileName: configFileName,
		viper:          viperOrigin,
	}
	viperOrigin.SafeWriteConfig()

	return cfg, nil
}
