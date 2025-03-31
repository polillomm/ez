package entity

import "github.com/goinfinite/ez/src/domain/valueObject"

type Mapping struct {
	Id          valueObject.MappingId       `json:"id"`
	AccountId   valueObject.AccountId       `json:"accountId"`
	AccountName valueObject.UnixUsername    `json:"accountName"`
	Hostname    *valueObject.Fqdn           `json:"hostname"`
	PublicPort  valueObject.NetworkPort     `json:"publicPort"`
	Protocol    valueObject.NetworkProtocol `json:"protocol"`
	Targets     []MappingTarget             `json:"targets"`
	CreatedAt   valueObject.UnixTime        `json:"createdAt"`
	UpdatedAt   valueObject.UnixTime        `json:"updatedAt"`
}

func NewMapping(
	id valueObject.MappingId,
	accountId valueObject.AccountId,
	accountName valueObject.UnixUsername,
	hostname *valueObject.Fqdn,
	publicPort valueObject.NetworkPort,
	protocol valueObject.NetworkProtocol,
	targets []MappingTarget,
	createdAt valueObject.UnixTime,
	updatedAt valueObject.UnixTime,
) Mapping {
	return Mapping{
		Id:          id,
		AccountId:   accountId,
		AccountName: accountName,
		Hostname:    hostname,
		PublicPort:  publicPort,
		Protocol:    protocol,
		Targets:     targets,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}
