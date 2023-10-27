package useCase

import (
	"errors"
	"log"

	"github.com/goinfinite/fleet/src/domain/dto"
	"github.com/goinfinite/fleet/src/domain/repository"
	"github.com/goinfinite/fleet/src/domain/valueObject"
)

func updateContainersWithProfileId(
	containerQueryRepo repository.ContainerQueryRepo,
	containerCmdRepo repository.ContainerCmdRepo,
	profileId valueObject.ContainerProfileId,
) error {
	containers, err := containerQueryRepo.Get()
	if err != nil {
		log.Printf("GetContainersError: %s", err)
		return errors.New("GetContainersInfraError")
	}

	for _, container := range containers {
		if container.ProfileId != profileId {
			continue
		}

		updateContainerDto := dto.NewUpdateContainer(
			container.AccountId,
			container.Id,
			&container.Status,
			&profileId,
		)

		err := containerCmdRepo.Update(container, updateContainerDto)
		if err != nil {
			log.Printf("UpdateContainersWithProfileIdError: %s", err)
			continue
		}
	}

	return nil
}

func UpdateContainerProfile(
	containerProfileQueryRepo repository.ContainerProfileQueryRepo,
	containerProfileCmdRepo repository.ContainerProfileCmdRepo,
	containerQueryRepo repository.ContainerQueryRepo,
	containerCmdRepo repository.ContainerCmdRepo,
	updateContainerProfileDto dto.UpdateContainerProfile,
) error {
	_, err := containerProfileQueryRepo.GetById(updateContainerProfileDto.Id)
	if err != nil {
		return errors.New("ContainerProfileNotFound")
	}

	err = containerProfileCmdRepo.Update(updateContainerProfileDto)
	if err != nil {
		log.Printf("UpdateContainerProfileError: %s", err)
		return errors.New("UpdateContainerProfileInfraError")
	}

	shouldUpdateContainers := updateContainerProfileDto.BaseSpecs != nil
	if !shouldUpdateContainers {
		return nil
	}

	err = updateContainersWithProfileId(
		containerQueryRepo,
		containerCmdRepo,
		updateContainerProfileDto.Id,
	)
	if err != nil {
		log.Printf("UpdateContainersAfterProfileUpdateError: %s", err)
		return errors.New("UpdateContainersAfterProfileUpdate")
	}

	return nil
}
