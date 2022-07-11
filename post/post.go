package post

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/yuanyp8/alertbot/conf"
	"net/http"
)

func Post(param string) error {

	fmt.Println(param)
	request := gorequest.New().Post(conf.C().Push.OMS.Addr)
	request.QueryData.Set("msg", param)
	resp, _, errs := request.End()
	fmt.Println(request.Url)

	if errs != nil {
		return errs[0]
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", resp.Status)
	}

	return nil
}
