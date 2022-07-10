package apps

import (
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
	ret := make([]*alert.Alert, 0, 300)
	if err := c.Bind(ret); err != nil {
		//
		if err := post.Post("alert metrics parse failed"); err != nil {
			log.Fatalf("alert metrics push failed, error: %s", err)
		}
	}
	for i, _ := range ret {
		if err := post.Post(ret[i].ToString()); err != nil {
			log.Fatalf("alert metrics push failed, error: %s", err)
		}
	}

}

func H() *Handler {
	return handler
}
