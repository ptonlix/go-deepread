package deepread

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/beego/beego/v2/adapter/logs"
)

type DeepReadApp struct {
	// CorpSecret 应用的凭证密钥，必填
	opts        options
	accessToken string
}

// New 构造一个 Workwx 客户端对象，需要提供企业 ID
func New(accessToken string, opts ...CtorOption) *DeepReadApp {
	optionsObj := defaultOptions()

	for _, o := range opts {
		o.applyTo(&optionsObj)
	}

	return &DeepReadApp{
		opts:        optionsObj,
		accessToken: accessToken,
	}
}

func (d *DeepReadApp) composeDrapiURL(path string, req interface{}) (*url.URL, error) {
	values := url.Values{}
	if valuer, ok := req.(urlValuer); ok {
		values = valuer.intoURLValues()
	}

	// TODO: refactor
	base, err := url.Parse(d.opts.DRAPIHost)
	if err != nil {
		return nil, fmt.Errorf("drapiHost invalid: host=%s err=%w", d.opts.DRAPIHost, err)
	}

	base.Path = path
	base.RawQuery = values.Encode()

	return base, nil
}

func (d *DeepReadApp) composeDrapiURLWithToken(path string, req interface{}, withAccessToken bool) (*url.URL, error) {
	url, err := d.composeDrapiURL(path, req)
	if err != nil {
		return nil, err
	}

	if !withAccessToken {
		return url, nil
	}

	tok := d.accessToken

	q := url.Query()
	q.Set("access_token", tok)
	url.RawQuery = q.Encode()

	return url, nil
}

func (d *DeepReadApp) executeDrapiGet(path string, req urlValuer, respObj interface{}, withAccessToken bool) error {
	url, err := d.composeDrapiURL(path, req)
	if err != nil {
		return err
	}
	urlStr := url.String()

	request, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return err
	}
	if withAccessToken {
		request.Header.Add("wx-cp-auth", d.accessToken)
	}
	resp, err := d.opts.HTTP.Do(request)

	if err != nil {
		return makeRequestErr(err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(respObj)
	if err != nil {
		return makeRespUnmarshalErr(err)
	}

	return nil
}

func (d *DeepReadApp) executeDrapiJSONPost(path string, req bodyer, respObj interface{}, withAccessToken bool) error {
	url, err := d.composeDrapiURL(path, req)
	if err != nil {
		return err
	}
	urlStr := url.String()

	body, err := req.intoBody()
	if err != nil {
		return makeReqMarshalErr(err)
	}

	logs.Debug("req: ", string(body))

	request, err := http.NewRequest("POST", urlStr, bytes.NewReader(body))
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json")
	if withAccessToken {
		request.Header.Add("wx-cp-auth", d.accessToken)
	}
	resp, err := d.opts.HTTP.Do(request)
	if err != nil {
		return makeRequestErr(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body) // 读取响应体数据
	responseString := string(bodyBytes)

	logs.Debug("resp: ", responseString)
	readio := strings.NewReader(responseString)
	//decoder := json.NewDecoder(resp.Body)
	decoder := json.NewDecoder(readio)
	err = decoder.Decode(respObj)
	if err != nil {
		return makeRespUnmarshalErr(err)
	}

	return nil
}
