package types

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

const (
	CodeAuthenticationFailure = 100

	CodeInvalidParameters    = 1200
	CodeWrongEmailOrPassword = 1201
	CodeDeletedUser          = 1202

	CodeExpiredVerificationCode              = 1100
	CodeWrongVerificationCode                = 1101
	CodeTooFrequentlySendingVerificationCode = 1102
	CodeNeedingResendingVerificationCode     = 1103
	CodeMaximumAttemptsOnVerificationCode    = 1104
)

const (
	MsgAuthenticationFailure = "authentication failure"

	MsgInvalidParamters     = "invalid parameter(s)"
	MsgWrongEmailOrPassword = "wrong email or password"

	MsgExpiredVerificationCode              = "expired verification code"
	MsgWrongVerificationCode                = "wrong verification code"
	MsgDeletedUser                          = "deleted user"
	MsgTooFrequentlySendingVerificationCode = "too frequently sending code"
	MsgNeedingResendingVerificationCode     = "needing resending verification code"
	MsgMaximumAttemptsOnVerificationCode    = "maximum attempts on verification code"
)

const (
	TaskPending    = "Pending"
	TaskInProgress = "InProgress"
	TaskDone       = "Done"
	TaskCanceled   = "Canceled"
	TaskFailed     = "Failed"
	TaskTimedOut   = "TimedOut"
)

const (
	PipelineIdle = "Idle"
	PipelineBusy = "Busy"
)

const (
	EventBuild  = "Build"
	EventDeploy = "Deploy"
)

type TaskType string

const (
	TestTask   TaskType = "test"
	BuildTask  TaskType = "build"
	DeployTask TaskType = "deploy"
)

type Role string

const (
	RoleOwner  Role = "Owner"
	RoleAdmin  Role = "Admin"
	RoleMember Role = "Member"
)
