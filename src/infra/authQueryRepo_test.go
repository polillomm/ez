package infra

import (
	"os"
	"testing"
	"time"

	testHelpers "github.com/speedianet/control/src/devUtils"
	"github.com/speedianet/control/src/domain/dto"
	"github.com/speedianet/control/src/domain/valueObject"
)

func TestAuthQueryRepo(t *testing.T) {
	testHelpers.LoadEnvVars()
	persistentDbSvc := testHelpers.GetPersistentDbSvc()
	authQueryRepo := NewAuthQueryRepo(persistentDbSvc)
	accCmdRepo := NewAccCmdRepo(persistentDbSvc)

	t.Run("ValidLoginCredentials", func(t *testing.T) {
		username, _ := valueObject.NewUsername(os.Getenv("DUMMY_USER_NAME"))
		password, _ := valueObject.NewPassword(os.Getenv("DUMMY_USER_PASS"))

		login := dto.NewLogin(username, password)
		isValid := authQueryRepo.IsLoginValid(login)
		if !isValid {
			t.Error("Expected valid login credentials, but got invalid")
		}
	})

	t.Run("InvalidLoginCredentials", func(t *testing.T) {
		username, _ := valueObject.NewUsername(os.Getenv("DUMMY_USER_NAME"))
		password, _ := valueObject.NewPassword("wrongPassword")

		login := dto.NewLogin(username, password)
		isValid := authQueryRepo.IsLoginValid(login)
		if isValid {
			t.Error("Expected invalid login credentials, but got valid")
		}
	})

	t.Run("ValidSessionAccessToken", func(t *testing.T) {
		authCmdRepo := AuthCmdRepo{}

		token, _ := authCmdRepo.GenerateSessionToken(
			valueObject.AccountId(1000),
			valueObject.NewUnixTimeAfterNow(3*time.Hour),
			valueObject.NewIpAddressPanic("127.0.0.1"),
		)

		_, err := authQueryRepo.GetAccessTokenDetails(token.TokenStr)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("InvalidSessionAccessToken", func(t *testing.T) {
		invalidToken, _ := valueObject.NewAccessTokenValue(
			"invalidTokenInvalidTokenInvalidTokenInvalidTokenInvalidToken",
		)
		_, err := authQueryRepo.GetAccessTokenDetails(invalidToken)
		if err == nil {
			t.Error("ExpectingError")
		}
	})

	t.Run("ValidAccountApiKey", func(t *testing.T) {
		apiKey, err := accCmdRepo.UpdateApiKey(
			valueObject.NewAccountIdPanic(os.Getenv("DUMMY_USER_ID")),
		)
		if err != nil {
			t.Error(err)
		}

		_, err = authQueryRepo.GetAccessTokenDetails(apiKey)
		if err != nil {
			t.Error(err)
		}
	})
}
