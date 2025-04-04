package presenter

import (
	"log/slog"
	"maps"
	"net/http"

	"github.com/goinfinite/ez/src/domain/dto"
	"github.com/goinfinite/ez/src/domain/entity"
	"github.com/goinfinite/ez/src/domain/valueObject"
	"github.com/goinfinite/ez/src/infra/db"
	"github.com/goinfinite/ez/src/presentation/service"
	uiHelper "github.com/goinfinite/ez/src/presentation/ui/helper"
	pageBackup "github.com/goinfinite/ez/src/presentation/ui/page/backup"
	presenterHelper "github.com/goinfinite/ez/src/presentation/ui/presenter/helper"
	"github.com/labstack/echo/v4"
)

type BackupPresenter struct {
	persistentDbSvc *db.PersistentDatabaseService
	transientDbSvc  *db.TransientDatabaseService
	trailDbSvc      *db.TrailDatabaseService
}

func NewBackupPresenter(
	persistentDbSvc *db.PersistentDatabaseService,
	transientDbSvc *db.TransientDatabaseService,
	trailDbSvc *db.TrailDatabaseService,
) *BackupPresenter {
	return &BackupPresenter{
		persistentDbSvc: persistentDbSvc,
		transientDbSvc:  transientDbSvc,
		trailDbSvc:      trailDbSvc,
	}
}

func (presenter *BackupPresenter) ReadTasks(
	echoContext echo.Context,
	backupService *service.BackupService,
) (readRequestDto dto.ReadBackupTasksRequest, readResponseDto dto.ReadBackupTasksResponse) {
	entityNamePrefix := "backupTasks"
	paginationMap := uiHelper.PaginationParser(echoContext, entityNamePrefix, "id")
	if echoContext.QueryParam(entityNamePrefix+"SortDirection") == "" {
		paginationMap["sortDirection"] = valueObject.PaginationSortDirectionDesc.String()
	}
	requestParamsMap := uiHelper.ReadRequestParser(
		echoContext, entityNamePrefix, dto.ReadBackupTasksRequest{},
	)
	serviceRequestBody := paginationMap
	maps.Copy(serviceRequestBody, requestParamsMap)

	readRequestDto, err := backupService.ReadTaskRequestFactory(serviceRequestBody)
	if err != nil {
		slog.Debug("ReadTaskRequestFactoryFailure", slog.Any("error", err))
		return readRequestDto, readResponseDto
	}

	readBackupTasksServiceOutput := backupService.ReadTask(serviceRequestBody)
	if readBackupTasksServiceOutput.Status != service.Success {
		slog.Debug("ReadBackupTasksFailure", slog.Any("serviceOutput", readBackupTasksServiceOutput))
		return readRequestDto, readResponseDto
	}

	var assertOk bool
	readResponseDto, assertOk = readBackupTasksServiceOutput.Body.(dto.ReadBackupTasksResponse)
	if !assertOk {
		slog.Debug("AssertBackupTasksResponseFailure")
		return readRequestDto, readResponseDto
	}

	return readRequestDto, readResponseDto
}

func (presenter *BackupPresenter) ReadTaskArchives(
	echoContext echo.Context,
	backupService *service.BackupService,
) (readRequestDto dto.ReadBackupTaskArchivesRequest, readResponseDto dto.ReadBackupTaskArchivesResponse) {
	entityNamePrefix := "backupTaskArchives"
	paginationMap := uiHelper.PaginationParser(echoContext, entityNamePrefix, "createdAt")
	if echoContext.QueryParam(entityNamePrefix+"SortDirection") == "" {
		paginationMap["sortDirection"] = valueObject.PaginationSortDirectionDesc.String()
	}
	requestParamsMap := uiHelper.ReadRequestParser(
		echoContext, entityNamePrefix, dto.ReadBackupTaskArchivesRequest{},
	)
	serviceRequestBody := paginationMap
	maps.Copy(serviceRequestBody, requestParamsMap)

	readRequestDto, err := backupService.ReadTaskArchiveRequestFactory(serviceRequestBody)
	if err != nil {
		slog.Debug("ReadTaskArchiveRequestFactoryFailure", slog.Any("error", err))
		return readRequestDto, readResponseDto
	}

	readBackupTaskArchivesServiceOutput := backupService.ReadTaskArchive(
		serviceRequestBody, &echoContext.Request().Host,
	)
	if readBackupTaskArchivesServiceOutput.Status != service.Success {
		slog.Debug(
			"ReadBackupTaskArchivesFailure",
			slog.Any("serviceOutput", readBackupTaskArchivesServiceOutput),
		)
		return readRequestDto, readResponseDto
	}

	var assertOk bool
	readResponseDto, assertOk = readBackupTaskArchivesServiceOutput.Body.(dto.ReadBackupTaskArchivesResponse)
	if !assertOk {
		slog.Debug("AssertBackupTaskArchivesResponseFailure")
		return readRequestDto, readResponseDto
	}

	return readRequestDto, readResponseDto
}

func (presenter *BackupPresenter) ReadJobs(
	echoContext echo.Context,
	backupService *service.BackupService,
) (readRequestDto dto.ReadBackupJobsRequest, readResponseDto dto.ReadBackupJobsResponse) {
	paginationMap := uiHelper.PaginationParser(echoContext, "backupJobs", "id")
	requestParamsMap := uiHelper.ReadRequestParser(
		echoContext, "backupJobs", dto.ReadBackupJobsRequest{},
	)
	serviceRequestBody := paginationMap
	maps.Copy(serviceRequestBody, requestParamsMap)

	readRequestDto, err := backupService.ReadJobRequestFactory(serviceRequestBody)
	if err != nil {
		slog.Debug("ReadJobRequestFactoryFailure", slog.Any("error", err))
		return readRequestDto, readResponseDto
	}

	readBackupJobsServiceOutput := backupService.ReadJob(serviceRequestBody)
	if readBackupJobsServiceOutput.Status != service.Success {
		slog.Debug(
			"ReadBackupJobsFailure",
			slog.Any("serviceOutput", readBackupJobsServiceOutput),
		)
		return readRequestDto, readResponseDto
	}

	var assertOk bool
	readResponseDto, assertOk = readBackupJobsServiceOutput.Body.(dto.ReadBackupJobsResponse)
	if !assertOk {
		slog.Debug("AssertBackupJobsResponseFailure")
		return readRequestDto, readResponseDto
	}

	return readRequestDto, readResponseDto
}

func (presenter *BackupPresenter) ReadDestinations(
	echoContext echo.Context,
	backupService *service.BackupService,
) (
	readRequestDto dto.ReadBackupDestinationsRequest,
	readResponseDto pageBackup.BackupDestinationModifiedResponseDto,
) {
	paginationMap := uiHelper.PaginationParser(echoContext, "backupDestinations", "id")
	requestParamsMap := uiHelper.ReadRequestParser(
		echoContext, "backupDestinations", dto.ReadBackupDestinationsRequest{},
	)
	serviceRequestBody := paginationMap
	maps.Copy(serviceRequestBody, requestParamsMap)

	readRequestDto, err := backupService.ReadDestinationRequestFactory(serviceRequestBody)
	if err != nil {
		slog.Debug("ReadDestinationRequestFactoryFailure", slog.Any("error", err))
		return readRequestDto, readResponseDto
	}

	readBackupDestinationsServiceOutput := backupService.ReadDestination(serviceRequestBody)
	if readBackupDestinationsServiceOutput.Status != service.Success {
		slog.Debug(
			"ReadBackupDestinationsFailure",
			slog.Any("serviceOutput", readBackupDestinationsServiceOutput),
		)
		return readRequestDto, readResponseDto
	}

	var assertOk bool
	originalDestinationsResponseDto, assertOk := readBackupDestinationsServiceOutput.Body.(dto.ReadBackupDestinationsResponse)
	if !assertOk {
		slog.Debug("AssertBackupDestinationsResponseFailure")
		return readRequestDto, readResponseDto
	}

	for _, iDestinationEntity := range originalDestinationsResponseDto.Destinations {
		destinationUnifiedEntity := pageBackup.BackupDestinationUnifiedEntity{}

		switch destinationEntity := iDestinationEntity.(type) {
		case entity.BackupDestinationLocal:
			destinationUnifiedEntity = pageBackup.BackupDestinationUnifiedEntity{
				BackupDestinationBase:  destinationEntity.BackupDestinationBase,
				BackupDestinationLocal: destinationEntity,
			}
		case entity.BackupDestinationObjectStorage:
			destinationUnifiedEntity = pageBackup.BackupDestinationUnifiedEntity{
				BackupDestinationBase:          destinationEntity.BackupDestinationBase,
				BackupDestinationRemoteBase:    destinationEntity.BackupDestinationRemoteBase,
				BackupDestinationObjectStorage: destinationEntity,
			}
		case entity.BackupDestinationRemoteHost:
			destinationUnifiedEntity = pageBackup.BackupDestinationUnifiedEntity{
				BackupDestinationBase:       destinationEntity.BackupDestinationBase,
				BackupDestinationRemoteBase: destinationEntity.BackupDestinationRemoteBase,
				BackupDestinationRemoteHost: destinationEntity,
			}
		}

		readResponseDto.Destinations = append(
			readResponseDto.Destinations, destinationUnifiedEntity,
		)
	}

	readResponseDto.Pagination = originalDestinationsResponseDto.Pagination

	return readRequestDto, readResponseDto
}

func (presenter *BackupPresenter) Handler(c echo.Context) (err error) {
	backupService := service.NewBackupService(
		presenter.persistentDbSvc, presenter.trailDbSvc,
	)

	tasksReadRequestDto, tasksReadResponseDto := presenter.ReadTasks(c, backupService)

	archivesReadRequestDto, archivesReadResponseDto := presenter.ReadTaskArchives(c, backupService)

	jobsReadRequestDto, jobsReadResponseDto := presenter.ReadJobs(c, backupService)
	accountSelectPairs := presenterHelper.ReadAccountSelectLabelValuePairs(
		presenter.persistentDbSvc, presenter.trailDbSvc,
	)
	containerSelectPairs := presenterHelper.ReadContainerSelectLabelValuePairs(
		presenter.persistentDbSvc, presenter.trailDbSvc,
	)
	backupDestinationSelectPairs := presenterHelper.ReadBackupDestinationSelectLabelValuePairs(
		presenter.persistentDbSvc, presenter.trailDbSvc,
	)
	createJobModalDto := pageBackup.CreateBackupJobModalDto{
		AccountSelectLabelValuePairs:           accountSelectPairs,
		ContainerSelectLabelValuePairs:         containerSelectPairs,
		BackupDestinationSelectLabelValuePairs: backupDestinationSelectPairs,
	}

	destinationsReadRequestDto, destinationsReadResponseDto := presenter.ReadDestinations(c, backupService)
	createDestinationModalDto := pageBackup.CreateBackupDestinationModalDto{
		AccountSelectLabelValuePairs: accountSelectPairs,
	}

	pageContent := pageBackup.BackupIndex(
		tasksReadRequestDto, tasksReadResponseDto,
		archivesReadRequestDto, archivesReadResponseDto,
		jobsReadRequestDto, jobsReadResponseDto, createJobModalDto,
		destinationsReadRequestDto, destinationsReadResponseDto, createDestinationModalDto,
	)

	return presenterHelper.Render(c, pageContent, http.StatusOK)
}
