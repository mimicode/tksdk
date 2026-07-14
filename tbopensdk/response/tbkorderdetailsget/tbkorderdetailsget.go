package tbkorderdetailsget

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

// taobao.tbk.order.details.get( 淘宝客-推广者-所有订单查询 )
type Response struct {
	response.TopResponse
	TbkOrderDetailsGetResponse Result `json:"tbk_order_details_get_response"`
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

type Result struct {
	Data      Data   `json:"data"`       //订单分页数据
	RequestID string `json:"request_id"` //平台颁发的每次请求访问的唯一标识
}

type Data struct {
	HasNext       bool    `json:"has_next"`       //是否还有下一页
	HasPre        bool    `json:"has_pre"`        //是否还有上一页
	PageNo        int64   `json:"page_no"`        //页码
	PageSize      int64   `json:"page_size"`      //页大小
	PositionIndex string  `json:"position_index"` //位点字段，由调用方原样传递
	Results       Results `json:"results"`        //订单明细列表
}

type Results struct {
	PublisherOrderDto []PublisherOrderDto `json:"publisher_order_dto"` //订单明细集合
}

type PublisherOrderDto struct {
	AdzoneID                    int64                        `json:"adzone_id"`                                 //推广位ID，pid中的最后一段数字，如pid=mm_1_2_3中的"3"
	AdzoneName                  string                       `json:"adzone_name"`                               //推广位名称
	AlimamaRate                 string                       `json:"alimama_rate"`                              //平台技术服务费比率
	AlimamaShareFee             string                       `json:"alimama_share_fee"`                         //平台技术服务费
	AlipayTotalPrice            string                       `json:"alipay_total_price"`                        //付款金额，买家拍下并付款金额（不含运费）
	AlimmShareInfoDto           *AlimmShareInfoDTO           `json:"alimm_share_info_dto,omitempty"`            //平台技术服务费明细节点
	ClickTime                   string                       `json:"click_time"`                                //点击时间，通过推广链接达到商品、店铺详情页的点击时间
	DepositPrice                string                       `json:"deposit_price"`                             //定金付款金额，预售订单支付的定金金额
	ExtraMktID                  string                       `json:"extra_mkt_id"`                              //管理member新商品ID后段
	FlowSource                  string                       `json:"flow_source"`                               //产品类型，如"联盟超级活动"、"品牌精选推广"等
	IncomeRate                  string                       `json:"income_rate"`                               //收入比率，收入比率=佣金比率+补贴比率
	ItemCategoryName            string                       `json:"item_category_name"`                        //类目名称，商品所属的一级类目名称
	ItemID                      interface{}                  `json:"item_id,omitempty"`                         //商品id
	ItemImg                     string                       `json:"item_img"`                                  //商品图片
	ItemLink                    *string                      `json:"item_link,omitempty"`                       //商品链接
	ItemNum                     int64                        `json:"item_num"`                                  //商品数量
	ItemPrice                   *string                      `json:"item_price,omitempty"`                      //商品单价
	ItemTitle                   string                       `json:"item_title"`                                //商品标题
	MarketingType               string                       `json:"marketing_type"`                            //营销类型，如联盟超级活动-优品、特价版客户端锁粉等
	ModifiedTime                string                       `json:"modified_time"`                             //订单更新时间
	OrderType                   string                       `json:"order_type"`                                //平台类型，订单所属平台类型，包括天猫、淘宝、聚划算、口碑等
	PayPrice                    *string                      `json:"pay_price,omitempty"`                       //结算金额，买家确认收货的付款金额（不含运费）
	PlatformSpecialShareInfoDto *PlatformSpecialShareInfoDTO `json:"platform_special_share_info_dto,omitempty"` //平台专项服务费明细节点
	PlatformSpecialServiceFee   string                       `json:"platform_special_service_fee"`              //平台专项服务费
	PlatformSpecialServiceRate  string                       `json:"platform_special_service_rate"`             //平台专项服务费比率
	PubID                       int64                        `json:"pub_id"`                                    //推广者的账号id
	PubShareFee                 string                       `json:"pub_share_fee"`                             //结算预估收入，结算预估佣金收入+结算预估补贴收入
	PubShareFeeForCommission    string                       `json:"pub_share_fee_for_commission"`              //结算预估佣金收入
	PubShareFeeForSdy           string                       `json:"pub_share_fee_for_sdy"`                     //结算预估补贴收入
	PubSharePreFee              string                       `json:"pub_share_pre_fee"`                         //付款预估收入，付款预估佣金收入+付款预估补贴收入
	PubSharePreFeeForCommission string                       `json:"pub_share_pre_fee_for_commission"`          //付款预估佣金收入
	PubSharePreFeeForSdy        string                       `json:"pub_share_pre_fee_for_sdy"`                 //付款预估补贴收入
	PubShareRate                string                       `json:"pub_share_rate"`                            //佣金分成比率，从佣金中分得的收益比率（含平台技术服务费比率）
	PubShareRateForSdy          string                       `json:"pub_share_rate_for_sdy"`                    //补贴分成比率，从补贴中分得的收益比率
	RefundTag                   int64                        `json:"refund_tag"`                                //维权标签，0-非维权，1-维权订单
	RelationID                  int64                        `json:"relation_id"`                               //渠道关系id，渠道管理功能中的渠道关系ID
	RightsID                    string                       `json:"rights_id"`                                 //淘礼金ID，每次创建的淘礼金有唯一识别ID
	SellerNick                  string                       `json:"seller_nick"`                               //掌柜旺旺
	SellerShopTitle             string                       `json:"seller_shop_title"`                         //店铺名称
	ServiceFeeDtoList           *ServiceFeeDtoList           `json:"service_fee_dto_list,omitempty"`            //服务费信息（字段已废弃）
	SiteID                      int64                        `json:"site_id"`                                   //媒体ID，pid中的第二段数字，如pid=mm_1_2_3中的"2"
	SiteName                    string                       `json:"site_name"`                                 //媒体名称
	SpecialID                   int64                        `json:"special_id"`                                //会员运营ID
	SubsidyFee                  string                       `json:"subsidy_fee"`                               //补贴金额，各类补贴的补贴金额总和
	SubsidyInfoDtoList          *SubsidyInfoDtoList          `json:"subsidy_info_dto_list,omitempty"`           //补贴金额明细节点
	SubsidyRate                 string                       `json:"subsidy_rate"`                              //补贴比率，各类型补贴的补贴比率总和
	SubsidyType                 string                       `json:"subsidy_type"`                              //补贴类型
	TBDepositTime               string                       `json:"tb_deposit_time"`                           //定金淘宝付款时间，预售订单在淘宝支付定金的付款时间
	TBPaidTime                  string                       `json:"tb_paid_time"`                              //淘宝付款时间，订单在淘宝付款的时间
	TBGmvTotalPrice             string                       `json:"tb_gmv_total_price"`                        //买家拍下金额（不含运费金额）
	TalentPid                   string                       `json:"talent_pid"`                                //专用（不对外开放）
	TerminalType                string                       `json:"terminal_type"`                             //成交平台，成交来自于PC或无线
	TkCreateTime                string                       `json:"tk_create_time"`                            //创建时间，订单创建的时间
	TkDepositTime               string                       `json:"tk_deposit_time"`                           //定金付款时间，预售订单支付定金的付款时间
	TkEarningTime               *string                      `json:"tk_earning_time,omitempty"`                 //结算时间，订单确认收货后且商家完成佣金支付的时间
	TkOrderRole                 int64                        `json:"tk_order_role"`                             //推广者身份，2-二方，3-三方
	TkPaidTime                  string                       `json:"tk_paid_time"`                              //付款时间，订单付款的时间，同步淘宝
	TkStatus                    int64                        `json:"tk_status"`                                 //订单状态，3-订单结算，12-订单付款，13-订单失效，14-订单成功
	TkTotalRate                 string                       `json:"tk_total_rate"`                             //佣金提成，佣金比率*佣金分成比率
	TkTotalRateForSdy           string                       `json:"tk_total_rate_for_sdy"`                     //补贴提成，补贴比率*补贴分成比率
	TotalCommissionFee          string                       `json:"total_commission_fee"`                      //佣金金额
	TotalCommissionRate         string                       `json:"total_commission_rate"`                     //佣金比率
	TradeID                     string                       `json:"trade_id"`                                  //子订单号
	TradeParentID               string                       `json:"trade_parent_id"`                           //订单编号
	TpOrderID                   string                       `json:"tp_order_id"`                               //非电商淘系子订单号
	Unid                        string                       `json:"unid"`                                      //unid（不对外开放）
	Untts                       string                       `json:"untts"`                                     //流量通untts（默认无，限定开放）
	Uvid                        string                       `json:"uvid"`                                      //加密标识（默认不开放）
}

type ServiceFeeDtoList struct {
	ServiceFeeDto []ServiceFeeDto `json:"service_fee_dto"` //服务费明细集合（字段已废弃）
}

type ServiceFeeDto struct {
	SharePreFee       string `json:"share_pre_fee"`       //预估专项服务费（字段已废弃）
	ShareFee          string `json:"share_fee"`           //结算专项服务费（字段已废弃）
	ShareRelativeRate string `json:"share_relative_rate"` //专项服务费率（字段已废弃）
	TkShareRoleType   int64  `json:"tk_share_role_type"`  //专项服务费来源，122-渠道（字段已废弃）
}

type SubsidyInfoDtoList struct {
	SubsidyDetailDTO []SubsidyDetailDTO `json:"subsidy_detail_d_t_o"` //补贴明细集合
}

type SubsidyDetailDTO struct {
	SubsidyFee        string `json:"subsidy_fee"`         //对应补贴类型的补贴金额
	SubsidyRate       string `json:"subsidy_rate"`        //补贴比率
	SubsidyShareRate  string `json:"subsidy_share_rate"`  //补贴分成比率
	SubsidyType       string `json:"subsidy_type"`        //该笔订单包含的补贴类型
	SubsidyUpperLimit string `json:"subsidy_upper_limit"` //单笔订单补贴上限
}

type AlimmShareInfoDTO struct {
	AlimmAgentServiceFee    string `json:"alimm_agent_service_fee"`     //结算渠道专项服务费
	AlimmAgentServicePreFee string `json:"alimm_agent_service_pre_fee"` //预估渠道专项服务费
	AlimmAgentServiceRate   string `json:"alimm_agent_service_rate"`    //渠道专项服务费比率
	AlimmTechServiceFee     string `json:"alimm_tech_service_fee"`      //结算技术服务费
	AlimmTechServicePreFee  string `json:"alimm_tech_service_pre_fee"`  //预估技术服务费
	AlimmTechServiceRate    string `json:"alimm_tech_service_rate"`     //技术服务费比率
}

type PlatformSpecialShareInfoDTO struct {
	ContentTechServiceFee    string `json:"content_tech_service_fee"`     //结算内容专项服务费
	ContentTechServicePreFee string `json:"content_tech_service_pre_fee"` //预估内容专项服务费
	ContentTechServiceRate   string `json:"content_tech_service_rate"`    //内容专项服务费比率
	TrafficTechServiceFee    string `json:"traffic_tech_service_fee"`     //结算流量专项服务费（默认无，限定开放）
	TrafficTechServicePreFee string `json:"traffic_tech_service_pre_fee"` //预估流量专项服务费（默认无，限定开放）
	TrafficTechServiceRate   string `json:"traffic_tech_service_rate"`    //流量专项服务费比率（默认无，限定开放）
}
