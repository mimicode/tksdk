package alibabacpsgensearchpjjxintroduceurlbykeyword

import (
	"encoding/json"

	"github.com/mimicode/tksdk/i688opensdk/response"
)

// Response 生成平价精选搜索页推广链接响应
type Response struct {
	response.TopResponse
	AlibabaCpsGenSearchPjjxIntroduceUrlByKeywordResponse *ResultVO `json:"result"`
}

type AlibabaCpsGenSearchPjjxIntroduceUrlByKeywordResponse struct {
	Response
}

// ResultVO 结果对象
type ResultVO struct {
	Result  *PjjxUrlResult `json:"result"`
	Success bool           `json:"success"`
}

// PjjxUrlResult 平价精选URL结果
type PjjxUrlResult struct {
	PcPjjxCpsUrl  string `json:"pcPjjxCpsUrl"`  // PC端推广链接
	WapPjjxCpsUrl string `json:"wapPjjxCpsUrl"` // 移动端推广链接
}

func (r *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), r)
	//保存原始信息
	r.Body = result
	//解析错误
	if unmarshal != nil {
		r.ErrorResponse.Code = -1
		r.ErrorResponse.Msg = unmarshal.Error()
	}
}