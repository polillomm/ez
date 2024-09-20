package useCase

import (
	"errors"
	"log/slog"

	"github.com/speedianet/control/src/domain/dto"
	"github.com/speedianet/control/src/domain/repository"
)

func CreateContainerSnapshotImage(
	containerImageQueryRepo repository.ContainerImageQueryRepo,
	containerImageCmdRepo repository.ContainerImageCmdRepo,
	containerQueryRepo repository.ContainerQueryRepo,
	accountQueryRepo repository.AccountQueryRepo,
	activityRecordCmdRepo repository.ActivityRecordCmdRepo,
	createSnapshotDto dto.CreateContainerSnapshotImage,
) error {
	containerEntityWithMetrics, err := containerQueryRepo.ReadWithMetricsById(
		createSnapshotDto.AccountId, createSnapshotDto.ContainerId,
	)
	if err != nil {
		return errors.New("ContainerNotFound")
	}

	accountEntity, err := accountQueryRepo.ReadById(createSnapshotDto.AccountId)
	if err != nil {
		slog.Error("ReadAccountInfoInfraError", slog.Any("error", err))
		return errors.New("ReadAccountInfoInfraError")
	}

	storageAvailable := accountEntity.Quota.StorageBytes - accountEntity.QuotaUsage.StorageBytes
	isThereStorageAvailable := storageAvailable >= containerEntityWithMetrics.Metrics.StorageSpaceBytes
	if !isThereStorageAvailable {
		return errors.New("AccountStorageQuotaUsageExceeded")
	}

	imageId, err := containerImageCmdRepo.CreateSnapshot(createSnapshotDto)
	if err != nil {
		slog.Error("CreateContainerSnapshotImageInfraError", slog.Any("error", err))
		return errors.New("CreateContainerSnapshotImageInfraError")
	}

	NewCreateSecurityActivityRecord(activityRecordCmdRepo).
		CreateContainerSnapshotImage(createSnapshotDto, imageId)

	if createSnapshotDto.ShouldCreateArchive == nil {
		return nil
	}

	if !*createSnapshotDto.ShouldCreateArchive {
		return nil
	}

	createArchiveDto := dto.NewCreateContainerImageArchiveFile(
		createSnapshotDto.AccountId, imageId, createSnapshotDto.ArchiveCompressionFormat,
		createSnapshotDto.OperatorAccountId, createSnapshotDto.OperatorIpAddress,
	)
	_, err = CreateContainerImageArchiveFile(
		containerImageQueryRepo, containerImageCmdRepo,
		accountQueryRepo, activityRecordCmdRepo, createArchiveDto,
	)
	if err != nil {
		return err
	}

	if createSnapshotDto.ShouldDiscardImage == nil {
		return nil
	}

	if !*createSnapshotDto.ShouldDiscardImage {
		return nil
	}

	deleteImageDto := dto.NewDeleteContainerImage(
		createSnapshotDto.AccountId, imageId,
		createSnapshotDto.OperatorAccountId, createSnapshotDto.OperatorIpAddress,
	)
	err = containerImageCmdRepo.Delete(deleteImageDto)
	if err != nil {
		slog.Error("DeleteContainerSnapshotImageInfraError", slog.Any("error", err))
		return errors.New("DeleteContainerSnapshotImageInfraError")
	}

	return nil
}
