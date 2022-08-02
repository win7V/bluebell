package controller

import (
	"bluebell/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// ---- 跟社区相关的 ----

func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区 (community_id,community_name) 以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandler 社区分类详情
func CommunityDetailHandler(c *gin.Context) {
	//获取社区ID,与冒号后面对应上
	communityID := c.Param("id") // 获取URL参数
	id, err := strconv.ParseInt(communityID, 10, 64)
	//如果不是字符类型,则报类型错误
	if err != nil {
		ResponseError(c, CodeInvalidParam)
	}
	//	根据id获取社区详情
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		// 不轻易把服务端报错报给用户
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
