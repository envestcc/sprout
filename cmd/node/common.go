package main

type errResp struct {
	Error string `json:"error,omitempty"`
}

func newErrResp(err error) *errResp {
	return &errResp{Error: err.Error()}
}

type msgReq struct {
	ProjectID      string `json:"projectID"        binding:"required"`
	ProjectVersion string `json:"projectVersion"   binding:"required"`
	Data           string `json:"data"             binding:"required"`
}
