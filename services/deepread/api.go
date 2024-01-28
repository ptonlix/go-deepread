package deepread

// execUserAdd 添加企业微信用户
func (d *DeepReadApp) execUserAdd(req reqUserAdd) (respUserAdd, error) {
	var resp respUserAdd
	err := d.executeDrapiJSONPost("/wework/callback/user/add", req, &resp, true)
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
	var resp respUserDelete
	err := d.executeDrapiJSONPost("/wework/callback/user/delete", req, &resp, true)
	if err != nil {
		return respUserDelete{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respUserDelete{}, bizErr
	}

	return resp, nil
}
