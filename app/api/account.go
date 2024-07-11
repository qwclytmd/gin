package api

import (
	"bcw/app/common/message"
	"bcw/app/services/account_service"
	"github.com/gin-gonic/gin"
)

// 列表
func AccountList(ctx *gin.Context) {

	input := account_service.AccountReq{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		message.ResponseAPIResult(ctx, &message.StatusParameterError, nil)
		return
	}

	s := account_service.AccountService{}
	result, err := s.GetAccountList(input)
	if err != nil {
		message.ResponseAPIResult(ctx, err, nil)
		return
	}

	message.ResponseAPIResult(ctx, &message.StatusOK, result)
}

// 添加管理员
//func (m ManagersAccount) AddManager(c *gin.Context) {
//
//	if err := c.ShouldBind(&m.RequestAdd); err != nil {
//		message.ResponseResult(c, 0, "参数错误", nil)
//		return
//	}
//
//	err := managersrv.InsertManager(c, m.RequestAdd)
//	if err != nil {
//		message.ResponseResult(c, 0, err.Error(), nil)
//		return
//	}
//
//	message.ResponseResult(c, 1, "success", nil)
//}
//
//// 编辑管理员
//func (m ManagersAccount) ChangeManager(c *gin.Context) {
//	if err := c.ShouldBind(&m.RequestUpdate); err != nil {
//		message.ResponseResult(c, 0, "参数错误", nil)
//		return
//	}
//
//	err := managersrv.UpdateManager(m.RequestUpdate)
//	if err != nil {
//		message.ResponseResult(c, 0, "系统错误", nil)
//		return
//	}
//
//	message.ResponseResult(c, 1, "success", nil)
//}
//
//// 删除管理员
//func (m ManagersAccount) DestroyManager(c *gin.Context) {
//
//	if err := c.ShouldBind(&m.RequestDelId); err != nil {
//
//		message.ResponseResult(c, 0, err.Error(), nil)
//		return
//	}
//
//	err := managersrv.DeleteManager(m.RequestDelId.Id)
//	if err != nil {
//		message.ResponseResult(c, 0, err.Error(), nil)
//		return
//	}
//
//	message.ResponseResult(c, 1, "success", nil)
//}
//
//// 密码修改
//func (m ManagersAccount) ChangePassword(c *gin.Context) {
//	input := dto.UpdatePasswordRequest{}
//	if err := c.ShouldBind(&input); err != nil {
//		message.ResponseResult(c, 0, err.Error(), nil)
//		return
//	}
//
//	currentId := c.GetInt("manager_id")
//
//	if currentId != input.Id {
//		message.ResponseResult(c, 0, "参数错误", nil)
//		return
//	}
//
//	err := managersrv.UpdatePassword(input)
//	if err != nil {
//		message.ResponseResult(c, 0, err.Error(), nil)
//		return
//	}
//
//	message.ResponseResult(c, 1, "success", nil)
//}
