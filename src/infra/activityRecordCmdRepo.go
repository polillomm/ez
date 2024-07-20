package infra

import (
	"github.com/speedianet/control/src/domain/dto"
	"github.com/speedianet/control/src/infra/db"
	dbModel "github.com/speedianet/control/src/infra/db/model"
)

type ActivityRecordCmdRepo struct {
	trailDbSvc *db.TrailDatabaseService
}

func NewActivityRecordCmdRepo(
	trailDbSvc *db.TrailDatabaseService,
) *ActivityRecordCmdRepo {
	return &ActivityRecordCmdRepo{trailDbSvc: trailDbSvc}
}

func (repo *ActivityRecordCmdRepo) Create(createDto dto.CreateActivityRecord) error {
	var codePtr *string
	if createDto.Code != nil {
		code := createDto.Code.String()
		codePtr = &code
	}

	var messagePtr *string
	if createDto.Message != nil {
		message := createDto.Message.String()
		messagePtr = &message
	}

	var ipAddressPtr *string
	if createDto.IpAddress != nil {
		ipAddress := createDto.IpAddress.String()
		ipAddressPtr = &ipAddress
	}

	var operatorAccountIdPtr *uint64
	if createDto.OperatorAccountId != nil {
		operatorAccountId := createDto.OperatorAccountId.Uint64()
		operatorAccountIdPtr = &operatorAccountId
	}

	var targetAccountIdPtr *uint64
	if createDto.TargetAccountId != nil {
		targetAccountId := createDto.TargetAccountId.Uint64()
		targetAccountIdPtr = &targetAccountId
	}

	var usernamePtr *string
	if createDto.Username != nil {
		username := createDto.Username.String()
		usernamePtr = &username
	}

	var containerIdPtr *string
	if createDto.ContainerId != nil {
		containerId := createDto.ContainerId.String()
		containerIdPtr = &containerId
	}

	var containerProfileIdPtr *uint64
	if createDto.ContainerProfileId != nil {
		containerProfileId := createDto.ContainerProfileId.Uint64()
		containerProfileIdPtr = &containerProfileId
	}

	var mappingIdPtr *uint64
	if createDto.MappingId != nil {
		mappingId := createDto.MappingId.Uint64()
		mappingIdPtr = &mappingId
	}

	securityEventModel := dbModel.NewActivityRecord(
		0, createDto.Level.String(), codePtr, messagePtr, ipAddressPtr,
		operatorAccountIdPtr, targetAccountIdPtr, usernamePtr, containerIdPtr,
		containerProfileIdPtr, mappingIdPtr,
	)

	return repo.trailDbSvc.Handler.Create(&securityEventModel).Error
}

func (repo *ActivityRecordCmdRepo) Delete(deleteDto dto.DeleteActivityRecords) error {
	deleteModel := dbModel.ActivityRecord{}
	if deleteDto.Id != nil {
		deleteModel.ID = deleteDto.Id.Uint64()
	}

	if deleteDto.Level != nil {
		deleteModel.Level = deleteDto.Level.String()
	}

	if deleteDto.Code != nil {
		codeStr := deleteDto.Code.String()
		deleteModel.Code = &codeStr
	}

	if deleteDto.Message != nil {
		messageStr := deleteDto.Message.String()
		deleteModel.Message = &messageStr
	}

	if deleteDto.IpAddress != nil {
		ipAddressStr := deleteDto.IpAddress.String()
		deleteModel.IpAddress = &ipAddressStr
	}

	if deleteDto.OperatorAccountId != nil {
		operatorAccountId := deleteDto.OperatorAccountId.Uint64()
		deleteModel.OperatorAccountId = &operatorAccountId
	}

	if deleteDto.TargetAccountId != nil {
		targetAccountId := deleteDto.TargetAccountId.Uint64()
		deleteModel.TargetAccountId = &targetAccountId
	}

	if deleteDto.Username != nil {
		usernameStr := deleteDto.Username.String()
		deleteModel.Username = &usernameStr
	}

	if deleteDto.ContainerId != nil {
		containerIdStr := deleteDto.ContainerId.String()
		deleteModel.ContainerId = &containerIdStr
	}

	if deleteDto.ContainerProfileId != nil {
		containerProfileId := deleteDto.ContainerProfileId.Uint64()
		deleteModel.ContainerProfileId = &containerProfileId
	}

	if deleteDto.MappingId != nil {
		mappingId := deleteDto.MappingId.Uint64()
		deleteModel.MappingId = &mappingId
	}

	dbQuery := repo.trailDbSvc.Handler.Where(&deleteModel)

	if deleteDto.CreatedAt != nil {
		dbQuery.Where("created_at >= ?", deleteDto.CreatedAt.GetAsGoTime())
	}

	return dbQuery.Delete(&dbModel.ActivityRecord{}).Error
}
