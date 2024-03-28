package types

type Task struct {
	Id             ObjectId    `json:"id"`
	Name           string      `json:"name"`
	CreatedAt      Datetime    `json:"createdAt"`
	UpdatedAt      Datetime    `json:"updatedAt"`
	ExecutedAt     Datetime    `json:"executedAt"`
	StoppedAt      Datetime    `json:"stoppedAt"`
	ScheduledAt    Datetime    `json:"scheduledAt"`
	Status         string      `json:"status"`
	UpstreamTaskId ObjectId    `json:"upstreamTaskId" bson:",omitempty"`
	StreamWebhook  string      `json:"streamWebhook" bson:",omitempty"`
	LogUrl         string      `json:"logUrl" bson:",omitempty"`
	Config         interface{} `json:"config"`
	Remarks        string      `json:"remarks"`
	AutoRun        bool        `json:"autoRun"`
	Timeout        int64       `json:"timeout"` // minutes
	Type           string      `json:"type"`
}
