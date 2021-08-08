package tbkscpublisherinfoget

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.tbk.sc.publisher.info.get( 淘宝客信息查询 - 社交 )
type Response struct {
	response.TopResponse
	TbkScPublisherInfoGetResult ResponseItem `json:"tbk_sc_publisher_info_get_response"`
}

//解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

type ResponseItem struct {
	Data      ResponseItemData `json:"data"`
	RequestID string           `json:"request_id"`
}

type ResponseItemData struct {
	InviterList        ResponseItemInviterList        `json:"inviter_list"`
	RootPIDChannelList ResponseItemRootPIDChannelList `json:"root_pid_channel_list"`
	TotalCount         int64                          `json:"total_count"`
}

type ResponseItemInviterList struct {
	MapData []ResponseItemMapDatum `json:"map_data"`
}

type ResponseItemMapDatum struct {
	AccountName  string `json:"account_name"`
	CreateDate   string `json:"create_date"`
	Note         string `json:"note"`
	OfflineScene string `json:"offline_scene"`
	OnlineScene  string `json:"online_scene"`
	RealName     string `json:"real_name"`
	RelationApp  string `json:"relation_app"`
	RelationID   int64  `json:"relation_id"`
	RootPID      string `json:"root_pid"`
	Rtag         string `json:"rtag"`
}

type ResponseItemRootPIDChannelList struct {
	String []string `json:"string"`
}
