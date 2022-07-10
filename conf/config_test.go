package conf_test

import (
	"fmt"
	"github.com/yuanyp8/alertbot/conf"
	"os"
	"testing"
)

func TestLoadConfOnlyFile(t *testing.T) {
	fmt.Printf("app http_addr: %+v\n", conf.C().App.HttpApi)
}

func TestLoadConfWithEnv(t *testing.T) {
	_ = os.Setenv("ALERT_ADDR", "http://dsdasa")
	fmt.Printf("AlertManager: %+v\n", conf.C().AlertManager.HttpAddr())
}

func TestConfig_Validate(t *testing.T) {
	conf.C().AlertManager.API = ""
	if err := conf.C().Validate(); err != nil {
		fmt.Printf("缺少必要字段\n")
	}
}

func init() {
	if err := conf.C().LoadConf("../etc/config.yaml"); err != nil {
		panic(err)
	}
}
