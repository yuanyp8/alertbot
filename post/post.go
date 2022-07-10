package post

import (
	"fmt"
	"github.com/yuanyp8/alertbot/conf"
	"net/http"
)

func Post(param string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", conf.C().Push.OMS.Addr, param), nil)

	if err != nil {
		return err
	}

	response, err := client.Do(req)

	defer response.Body.Close()

	if err != nil {
		return err
	}
	return nil
}
