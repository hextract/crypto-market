package dto

type JsonOrderResult struct {
	Order           *JsonOrder           `json:"order"`
	ExecutionResult *JsonExecutionResult `json:"execution_result"`
}
