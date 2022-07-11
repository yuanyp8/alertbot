package apps

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yuanyp8/alertbot/alert"
	"github.com/yuanyp8/alertbot/post"
	"log"
)

var handler = &Handler{}

type Handler struct{}

func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/alerts", h.Post)
}

func (h *Handler) Post(c *gin.Context) {
	ret := alert.NewNotification()
	if err := c.BindJSON(ret); err != nil {

		if err := post.Post("alert metrics parse failed"); err != nil {
			fmt.Println("----------")
			log.Fatalf("alert metrics push failed, error: %s", err)
		}
	}
	slc := ret.ToString()
	for i, _ := range slc {
		if err := post.Post(slc[i]); err != nil {
			fmt.Println("++++++++")
			log.Fatalf("alert metrics push failed, error: %s", err)
		}
	}

}

func H() *Handler {
	return handler
}
