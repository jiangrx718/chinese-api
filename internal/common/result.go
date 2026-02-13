package common

func NewCRMServiceResult() *CRMServiceResult {
	return &CRMServiceResult{}
}

type CRMServiceResult struct {
	BaseServiceResult
	Offset int64 `json:"offset,omitempty"`
	Limit  int64 `json:"limit,omitempty"`
	Count  int64 `json:"count,omitempty"`
}
