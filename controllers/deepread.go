package controllers

import (
	"go-deepread/init/wework"
	"go-deepread/services"
	"go-deepread/services/deepread"
	"net/http"

	"github.com/beego/beego/v2/adapter/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/xen0n/go-workwx"
)

type DeepReadController struct {
	DrServer *services.DeepReadService
	beego.Controller
}

// type dummyRxMessageHandler struct{}

// var _ workwx.RxMessageHandler = dummyRxMessageHandler{}

// OnIncomingMessage 一条消息到来时的回调。
func (dr *DeepReadController) OnIncomingMessage(msg *workwx.RxMessage) error {
	// You can do much more!
	logs.Debug("incoming message: %s\n", msg)
	switch msg.ChangeType {
	case workwx.ChangeTypeAddExternalContact:
		// 进入发送欢迎语流程
		if data, flag := msg.EventAddExternalContact(); flag {
			go dr.WelcomeFlow(data.GetUserID(), data.GetExternalUserID(), data.GetWelcomeCode())
		}
	case workwx.ChangeTypeDelExternalContact, workwx.ChangeTypeDelFollowUser:
		// 进入删除用户流程
		if data, flag := msg.EventDelExternalContact(); flag {
			go dr.DeleteUserFlow(data.GetExternalUserID())
		}
		if data, flag := msg.EventDelFollowUser(); flag {
			go dr.DeleteUserFlow(data.GetExternalUserID())
		}
	}

	return nil
}

func (dr *DeepReadController) NewHttpHandler() (http.Handler, error) {
	hh, err := workwx.NewHTTPHandler(wework.WeworkConf.PToken, wework.WeworkConf.PKey, dr)
	if err != nil {
		panic(err)
	}
	return hh, err
}

func (dr *DeepReadController) WelcomeFlow(userId, externalUserID, welcomeCode string) error {
	// 查找用户详情
	userInfo, err := dr.DrServer.WwClient.GetUserInfo(externalUserID)
	if err != nil {
		logs.Error("WelcomeFlow GetUserInfo failed: ", err)
		return err
	}
	// 添加用户
	welcomeData, err := dr.DrServer.DrClient.AddUser(&deepread.User{
		ExternalUserid: userInfo.ExternalContact.ExternalUserid,
		Name:           userInfo.ExternalContact.Name,
		Avatar:         userInfo.ExternalContact.Avatar,
		Gender:         int(userInfo.ExternalContact.Gender),
		Unionid:        userInfo.ExternalContact.Unionid,
		MsgType:        "subscribe",
	})
	if err != nil {
		logs.Error("WelcomeFlow AddUser failed: ", err)
		return err
	}
	// 更新用户备注
	go func() {
		if err := dr.DrServer.WwClient.UpdateUser(userId, externalUserID, userInfo.ExternalContact.Unionid); err != nil {
			logs.Error("WelcomeFlow UpdateUser failed: ", err)
		}
	}()

	// 发送欢迎语
	err = dr.DrServer.WwClient.SendWelcome(welcomeCode, welcomeData.Text, welcomeData.Title, welcomeData.PicURL, welcomeData.Desc, welcomeData.URL)
	if err != nil {
		logs.Error("WelcomeFlow SendWelcome failed: ", err)
		return err
	}
	logs.Info("WelcomeFlow Successfully : ", externalUserID)
	return nil
}

func (dr *DeepReadController) DeleteUserFlow(externalUserid string) error {
	// 查找用户详情
	userInfo, err := dr.DrServer.WwClient.GetUserInfo(externalUserid)
	if err != nil {
		logs.Error("WelcomeFlow GetUserInfo failed: ", err)
		return err
	}

	if err := dr.DrServer.DrClient.DeleteUser(userInfo.ExternalContact.Unionid); err != nil {
		logs.Error("DeleteUserFlow DeleteUser failed: ", err)
		return err
	}
	logs.Info("DeleteUserFlow Successfully failed: ", userInfo.ExternalContact.Unionid)
	return nil
}
