package types

type Project struct {
	Id          ObjectId `json:"id" bson:"_id"`
	Name        string   `json:"name"`
	AvatarUrl   string   `json:"avatarUrl"`
	OwnerUserId ObjectId `json:"ownerUserId"`
	Members     []Member `json:"members"`
	// Pipelines   []Pipeline         `json:"pipelines"`
	CreatedAt Datetime `json:"createdAt"`
	UpdatedAt Datetime `json:"updatedAt"`
}

type CreateProjectInput struct {
	Name   string
	UserId ObjectId
}

type UpdateProject struct {
	Name      *string `bson:",omitempty"`
	AvatarUrl *string `bson:",omitempty"`
}

type UpdateProjectInput struct {
	Id      ObjectId
	UserId  ObjectId
	Project UpdateProject
}

type DeleteProjectInput struct {
	Id     ObjectId
	UserId ObjectId
}

type GetProjectsInput struct {
	UserId ObjectId
}

type GetProjectInput struct {
	UserId ObjectId
	Id     ObjectId
}

type GetProjectsOutput struct {
	Items      []Project `json:"items"`
	TotalCount int64     `json:"totalCount"`
}
