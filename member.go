package types

type Member struct {
	UserId    ObjectId `json:"userId"`
	Role      Role     `json:"role"`
	CreatedAt Datetime `json:"datetimeCreated"`
}

type CreateMemberInput struct {
	ProjectId ObjectId
	Member    struct {
		UserId ObjectId
		Role   Role
	}
}

type DeleteMemberInput struct {
	ProjectId ObjectId
	UserId    ObjectId
}

type UpdateMember struct {
	Role *Role
}
type UpdateMemberInput struct {
	ProjectId ObjectId
	UserId    ObjectId
	Member    UpdateMember
}
