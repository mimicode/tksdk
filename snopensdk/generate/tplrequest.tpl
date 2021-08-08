package request

//--APIDESC--
//--APIURL--
type --APINAME--Request struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *--APINAME--Request) CheckParameters() {
	--APIPARAMCHECK--
}

//添加请求参数
func (tk *--APINAME--Request) AddParameter(key string,val interface{} ) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *--APINAME--Request) GetApiName() string {
	return "--APIORGNAME--"
}

//方法名称
func (tk *--APINAME--Request) GetVersion() string {
	return "--APIVERSION--"
}

//返回请求参数
func (tk *--APINAME--Request) GetParameters() map[string]interface{} {
	return tk.Parameters
}
