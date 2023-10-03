package infra

import (
	"os"
	"testing"

	testHelpers "github.com/speedianet/sfm/src/devUtils"
	"github.com/speedianet/sfm/src/domain/dto"
	"github.com/speedianet/sfm/src/domain/valueObject"
)

func TestContainerCmdRepo(t *testing.T) {
	testHelpers.LoadEnvVars()

	t.Run("AddContainer", func(t *testing.T) {
		repo := ContainerCmdRepo{}

		portBindings := []valueObject.PortBinding{
			valueObject.NewPortBinding(
				valueObject.NewNetworkProtocolPanic("tcp"),
				8080,
				8080,
			),
			valueObject.NewPortBinding(
				valueObject.NewNetworkProtocolPanic("tcp"),
				8443,
				8443,
			),
		}

		restartPolicy := valueObject.NewContainerRestartPolicyPanic("unless-stopped")

		resourceProfileId := valueObject.NewResourceProfileIdPanic(0)

		envs := []valueObject.ContainerEnv{
			valueObject.NewContainerEnvPanic("SFM_ENV1=testing"),
			valueObject.NewContainerEnvPanic("SFM_ENV2=testing"),
		}

		accountId := valueObject.NewAccountIdPanic(os.Getenv("DUMMY_USER_ID"))

		addContainer := dto.NewAddContainer(
			accountId,
			valueObject.NewFqdnPanic("speedia.net"),
			valueObject.NewContainerImgAddressPanic("docker.io/speedia/sam:latest"),
			portBindings,
			&restartPolicy,
			nil,
			&resourceProfileId,
			envs,
		)

		err := repo.Add(addContainer)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("DeleteContainer", func(t *testing.T) {
		accId := valueObject.NewAccountIdPanic(os.Getenv("DUMMY_USER_ID"))
		containers, err := ContainerQueryRepo{}.GetByAccId(accId)
		if err != nil {
			t.Errorf("GetContainersFailed: %v", err)
		}

		if len(containers) == 0 {
			t.Errorf("NoContainersFound: %v", err)
		}

		err = ContainerCmdRepo{}.Delete(
			accId,
			containers[0].Id,
		)
		if err != nil {
			t.Errorf("DeleteContainerFailed: %v", err)
		}
	})

	t.Run("UpdateContainer", func(t *testing.T) {
		accId := valueObject.NewAccountIdPanic(os.Getenv("DUMMY_USER_ID"))
		containers, err := ContainerQueryRepo{}.GetByAccId(accId)
		if err != nil {
			t.Errorf("GetContainersFailed: %v", err)
		}

		if len(containers) == 0 {
			t.Errorf("NoContainersFound: %v", err)
		}

		updateContainer := dto.NewUpdateContainer(
			accId,
			containers[0].Id,
			false,
			nil,
		)

		err = ContainerCmdRepo{}.Update(updateContainer)
		if err != nil {
			t.Errorf("UpdateContainerFailed: %v", err)
		}
	})
}
