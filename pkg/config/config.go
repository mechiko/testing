package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

// в этом модуле подогнано все с колен :) но не охота переделывать... потом

type IConfiguration interface{}
type Config struct {
	Configuration  IConfiguration
	ConfigFileName string
	Viper          *viper.Viper
}

var cfg *Config = nil

// var ConfigFileName = ""

var TomlConfig []byte

var doOnce sync.Once

const defaultConfigName = "config"

// меняем каталог на каталог запуска это не помню от куда велосипед
func init() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	os.Chdir(dir)
}

// на уровне модуля доступ к глобальным данным модуля
func GetViper() (*viper.Viper, error) {
	if cfg.Viper == nil {
		return nil, fmt.Errorf("%s", "get viper error")
	}
	return cfg.Viper, nil
}

// возвращаем указатель на структуру
// type config struct {
// 	*Configuration
// 	viper *viper.Viper
// }
// на уровне модуля доступ к глобальным данным модуля
// вызывается в application.initConfig()
func GetInstance(cfgName string, configuration IConfiguration, dbg ...string) (*Config, error) {
	var err error
	doOnce.Do(func() {
		defer recoverFunc("GetInstance")
		if len(dbg) > 0 {
			fmt.Printf("package:Config file:configuration.go func:GetInstance() msg:doOnce invoked by %v\n", dbg[0])
		}
		err = initConfiguration(cfgName, configuration)
	})
	if err != nil {
		return nil, fmt.Errorf("GetInstance() %w", err)
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

// приватно для модуля
func initConfiguration(cfgName string, configuration IConfiguration) error {
	defer recoverFunc("InitConfiguration")

	configName := defaultConfigName
	if cfgName != "" {
		configName = cfgName
	}
	// fmt.Printf("comfigName = %v\n", configName)
	viperOrigin := viper.GetViper()
	// здесь зачем то беру путь запуска... легче брать локальную папку
	// pwd, _ := os.Getwd()
	// pwd := "."
	// ConfigFileName = filepath.Join(pwd, configName+".toml")
	configFileName := configName + ".toml"
	// fmt.Printf("ConfigFileName = %v\n", ConfigFileName)

	viperOrigin.SetConfigName(configName)
	viperOrigin.SetConfigType("toml")
	viperOrigin.AddConfigPath(".")

	if err := viperOrigin.MergeConfig(strings.NewReader(string(TomlConfig))); err != nil {
		return fmt.Errorf("viperOrigin.MergeConfig() %w", err)
	}
	// сливаем конфиг из переменной пакета tomlConfig и файла конфига
	// если файла нет пишем в консоль сообщение отладки
	// если другая ошибка то возвращаем ее
	if err := viper.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("viper.ConfigFileNotFoundError %w", err)
		} else {
			fmt.Printf("Config file ('%s') not found\n\r", configFileName)
		}
	}
	// парсим в структуру содержимое конфига и возвращаем ошибку
	// var conf Configuration
	if err := viper.Unmarshal(configuration); err != nil {
		return fmt.Errorf("viper.Unmarshal(configuration) %w", err)
	}

	cfg = &Config{
		Configuration:  configuration,
		ConfigFileName: configFileName,
		Viper:          nil,
	}
	// записываем конфиг если файла нет (SafeWriteConfig() Записывает если его нет)
	cfg.Viper = viperOrigin
	viperOrigin.SafeWriteConfig()
	// if err := viperOrigin.SafeWriteConfig(); err != nil {
	// 	// пропустим
	// }
	return nil
}

func Get() IConfiguration {
	var err error
	if cfg == nil {
		// init default
		TomlConfig = tomlConfig2
		conf := &DefaultConfiguration{}
		if cfg, err = GetInstance(defaultConfigName, conf); err != nil {
			fmt.Printf("default config module err = %v\n", err.Error())
		}
	}
	return cfg.Configuration
}

// на уровне модуля доступ к глобальным данным модуля
// type Config struct {
// 	Configuration IConfiguration
// 	Viper         *viper.Viper
// }
func GetConfig() *Config {
	var err error
	if cfg == nil {
		// init default
		TomlConfig = tomlConfig2
		conf := &DefaultConfiguration{}
		if cfg, err = GetInstance(defaultConfigName, conf); err != nil {
			fmt.Printf("default config module err = %v\n", err.Error())
		}
	}
	return cfg
}

func recoverFunc(s string) {
	if r := recover(); r != nil {
		err := fmt.Sprintf("%s %v", s, r)
		d := []byte(err)
		_ = ioutil.WriteFile("error.txt", d, 0644)
		// messageErr(err)
		os.Exit(1)
	}
}
