package taobaotbkscgenerallinkparse

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

// taobao.tbk.sc.general.link.parse( 淘宝客-服务商-万能解析 )
type Response struct {
	response.TopResponse
	TbkScGeneralLinkParseResponse TbkScGeneralLinkParseResponse `json:"tbk_sc_general_link_parse_response"`
}

// 解析输出结果
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

type TbkScGeneralLinkParseResponse struct {
	Data      Data   `json:"data"`
	RequestID string `json:"request_id"`
}

type Data struct {
	MaterialURLList DataMaterialURLList `json:"material_url_list"`
}

type DataMaterialURLList struct {
	MaterialURLList []MaterialURLListElement `json:"material_url_list"`
}

type MaterialURLListElement struct {
	InputMaterialURL string      `json:"input_material_url"` //入参物料链接
	ExtraInfo        string      `json:"extra_info"`         //解析成功的场景下，需要补充说明的信息
	LinkInfoDto      LinkInfoDto `json:"link_info_dto"`
}

type LinkInfoDto struct {
	MaterialID    string `json:"material_id"`     //物料ID，需根据material_type判断物料类型， 可能为商品ID、卖家ID、会场ID
	MaterialType  int64  `json:"material_type"`   //1—商品 2—店铺 3—会场 4-承接开放 5-优惠券 6-直播间
	TkBizType     int64  `json:"tk_biz_type"`     //1-联盟口令，2-主站口令。入参物料不为淘口令时，此字段返回null
	TpwdOriginURL string `json:"tpwd_origin_url"` //入参口令/链接对应的原链接
	IsvMktid      string `json:"isv_mktid"`       //根据入参工具服务商账号信息生成的下一跳新商品ID
	OriginPid     string `json:"origin_pid"`      //入参推广链接中的pid，需配合入参fields使用，如果不属于当前调用的推广者，则不返回
}
