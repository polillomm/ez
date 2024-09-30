package useCase

import (
	"errors"
	"log/slog"

	"github.com/goinfinite/ez/src/domain/dto"
	"github.com/goinfinite/ez/src/domain/repository"
	"github.com/goinfinite/ez/src/domain/valueObject"
)

func StopAllContainers(
	containerQueryRepo repository.ContainerQueryRepo,
	containerCmdRepo repository.ContainerCmdRepo,
) error {
	containers, err := containerQueryRepo.Read()
	if err != nil {
		return errors.New("ReadContainersError: " + err.Error())
	}

	for _, currentContainer := range containers {
		newContainerStatus := false
		updateDto := dto.NewUpdateContainer(
			currentContainer.AccountId, currentContainer.Id, &newContainerStatus,
			&currentContainer.ProfileId, valueObject.SystemAccountId,
			valueObject.SystemIpAddress,
		)

		err = containerCmdRepo.Update(updateDto)
		if err != nil {
			slog.Error(
				"StopContainerError",
				slog.String("containerId", currentContainer.Id.String()),
				slog.Any("error", err),
			)
			continue
		}
	}

	return nil
}
