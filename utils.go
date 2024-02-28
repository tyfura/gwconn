package gwconn

func GeneralRespOk() *GeneralResp {
	return &GeneralResp{Ok: true}
}

func GeneralRespFalse() *GeneralResp {
	return &GeneralResp{Ok: false}
}
