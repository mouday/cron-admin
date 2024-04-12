package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mouday/cron-admin/src/config"
	"github.com/mouday/cron-admin/src/model"
	"github.com/mouday/cron-admin/src/utils"
	"github.com/mouday/cron-admin/src/vo"
)

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

/*
 * 登录
 */
func Login(ctx *gin.Context) {
	loginForm := LoginForm{}
	ctx.BindJSON(&loginForm)

	db := config.GetDB()
	userRow := &model.UserModel{}
	db.Model(&model.UserModel{}).Where("username = ? and status = 1", loginForm.Username).First(&userRow)

	if userRow.UserId != "" && utils.VerifyPassword(loginForm.Password, userRow.Password) {
		vo.Success(ctx, gin.H{
			"token": utils.CreateToken(userRow.UserId, config.SCERET),
		})
	} else {
		vo.Error(ctx, -1, "用户名或密码错误")
	}

}
