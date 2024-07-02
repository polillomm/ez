package cliController

import (
	"github.com/speedianet/control/src/domain/useCase"
	"github.com/speedianet/control/src/infra"
	"github.com/speedianet/control/src/infra/db"
	cliHelper "github.com/speedianet/control/src/presentation/cli/helper"
	"github.com/spf13/cobra"
)

type LicenseController struct {
	persistentDbSvc *db.PersistentDatabaseService
	transientDbSvc  *db.TransientDatabaseService
}

func NewLicenseController(
	persistentDbSvc *db.PersistentDatabaseService,
	transientDbSvc *db.TransientDatabaseService,
) *LicenseController {
	return &LicenseController{
		persistentDbSvc: persistentDbSvc,
		transientDbSvc:  transientDbSvc,
	}
}

func (controller *LicenseController) ReadLicenseInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "ReadLicenseInfo",
		Run: func(cmd *cobra.Command, args []string) {
			licenseQueryRepo := infra.NewLicenseQueryRepo(
				controller.persistentDbSvc,
				controller.transientDbSvc,
			)
			licenseStatus, err := useCase.ReadLicenseInfo(licenseQueryRepo)
			if err != nil {
				cliHelper.ResponseWrapper(false, err.Error())
			}

			cliHelper.ResponseWrapper(true, licenseStatus)
		},
	}

	return cmd
}

func (controller *LicenseController) RefreshLicense() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "refresh",
		Short: "RefreshLicense",
		Run: func(cmd *cobra.Command, args []string) {
			licenseQueryRepo := infra.NewLicenseQueryRepo(
				controller.persistentDbSvc,
				controller.transientDbSvc,
			)
			licenseCmdRepo := infra.NewLicenseCmdRepo(
				controller.persistentDbSvc,
				controller.transientDbSvc,
			)

			err := useCase.LicenseValidation(licenseQueryRepo, licenseCmdRepo)
			if err != nil {
				cliHelper.ResponseWrapper(false, err.Error())
			}

			cliHelper.ResponseWrapper(true, "LicenseRefreshed")
		},
	}

	return cmd
}
