package deepread

// execUserAdd 添加企业微信用户
func (d *DeepReadApp) execUserAdd(req reqUserAdd) (respUserAdd, error) {
	resp := respUserAdd{}
	err := d.executeDrapiJSONPost("/read/wx/cp/receive/message", req, &resp, true)
	if err != nil {
		return respUserAdd{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respUserAdd{}, bizErr
	}

	return resp, nil
}

// execUserDelete 删除企业微信用户
func (d *DeepReadApp) execUserDelete(req reqUserDelete) (respUserDelete, error) {
	resp := respUserDelete{}
	err := d.executeDrapiJSONPost("/read/wx/cp/receive/message", req, &resp, true)
	if err != nil {
		return respUserDelete{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respUserDelete{}, bizErr
	}

	return resp, nil
}
