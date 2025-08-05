package alibabacpslistcpssettleinfodetail

import (
	"encoding/json"
	"github.com/mimicode/tksdk/i688opensdk/response"
)

// Response 获取联盟结算明细列表(按月)响应
type Response struct {
	response.TopResponse
	// 结算明细列表
	Result []OpenCPSSettleDetailDTO `json:"result"`
	// 总记录数
	TotalRow int `json:"totalRow"`
}

// OpenCPSSettleDetailDTO 联盟结算明细DTO
type OpenCPSSettleDetailDTO struct {
	// 结算月份
	SettleMonth string `json:"settleMonth"`
	// 预估收入
	EstimateIncome float64 `json:"estimateIncome"`
	// 维权返款
	RefundSuccAmount float64 `json:"refundSuccAmount"`
	// 维权金额
	RefundAmount float64 `json:"refundAmount"`
	// 冻结金额
	FreezeAmount float64 `json:"freezeAmount"`
	// 补扣金额
	RedeductAmount float64 `json:"redeductAmount"`
	// 月结发钱金额
	RealAmount float64 `json:"realAmount"`
	// 税费代扣金额
	TaxAmount float64 `json:"taxAmount"`
	// 技术服务费
	ServiceAmount float64 `json:"serviceAmount"`
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