package repository

import (
	"github.com/speedianet/sfm/src/domain/dto"
	"github.com/speedianet/sfm/src/domain/entity"
	"github.com/speedianet/sfm/src/domain/valueObject"
)

type ContainerCmdRepo interface {
	Add(addContainer dto.AddContainer) error
	Update(
		currentContainer entity.Container,
		updateContainer dto.UpdateContainer,
	) error
	Delete(
		accId valueObject.AccountId,
		containerId valueObject.ContainerId,
	) error
}
