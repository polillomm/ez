package dbModel

import (
	"errors"
	"time"

	"github.com/speedianet/sfm/src/domain/entity"
	"github.com/speedianet/sfm/src/domain/valueObject"
	"gorm.io/gorm"
)

type Account struct {
	ID         uint   `gorm:"primarykey"`
	GroupID    uint   `gorm:"not null"`
	Username   string `gorm:"not null"`
	KeyHash    *string
	Quota      AccountQuota
	QuotaUsage AccountQuotaUsage
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (Account) TableName() string {
	return "accounts"
}

func (Account) ToModel(entity entity.Account) (Account, error) {
	accId := uint(entity.Id.Get())
	quota, err := AccountQuota{}.ToModel(entity.Quota, accId)
	if err != nil {
		return Account{}, err
	}

	quotaUsage, err := AccountQuotaUsage{}.ToModel(entity.QuotaUsage, accId)
	if err != nil {
		return Account{}, err
	}

	return Account{
		ID:         accId,
		GroupID:    uint(entity.GroupId.Get()),
		Username:   entity.Username.String(),
		KeyHash:    nil,
		Quota:      quota,
		QuotaUsage: quotaUsage,
	}, nil
}

func (model Account) ToEntity() (entity.Account, error) {
	accId, err := valueObject.NewAccountId(model.ID)
	if err != nil {
		return entity.Account{}, err
	}

	groupId, err := valueObject.NewGroupId(model.GroupID)
	if err != nil {
		return entity.Account{}, err
	}

	username, err := valueObject.NewUsername(model.Username)
	if err != nil {
		return entity.Account{}, err
	}

	quota, err := model.Quota.ToValueObject()
	if err != nil {
		return entity.Account{}, err
	}

	quotaUsage, err := model.QuotaUsage.ToValueObject()
	if err != nil {
		return entity.Account{}, err
	}

	return entity.NewAccount(
		accId,
		groupId,
		username,
		quota,
		quotaUsage,
		valueObject.UnixTime(model.CreatedAt.Unix()),
		valueObject.UnixTime(model.UpdatedAt.Unix()),
	), nil
}

func (model Account) Delete(dbSvc *gorm.DB) error {
	relatedTables := []string{
		AccountQuota{}.TableName(),
		AccountQuotaUsage{}.TableName(),
	}

	for _, tableName := range relatedTables {
		err := dbSvc.Exec(
			"DELETE FROM "+tableName+" WHERE account_id = ?", model.ID,
		).Error
		if err != nil {
			return errors.New("DeleteAccRelatedTablesDbError")
		}
	}

	err := dbSvc.Delete(model, model.ID).Error
	if err != nil {
		return errors.New("DeleteAccDbError")
	}

	return nil
}
