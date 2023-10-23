package useCase

import (
	"errors"
	"log"

	"github.com/speedianet/sfm/src/domain/dto"
	"github.com/speedianet/sfm/src/domain/repository"
)

func UpdateContainer(
	containerQueryRepo repository.ContainerQueryRepo,
	containerCmdRepo repository.ContainerCmdRepo,
	accQueryRepo repository.AccQueryRepo,
	accCmdRepo repository.AccCmdRepo,
	containerProfileQueryRepo repository.ContainerProfileQueryRepo,
	updateContainer dto.UpdateContainer,
) error {
	currentContainer, err := containerQueryRepo.GetById(
		updateContainer.AccountId,
		updateContainer.ContainerId,
	)
	if err != nil {
		return errors.New("ContainerNotFound")
	}

	shouldUpdateQuota := updateContainer.ProfileId != nil
	if shouldUpdateQuota {
		err = CheckAccountQuota(
			accQueryRepo,
			updateContainer.AccountId,
			containerProfileQueryRepo,
			*updateContainer.ProfileId,
		)
		if err != nil {
			return err
		}
	}

	err = containerCmdRepo.Update(currentContainer, updateContainer)
	if err != nil {
		log.Printf("UpdateContainerError: %s", err)
		return errors.New("UpdateContainerInfraError")
	}

	if shouldUpdateQuota {
		err = accCmdRepo.UpdateQuotaUsage(updateContainer.AccountId)
		if err != nil {
			log.Printf("UpdateAccountQuotaError: %s", err)
			return errors.New("UpdateAccountQuotaError")
		}
	}

	return nil
}
