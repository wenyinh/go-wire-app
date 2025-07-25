package user

import (
	"github.com/gin-gonic/gin"
	"github.com/wenyinh/go-wire-app/pkg/typed/model"
	"github.com/wenyinh/go-wire-app/pkg/utils"
)

const GetUserUri = "/:userId"

func (ctrl *Controller) GetUser(c *gin.Context) {
	var req model.GetUserRequest
	// 1. 参数绑定
	if err := c.ShouldBindUri(&req); err != nil {
		utils.Fail(c, "Invalid request")
		return
	}
	// 2. 调用 Service 层查询用户
	resp, err := ctrl.Service.GetUser(c.Request.Context(), req)
	if err != nil {
		utils.Fail(c, "Failed to Get user")
		return
	}
	// 3. 返回成功
	utils.SuccessWithData(c, resp)
}
