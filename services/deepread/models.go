package deepread

import (
	"encoding/json"
)

func marshalIntoJSONBody(x interface{}) ([]byte, error) {
	y, err := json.Marshal(x)
	if err != nil {
		// should never happen unless OOM or similar bad things
		return nil, makeReqMarshalErr(err)
	}

	return y, nil
}

type respCommon struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

// IsOK 响应体是否为一次成功请求的响应
func (x *respCommon) IsOK() bool {
	return x.Code == 0
}

func (x *respCommon) TryIntoErr() error {
	if x.IsOK() {
		return nil
	}

	return &DeepReadClientError{
		Code: x.Code,
		Msg:  x.Msg,
	}
}

// reqUserUpdate 更新成员请求
type reqUserAdd struct {
	User *User
}

var _ bodyer = reqUserAdd{}

func (x reqUserAdd) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x.User)
}

// respUserUpdate 更新成员响应
type respUserAdd struct {
	respCommon
	Data Welcome `json:"data"`
}

// reqUserUpdate 更新成员请求
type reqUserDelete struct {
	Unionid string `json:"unionid"`
}

var _ bodyer = reqUserAdd{}

func (x reqUserDelete) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x.Unionid)
}

// respUserUpdate 更新成员响应
type respUserDelete struct {
	respCommon
	Data string `json:"data"`
}
