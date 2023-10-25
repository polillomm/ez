package repository

import (
	"github.com/goinfinite/fleet/src/domain/dto"
	"github.com/goinfinite/fleet/src/domain/valueObject"
)

type AccCmdRepo interface {
	Add(addAccount dto.AddAccount) error
	Delete(accId valueObject.AccountId) error
	UpdatePassword(accId valueObject.AccountId, password valueObject.Password) error
	UpdateApiKey(accId valueObject.AccountId) (valueObject.AccessTokenStr, error)
	UpdateQuota(accId valueObject.AccountId, quota valueObject.AccountQuota) error
	UpdateQuotaUsage(accId valueObject.AccountId) error
}
