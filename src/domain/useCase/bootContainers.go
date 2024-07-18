package useCase

import (
	"log"
	"log/slog"

	"github.com/speedianet/control/src/domain/dto"
	"github.com/speedianet/control/src/domain/repository"
)

func BootContainers(
	containerQueryRepo repository.ContainerQueryRepo,
	containerCmdRepo repository.ContainerCmdRepo,
) {
	containers, err := containerQueryRepo.Read()
	if err != nil {
		log.Printf("ReadContainersError: %v", err)
		return
	}

	for _, currentContainer := range containers {
		restartPolicyStr := currentContainer.RestartPolicy.String()
		shouldBoot := restartPolicyStr == "always" || restartPolicyStr == "on-failure"
		if !shouldBoot {
			continue
		}

		newContainerStatus := true
		updateDto := dto.NewUpdateContainer(
			currentContainer.AccountId,
			currentContainer.Id,
			&newContainerStatus,
			nil,
		)

		err = containerCmdRepo.Update(updateDto)
		if err != nil {
			slog.Error(
				"UpdateContainerInfraError",
				slog.String("containerId", currentContainer.Id.String()),
				slog.Any("error", err),
			)
			continue
		}
	}
}
