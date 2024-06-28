package infra

import (
	"os"
	"testing"

	testHelpers "github.com/speedianet/control/src/devUtils"
	"github.com/speedianet/control/src/domain/dto"
	"github.com/speedianet/control/src/domain/valueObject"
	infraHelper "github.com/speedianet/control/src/infra/helper"
)

func addDummyUser() error {
	username, _ := valueObject.NewUsername(os.Getenv("DUMMY_USER_NAME"))
	password, _ := valueObject.NewPassword(os.Getenv("DUMMY_USER_PASS"))
	quota := valueObject.NewAccountQuotaWithDefaultValues()
	addAccount := dto.NewAddAccount(username, password, &quota)

	accCmdRepo := NewAccCmdRepo(testHelpers.GetPersistentDbSvc())
	err := accCmdRepo.Add(addAccount)
	if err != nil {
		return err
	}

	return nil
}

func deleteDummyUser() error {
	accountId := valueObject.NewAccountIdPanic(os.Getenv("DUMMY_USER_ID"))

	accCmdRepo := NewAccCmdRepo(testHelpers.GetPersistentDbSvc())
	err := accCmdRepo.Delete(accountId)
	if err != nil {
		return err
	}

	return nil
}

func resetDummyUser() {
	_ = addDummyUser()
	_ = deleteDummyUser()
	_ = addDummyUser()
}

func TestAccCmdRepo(t *testing.T) {
	testHelpers.LoadEnvVars()
	persistentDbSvc := testHelpers.GetPersistentDbSvc()
	accQueryRepo := NewAccQueryRepo(persistentDbSvc)
	accCmdRepo := NewAccCmdRepo(persistentDbSvc)

	t.Run("AddValidAccount", func(t *testing.T) {
		err := addDummyUser()
		if err != nil {
			t.Errorf("UnexpectedError: %v", err)
		}
	})

	t.Run("AddInvalidAccount", func(t *testing.T) {
		username, _ := valueObject.NewUsername("root")
		password, _ := valueObject.NewPassword("invalid")
		quota := valueObject.NewAccountQuotaWithDefaultValues()
		addAccount := dto.NewAddAccount(username, password, &quota)

		err := accCmdRepo.Add(addAccount)
		if err == nil {
			t.Error("AccountShouldNotBeAdded")
		}
	})

	t.Run("DeleteValidAccount", func(t *testing.T) {
		err := deleteDummyUser()
		if err != nil {
			t.Errorf("UnexpectedError: %v", err)
		}
	})

	t.Run("UpdatePasswordValidAccount", func(t *testing.T) {
		resetDummyUser()

		accountId := valueObject.NewAccountIdPanic(os.Getenv("DUMMY_USER_ID"))
		newPassword, _ := valueObject.NewPassword("newPassword")

		err := accCmdRepo.UpdatePassword(accountId, newPassword)
		if err != nil {
			t.Errorf("UnexpectedError: %v", err)
		}
	})

	t.Run("UpdateApiKeyValidAccount", func(t *testing.T) {
		resetDummyUser()

		accountId := valueObject.NewAccountIdPanic(os.Getenv("DUMMY_USER_ID"))

		_, err := accCmdRepo.UpdateApiKey(accountId)
		if err != nil {
			t.Errorf("UnexpectedError: %v", err)
		}
	})

	t.Run("UpdateQuotaValidAccount", func(t *testing.T) {
		resetDummyUser()

		accountId := valueObject.NewAccountIdPanic(os.Getenv("DUMMY_USER_ID"))
		quota := valueObject.NewAccountQuotaWithDefaultValues()
		quota.CpuCores = valueObject.NewCpuCoresCountPanic(1)
		quota.DiskBytes = valueObject.NewBytePanic(1073741824)

		err := accCmdRepo.UpdateQuota(accountId, quota)
		if err != nil {
			t.Errorf("UnexpectedError: %v", err)
		}
	})

	t.Run("UpdateQuotasUsage", func(t *testing.T) {
		resetDummyUser()

		testFilePath := "/var/data/" + os.Getenv("DUMMY_USER_NAME") + "/test.file"

		_, err := infraHelper.RunCmd("fallocate", "-l", "100M", testFilePath)
		if err != nil {
			t.Error(err)
		}

		accId := valueObject.NewAccountIdPanic(os.Getenv("DUMMY_USER_ID"))
		os.Chown(testFilePath, int(accId.Get()), int(accId.Get()))

		err = accCmdRepo.UpdateQuotaUsage(accId)
		if err != nil {
			t.Error(err)
		}

		accEntity, err := accQueryRepo.GetById(accId)
		if err != nil {
			t.Error(err)
		}
		if accEntity.QuotaUsage.DiskBytes.Get() < 100000000 {
			t.Error("QuotaUsageNotUpdated")
		}

		_, err = infraHelper.RunCmd("rm", "-f", testFilePath)
		if err != nil {
			t.Error(err)
		}
	})
}
