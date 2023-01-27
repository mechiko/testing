package logger

import (
	"os"
	"sync"
)

var once sync.Once

var instance *Logger = nil

func InitLogger(config *ConfigZero) *Logger {
	// defer recoverFunc("initLogger")
	once.Do(func() {
		if config.Truncate {
			os.Remove(config.Filename)
		}
		instance = config.configure()
	})
	return instance
}
