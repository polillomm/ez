package entity

import (
	"encoding/json"

	"github.com/goinfinite/ez/src/domain/valueObject"
)

type ContainerImageArchiveFile struct {
	ImageId      valueObject.ContainerImageId `json:"imageId"`
	AccountId    valueObject.AccountId        `json:"accountId"`
	UnixFilePath valueObject.UnixFilePath     `json:"unixFilePath"`
	SizeBytes    valueObject.Byte             `json:"sizeBytes"`
	DownloadUrl  *valueObject.Url             `json:"downloadUrl,omitempty"`
	ContainerId  *valueObject.ContainerId     `json:"containerId,omitempty"`
	CreatedAt    valueObject.UnixTime         `json:"createdAt"`
}

func NewContainerImageArchiveFile(
	imageId valueObject.ContainerImageId,
	accountId valueObject.AccountId,
	unixFilePath valueObject.UnixFilePath,
	sizeBytes valueObject.Byte,
	downloadUrl *valueObject.Url,
	containerId *valueObject.ContainerId,
	createdAt valueObject.UnixTime,
) ContainerImageArchiveFile {
	return ContainerImageArchiveFile{
		ImageId:      imageId,
		AccountId:    accountId,
		UnixFilePath: unixFilePath,
		SizeBytes:    sizeBytes,
		DownloadUrl:  downloadUrl,
		ContainerId:  containerId,
		CreatedAt:    createdAt,
	}
}

func (entity ContainerImageArchiveFile) JsonSerialize() string {
	jsonBytes, _ := json.Marshal(entity)
	return string(jsonBytes)
}
