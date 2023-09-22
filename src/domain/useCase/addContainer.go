package useCase

import (
	"errors"
	"log"

	"github.com/speedianet/sfm/src/domain/dto"
	"github.com/speedianet/sfm/src/domain/repository"
	"github.com/speedianet/sfm/src/domain/valueObject"
)

func AddContainer(
	containerCmdRepo repository.ContainerCmdRepo,
	accQueryRepo repository.AccQueryRepo,
	addContainer dto.AddContainer,
) error {
	defaultSpecs := valueObject.NewContainerSpecs(
		0.5,
		valueObject.Byte(512000000),
	)

	if addContainer.BaseSpecs == nil {
		addContainer.BaseSpecs = &defaultSpecs
	}

	err := CheckAccountQuota(
		accQueryRepo,
		addContainer.AccountId,
		*addContainer.BaseSpecs,
	)
	if err != nil {
		log.Printf("QuotaCheckError: %s", err)
		return err
	}

	err = containerCmdRepo.Add(addContainer)
	if err != nil {
		log.Printf("AddContainerError: %s", err)
		return errors.New("AddContainerInfraError")
	}

	return nil
}
