package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	ExHttpStatusAuthenticationFailure = 460
	ExHttpStatusBusinessLogicError    = 461
)

const (
	CodeClientError      = 1000
	CodePipelineBusy     = 1100
	CodePipelineNotFound = 1101
	CodeTaskBusy         = 1102
	CodeTaskNotFound     = 1103

	CodeServerError = 2000
)

const (
	MsgClientError      = "Client error"
	MsgPipelineBusy     = "Pipleline busy"
	MsgPipelineNotFound = "Pipleline not found"
	MsgTaskBusy         = "Task not found"
	MsgTaskNotFound     = "Task not found"

	MsgServerError = "Server error"
)

type PostPipelineResponsePayload struct {
	Id primitive.ObjectID
}
type PostPipelineResponse struct {
	Code    int                          `json:"code"`
	Msg     string                       `json:"msg"`
	Payload *PostPipelineResponsePayload `json:"payload"`
}

type GetPipelinesResponse struct {
	Code    int                 `json:"code"`
	Msg     string              `json:"msg"`
	Payload *GetPipelinesOutput `json:"payload"`
}

type GetPipelineResponsePayload struct {
	Pipeline Pipeline
}

type GetPipelineResponse struct {
	Code    int                         `json:"code"`
	Msg     string                      `json:"msg"`
	Payload *GetPipelineResponsePayload `json:"payload"`
}

type DeletePipelineResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type PatchPipelineResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type PutPipelineStatusResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type PostTaskResponsePayload struct {
	Id primitive.ObjectID
}
type PostTaskResponse struct {
	Code    int                      `json:"code"`
	Msg     string                   `json:"msg"`
	Payload *PostTaskResponsePayload `json:"payload"`
}

type GetTaskResponsePayload struct {
	Task Task
}
type GetTaskResponse struct {
	Code    int                     `json:"code"`
	Msg     string                  `json:"msg"`
	Payload *GetTaskResponsePayload `json:"payload"`
}

type DeleteTaskResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type PatchTaskResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type PutTaskStatusResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type WebhookResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type AuthenticationResponse struct {
	Msg     string                `json:"msg"`
	Code    int                   `json:"code"`
	Payload *AuthenticationOutput `json:"payload"`
}

type PostUserResponse struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type DeleteUserResponse struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type GetUserResponse struct {
	Msg     string `json:"msg"`
	Code    int    `json:"code"`
	Payload *User  `json:"payload"`
}

type GetUsersResponse struct {
	Msg     string          `json:"msg"`
	Code    int             `json:"code"`
	Payload *GetUsersOutput `json:"payload"`
}

type GetProjectsResponse struct {
	Msg     string             `json:"msg"`
	Code    int                `json:"code"`
	Payload *GetProjectsOutput `json:"payload"`
}

type GetProjectResponse struct {
	Msg     string   `json:"msg"`
	Code    int      `json:"code"`
	Payload *Project `json:"payload"`
}

type DeleteProjectResponse struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type PostProjectResponse struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type PatchProjectResponse struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type DeleteMemberResponse struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type PostMemberResponse struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type PatchMemberResponse struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type GetDiskInfoResponse struct {
	Msg     string    `json:"msg"`
	Code    int       `json:"code"`
	Payload *DiskInfo `json:"payload"`
}
