package config

import (
	"fmt"
	"os"
)

// это структура конфигурации потому используем пустой интерфейс он принимает все
func (c *Config) Get() IConfiguration {
	return cfg.Configuration
}

// тут функции надстройки возможно и не так полезные, хотя для записи значений в конфиг используются
// Viper пишет только то что прочитал изначально из конфига, если меняем что то
// то надо усатнавливать через Set а не просто в Configuration
func (c *Config) Save() error {
	defer recoverFunc("Save()")
	if _, err := os.Stat(c.ConfigFileName); os.IsNotExist(err) {
		if err := c.Viper.SafeWriteConfig(); err != nil {
			return fmt.Errorf("Viper.SafeWriteConfig() %w", err)
		}
	} else {
		if err := c.Viper.WriteConfig(); err != nil {
			return fmt.Errorf("Viper.WriteConfig() %w", err)
		}
	}
	return nil
}

func (c *Config) SaveAs(fn string) error {
	defer recoverFunc("SaveAs")
	err := c.Viper.WriteConfigAs(fn)
	if err != nil {
		return fmt.Errorf("Viper.WriteConfigAs() %w", err)
	}
	return nil
}

func (c *Config) SaveSafe() error {
	defer recoverFunc("SaveSafe()")

	err := c.Viper.SafeWriteConfig()
	if err != nil {
		return fmt.Errorf("Viper.SafeWriteConfig() %w", err)
	}
	return nil
}

// GetString("datastore.metric.host")
func (c *Config) GetByName(name string) interface{} {
	return c.Viper.Get(name)
}

// записываем ключ и его значение, затем обновляем структуру Config этими значениями
func (c *Config) Set(key string, value interface{}, save ...bool) error {
	c.Viper.Set(key, value)
	if err := c.Viper.Unmarshal(c.Configuration); err != nil {
		return fmt.Errorf("Viper.Unmarshal(c.Configuration) %w", err)
	}
	if len(save) > 0 {
		saving := save[0]
		if saving {
			if err := c.Save(); err != nil {
				return fmt.Errorf("Save() %w", err)
			}
		}
	}
	return nil
}

// func messageErr(str string) {
// 	dialog.Message("%s", str).Title("Ошибка Configuration").Info()
// }

// func fileNameWithoutExtension(fileName string) string {
// 	return strings.TrimSuffix(fileName, filepath.Ext(filepath.Base(fileName)))
// }
