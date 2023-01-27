package entity

// import "github.com/spf13/viper"

type IConfiguration interface{}

type ConfigInterface interface {
	Get() interface{}
	Save() error
	SaveAs(fn string) error
	SaveSafe() error
	GetByName(name string) interface{}
	Set(key string, value interface{}, save ...bool) error
	// GetViper() (*viper.Viper, error)
	Unmarshal(*Configuration) error
}
