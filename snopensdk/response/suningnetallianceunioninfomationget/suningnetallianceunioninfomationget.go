package suningnetallianceunioninfomationget

import (
	"encoding/json"
	response2 "github.com/mimicode/tksdk/snopensdk/response"
)

//suning.netalliance.unioninfomation.get 单个查询联盟商品信息
type Response struct {
	response2.TopResponse
	SnResponseContent SnResponseContent `json:"sn_responseContent"`
}

//解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.ErrorCode = "-1"
		t.ErrorResponse.ErrorMsg = unmarshal.Error()
	} else {
		if t.SnResponseContent.SnError != nil {
			t.ErrorResponse.ErrorCode = t.SnResponseContent.SnError.ErrorCode
			t.ErrorResponse.ErrorMsg = t.SnResponseContent.SnError.ErrorMsg
		}
	}
}

type SnResponseContent struct {
	SnBody  SnBody   `json:"sn_body"`
	SnError *SnError `json:"sn_error"`
}
type SnError struct {
	ErrorMsg  string `json:"error_msg"`
	ErrorCode string `json:"error_code"`
}

type SnBody struct {
	GetUnionInfomation []GetUnionInfomation `json:"getUnionInfomation"`
}

type GetUnionInfomation struct {
	UpdateDate          string  `json:"updateDate"`
	ThirdCategoryCode   string  `json:"thirdCategoryCode"`
	PromoDesc           string  `json:"promoDesc"`
	FourthCategoryCode  string  `json:"fourthCategoryCode"`
	CategoryName        string  `json:"categoryName"`
	SuningPrice         int64   `json:"suningPrice"`
	FirstCategoryName   string  `json:"firstCategoryName"`
	Rate                float64 `json:"rate"`
	OperatingModel      int64   `json:"operatingModel"`
	Price               int64   `json:"price"`
	GoodsName           string  `json:"goodsName"`
	BrandCode           string  `json:"brandCode"`
	SecondCategoryCode  string  `json:"secondCategoryCode"`
	WAPPrePayCommission float64 `json:"wapPrePayCommission"`
	MertCode            string  `json:"mertCode"`
	PictureURL          string  `json:"pictureUrl"`
	CmmdtyEanCode       string  `json:"cmmdtyEanCode"`
	ThirdCategoryName   string  `json:"thirdCategoryName"`
	VendorName          string  `json:"vendorName"`
	FourthCategoryName  string  `json:"fourthCategoryName"`
	PrePayCommission    float64 `json:"prePayCommission"`
	FirstCategoryCode   string  `json:"firstCategoryCode"`
	SecondCategoryName  string  `json:"secondCategoryName"`
	WAPRate             float64 `json:"wapRate"`
	GoodsCode           string  `json:"goodsCode"`
	UpStatus            string  `json:"upStatus"`
	ProductURL          string  `json:"productUrl"`
	ProductURLWAP       string  `json:"productUrlWap"`
}
