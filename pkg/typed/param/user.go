package param

import "github.com/wenyinh/go-wire-app/pkg/typed/entity"

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Age      int    `json:"age"`
}

type CreateUserResponse struct {
	UserId uint64 `json:"user_id"`
}

type GetUserRequest struct {
	UserId uint64 `json:"userId" uri:"userId" binding:"required"`
}
type GetUserResponse struct {
	User entity.UserEntity `json:"user"`
}
