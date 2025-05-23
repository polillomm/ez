package entity

import (
	"github.com/goinfinite/ez/src/domain/valueObject"
)

type BackupTask struct {
	TaskId                 valueObject.BackupTaskId               `json:"taskId"`
	AccountId              valueObject.AccountId                  `json:"accountId"`
	AccountUsername        valueObject.UnixUsername               `json:"accountUsername"`
	JobId                  valueObject.BackupJobId                `json:"jobId"`
	DestinationId          valueObject.BackupDestinationId        `json:"destinationId"`
	TaskStatus             valueObject.BackupTaskStatus           `json:"taskStatus"`
	RetentionStrategy      valueObject.BackupRetentionStrategy    `json:"retentionStrategy"`
	BackupSchedule         valueObject.CronSchedule               `json:"backupSchedule"`
	TimeoutSecs            valueObject.TimeDuration               `json:"timeoutSecs"`
	ContainerAccountIds    []valueObject.AccountId                `json:"containerAccountIds"`
	SuccessfulContainerIds []valueObject.ContainerId              `json:"successfulContainerIds"`
	FailedContainerIds     []valueObject.ContainerId              `json:"failedContainerIds"`
	ExecutionOutput        *valueObject.BackupTaskExecutionOutput `json:"executionOutput"`
	SizeBytes              *valueObject.Byte                      `json:"sizeBytes"`
	StartedAt              *valueObject.UnixTime                  `json:"startedAt"`
	FinishedAt             *valueObject.UnixTime                  `json:"finishedAt"`
	ElapsedSecs            *valueObject.TimeDuration              `json:"elapsedSecs"`
	CreatedAt              valueObject.UnixTime                   `json:"createdAt"`
	UpdatedAt              valueObject.UnixTime                   `json:"updatedAt"`
}

func NewBackupTask(
	taskId valueObject.BackupTaskId,
	accountId valueObject.AccountId,
	accountUsername valueObject.UnixUsername,
	jobId valueObject.BackupJobId,
	destinationId valueObject.BackupDestinationId,
	taskStatus valueObject.BackupTaskStatus,
	retentionStrategy valueObject.BackupRetentionStrategy,
	backupSchedule valueObject.CronSchedule,
	timeoutSecs valueObject.TimeDuration,
	containerAccountIds []valueObject.AccountId,
	successfulContainerIds, failedContainerIds []valueObject.ContainerId,
	executionOutput *valueObject.BackupTaskExecutionOutput,
	sizeBytes *valueObject.Byte,
	startedAt, finishedAt *valueObject.UnixTime,
	elapsedSecs *valueObject.TimeDuration,
	createdAt, updatedAt valueObject.UnixTime,
) BackupTask {
	return BackupTask{
		TaskId:                 taskId,
		AccountId:              accountId,
		AccountUsername:        accountUsername,
		JobId:                  jobId,
		DestinationId:          destinationId,
		TaskStatus:             taskStatus,
		RetentionStrategy:      retentionStrategy,
		BackupSchedule:         backupSchedule,
		TimeoutSecs:            timeoutSecs,
		ContainerAccountIds:    containerAccountIds,
		SuccessfulContainerIds: successfulContainerIds,
		FailedContainerIds:     failedContainerIds,
		ExecutionOutput:        executionOutput,
		SizeBytes:              sizeBytes,
		StartedAt:              startedAt,
		FinishedAt:             finishedAt,
		ElapsedSecs:            elapsedSecs,
		CreatedAt:              createdAt,
		UpdatedAt:              updatedAt,
	}
}
