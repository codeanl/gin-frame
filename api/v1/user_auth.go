package v1

import (
	"github.com/gin-gonic/gin"
	"rookieCode/model/req"
	"rookieCode/utils"
	"rookieCode/utils/r"
)

type UserAuth struct{}

// 根据 Token 获取用户信息
func (*User) GetInfo(c *gin.Context) {
	r.SuccessData(c, userService.GetInfo(utils.GetFromContext[int](c, "user_info_id")))
}

// 用户注册
func (*UserAuth) Register(c *gin.Context) {
	r.SendCode(c, userService.Register(utils.BindValidJson[req.Register](c)))
}

// 登录
func (*UserAuth) Login(c *gin.Context) {
	loginReq := utils.BindValidJson[req.Login](c)
	loginVo, code := userService.Login(c, loginReq.Username, loginReq.Password)
	r.SendData(c, code, loginVo)
}

// 退出登录
func (*UserAuth) Logout(c *gin.Context) {
	userService.Logout(c)
	r.Success(c)
}

// 发送邮件验证码
func (*UserAuth) SendCode(c *gin.Context) {
	r.SendCode(c, userService.SendCode(c.Query("email")))
}
