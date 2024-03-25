package types

type Pipeline struct {
	Id            ObjectId           `json:"id" bson:"_id"`
	Name          string             `json:"name"`
	CreatedAt     Datetime           `json:"createAt"`
	UpdatedAt     Datetime           `json:"updateAt"`
	ExecutedAt    Datetime           `json:"executedAt"`
	StoppedAt     Datetime           `json:"stoppedAt"`
	ScheduledAt   Datetime           `json:"scheduledAt"`
	Status        string             `json:"status"`
	Arguments     []string           `json:"arguments"`
	Labels        map[string]*string `json:"labels"`
	Tasks         []Task             `json:"tasks"`
	RepoWatched   string             `json:"repoWatched"`
	BranchWatched string             `json:"branchWatched"`
	AutoRun       bool               `json:"autoRun"`
	ProjectId     ObjectId           `json:"projectId"`
}

type CreatePipelineInput struct {
	Name          string
	Arguments     []string
	Labels        map[string]*string
	RepoWatched   string
	BranchWatched string
	AutoRun       bool
	ProjectId     ObjectId
}

type TaskFilter struct {
	UpstreamTaskId *ObjectId
	AutoRun        *bool
}

type GetPipelineInput struct {
	Id         ObjectId
	Name       string
	TaskFilter TaskFilter
}

type GetPipelinesInput struct {
	RepoWatched   *string `bson:",omitempty"`
	BranchWatched *string `bson:",omitempty"`
	AutoRun       *bool   `bson:",omitempty"`
	ProjectId     ObjectId
}

type GetPipelinesOutput struct {
	TotalCount int        `json:"totalCount"`
	Items      []Pipeline `json:"items"`
}

type PipelineUpdate struct {
	Name          *string            `bson:",omitempty"`
	ScheduledAt   *Datetime          `bson:",omitempty"`
	Arguments     []string           `bson:",omitempty"`
	Labels        map[string]*string `bson:",omitempty"`
	RepoWatched   *string            `bson:",omitempty"`
	BranchWatched *string            `bson:",omitempty"`
	AutoRun       *bool              `bson:",omitempty"`
	ProjectId     *ObjectId          `bson:",omitempty"`
}

type UpdatePipelineInput struct {
	Id       ObjectId
	Pipeline PipelineUpdate
}

type UpdatePipelineStatusInput struct {
	PipelineId ObjectId
	Pipeline   struct {
		Status string
	}
}
