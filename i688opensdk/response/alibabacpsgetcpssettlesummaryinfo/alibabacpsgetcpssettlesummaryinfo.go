package alibabacpsgetcpssettlesummaryinfo

import (
	"encoding/json"
	"github.com/mimicode/tksdk/i688opensdk/response"
)

// Response 获取联盟结算摘要账单响应
type Response struct {
	response.TopResponse
	// 响应结果
	Result OpenCPSSettleSummaryDTO `json:"result"`
}

// OpenCPSSettleSummaryDTO 联盟结算摘要账单响应结果
type OpenCPSSettleSummaryDTO struct {
	// 昨日预估收入
	YesterdayEstimateAmount float64 `json:"yesterdayEstimateAmount"`
	// 本月预估收入
	ThisMonthEstimateAmount float64 `json:"thisMonthEstimateAmount"`
	// 上月预估收入
	LastMonthEstimateAmount float64 `json:"lastMonthEstimateAmount"`
	// 预估未结算收入
	UnsettleEstimateAmount float64 `json:"unsettleEstimateAmount"`
	// 可提现余额
	SettleAmount float64 `json:"settleAmount"`
}

// WrapResult 包装结果
func (r *Response) WrapResult(result string) {
	// 保存原始结果到Body
	r.Body = result
	// 尝试解析具体的响应结构
	if err := json.Unmarshal([]byte(result), r); err != nil {
		// 如果解析失败，设置错误信息
		r.ErrorResponse.Code = 500
		r.ErrorResponse.Msg = "Failed to parse response: " + err.Error()
		return
	}
	// 检查是否有错误响应
	if r.ErrorResponse.Code != 0 || r.ErrorResponse.SubCode != "" {
		// 错误信息已经在ErrorResponse中，无需额外处理
		return
	}
}