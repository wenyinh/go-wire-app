package user

import (
	"github.com/gin-gonic/gin"
	"github.com/wenyinh/go-wire-app/pkg/service"
)

const (
	UserUriPrefix = "/user"
)

type Controller struct {
	UriPrefix string
	Service   service.UserService
}

func NewController(service service.UserService) (*Controller, error) {
	return &Controller{
		UriPrefix: UserUriPrefix,
		Service:   service,
	}, nil
}

func (ctrl *Controller) RegisterRoutes(g *gin.RouterGroup) {
	api := g.Group(ctrl.UriPrefix)

	// add chat API
	api.POST(CreateUserUri, ctrl.CreateUser)
	api.GET(GetUserUri, ctrl.GetUser)

}
