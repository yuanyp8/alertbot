package conf

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"sync"
)

var (
	config   *Config
	validate = validator.New()
	once     = sync.Once{}
)

type Config struct {
	AlertManager *AlertManager `mapstructure:"alertmanager"`
	Push         *Push         `mapstructure:"push"`
	App          *App          `mapstructure:"app"`
}

type AlertManager struct {
	Addr string `mapstructure:"addr" env:"ALERT_ADDR" validate:"required"`
	API  string `mapstructure:"api" validate:"required"`
}

type Push struct {
	OMS *OMS
}
type OMS struct {
	Addr string `mapstructure:"addr" env:"OMS_ADDR" validate:"required"`
}

type App struct {
	HttpApi string `mapstructure:"http_addr" env:"HttpApi"`
}

func NewOMS() *OMS {
	return &OMS{}
}

func NewConfig() *Config {
	return &Config{
		AlertManager: NewAlertManager(),
		Push:         NewPush(),
		App:          NewApp(),
	}
}

func NewAlertManager() *AlertManager {
	return &AlertManager{}
}

func NewPush() *Push {
	return &Push{
		OMS: NewOMS(),
	}
}

func NewApp() *App {
	return &App{}
}

func (a *AlertManager) HttpAddr() string {
	return fmt.Sprintf("http://%s:%s", a.Addr, a.API)
}

// Validate 验证配置文件必填字段是否合法
func (c *Config) Validate() error {
	return validate.Struct(c)
}

func (c *Config) LoadConf(pathname string) error {
	var e error
	once.Do(func() {
		// load config from the given pathname file
		vip := viper.New()
		vip.SetConfigFile(pathname)

		if err := vip.ReadInConfig(); err != nil {
			err = fmt.Errorf("Fatal error loading config file: %w \n", err)
			return
		}

		if err := vip.Unmarshal(c); err != nil {
			err = fmt.Errorf("Fatal error unmarshal config: %w \n", err)
			return
		}

		// load config from env
		if err := env.Parse(c); err != nil {
			err = fmt.Errorf("Fatal error parse env config: %w \n", err)
			return
		}
		e = c.Validate()
	})
	return e
}

func init() {
	config = NewConfig()
}

// C @Required: conf.LoadConf
func C() *Config {
	return config
}
