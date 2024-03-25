package types

type BuildConfig struct {
	ImageName  string             `json:"imageName"`
	ImageTag   string             `json:"imageTag" bson:",omitempty"`
	Args       map[string]*string `json:"args" bson:",omitempty"`
	Dockerfile string             `json:"dockerfile" bson:",omitempty"`
	RepoUrl    string             `json:"repoUrl"`
	RepoName   string             `json:"repoName"`
	RepoBranch string             `json:"repoBranch"`
}

type FileMountConfig struct {
	Content string `json:"content" bson:",omitempty"`
	Name    string `json:"name" bson:",omitempty"`
}

type RestartPolicy struct {
	Name              string `json:"name" bson:",omitempty"`
	MaximumRetryCount int    `json:"maxiumRetryCount" bson:",omitempty"`
}

type DeployConfig struct {
	ImageName     string            `json:"imageName"`
	ImageTag      string            `json:"imageTag" bson:",omitempty"`
	ServiceName   string            `json:"serviceName" bson:",omitempty"`
	MountSource   string            `json:"mountSource" bson:",omitempty"`
	MountTarget   string            `json:"mountTarget" bson:",omitempty"`
	FilesToMount  []FileMountConfig `json:"filesToMount" bson:",omitempty"`
	AutoRemove    bool              `json:"autoRemove"`
	RestartPolicy RestartPolicy     `json:"restartPolicy" bson:",omitempty"`
	Env           []string          `json:"env" bson:",omitempty"`
	HostPort      string            `json:"hostPort" bson:",omitempty"`
	ExposedPort   string            `json:"exposedPort" bson:",omitempty"`
	NetworkId     string            `json:"networkId" bson:",omitempty"`
	NetworkName   string            `json:"networkName" bson:",omitempty"`
}

type RestartConfig struct {
	ServiceName string `json:"serviceName"`
}

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

type UpdateTaskInputTask struct {
	Name           *string
	UpstreamTaskId *ObjectId
	StreamWebhook  *string
	LogUrl         *string
	ScheduledAt    *Datetime
	Config         *interface{}
	Remarks        *string
	AutoRun        *bool
	Timeout        *int64
	Type           *string
}

type UpdateTaskInput struct {
	PipelineId ObjectId
	Id         ObjectId
	Task       UpdateTaskInputTask
}

type UpdateTaskStatusInput struct {
	PipelineId ObjectId
	TaskId     ObjectId
	Task       struct {
		Status string
	}
}

type CreateTaskInputTask struct {
	Id             ObjectId
	Name           string
	ScheduledAt    Datetime `bson:",omitempty"`
	Config         interface{}
	UpstreamTaskId ObjectId `bson:",omitempty"`
	StreamWebhook  string
	LogUrl         string
	AutoRun        bool
	Timeout        int64
	Type           string
}
type CreateTaskInput struct {
	PipelineId ObjectId
	Task       CreateTaskInputTask
}

type GetTaskInput struct {
	PipelineId ObjectId
	Id         ObjectId
}

type GetTasksInput struct {
	PipelineId     ObjectId
	UpstreamTaskId *ObjectId
}

type DeleteTaskInput struct {
	PipelineId ObjectId
	Id         ObjectId
}
