package alibabaalscunionkbbbtitemstoredetailget

import (
	"encoding/json"
	"fmt"
	"github.com/mimicode/tksdk/alscopensdk/response"
)

// Response alibaba.alsc.union.kb.bbt.item.store.detail.get( 本地联盟爆爆团门店详情 )
type Response struct {
	response.TopResponse
	ResultResponse ResultResponse `json:"alibaba_alsc_union_kb_bbt_item_store_detail_get_response"`
}

// WrapResult 解析输出结果
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

// ResultResponse 响应结果
type ResultResponse struct {
	BizErrorDesc string `json:"biz_error_desc"`
	BizErrorCode int    `json:"biz_error_code"`
	RequestID    string `json:"request_id"`
	Data         Data   `json:"data"`
	ResultCode   int64  `json:"result_code"`
	Message      string `json:"message"`
}

func (rr ResultResponse) IsError() bool {
	return rr.BizErrorCode != 0
}

func (rr ResultResponse) Error() string {
	return fmt.Sprintf("biz_error_code:%d biz_error_desc:%s", rr.BizErrorCode, rr.BizErrorDesc)
}

type Data struct {
	BasicInfo  BasicInfo  `json:"basic_info"`
	Statistics Statistics `json:"statistics"`
}

type BasicInfo struct {
	StoreID        string         `json:"store_id"`
	Name           string         `json:"name"`
	Cover          string         `json:"cover"`
	Mobiles        Mobiles        `json:"mobiles"`
	Business       Business       `json:"business"`
	Location       Location       `json:"location"`
	Brand          Brand          `json:"brand"`
	Categories     Categories     `json:"categories"`
	Qualifications Qualifications `json:"qualifications"`
}

type Brand struct {
	BrandID   string `json:"brand_id"`
	BrandName string `json:"brand_name"`
}

type Business struct {
	BusinessStatus int64        `json:"business_status"`
	BusinessDesc   string       `json:"business_desc"`
	BusinessTime   BusinessTime `json:"business_time"`
	Promotion      string       `json:"promotion"`
}

type BusinessTime struct {
	TimeTexts Mobiles `json:"time_texts"`
}

type Mobiles struct {
	String []string `json:"string"`
}

type Categories struct {
	Category []Category `json:"category"`
}

type Category struct {
	CategoryID       string `json:"category_id"`
	Name             string `json:"name"`
	ParentCategoryID string `json:"parent_category_id"`
}

type Location struct {
	Address     string `json:"address"`
	AddressMemo string `json:"address_memo"`
	Region      Region `json:"region"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
}

type Region struct {
	ProvinceID   string `json:"province_id"`
	ProvinceName string `json:"province_name"`
	CityID       string `json:"city_id"`
	CityName     string `json:"city_name"`
	DistrictID   string `json:"district_id"`
	DistrictName string `json:"district_name"`
}

type Qualifications struct {
	ImageContent []ImageContent `json:"image_content"`
}

type ImageContent struct {
	Title string  `json:"title"`
	Desc  string  `json:"desc"`
	Urls  Mobiles `json:"urls"`
}

type Statistics struct {
	ServiceRating string `json:"service_rating"`
	AvgPrice      string `json:"avg_price"`
}
