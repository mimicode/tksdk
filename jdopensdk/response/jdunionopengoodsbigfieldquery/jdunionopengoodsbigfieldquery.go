package jdunionopengoodsbigfieldquery

import (
	"encoding/json"

	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.goods.bigfield.query 商品详情查询接口,查询商详大字段信息【支持用户授权】
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_goods_bigfield_query_responce"`
}

// WrapResult 解析输出结果
func (t *Response) WrapResult(result string) {
	err := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if err != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Message = err.Error()
	} else {
		//解析queryResult
		if t.Responce.QueryResultStr == "" {
			t.ErrorResponse.Code = -2
			t.ErrorResponse.Message = "empty queryResult"
		} else {
			if err = json.Unmarshal([]byte(t.Responce.QueryResultStr), &t.Responce.QueryResult); err != nil {
				t.ErrorResponse.Code = -1
				t.ErrorResponse.Message = err.Error()
			} else {
				t.ErrorResponse.Code = t.Responce.QueryResult.Code
				t.ErrorResponse.Message = t.Responce.QueryResult.Message
				t.ErrorResponse.RequestID = t.Responce.QueryResult.RequestID
			}
		}
	}
	t.Responce.QueryResultStr = ""
}

// Responce 响应结果
type Responce struct {
	Code           string `json:"code"`
	QueryResultStr string `json:"queryResult"`
	QueryResult    QueryResult
}

// QueryResult 具体内容
type QueryResult struct {
	Code      int64               `json:"code"`
	Data      []BigFieldGoodsResp `json:"data"`
	Message   string              `json:"message"`
	RequestID string              `json:"requestId"`
}

// BigFieldGoodsResp 商品大字段信息
type BigFieldGoodsResp struct {
	SkuID             int64             `json:"skuId"`
	SkuName           string            `json:"skuName"`
	CategoryInfo      CategoryInfo      `json:"categoryInfo"`
	ImageInfo         ImageInfo         `json:"imageInfo"`
	BaseBigFieldInfo  BaseBigFieldInfo  `json:"baseBigFieldInfo"`
	BookBigFieldInfo  BookBigFieldInfo  `json:"bookBigFieldInfo"`
	VideoBigFieldInfo VideoBigFieldInfo `json:"videoBigFieldInfo"`
	MainSkuID         int64             `json:"mainSkuId"`    //自营主skuId
	ProductID         int64             `json:"productId"`    //非自营商品Id
	SkuStatus         int64             `json:"skuStatus"`    //sku上下架状态 1：上架(可搜索，可购买)，0：下架(可通过skuid搜索，不可购买)，2：可上架（可通过skuid搜索，不可购买），10：pop删除(不可搜索，不可购买)
	Owner             string            `json:"owner"`        //g=自营，p=pop
	DetailImages      string            `json:"detailImages"` //商详图
	ItemID            string            `json:"itemId"`       //联盟商品ID(原始入参ItemId)
	CallerItemID      string            `json:"callerItemId"` //工具商联盟商品ID
}

// CategoryInfo 类目信息
type CategoryInfo struct {
	Cid1     int64  `json:"cid1"`
	Cid1Name string `json:"cid1Name"`
	Cid2     int64  `json:"cid2"`
	Cid2Name string `json:"cid2Name"`
	Cid3     int64  `json:"cid3"`
	Cid3Name string `json:"cid3Name"`
}

// ImageInfo 图片信息
type ImageInfo struct {
	ImageList  []UrlInfo `json:"imageList"`
	WhiteImage string    `json:"whiteImage"` //白底图【废弃】
}

// UrlInfo 图片合集
type UrlInfo struct {
	URL string `json:"url"`
}

// BaseBigFieldInfo 基础大字段信息
type BaseBigFieldInfo struct {
	Wdis       string `json:"wdis"`       //商品介绍
	PropCode   string `json:"propCode"`   //规格参数【废弃】
	WareQD     string `json:"wareQD"`     //包装清单(仅自营商品)
	PropGroups string `json:"propGroups"` //规格参数(JSON串)
}

// BookBigFieldInfo 图书大字段信息
type BookBigFieldInfo struct {
	Comments        string `json:"comments"`        //媒体评论
	Image           string `json:"image"`           //精彩文摘与插图(插图)
	ContentDesc     string `json:"contentDesc"`     //内容摘要(内容简介)
	RelatedProducts string `json:"relatedProducts"` //产品描述(相关商品)
	EditerDesc      string `json:"editerDesc"`      //编辑推荐
	Catalogue       string `json:"catalogue"`       //目录
	BookAbstract    string `json:"bookAbstract"`    //精彩摘要(精彩书摘)
	AuthorDesc      string `json:"authorDesc"`      //作者简介
	Introduction    string `json:"introduction"`    //前言(前言/序言)
	ProductFeatures string `json:"productFeatures"` //产品特色
}

// VideoBigFieldInfo 影音大字段信息
type VideoBigFieldInfo struct {
	Comments            string `json:"comments"`             //评论
	Image               string `json:"image"`                //商品描述(精彩剧照)
	ContentDesc         string `json:"contentDesc"`          //内容摘要(内容简介)
	EditerDesc          string `json:"editerDesc"`           //编辑推荐
	Catalogue           string `json:"catalogue"`            //目录
	BoxContents         string `json:"box_Contents"`         //包装清单
	MaterialDescription string `json:"material_Description"` //特殊说明
	Manual              string `json:"manual"`               //说明书
	ProductFeatures     string `json:"productFeatures"`      //产品特色
}
