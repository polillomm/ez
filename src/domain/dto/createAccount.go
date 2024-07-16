package dto

import "github.com/speedianet/control/src/domain/valueObject"

type CreateAccount struct {
	Username          valueObject.Username      `json:"username"`
	Password          valueObject.Password      `json:"password"`
	Quota             *valueObject.AccountQuota `json:"quota"`
	OperatorAccountId valueObject.AccountId     `json:"-"`
	IpAddress         valueObject.IpAddress     `json:"-"`
}

func NewCreateAccount(
	username valueObject.Username,
	password valueObject.Password,
	quota *valueObject.AccountQuota,
	operatorAccountId valueObject.AccountId,
	ipAddress valueObject.IpAddress,
) CreateAccount {
	return CreateAccount{
		Username:          username,
		Password:          password,
		Quota:             quota,
		OperatorAccountId: operatorAccountId,
		IpAddress:         ipAddress,
	}
}
