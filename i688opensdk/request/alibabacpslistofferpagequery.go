package request

import (
	"net/url"
	"strconv"
)

// AlibabaCpsListOfferPageQueryRequest 获取联盟offer列表 API请求
// https://open.1688.com/api/apidocdetail.htm?spm=1688open.solution-detail.0.0.1d472cceFMBCLf&id=com.alibaba.p4p%3Aalibaba.cps.listOfferPageQuery-1&aopApiCategory=category_new
type AlibabaCpsListOfferPageQueryRequest struct {
	Parameters *url.Values //请求参数
}

// NewAlibabaCpsListOfferPageQueryRequest 创建AlibabaCpsListOfferPageQueryRequest实例
func NewAlibabaCpsListOfferPageQueryRequest() *AlibabaCpsListOfferPageQueryRequest {
	return &AlibabaCpsListOfferPageQueryRequest{}
}

// CheckParameters 检查参数
func (r *AlibabaCpsListOfferPageQueryRequest) CheckParameters() {
	// 检查必需参数
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	
	// pageNo和pageSize是必需参数
	if r.Parameters.Get("pageNo") == "" {
		panic("pageNo is required")
	}
	if r.Parameters.Get("pageSize") == "" {
		panic("pageSize is required")
	}
	
	// categoryId和feedInfo至少输入一项
	if r.Parameters.Get("categoryId") == "" && r.Parameters.Get("feedInfo") == "" {
		panic("categoryId and feedInfo at least one is required")
	}
}

// AddParameter 添加参数
func (r *AlibabaCpsListOfferPageQueryRequest) AddParameter(key, val string) {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	r.Parameters.Add(key, val)
}

// SetCategoryId 设置类目id
// 类目id和feedInfo至少输入一项
// 1-农业 2-食品、饮料 4-纺织、皮革 5-电工电气 6-家用电器 7-数码、电脑 8-化工 9-冶金矿产 10-能源
// 12-交通运输 13-家装、建材 15-日用百货 17-工艺品、礼品 18-运动装备 51-代理 52-纸业 53-传媒、广电
// 54-服饰配件、饰品 55-橡塑 56-精细化学品 57-电子元器件 58-照明工业 59-五金、工具 64-环保
// 65-机械及行业设备 66-医药、保养 67-办公、文教 68-包装 69-商务服务 70-安全、防护 71-汽摩及配件
// 72-印刷 73-项目合作 96-家纺家饰 97-美容护肤/彩妆 311-童装 312-内衣 509-通信产品 1426-机床
// 1501-母婴用品 1813-玩具 2805-加工 2829-二手设备转让 3007-个人防护 10165-男装 10166-女装
// 10208-仪器仪表 1038378-鞋 1042954-箱包皮具 122916001-宠物及园艺 122916002-汽车用品
// 123614001-钢铁 127380009-运动服饰 130822002-餐饮生鲜 130822220-个护/家清 130823000-性保健品
func (r *AlibabaCpsListOfferPageQueryRequest) SetCategoryId(categoryId int64) {
	r.AddParameter("categoryId", strconv.FormatInt(categoryId, 10))
}

// SetFeedInfo 设置关键词或offerId
// 类目id和feedInfo至少输入一项
// 示例：Mp3 或 537357997672
func (r *AlibabaCpsListOfferPageQueryRequest) SetFeedInfo(feedInfo string) {
	r.AddParameter("feedInfo", feedInfo)
}

// SetDefineTags 设置属性标签，属性间以逗号分隔
// 含实力商家、深度认证、一件代发、进口货源、是否包含优惠券等标志
// 示例：slsj_flag=true,sdrd_flag=true,yjdf_flag=true,jkhy_flag=true,yhq_flag=true
func (r *AlibabaCpsListOfferPageQueryRequest) SetDefineTags(defineTags string) {
	r.AddParameter("defineTags", defineTags)
}

// SetFilterMinPrice 设置售价下限
func (r *AlibabaCpsListOfferPageQueryRequest) SetFilterMinPrice(price float64) {
	r.AddParameter("filterMinPrice", strconv.FormatFloat(price, 'f', -1, 64))
}

// SetFilterMaxPrice 设置售价上限
func (r *AlibabaCpsListOfferPageQueryRequest) SetFilterMaxPrice(price float64) {
	r.AddParameter("filterMaxPrice", strconv.FormatFloat(price, 'f', -1, 64))
}

// SetFilterQuantityBeginMin 设置起批量下限
func (r *AlibabaCpsListOfferPageQueryRequest) SetFilterQuantityBeginMin(quantity int) {
	r.AddParameter("filterQuantityBeginMin", strconv.Itoa(quantity))
}

// SetFilterQuantityBeginMax 设置起批量上限
func (r *AlibabaCpsListOfferPageQueryRequest) SetFilterQuantityBeginMax(quantity int) {
	r.AddParameter("filterQuantityBeginMax", strconv.Itoa(quantity))
}

// SetFilterSaleQuantityMin 设置销量下限
func (r *AlibabaCpsListOfferPageQueryRequest) SetFilterSaleQuantityMin(quantity int) {
	r.AddParameter("filterSaleQuantityMin", strconv.Itoa(quantity))
}

// SetFilterSaleQuantityMax 设置销量上限
func (r *AlibabaCpsListOfferPageQueryRequest) SetFilterSaleQuantityMax(quantity int) {
	r.AddParameter("filterSaleQuantityMax", strconv.Itoa(quantity))
}

// SetFilterRatioMin 设置佣金比例下限
func (r *AlibabaCpsListOfferPageQueryRequest) SetFilterRatioMin(ratio float64) {
	r.AddParameter("filterRatioMin", strconv.FormatFloat(ratio, 'f', -1, 64))
}

// SetFilterRatioMax 设置佣金比例上限
func (r *AlibabaCpsListOfferPageQueryRequest) SetFilterRatioMax(ratio float64) {
	r.AddParameter("filterRatioMax", strconv.FormatFloat(ratio, 'f', -1, 64))
}

// SetSortField 设置排序字段(desc降序，asc升序)
// 销量：saleQuantity^asc
// 价格：price^desc
// 推广量：tkCnt^desc
// 分佣比例：ratio^desc
// 月支出佣金：tkCommission^desc
// 优惠券面值：yhqDiscountFee^desc
// 老买家分佣比例：oldBuyerRatio^desc
func (r *AlibabaCpsListOfferPageQueryRequest) SetSortField(sortField string) {
	r.AddParameter("sortField", sortField)
}

// SetPageNo 设置页号
func (r *AlibabaCpsListOfferPageQueryRequest) SetPageNo(pageNo int) {
	r.AddParameter("pageNo", strconv.Itoa(pageNo))
}

// SetPageSize 设置页大小
func (r *AlibabaCpsListOfferPageQueryRequest) SetPageSize(pageSize int) {
	r.AddParameter("pageSize", strconv.Itoa(pageSize))
}

// SetFilterOldBuyerRatio 设置老买家佣金比例筛选
// 注意这里与ratio格式不同，用~分割上下区间
// 如0.1~:表示大于等于0.1; ~0.2:表示小于等于0.2; 0.2~0.3:表示0.2到0.3
// 示例：0.2~0.3
func (r *AlibabaCpsListOfferPageQueryRequest) SetFilterOldBuyerRatio(ratio string) {
	r.AddParameter("filterOldBuyerRatio", ratio)
}

// GetApiName 获取API名称
func (r *AlibabaCpsListOfferPageQueryRequest) GetApiName() string {
	return "alibaba.cps.listOfferPageQuery"
}

// GetApiVersion 获取API版本
func (r *AlibabaCpsListOfferPageQueryRequest) GetApiVersion() string {
	return "param2/1"
}

// GetBusinessModule 获取业务模块
func (r *AlibabaCpsListOfferPageQueryRequest) GetBusinessModule() string {
	return "com.alibaba.p4p"
}

// GetParameters 获取所有参数
func (r *AlibabaCpsListOfferPageQueryRequest) GetParameters() url.Values {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	return *r.Parameters
}