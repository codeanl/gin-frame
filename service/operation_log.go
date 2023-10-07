package service

import (
	"rookieCode/dao"
	"rookieCode/model"
	"rookieCode/model/req"
	"rookieCode/model/resp"
	"rookieCode/utils/r"
)

type OperationLog struct{}

func (*OperationLog) GetList(req req.PageQuery) resp.PageResult[[]model.OperationLog] {
	list, total := operationLogDao.GetList(req)
	return resp.PageResult[[]model.OperationLog]{
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		List:     list,
		Total:    total,
	}
}

func (*OperationLog) Delete(ids []int) (code int) {
	dao.Delete(model.OperationLog{}, "id in ?", ids)
	return r.OK
}
