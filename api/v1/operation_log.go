package v1

import (
	"github.com/gin-gonic/gin"
	"rookieCode/utils"
	"rookieCode/utils/r"
)

type OperationLog struct{}

func (*OperationLog) GetList(c *gin.Context) {
	r.SuccessData(c, operationLogService.GetList(utils.BindPageQuery(c)))
}

func (*OperationLog) Delete(c *gin.Context) {
	r.SendCode(c, operationLogService.Delete(utils.BindJson[[]int](c)))
}
