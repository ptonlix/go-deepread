package services

import (
	"go-deepread/init/server"
	"go-deepread/services/deepread"
)

type DeepReadService struct {
	DrClient *deepread.DeepReadApp
	WwClient *WeWorkService
}

func NewDeepReadService() *DeepReadService {
	drclient := deepread.New(server.ServerConf.AccessToken)
	wwclient := NewWeWorkService()
	return &DeepReadService{DrClient: drclient, WwClient: wwclient}
}

var DrServer *DeepReadService

func init() {
	DrServer = NewDeepReadService()
}
