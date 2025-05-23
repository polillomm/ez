package cliController

import (
	"github.com/goinfinite/ez/src/domain/valueObject"
	"github.com/goinfinite/ez/src/infra/db"
	cliHelper "github.com/goinfinite/ez/src/presentation/cli/helper"
	"github.com/goinfinite/ez/src/presentation/service"
	"github.com/spf13/cobra"
)

type AccountController struct {
	accountService *service.AccountService
}

func NewAccountController(
	persistentDbSvc *db.PersistentDatabaseService,
	trailDbSvc *db.TrailDatabaseService,
) *AccountController {
	return &AccountController{
		accountService: service.NewAccountService(persistentDbSvc, trailDbSvc),
	}
}

func (controller *AccountController) Read() *cobra.Command {
	var idUint64 uint64
	var usernameStr string
	var paginationPageNumberUint32 uint32
	var paginationItemsPerPageUint16 uint16
	var paginationSortByStr, paginationSortDirectionStr, paginationLastSeenIdStr string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "ReadAccounts",
		Run: func(cmd *cobra.Command, args []string) {
			requestBody := map[string]any{}

			if idUint64 != 0 {
				requestBody["id"] = idUint64
			}

			if usernameStr != "" {
				requestBody["username"] = usernameStr
			}

			requestBody = cliHelper.PaginationParser(
				requestBody, paginationPageNumberUint32, paginationItemsPerPageUint16,
				paginationSortByStr, paginationSortDirectionStr, paginationLastSeenIdStr,
			)

			cliHelper.ServiceResponseWrapper(
				controller.accountService.Read(requestBody),
			)
		},
	}

	cmd.Flags().Uint64VarP(&idUint64, "account-id", "i", 0, "AccountId")
	cmd.Flags().StringVarP(
		&usernameStr, "account-username", "n", "", "AccountUsername",
	)
	cmd.Flags().Uint32VarP(
		&paginationPageNumberUint32, "page-number", "p", 0, "PageNumber (Pagination)",
	)
	cmd.Flags().Uint16VarP(
		&paginationItemsPerPageUint16, "items-per-page", "m", 0,
		"ItemsPerPage (Pagination)",
	)
	cmd.Flags().StringVarP(
		&paginationSortByStr, "sort-by", "y", "", "SortBy (Pagination)",
	)
	cmd.Flags().StringVarP(
		&paginationSortDirectionStr, "sort-direction", "r", "",
		"SortDirection (Pagination)",
	)
	cmd.Flags().StringVarP(
		&paginationLastSeenIdStr, "last-seen-id", "l", "", "LastSeenId (Pagination)",
	)

	return cmd
}

func (controller *AccountController) Create() *cobra.Command {
	var usernameStr string
	var passwordStr string
	var quotaStr string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "CreateAccount",
		Run: func(cmd *cobra.Command, args []string) {
			requestBody := map[string]interface{}{
				"username": usernameStr,
				"password": passwordStr,
			}

			if quotaStr != "" {
				quota, err := valueObject.NewAccountQuotaFromString(quotaStr)
				if err != nil {
					cliHelper.ResponseWrapper(false, err.Error())
				}
				requestBody["quota"] = quota
			}

			cliHelper.ServiceResponseWrapper(
				controller.accountService.Create(requestBody),
			)
		},
	}

	cmd.Flags().StringVarP(&usernameStr, "username", "u", "", "Username")
	cmd.MarkFlagRequired("username")
	cmd.Flags().StringVarP(&passwordStr, "password", "p", "", "Password")
	cmd.MarkFlagRequired("password")
	cmd.Flags().StringVarP(
		&quotaStr, "quota", "q", "",
		"AccountQuota (cpu:memory:storage:inodes:storagePerformanceUnits)",
	)
	return cmd
}

func (controller *AccountController) Update() *cobra.Command {
	var accountIdStr, passwordStr, shouldUpdateApiKeyBoolStr, quotaStr string

	cmd := &cobra.Command{
		Use:   "update",
		Short: "UpdateAccount (pass or apiKey)",
		Run: func(cmd *cobra.Command, args []string) {
			requestBody := map[string]interface{}{
				"accountId":          accountIdStr,
				"shouldUpdateApiKey": shouldUpdateApiKeyBoolStr,
			}

			if passwordStr != "" {
				requestBody["password"] = passwordStr
			}

			if quotaStr != "" {
				quota, err := valueObject.NewAccountQuotaFromString(quotaStr)
				if err != nil {
					cliHelper.ResponseWrapper(false, err.Error())
				}
				requestBody["quota"] = quota
			}

			cliHelper.ServiceResponseWrapper(
				controller.accountService.Update(requestBody),
			)
		},
	}

	cmd.Flags().StringVarP(&accountIdStr, "id", "i", "", "AccountId")
	cmd.MarkFlagRequired("id")
	cmd.Flags().StringVarP(&passwordStr, "password", "p", "", "Password")
	cmd.Flags().StringVarP(
		&shouldUpdateApiKeyBoolStr, "update-api-key", "k", "false", "ShouldUpdateApiKey",
	)
	cmd.Flags().StringVarP(
		&quotaStr, "quota", "q", "",
		"AccountQuota (cpu:memory:storage:inodes:storagePerformanceUnits)",
	)
	return cmd
}

func (controller *AccountController) RefreshQuotas() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "refresh-quotas",
		Short: "RefreshAccountQuotas",
		Run: func(cmd *cobra.Command, args []string) {
			cliHelper.ServiceResponseWrapper(
				controller.accountService.RefreshQuotas(),
			)
		},
	}

	return cmd
}

func (controller *AccountController) Delete() *cobra.Command {
	var accountIdStr string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "DeleteAccount",
		Run: func(cmd *cobra.Command, args []string) {
			requestBody := map[string]interface{}{
				"accountId": accountIdStr,
			}

			cliHelper.ServiceResponseWrapper(
				controller.accountService.Delete(requestBody),
			)
		},
	}

	cmd.Flags().StringVarP(&accountIdStr, "id", "i", "", "AccountId")
	cmd.MarkFlagRequired("id")
	return cmd
}
