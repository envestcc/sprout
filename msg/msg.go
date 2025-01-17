package msg

import "fmt"

type Msg struct {
	ProjectID      string `json:"projectID"`
	ProjectVersion string `json:"projectVersion"`
	Data           string `json:"data"`
}

type MsgKey string

func (m *Msg) Key() MsgKey {
	return MsgKey(fmt.Sprintf("%s:%s", m.ProjectID, m.ProjectVersion))
}

type FetchStrategy string

const (
	FIFO FetchStrategy = "fifo"
)
