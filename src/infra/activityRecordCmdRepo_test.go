package infra

import (
	"testing"

	testHelpers "github.com/speedianet/control/src/devUtils"
	"github.com/speedianet/control/src/domain/dto"
	"github.com/speedianet/control/src/domain/valueObject"
)

func TestActivityRecordCmdRepo(t *testing.T) {
	testHelpers.LoadEnvVars()
	trailDbSvc := testHelpers.GetTrailDbSvc()
	activityRecordCmdRepo := NewActivityRecordCmdRepo(trailDbSvc)
	level, _ := valueObject.NewActivityRecordLevel("SEC")
	recordCode, _ := valueObject.NewActivityRecordCode("LoginFailed")
	operatorIpAddress := valueObject.NewLocalhostIpAddress()

	t.Run("CreateActivityRecord", func(t *testing.T) {
		createDto := dto.CreateActivityRecord{
			RecordLevel:       level,
			RecordCode:        recordCode,
			OperatorIpAddress: &operatorIpAddress,
		}

		err := activityRecordCmdRepo.Create(createDto)
		if err != nil {
			t.Errorf("ExpectedNoErrorButGot: %v", err)
		}
	})

	t.Run("DeleteActivityRecords", func(t *testing.T) {
		ipAddress := valueObject.NewLocalhostIpAddress()
		deleteDto := dto.NewDeleteActivityRecords(
			nil, &level, &recordCode, nil, &ipAddress, nil, nil, nil,
			nil, nil, nil, nil, nil, nil,
		)

		err := activityRecordCmdRepo.Delete(deleteDto)
		if err != nil {
			t.Errorf("ExpectedNoErrorButGot: %v", err)
		}
	})
}
