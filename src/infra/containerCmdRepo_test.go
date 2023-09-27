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

		baseSpecs := valueObject.NewContainerSpecs(
			valueObject.NewCpuCoresCountPanic(1),
			valueObject.Byte(1073741824),
		)
		maxSpecs := valueObject.NewContainerSpecs(
			valueObject.NewCpuCoresCountPanic(2),
			valueObject.Byte(2147483648),
		)

		envs := []valueObject.ContainerEnv{
			valueObject.NewContainerEnvPanic("SFM_ENV1=testing"),
			valueObject.NewContainerEnvPanic("SFM_ENV2=testing"),
		}

		accountId := valueObject.NewAccountIdPanic(os.Getenv("DUMMY_USER_ID"))

		addContainer := dto.NewAddContainer(
			accountId,
			valueObject.NewFqdnPanic("speedia.net"),
			valueObject.NewContainerImgAddressPanic("docker.io/speedia/sam:latest"),
			&portBindings,
			&restartPolicy,
			nil,
			&baseSpecs,
			&maxSpecs,
			&envs,
		)

		err := repo.Add(addContainer)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
}
