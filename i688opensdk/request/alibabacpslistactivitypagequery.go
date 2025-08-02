package request

import (
	"net/url"
)

// alibaba.cps.listActivityPageQuery( 获取联盟活动列表 )
// https://open.1688.com/solution/solutionDetail.htm?spm=a260s.develop-solution-purchased.0.0.672e46cbWG53ww&solutionKey=1515046370581#apiAndMessageList
type AlibabaCpsListActivityPageQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (r *AlibabaCpsListActivityPageQueryRequest) CheckParameters() {
	// 该API无必填参数
}

// 添加请求参数
func (r *AlibabaCpsListActivityPageQueryRequest) AddParameter(key, val string) {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	r.Parameters.Add(key, val)
}

// /////////////////////// 应用级参数 ///////////////////////////

// /////////////////////// 接口信息 ///////////////////////////
// 返回接口名称
func (r *AlibabaCpsListActivityPageQueryRequest) GetApiName() string {
	return "alibaba.cps.listActivityPageQuery"
}

// 返回接口版本
func (r *AlibabaCpsListActivityPageQueryRequest) GetApiVersion() string {
	return "param2/1"
}

// 返回业务模块
func (r *AlibabaCpsListActivityPageQueryRequest) GetBusinessModule() string {
	return "com.alibaba.p4p"
}

// 返回请求参数
func (r *AlibabaCpsListActivityPageQueryRequest) GetParameters() url.Values {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	return *r.Parameters
}
