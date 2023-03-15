package Lphone

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PhoneNumRouter struct {
}

// InitPhoneNumRouter 初始化 PhoneNum 路由信息
func (s *PhoneNumRouter) InitPhoneNumRouter(Router *gin.RouterGroup) {
	phoneNumRouter := Router.Group("phoneNum").Use(middleware.OperationRecord())
	phoneNumRouterWithoutRecord := Router.Group("phoneNum")
	var phoneNumApi = v1.ApiGroupApp.LphoneApiGroup.PhoneNumApi
	{
		phoneNumRouter.POST("createPhoneNum", phoneNumApi.CreatePhoneNum)             // 新建PhoneNum
		phoneNumRouter.DELETE("deletePhoneNum", phoneNumApi.DeletePhoneNum)           // 删除PhoneNum
		phoneNumRouter.DELETE("deletePhoneNumByIds", phoneNumApi.DeletePhoneNumByIds) // 批量删除PhoneNum
		phoneNumRouter.PUT("updatePhoneNum", phoneNumApi.UpdatePhoneNum)              // 更新PhoneNum
	}
	{
		phoneNumRouterWithoutRecord.GET("findPhoneNum", phoneNumApi.FindPhoneNum)       // 根据ID获取PhoneNum
		phoneNumRouterWithoutRecord.GET("getPhoneNumList", phoneNumApi.GetPhoneNumList) // 获取PhoneNum列表
	}
}
