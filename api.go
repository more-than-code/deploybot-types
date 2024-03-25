package types

type UpdateTaskStatusInput struct {
	PipelineId ObjectId
	TaskId     ObjectId
	Task       struct {
		Status string
	}
}

type WebhookResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type GetTaskResponse struct {
	Code    int                     `json:"code"`
	Msg     string                  `json:"msg"`
	Payload *GetTaskResponsePayload `json:"payload"`
}

type GetTaskResponsePayload struct {
	Task Task
}
