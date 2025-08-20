package user

import (
	"github.com/gin-gonic/gin"
	"github.com/wenyinh/go-wire-app/pkg/typed/param"
	"github.com/wenyinh/go-wire-app/pkg/utils"
)

const CreateUserUri = ""

func (ctrl *Controller) CreateUser(c *gin.Context) {
	var req param.CreateUserRequest
	// 1. 参数绑定
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, "Invalid request")
		return
	}
	// 2. 调用 Service 层创建用户
	resp, err := ctrl.Service.CreateUser(c.Request.Context(), req)
	if err != nil {
		utils.Fail(c, "Failed to create user")
		return
	}
	// 3. 返回成功
	utils.SuccessWithData(c, resp)
}
