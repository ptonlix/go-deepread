package services

import (
	"context"
	"go-deepread/init/wework"

	"github.com/beego/beego/v2/core/logs"
	"github.com/xen0n/go-workwx"
)

type WeWorkService struct {
	app *workwx.WorkwxApp
}

func NewWeWorkService() *WeWorkService {

	client := workwx.New(wework.WeworkConf.CorpId)

	// work with individual apps
	app := client.WithApp(wework.WeworkConf.CorpSecret, int64(wework.WeworkConf.AgentId))
	app.SpawnAccessTokenRefresher()
	return &WeWorkService{app: app}
}

func NewWithCtx(ctx context.Context) *WeWorkService {

	client := workwx.New(wework.WeworkConf.CorpId)

	// work with individual apps
	app := client.WithApp(wework.WeworkConf.CorpSecret, int64(wework.WeworkConf.AgentId))
	app.SpawnAccessTokenRefresherWithContext(ctx)
	return &WeWorkService{app: app}
}

func (w *WeWorkService) UpdateUser(userId, externalUserId, unionId string) error {
	if err := w.app.RemarkExternalContact(&workwx.ExternalContactRemark{Userid: userId, ExternalUserid: externalUserId, Description: unionId}); err != nil {
		logs.Error("UpdateUser is failed!:", err)
		return err
	}

	return nil
}

func (w *WeWorkService) SendWelcome(welcomeCode, welcomeText, title, picurl, desc, url string) error {

	att := workwx.Attachments{Link: workwx.Link{Title: title, PicURL: picurl, Desc: desc, URL: url}}

	if err := w.app.SendWelcomeMsg(welcomeCode, workwx.Text{Content: welcomeText}, []workwx.Attachments{att}); err != nil {
		logs.Error("SendWelcomeMsg is failed!:", err)
		return err
	}

	return nil
}

func (w *WeWorkService) GetUserInfo(externalUserId string) (*workwx.ExternalContactInfo, error) {
	resp, err := w.app.GetExternalContact(externalUserId)
	if err != nil {
		logs.Error("GetUser is failed!:", err)
		return nil, err
	}

	return resp, nil
}
