package tbkdgvegastljcreateresult

type TbkDgVegasTljCreateResponse struct {
	Result    Result `json:"result"`
	RequestID string `json:"request_id"`
}

type Result struct {
	MsgCode string `json:"msg_code"`
	MsgInfo string `json:"msg_info"`
	Success bool   `json:"success"`
	Model   *Model `json:"model"`
}

type Model struct {
	RightsID  string `json:"rights_id"`
	SendURL   string `json:"send_url"`
	VegasCode string `json:"vegas_code"`
}
