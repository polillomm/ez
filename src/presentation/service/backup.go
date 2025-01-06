package service

import (
	"errors"

	"github.com/goinfinite/ez/src/domain/dto"
	"github.com/goinfinite/ez/src/domain/useCase"
	"github.com/goinfinite/ez/src/domain/valueObject"
	voHelper "github.com/goinfinite/ez/src/domain/valueObject/helper"
	"github.com/goinfinite/ez/src/infra"
	backupInfra "github.com/goinfinite/ez/src/infra/backup"
	"github.com/goinfinite/ez/src/infra/db"
	serviceHelper "github.com/goinfinite/ez/src/presentation/service/helper"
)

type BackupService struct {
	persistentDbSvc       *db.PersistentDatabaseService
	backupQueryRepo       *backupInfra.BackupQueryRepo
	activityRecordCmdRepo *infra.ActivityRecordCmdRepo
}

func NewBackupService(
	persistentDbSvc *db.PersistentDatabaseService,
	trailDbSvc *db.TrailDatabaseService,
) *BackupService {
	return &BackupService{
		persistentDbSvc:       persistentDbSvc,
		backupQueryRepo:       backupInfra.NewBackupQueryRepo(persistentDbSvc),
		activityRecordCmdRepo: infra.NewActivityRecordCmdRepo(trailDbSvc),
	}
}

func (service *BackupService) ReadDestination(input map[string]interface{}) ServiceOutput {
	var destinationIdPtr *valueObject.BackupDestinationId
	if input["destinationId"] != nil {
		destinationId, err := valueObject.NewBackupDestinationId(input["destinationId"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		destinationIdPtr = &destinationId
	}

	var accountIdPtr *valueObject.AccountId
	if input["accountId"] != nil {
		accountId, err := valueObject.NewAccountId(input["accountId"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		accountIdPtr = &accountId
	}

	var destinationNamePtr *valueObject.BackupDestinationName
	if input["destinationName"] != nil {
		destinationName, err := valueObject.NewBackupDestinationName(input["destinationName"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		destinationNamePtr = &destinationName
	}

	var destinationTypePtr *valueObject.BackupDestinationType
	if input["destinationType"] != nil {
		destinationType, err := valueObject.NewBackupDestinationType(input["destinationType"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		destinationTypePtr = &destinationType
	}

	var objectStorageProviderPtr *valueObject.ObjectStorageProvider
	if input["objectStorageProvider"] != nil {
		objectStorageProvider, err := valueObject.NewObjectStorageProvider(input["objectStorageProvider"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		objectStorageProviderPtr = &objectStorageProvider
	}

	var remoteHostnamePtr *valueObject.Fqdn
	if input["remoteHostname"] != nil {
		remoteHostname, err := valueObject.NewFqdn(input["remoteHostname"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		remoteHostnamePtr = &remoteHostname
	}

	var remoteHostTypePtr *valueObject.BackupDestinationRemoteHostType
	if input["remoteHostType"] != nil {
		remoteHostType, err := valueObject.NewBackupDestinationRemoteHostType(input["remoteHostType"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		remoteHostTypePtr = &remoteHostType
	}

	timeParamNames := []string{"createdBeforeAt", "createdAfterAt"}
	timeParamPtrs := serviceHelper.TimeParamsParser(timeParamNames, input)

	requestPagination, err := serviceHelper.PaginationParser(
		input, useCase.BackupDestinationsDefaultPagination,
	)
	if err != nil {
		return NewServiceOutput(UserError, err)
	}

	readDto := dto.ReadBackupDestinationsRequest{
		Pagination:            requestPagination,
		DestinationId:         destinationIdPtr,
		AccountId:             accountIdPtr,
		DestinationName:       destinationNamePtr,
		DestinationType:       destinationTypePtr,
		ObjectStorageProvider: objectStorageProviderPtr,
		RemoteHostType:        remoteHostTypePtr,
		RemoteHostname:        remoteHostnamePtr,
		CreatedBeforeAt:       timeParamPtrs["createdBeforeAt"],
		CreatedAfterAt:        timeParamPtrs["createdAfterAt"],
	}

	responseDto, err := useCase.ReadBackupDestinations(service.backupQueryRepo, readDto)
	if err != nil {
		return NewServiceOutput(InfraError, err.Error())
	}

	return NewServiceOutput(Success, responseDto)
}

func (service *BackupService) CreateDestination(
	input map[string]interface{},
) ServiceOutput {
	requiredParams := []string{"accountId", "destinationName", "destinationType"}

	err := serviceHelper.RequiredParamsInspector(input, requiredParams)
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	accountId, err := valueObject.NewAccountId(input["accountId"])
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	destinationName, err := valueObject.NewBackupDestinationName(input["destinationName"])
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	var destinationDescriptionPtr *valueObject.BackupDestinationDescription
	if input["destinationDescription"] != nil {
		destinationDescription, err := valueObject.NewBackupDestinationDescription(
			input["destinationDescription"],
		)
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		destinationDescriptionPtr = &destinationDescription
	}

	destinationType, err := valueObject.NewBackupDestinationType(input["destinationType"])
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	var destinationPathPtr *valueObject.UnixFilePath
	if input["destinationPath"] != nil {
		destinationPath, err := valueObject.NewUnixFilePath(input["destinationPath"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		destinationPathPtr = &destinationPath
	}

	var minLocalStorageFreePercentPtr *uint8
	if input["minLocalStorageFreePercent"] != nil {
		minLocalStorageFreePercent, err := voHelper.InterfaceToUint8(
			input["minLocalStorageFreePercent"],
		)
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidMinLocalStorageFreePercent"))
		}
		minLocalStorageFreePercentPtr = &minLocalStorageFreePercent
	}

	var maxDestinationStorageUsagePercentPtr *uint8
	if input["maxDestinationStorageUsagePercent"] != nil {
		maxDestinationStorageUsagePercent, err := voHelper.InterfaceToUint8(
			input["maxDestinationStorageUsagePercent"],
		)
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidMaxDestinationStorageUsagePercent"))
		}
		maxDestinationStorageUsagePercentPtr = &maxDestinationStorageUsagePercent
	}

	var maxConcurrentConnectionsPtr *uint16
	if input["maxConcurrentConnections"] != nil {
		maxConcurrentConnections, err := voHelper.InterfaceToUint16(
			input["maxConcurrentConnections"],
		)
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidMaxConcurrentConnections"))
		}
		maxConcurrentConnectionsPtr = &maxConcurrentConnections
	}

	var downloadBytesSecRateLimitPtr *uint64
	if input["downloadBytesSecRateLimit"] != nil {
		downloadBytesSecRateLimit, err := voHelper.InterfaceToUint64(
			input["downloadBytesSecRateLimit"],
		)
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidDownloadBytesSecRateLimit"))
		}
		downloadBytesSecRateLimitPtr = &downloadBytesSecRateLimit
	}

	var uploadBytesSecRateLimitPtr *uint64
	if input["uploadBytesSecRateLimit"] != nil {
		uploadBytesSecRateLimit, err := voHelper.InterfaceToUint64(
			input["uploadBytesSecRateLimit"],
		)
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidUploadBytesSecRateLimit"))
		}
		uploadBytesSecRateLimitPtr = &uploadBytesSecRateLimit
	}

	var skipCertificateVerificationPtr *bool
	if input["skipCertificateVerification"] != nil {
		skipCertificateVerification, err := voHelper.InterfaceToBool(
			input["skipCertificateVerification"],
		)
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidSkipCertificateVerification"))
		}
		skipCertificateVerificationPtr = &skipCertificateVerification
	}

	var objectStorageProviderPtr *valueObject.ObjectStorageProvider
	if input["objectStorageProvider"] != nil {
		objectStorageProvider, err := valueObject.NewObjectStorageProvider(
			input["objectStorageProvider"],
		)
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		objectStorageProviderPtr = &objectStorageProvider
	}

	var objectStorageProviderRegionPtr *valueObject.ObjectStorageProviderRegion
	if input["objectStorageProviderRegion"] != nil {
		objectStorageProviderRegion, err := valueObject.NewObjectStorageProviderRegion(
			input["objectStorageProviderRegion"],
		)
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		objectStorageProviderRegionPtr = &objectStorageProviderRegion
	}

	var objectStorageProviderAccessKeyIdPtr *valueObject.ObjectStorageProviderAccessKeyId
	if input["objectStorageProviderAccessKeyId"] != nil {
		objectStorageProviderAccessKeyId, err := valueObject.NewObjectStorageProviderAccessKeyId(
			input["objectStorageProviderAccessKeyId"],
		)
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		objectStorageProviderAccessKeyIdPtr = &objectStorageProviderAccessKeyId
	}

	var objectStorageProviderSecretAccessKeyPtr *valueObject.ObjectStorageProviderSecretAccessKey
	if input["objectStorageProviderSecretAccessKey"] != nil {
		objectStorageProviderSecretAccessKey, err := valueObject.NewObjectStorageProviderSecretAccessKey(
			input["objectStorageProviderSecretAccessKey"],
		)
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		objectStorageProviderSecretAccessKeyPtr = &objectStorageProviderSecretAccessKey
	}

	var objectStorageEndpointUrlPtr *valueObject.Url
	if input["objectStorageEndpointUrl"] != nil {
		objectStorageEndpointUrl, err := valueObject.NewUrl(input["objectStorageEndpointUrl"])
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidObjectStorageEndpointUrl"))
		}
		objectStorageEndpointUrlPtr = &objectStorageEndpointUrl
	}

	var objectStorageBucketNamePtr *valueObject.ObjectStorageBucketName
	if input["objectStorageBucketName"] != nil {
		objectStorageBucketName, err := valueObject.NewObjectStorageBucketName(
			input["objectStorageBucketName"],
		)
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		objectStorageBucketNamePtr = &objectStorageBucketName
	}

	var remoteHostTypePtr *valueObject.BackupDestinationRemoteHostType
	if input["remoteHostType"] != nil {
		remoteHostType, err := valueObject.NewBackupDestinationRemoteHostType(
			input["remoteHostType"],
		)
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		remoteHostTypePtr = &remoteHostType
	}

	var remoteHostnamePtr *valueObject.NetworkHost
	if input["remoteHostname"] != nil {
		remoteHostname, err := valueObject.NewNetworkHost(input["remoteHostname"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		remoteHostnamePtr = &remoteHostname
	}

	var remoteHostNetworkPortPtr *valueObject.NetworkPort
	if input["remoteHostNetworkPort"] != nil {
		remoteHostNetworkPort, err := valueObject.NewNetworkPort(
			input["remoteHostNetworkPort"],
		)
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		remoteHostNetworkPortPtr = &remoteHostNetworkPort
	}

	var remoteHostUsernamePtr *valueObject.UnixUsername
	if input["remoteHostUsername"] != nil {
		remoteHostUsername, err := valueObject.NewUnixUsername(input["remoteHostUsername"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		remoteHostUsernamePtr = &remoteHostUsername
	}

	var remoteHostPasswordPtr *valueObject.Password
	if input["remoteHostPassword"] != nil {
		remoteHostPassword, err := valueObject.NewPassword(input["remoteHostPassword"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		remoteHostPasswordPtr = &remoteHostPassword
	}

	var remoteHostPrivateKeyFilePathPtr *valueObject.UnixFilePath
	if input["remoteHostPrivateKeyFilePath"] != nil {
		remoteHostPrivateKeyFilePath, err := valueObject.NewUnixFilePath(
			input["remoteHostPrivateKeyFilePath"],
		)
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidRemoteHostPrivateKeyFilePath"))
		}
		remoteHostPrivateKeyFilePathPtr = &remoteHostPrivateKeyFilePath
	}

	var remoteHostConnectionTimeoutSecsPtr *uint16
	if input["remoteHostConnectionTimeoutSecs"] != nil {
		remoteHostConnectionTimeoutSecs, err := voHelper.InterfaceToUint16(
			input["remoteHostConnectionTimeoutSecs"],
		)
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidRemoteHostConnectionTimeoutSecs"))
		}
		remoteHostConnectionTimeoutSecsPtr = &remoteHostConnectionTimeoutSecs
	}

	var remoteHostConnectionRetrySecsPtr *uint16
	if input["remoteHostConnectionRetrySecs"] != nil {
		remoteHostConnectionRetrySecs, err := voHelper.InterfaceToUint16(
			input["remoteHostConnectionRetrySecs"],
		)
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidRemoteHostConnectionRetrySecs"))
		}
		remoteHostConnectionRetrySecsPtr = &remoteHostConnectionRetrySecs
	}

	operatorAccountId := LocalOperatorAccountId
	if input["operatorAccountId"] != nil {
		operatorAccountId, err = valueObject.NewAccountId(input["operatorAccountId"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
	}

	operatorIpAddress := LocalOperatorIpAddress
	if input["operatorIpAddress"] != nil {
		operatorIpAddress, err = valueObject.NewIpAddress(input["operatorIpAddress"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
	}

	createDto := dto.CreateBackupDestination{
		AccountId:                            accountId,
		DestinationName:                      destinationName,
		DestinationDescription:               destinationDescriptionPtr,
		DestinationType:                      destinationType,
		DestinationPath:                      destinationPathPtr,
		MinLocalStorageFreePercent:           minLocalStorageFreePercentPtr,
		MaxDestinationStorageUsagePercent:    maxDestinationStorageUsagePercentPtr,
		MaxConcurrentConnections:             maxConcurrentConnectionsPtr,
		DownloadBytesSecRateLimit:            downloadBytesSecRateLimitPtr,
		UploadBytesSecRateLimit:              uploadBytesSecRateLimitPtr,
		SkipCertificateVerification:          skipCertificateVerificationPtr,
		ObjectStorageProvider:                objectStorageProviderPtr,
		ObjectStorageProviderRegion:          objectStorageProviderRegionPtr,
		ObjectStorageProviderAccessKeyId:     objectStorageProviderAccessKeyIdPtr,
		ObjectStorageProviderSecretAccessKey: objectStorageProviderSecretAccessKeyPtr,
		ObjectStorageEndpointUrl:             objectStorageEndpointUrlPtr,
		ObjectStorageBucketName:              objectStorageBucketNamePtr,
		RemoteHostType:                       remoteHostTypePtr,
		RemoteHostname:                       remoteHostnamePtr,
		RemoteHostNetworkPort:                remoteHostNetworkPortPtr,
		RemoteHostUsername:                   remoteHostUsernamePtr,
		RemoteHostPassword:                   remoteHostPasswordPtr,
		RemoteHostPrivateKeyFilePath:         remoteHostPrivateKeyFilePathPtr,
		RemoteHostConnectionTimeoutSecs:      remoteHostConnectionTimeoutSecsPtr,
		RemoteHostConnectionRetrySecs:        remoteHostConnectionRetrySecsPtr,
		OperatorAccountId:                    operatorAccountId,
		OperatorIpAddress:                    operatorIpAddress,
	}

	backupCmdRepo := backupInfra.NewBackupCmdRepo(service.persistentDbSvc)
	err = useCase.CreateBackupDestination(backupCmdRepo, service.activityRecordCmdRepo, createDto)
	if err != nil {
		return NewServiceOutput(InfraError, err.Error())
	}

	return NewServiceOutput(Created, "BackupDestinationCreated")
}

func (service *BackupService) UpdateDestination(input map[string]interface{}) ServiceOutput {
	requiredParams := []string{"destinationId", "accountId"}

	err := serviceHelper.RequiredParamsInspector(input, requiredParams)
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	destinationId, err := valueObject.NewBackupDestinationId(input["destinationId"])
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	accountId, err := valueObject.NewAccountId(input["accountId"])
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	var destinationNamePtr *valueObject.BackupDestinationName
	if input["destinationName"] != nil {
		destinationName, err := valueObject.NewBackupDestinationName(input["destinationName"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		destinationNamePtr = &destinationName
	}

	var destinationDescriptionPtr *valueObject.BackupDestinationDescription
	if input["destinationDescription"] != nil {
		destinationDescription, err := valueObject.NewBackupDestinationDescription(
			input["destinationDescription"],
		)
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		destinationDescriptionPtr = &destinationDescription
	}

	var destinationTypePtr *valueObject.BackupDestinationType
	if input["destinationType"] != nil {
		destinationType, err := valueObject.NewBackupDestinationType(input["destinationType"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		destinationTypePtr = &destinationType
	}

	var destinationPathPtr *valueObject.UnixFilePath
	if input["destinationPath"] != nil {
		destinationPath, err := valueObject.NewUnixFilePath(input["destinationPath"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		destinationPathPtr = &destinationPath
	}

	var minLocalStorageFreePercentPtr *uint8
	if input["minLocalStorageFreePercent"] != nil {
		minLocalStorageFreePercent, err := voHelper.InterfaceToUint8(
			input["minLocalStorageFreePercent"],
		)
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidMinLocalStorageFreePercent"))
		}
		minLocalStorageFreePercentPtr = &minLocalStorageFreePercent
	}

	var maxDestinationStorageUsagePercentPtr *uint8
	if input["maxDestinationStorageUsagePercent"] != nil {
		maxDestinationStorageUsagePercent, err := voHelper.InterfaceToUint8(
			input["maxDestinationStorageUsagePercent"],
		)
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidMaxDestinationStorageUsagePercent"))
		}
		maxDestinationStorageUsagePercentPtr = &maxDestinationStorageUsagePercent
	}

	var maxConcurrentConnectionsPtr *uint16
	if input["maxConcurrentConnections"] != nil {
		maxConcurrentConnections, err := voHelper.InterfaceToUint16(
			input["maxConcurrentConnections"],
		)
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidMaxConcurrentConnections"))
		}
		maxConcurrentConnectionsPtr = &maxConcurrentConnections
	}

	var downloadBytesSecRateLimitPtr *uint64
	if input["downloadBytesSecRateLimit"] != nil {
		downloadBytesSecRateLimit, err := voHelper.InterfaceToUint64(
			input["downloadBytesSecRateLimit"],
		)
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidDownloadBytesSecRateLimit"))
		}
		downloadBytesSecRateLimitPtr = &downloadBytesSecRateLimit
	}

	var uploadBytesSecRateLimitPtr *uint64
	if input["uploadBytesSecRateLimit"] != nil {
		uploadBytesSecRateLimit, err := voHelper.InterfaceToUint64(
			input["uploadBytesSecRateLimit"],
		)
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidUploadBytesSecRateLimit"))
		}
		uploadBytesSecRateLimitPtr = &uploadBytesSecRateLimit
	}

	var skipCertificateVerificationPtr *bool
	if input["skipCertificateVerification"] != nil {
		skipCertificateVerification, err := voHelper.InterfaceToBool(
			input["skipCertificateVerification"],
		)
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidSkipCertificateVerification"))
		}
		skipCertificateVerificationPtr = &skipCertificateVerification
	}

	var objectStorageProviderPtr *valueObject.ObjectStorageProvider
	if input["objectStorageProvider"] != nil {
		objectStorageProvider, err := valueObject.NewObjectStorageProvider(
			input["objectStorageProvider"],
		)
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		objectStorageProviderPtr = &objectStorageProvider
	}

	var objectStorageProviderRegionPtr *valueObject.ObjectStorageProviderRegion
	if input["objectStorageProviderRegion"] != nil {
		objectStorageProviderRegion, err := valueObject.NewObjectStorageProviderRegion(
			input["objectStorageProviderRegion"],
		)
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		objectStorageProviderRegionPtr = &objectStorageProviderRegion
	}

	var objectStorageProviderAccessKeyIdPtr *valueObject.ObjectStorageProviderAccessKeyId
	if input["objectStorageProviderAccessKeyId"] != nil {
		objectStorageProviderAccessKeyId, err := valueObject.NewObjectStorageProviderAccessKeyId(
			input["objectStorageProviderAccessKeyId"],
		)
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		objectStorageProviderAccessKeyIdPtr = &objectStorageProviderAccessKeyId
	}

	var objectStorageProviderSecretAccessKeyPtr *valueObject.ObjectStorageProviderSecretAccessKey
	if input["objectStorageProviderSecretAccessKey"] != nil {
		objectStorageProviderSecretAccessKey, err := valueObject.NewObjectStorageProviderSecretAccessKey(
			input["objectStorageProviderSecretAccessKey"],
		)
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		objectStorageProviderSecretAccessKeyPtr = &objectStorageProviderSecretAccessKey
	}

	var objectStorageEndpointUrlPtr *valueObject.Url
	if input["objectStorageEndpointUrl"] != nil {
		objectStorageEndpointUrl, err := valueObject.NewUrl(input["objectStorageEndpointUrl"])
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidObjectStorageEndpointUrl"))
		}
		objectStorageEndpointUrlPtr = &objectStorageEndpointUrl
	}

	var objectStorageBucketNamePtr *valueObject.ObjectStorageBucketName
	if input["objectStorageBucketName"] != nil {
		objectStorageBucketName, err := valueObject.NewObjectStorageBucketName(
			input["objectStorageBucketName"],
		)
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		objectStorageBucketNamePtr = &objectStorageBucketName
	}

	var remoteHostTypePtr *valueObject.BackupDestinationRemoteHostType
	if input["remoteHostType"] != nil {
		remoteHostType, err := valueObject.NewBackupDestinationRemoteHostType(
			input["remoteHostType"],
		)
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		remoteHostTypePtr = &remoteHostType
	}

	var remoteHostnamePtr *valueObject.NetworkHost
	if input["remoteHostname"] != nil {
		remoteHostname, err := valueObject.NewNetworkHost(input["remoteHostname"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		remoteHostnamePtr = &remoteHostname
	}

	var remoteHostNetworkPortPtr *valueObject.NetworkPort
	if input["remoteHostNetworkPort"] != nil {
		remoteHostNetworkPort, err := valueObject.NewNetworkPort(
			input["remoteHostNetworkPort"],
		)
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		remoteHostNetworkPortPtr = &remoteHostNetworkPort
	}

	var remoteHostUsernamePtr *valueObject.UnixUsername
	if input["remoteHostUsername"] != nil {
		remoteHostUsername, err := valueObject.NewUnixUsername(input["remoteHostUsername"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		remoteHostUsernamePtr = &remoteHostUsername
	}

	var remoteHostPasswordPtr *valueObject.Password
	if input["remoteHostPassword"] != nil {
		remoteHostPassword, err := valueObject.NewPassword(input["remoteHostPassword"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		remoteHostPasswordPtr = &remoteHostPassword
	}

	var remoteHostPrivateKeyFilePathPtr *valueObject.UnixFilePath
	if input["remoteHostPrivateKeyFilePath"] != nil {
		remoteHostPrivateKeyFilePath, err := valueObject.NewUnixFilePath(
			input["remoteHostPrivateKeyFilePath"],
		)
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidRemoteHostPrivateKeyFilePath"))
		}
		remoteHostPrivateKeyFilePathPtr = &remoteHostPrivateKeyFilePath
	}

	var remoteHostConnectionTimeoutSecsPtr *uint16
	if input["remoteHostConnectionTimeoutSecs"] != nil {
		remoteHostConnectionTimeoutSecs, err := voHelper.InterfaceToUint16(
			input["remoteHostConnectionTimeoutSecs"],
		)
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidRemoteHostConnectionTimeoutSecs"))
		}
		remoteHostConnectionTimeoutSecsPtr = &remoteHostConnectionTimeoutSecs
	}

	var remoteHostConnectionRetrySecsPtr *uint16
	if input["remoteHostConnectionRetrySecs"] != nil {
		remoteHostConnectionRetrySecs, err := voHelper.InterfaceToUint16(
			input["remoteHostConnectionRetrySecs"],
		)
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidRemoteHostConnectionRetrySecs"))
		}
		remoteHostConnectionRetrySecsPtr = &remoteHostConnectionRetrySecs
	}

	operatorAccountId := LocalOperatorAccountId
	if input["operatorAccountId"] != nil {
		operatorAccountId, err = valueObject.NewAccountId(input["operatorAccountId"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
	}

	operatorIpAddress := LocalOperatorIpAddress
	if input["operatorIpAddress"] != nil {
		operatorIpAddress, err = valueObject.NewIpAddress(input["operatorIpAddress"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
	}

	updateDto := dto.UpdateBackupDestination{
		DestinationId:                        destinationId,
		AccountId:                            accountId,
		DestinationName:                      destinationNamePtr,
		DestinationDescription:               destinationDescriptionPtr,
		DestinationType:                      destinationTypePtr,
		DestinationPath:                      destinationPathPtr,
		MinLocalStorageFreePercent:           minLocalStorageFreePercentPtr,
		MaxDestinationStorageUsagePercent:    maxDestinationStorageUsagePercentPtr,
		MaxConcurrentConnections:             maxConcurrentConnectionsPtr,
		DownloadBytesSecRateLimit:            downloadBytesSecRateLimitPtr,
		UploadBytesSecRateLimit:              uploadBytesSecRateLimitPtr,
		SkipCertificateVerification:          skipCertificateVerificationPtr,
		ObjectStorageProvider:                objectStorageProviderPtr,
		ObjectStorageProviderRegion:          objectStorageProviderRegionPtr,
		ObjectStorageProviderAccessKeyId:     objectStorageProviderAccessKeyIdPtr,
		ObjectStorageProviderSecretAccessKey: objectStorageProviderSecretAccessKeyPtr,
		ObjectStorageEndpointUrl:             objectStorageEndpointUrlPtr,
		ObjectStorageBucketName:              objectStorageBucketNamePtr,
		RemoteHostType:                       remoteHostTypePtr,
		RemoteHostname:                       remoteHostnamePtr,
		RemoteHostNetworkPort:                remoteHostNetworkPortPtr,
		RemoteHostUsername:                   remoteHostUsernamePtr,
		RemoteHostPassword:                   remoteHostPasswordPtr,
		RemoteHostPrivateKeyFilePath:         remoteHostPrivateKeyFilePathPtr,
		RemoteHostConnectionTimeoutSecs:      remoteHostConnectionTimeoutSecsPtr,
		RemoteHostConnectionRetrySecs:        remoteHostConnectionRetrySecsPtr,
		OperatorAccountId:                    operatorAccountId,
		OperatorIpAddress:                    operatorIpAddress,
	}

	backupQueryRepo := backupInfra.NewBackupQueryRepo(service.persistentDbSvc)
	backupCmdRepo := backupInfra.NewBackupCmdRepo(service.persistentDbSvc)
	err = useCase.UpdateBackupDestination(
		backupQueryRepo, backupCmdRepo, service.activityRecordCmdRepo, updateDto,
	)
	if err != nil {
		return NewServiceOutput(InfraError, err.Error())
	}

	return NewServiceOutput(Success, "BackupDestinationUpdated")
}

func (service *BackupService) DeleteDestination(input map[string]interface{}) ServiceOutput {
	requiredParams := []string{"destinationId", "accountId"}

	err := serviceHelper.RequiredParamsInspector(input, requiredParams)
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	destinationId, err := valueObject.NewBackupDestinationId(input["destinationId"])
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	accountId, err := valueObject.NewAccountId(input["accountId"])
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	operatorAccountId := LocalOperatorAccountId
	if input["operatorAccountId"] != nil {
		operatorAccountId, err = valueObject.NewAccountId(input["operatorAccountId"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
	}

	operatorIpAddress := LocalOperatorIpAddress
	if input["operatorIpAddress"] != nil {
		operatorIpAddress, err = valueObject.NewIpAddress(input["operatorIpAddress"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
	}

	deleteDto := dto.NewDeleteBackupDestination(
		destinationId, accountId, operatorAccountId, operatorIpAddress,
	)

	backupQueryRepo := backupInfra.NewBackupQueryRepo(service.persistentDbSvc)
	backupCmdRepo := backupInfra.NewBackupCmdRepo(service.persistentDbSvc)
	err = useCase.DeleteBackupDestination(
		backupQueryRepo, backupCmdRepo, service.activityRecordCmdRepo, deleteDto,
	)
	if err != nil {
		return NewServiceOutput(InfraError, err.Error())
	}

	return NewServiceOutput(Success, "BackupDestinationDeleted")
}

func (service *BackupService) ReadJob(input map[string]interface{}) ServiceOutput {
	var jobIdPtr *valueObject.BackupJobId
	if input["jobId"] != nil {
		jobId, err := valueObject.NewBackupJobId(input["jobId"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		jobIdPtr = &jobId
	}

	var jobStatusPtr *bool
	if input["jobStatus"] != nil {
		jobStatus, err := voHelper.InterfaceToBool(input["jobStatus"])
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidJobStatus"))
		}
		jobStatusPtr = &jobStatus
	}

	var accountIdPtr *valueObject.AccountId
	if input["accountId"] != nil {
		accountId, err := valueObject.NewAccountId(input["accountId"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		accountIdPtr = &accountId
	}

	var destinationIdPtr *valueObject.BackupDestinationId
	if input["destinationId"] != nil {
		destinationId, err := valueObject.NewBackupDestinationId(input["destinationId"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		destinationIdPtr = &destinationId
	}

	var retentionStrategyPtr *valueObject.BackupRetentionStrategy
	if input["retentionStrategy"] != nil {
		retentionStrategy, err := valueObject.NewBackupRetentionStrategy(input["retentionStrategy"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		retentionStrategyPtr = &retentionStrategy
	}

	var archiveCompressionFormatPtr *valueObject.CompressionFormat
	if input["archiveCompressionFormat"] != nil {
		archiveCompressionFormat, err := valueObject.NewCompressionFormat(input["archiveCompressionFormat"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		archiveCompressionFormatPtr = &archiveCompressionFormat
	}

	var lastRunStatusPtr *valueObject.BackupTaskStatus
	if input["lastRunStatus"] != nil {
		lastRunStatus, err := valueObject.NewBackupTaskStatus(input["lastRunStatus"])
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidLastRunStatus"))
		}
		lastRunStatusPtr = &lastRunStatus
	}

	timeParamNames := []string{
		"lastRunBeforeAt", "lastRunAfterAt", "nextRunBeforeAt", "nextRunAfterAt",
		"createdBeforeAt", "createdAfterAt",
	}
	timeParamPtrs := serviceHelper.TimeParamsParser(timeParamNames, input)

	requestPagination, err := serviceHelper.PaginationParser(
		input, useCase.BackupJobsDefaultPagination,
	)
	if err != nil {
		return NewServiceOutput(UserError, err)
	}

	readDto := dto.ReadBackupJobsRequest{
		Pagination:               requestPagination,
		JobId:                    jobIdPtr,
		JobStatus:                jobStatusPtr,
		AccountId:                accountIdPtr,
		DestinationId:            destinationIdPtr,
		RetentionStrategy:        retentionStrategyPtr,
		ArchiveCompressionFormat: archiveCompressionFormatPtr,
		LastRunStatus:            lastRunStatusPtr,
		LastRunBeforeAt:          timeParamPtrs["lastRunBeforeAt"],
		LastRunAfterAt:           timeParamPtrs["lastRunAfterAt"],
		NextRunBeforeAt:          timeParamPtrs["nextRunBeforeAt"],
		NextRunAfterAt:           timeParamPtrs["nextRunAfterAt"],
		CreatedBeforeAt:          timeParamPtrs["createdBeforeAt"],
		CreatedAfterAt:           timeParamPtrs["createdAfterAt"],
	}

	responseDto, err := useCase.ReadBackupJobs(service.backupQueryRepo, readDto)
	if err != nil {
		return NewServiceOutput(InfraError, err.Error())
	}

	return NewServiceOutput(Success, responseDto)
}

func (service *BackupService) CreateJob(input map[string]interface{}) ServiceOutput {
	requiredParams := []string{"accountId", "destinationIds", "backupSchedule"}

	err := serviceHelper.RequiredParamsInspector(input, requiredParams)
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	accountId, err := valueObject.NewAccountId(input["accountId"])
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	var jobDescriptionPtr *valueObject.BackupJobDescription
	if input["jobDescription"] != nil {
		jobDescription, err := valueObject.NewBackupJobDescription(input["jobDescription"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		jobDescriptionPtr = &jobDescription
	}

	destinationIds, assertOk := input["destinationIds"].([]valueObject.BackupDestinationId)
	if !assertOk {
		return NewServiceOutput(UserError, errors.New("InvalidDestinationIds"))
	}

	backupSchedule, err := valueObject.NewCronSchedule(input["backupSchedule"])
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	var retentionStrategyPtr *valueObject.BackupRetentionStrategy
	if input["retentionStrategy"] != nil {
		retentionStrategy, err := valueObject.NewBackupRetentionStrategy(
			input["retentionStrategy"],
		)
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		retentionStrategyPtr = &retentionStrategy
	}

	var archiveCompressionFormatPtr *valueObject.CompressionFormat
	if input["archiveCompressionFormat"] != nil {
		archiveCompressionFormat, err := valueObject.NewCompressionFormat(
			input["archiveCompressionFormat"],
		)
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		archiveCompressionFormatPtr = &archiveCompressionFormat
	}

	var timeoutSecsPtr *uint64
	if input["timeoutSecs"] != nil {
		timeoutSecs, err := voHelper.InterfaceToUint64(input["timeoutSecs"])
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidTimeoutSecs"))
		}
		timeoutSecsPtr = &timeoutSecs
	}

	var maxTaskRetentionCountPtr *uint16
	if input["maxTaskRetentionCount"] != nil {
		maxTaskRetentionCount, err := voHelper.InterfaceToUint16(
			input["maxTaskRetentionCount"],
		)
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidMaxTaskRetentionCount"))
		}
		maxTaskRetentionCountPtr = &maxTaskRetentionCount
	}

	var maxTaskRetentionDaysPtr *uint16
	if input["maxTaskRetentionDays"] != nil {
		maxTaskRetentionDays, err := voHelper.InterfaceToUint16(input["maxTaskRetentionDays"])
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidMaxTaskRetentionDays"))
		}
		maxTaskRetentionDaysPtr = &maxTaskRetentionDays
	}

	var maxConcurrentCpuCoresPtr *uint16
	if input["maxConcurrentCpuCores"] != nil {
		maxConcurrentCpuCores, err := voHelper.InterfaceToUint16(input["maxConcurrentCpuCores"])
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidMaxConcurrentCpuCores"))
		}
		maxConcurrentCpuCoresPtr = &maxConcurrentCpuCores
	}

	containerAccountIds := []valueObject.AccountId{}
	if input["containerAccountIds"] != nil {
		containerAccountIds, assertOk = input["containerAccountIds"].([]valueObject.AccountId)
		if !assertOk {
			return NewServiceOutput(UserError, errors.New("InvalidContainerAccountIds"))
		}
	}

	containerIds := []valueObject.ContainerId{}
	if input["containerIds"] != nil {
		containerIds, assertOk = input["containerIds"].([]valueObject.ContainerId)
		if !assertOk {
			return NewServiceOutput(UserError, errors.New("InvalidContainerIds"))
		}
	}

	ignoreContainerAccountIds := []valueObject.AccountId{}
	if input["ignoreContainerAccountIds"] != nil {
		ignoreContainerAccountIds, assertOk = input["ignoreContainerAccountIds"].([]valueObject.AccountId)
		if !assertOk {
			return NewServiceOutput(UserError, errors.New("InvalidIgnoreContainerAccountIds"))
		}
	}

	ignoreContainerIds := []valueObject.ContainerId{}
	if input["ignoreContainerIds"] != nil {
		ignoreContainerIds, assertOk = input["ignoreContainerIds"].([]valueObject.ContainerId)
		if !assertOk {
			return NewServiceOutput(UserError, errors.New("InvalidIgnoreContainerIds"))
		}
	}

	operatorAccountId := LocalOperatorAccountId
	if input["operatorAccountId"] != nil {
		operatorAccountId, err = valueObject.NewAccountId(input["operatorAccountId"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
	}

	operatorIpAddress := LocalOperatorIpAddress
	if input["operatorIpAddress"] != nil {
		operatorIpAddress, err = valueObject.NewIpAddress(input["operatorIpAddress"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
	}

	createDto := dto.NewCreateBackupJob(
		accountId, jobDescriptionPtr, destinationIds, retentionStrategyPtr, backupSchedule,
		archiveCompressionFormatPtr, timeoutSecsPtr, maxTaskRetentionCountPtr,
		maxTaskRetentionDaysPtr, maxConcurrentCpuCoresPtr, containerAccountIds,
		containerIds, ignoreContainerAccountIds, ignoreContainerIds,
		operatorAccountId, operatorIpAddress,
	)

	backupCmdRepo := backupInfra.NewBackupCmdRepo(service.persistentDbSvc)
	err = useCase.CreateBackupJob(backupCmdRepo, service.activityRecordCmdRepo, createDto)
	if err != nil {
		return NewServiceOutput(InfraError, err.Error())
	}

	return NewServiceOutput(Created, "BackupJobCreated")
}

func (service *BackupService) UpdateJob(input map[string]interface{}) ServiceOutput {
	requiredParams := []string{"jobId", "accountId"}

	err := serviceHelper.RequiredParamsInspector(input, requiredParams)
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	jobId, err := valueObject.NewBackupJobId(input["jobId"])
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	accountId, err := valueObject.NewAccountId(input["accountId"])
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	var jobStatusPtr *bool
	if input["jobStatus"] != nil {
		jobStatus, err := voHelper.InterfaceToBool(input["jobStatus"])
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidJobStatus"))
		}
		jobStatusPtr = &jobStatus
	}

	var jobDescriptionPtr *valueObject.BackupJobDescription
	if input["jobDescription"] != nil {
		jobDescription, err := valueObject.NewBackupJobDescription(input["jobDescription"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		jobDescriptionPtr = &jobDescription
	}

	var destinationIds []valueObject.BackupDestinationId
	var assertOk bool
	if input["destinationIds"] != nil {
		destinationIds, assertOk = input["destinationIds"].([]valueObject.BackupDestinationId)
		if !assertOk {
			return NewServiceOutput(UserError, errors.New("InvalidDestinationIds"))
		}
	}

	var backupSchedulePtr *valueObject.CronSchedule
	if input["backupSchedule"] != nil {
		backupSchedule, err := valueObject.NewCronSchedule(input["backupSchedule"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		backupSchedulePtr = &backupSchedule
	}

	var timeoutSecsPtr *uint64
	if input["timeoutSecs"] != nil {
		timeoutSecs, err := voHelper.InterfaceToUint64(input["timeoutSecs"])
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidTimeoutSecs"))
		}
		timeoutSecsPtr = &timeoutSecs
	}

	var maxTaskRetentionCountPtr *uint16
	if input["maxTaskRetentionCount"] != nil {
		maxTaskRetentionCount, err := voHelper.InterfaceToUint16(
			input["maxTaskRetentionCount"],
		)
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidMaxTaskRetentionCount"))
		}
		maxTaskRetentionCountPtr = &maxTaskRetentionCount
	}

	var maxTaskRetentionDaysPtr *uint16
	if input["maxTaskRetentionDays"] != nil {
		maxTaskRetentionDays, err := voHelper.InterfaceToUint16(input["maxTaskRetentionDays"])
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidMaxTaskRetentionDays"))
		}
		maxTaskRetentionDaysPtr = &maxTaskRetentionDays
	}

	var maxConcurrentCpuCoresPtr *uint16
	if input["maxConcurrentCpuCores"] != nil {
		maxConcurrentCpuCores, err := voHelper.InterfaceToUint16(input["maxConcurrentCpuCores"])
		if err != nil {
			return NewServiceOutput(UserError, errors.New("InvalidMaxConcurrentCpuCores"))
		}
		maxConcurrentCpuCoresPtr = &maxConcurrentCpuCores
	}

	var containerAccountIds []valueObject.AccountId
	if input["containerAccountIds"] != nil {
		containerAccountIds, assertOk = input["containerAccountIds"].([]valueObject.AccountId)
		if !assertOk {
			return NewServiceOutput(UserError, errors.New("InvalidContainerAccountIds"))
		}
	}

	var containerIds []valueObject.ContainerId
	if input["containerIds"] != nil {
		containerIds, assertOk = input["containerIds"].([]valueObject.ContainerId)
		if !assertOk {
			return NewServiceOutput(UserError, errors.New("InvalidContainerIds"))
		}
	}

	var ignoreContainerAccountIds []valueObject.AccountId
	if input["ignoreContainerAccountIds"] != nil {
		ignoreContainerAccountIds, assertOk = input["ignoreContainerAccountIds"].([]valueObject.AccountId)
		if !assertOk {
			return NewServiceOutput(UserError, errors.New("InvalidIgnoreContainerAccountIds"))
		}
	}

	var ignoreContainerIds []valueObject.ContainerId
	if input["ignoreContainerIds"] != nil {
		ignoreContainerIds, assertOk = input["ignoreContainerIds"].([]valueObject.ContainerId)
		if !assertOk {
			return NewServiceOutput(UserError, errors.New("InvalidIgnoreContainerIds"))
		}
	}

	operatorAccountId := LocalOperatorAccountId
	if input["operatorAccountId"] != nil {
		operatorAccountId, err = valueObject.NewAccountId(input["operatorAccountId"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
	}

	operatorIpAddress := LocalOperatorIpAddress
	if input["operatorIpAddress"] != nil {
		operatorIpAddress, err = valueObject.NewIpAddress(input["operatorIpAddress"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
	}

	updateDto := dto.UpdateBackupJob{
		JobId:                     jobId,
		AccountId:                 accountId,
		JobStatus:                 jobStatusPtr,
		JobDescription:            jobDescriptionPtr,
		DestinationIds:            destinationIds,
		BackupSchedule:            backupSchedulePtr,
		TimeoutSecs:               timeoutSecsPtr,
		MaxTaskRetentionCount:     maxTaskRetentionCountPtr,
		MaxTaskRetentionDays:      maxTaskRetentionDaysPtr,
		MaxConcurrentCpuCores:     maxConcurrentCpuCoresPtr,
		ContainerAccountIds:       containerAccountIds,
		ContainerIds:              containerIds,
		IgnoreContainerAccountIds: ignoreContainerAccountIds,
		IgnoreContainerIds:        ignoreContainerIds,
		OperatorAccountId:         operatorAccountId,
		OperatorIpAddress:         operatorIpAddress,
	}

	backupQueryRepo := backupInfra.NewBackupQueryRepo(service.persistentDbSvc)
	backupCmdRepo := backupInfra.NewBackupCmdRepo(service.persistentDbSvc)
	err = useCase.UpdateBackupJob(
		backupQueryRepo, backupCmdRepo, service.activityRecordCmdRepo, updateDto,
	)
	if err != nil {
		return NewServiceOutput(InfraError, err.Error())
	}

	return NewServiceOutput(Success, "BackupJobUpdated")
}

func (service *BackupService) ReadTask(input map[string]interface{}) ServiceOutput {
	var taskIdPtr *valueObject.BackupTaskId
	if input["taskId"] != nil {
		taskId, err := valueObject.NewBackupTaskId(input["taskId"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		taskIdPtr = &taskId
	}

	var accountIdPtr *valueObject.AccountId
	if input["accountId"] != nil {
		accountId, err := valueObject.NewAccountId(input["accountId"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		accountIdPtr = &accountId
	}

	var jobIdPtr *valueObject.BackupJobId
	if input["jobId"] != nil {
		jobId, err := valueObject.NewBackupJobId(input["jobId"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		jobIdPtr = &jobId
	}

	var destinationIdPtr *valueObject.BackupDestinationId
	if input["destinationId"] != nil {
		destinationId, err := valueObject.NewBackupDestinationId(input["destinationId"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		destinationIdPtr = &destinationId
	}

	var taskStatusPtr *valueObject.BackupTaskStatus
	if input["taskStatus"] != nil {
		taskStatus, err := valueObject.NewBackupTaskStatus(input["taskStatus"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		taskStatusPtr = &taskStatus
	}

	var retentionStrategyPtr *valueObject.BackupRetentionStrategy
	if input["retentionStrategy"] != nil {
		retentionStrategy, err := valueObject.NewBackupRetentionStrategy(input["retentionStrategy"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		retentionStrategyPtr = &retentionStrategy
	}

	var containerIdPtr *valueObject.ContainerId
	if input["containerId"] != nil {
		containerId, err := valueObject.NewContainerId(input["containerId"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		containerIdPtr = &containerId
	}

	timeParamNames := []string{
		"startedBeforeAt", "startedAfterAt", "finishedBeforeAt", "finishedAfterAt",
		"createdBeforeAt", "createdAfterAt",
	}
	timeParamPtrs := serviceHelper.TimeParamsParser(timeParamNames, input)

	requestPagination, err := serviceHelper.PaginationParser(
		input, useCase.BackupTasksDefaultPagination,
	)
	if err != nil {
		return NewServiceOutput(UserError, err)
	}

	readDto := dto.ReadBackupTasksRequest{
		Pagination:        requestPagination,
		TaskId:            taskIdPtr,
		AccountId:         accountIdPtr,
		JobId:             jobIdPtr,
		DestinationId:     destinationIdPtr,
		TaskStatus:        taskStatusPtr,
		RetentionStrategy: retentionStrategyPtr,
		ContainerId:       containerIdPtr,
		StartedBeforeAt:   timeParamPtrs["startedBeforeAt"],
		StartedAfterAt:    timeParamPtrs["startedAfterAt"],
		FinishedBeforeAt:  timeParamPtrs["finishedBeforeAt"],
		FinishedAfterAt:   timeParamPtrs["finishedAfterAt"],
		CreatedBeforeAt:   timeParamPtrs["createdBeforeAt"],
		CreatedAfterAt:    timeParamPtrs["createdAfterAt"],
	}

	responseDto, err := useCase.ReadBackupTasks(service.backupQueryRepo, readDto)
	if err != nil {
		return NewServiceOutput(InfraError, err.Error())
	}

	return NewServiceOutput(Success, responseDto)
}
