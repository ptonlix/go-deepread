package deepread

// User 微信成员信息
type User struct {
	// ExternalUserid 外部联系人的userid
	ExternalUserid string `json:"external_userid"`
	// Name 外部联系人的名称，如果外部联系人为微信用户，则返回外部联系人的名称为其微信昵称；如果外部联系人为企业微信用户，则会按照以下优先级顺序返回：此外部联系人或管理员设置的昵称、认证的实名和账号名称。
	Name string `json:"name"`
	// Avatar 外部联系人头像，第三方不可获取
	Avatar string `json:"avatar"`
	// Gender 外部联系人性别 0-未知 1-男性 2-女性
	Gender int `json:"gender"`
	// Unionid 外部联系人在微信开放平台的唯一身份标识（微信unionid），通过此字段企业可将外部联系人与公众号/小程序用户关联起来。仅当联系人类型是微信用户，且企业或第三方服务商绑定了微信开发者ID有此字段。查看绑定方法 关于返回的unionid，如果是第三方应用调用该接口，则返回的unionid是该第三方服务商所关联的微信开放者帐号下的unionid。也就是说，同一个企业客户，企业自己调用，与第三方服务商调用，所返回的unionid不同；不同的服务商调用，所返回的unionid也不同。
	Unionid string `json:"unionId"`

	MsgType string `json:"msgType"`
}

type Welcome struct {
	// Text 欢迎语介绍
	Text string `json:"text"`
	// Title 图文消息标题，最长为128字节
	Title string `json:"title"`
	// PicURL 图文消息封面的url
	PicURL string `json:"picurl"`
	// Desc 图文消息的描述，最长为512字节
	Desc string `json:"desc"`
	// URL 图文消息的链接
	URL string `json:"url"`
}

// AddUser 添加微信用户
func (d *DeepReadApp) AddUser(user *User) (*Welcome, error) {
	resp, err := d.execUserAdd(reqUserAdd{
		User: user})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// DeleteUser 删除微信用户
func (d *DeepReadApp) DeleteUser(unionid string) error {
	_, err := d.execUserDelete(reqUserDelete{
		Unionid: unionid,
		MsgType: "unsubscribe"})
	if err != nil {
		return err
	}
	return nil
}
