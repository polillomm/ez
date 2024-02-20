package dbModel

import (
	"time"

	"github.com/speedianet/control/src/domain/entity"
	"github.com/speedianet/control/src/domain/valueObject"
)

type LicenseInfo struct {
	ID          uint `gorm:"primarykey"`
	Method      string
	Status      string
	Fingerprint string
	ErrorCount  uint `gorm:"default:0"`
	ExpiresAt   time.Time
	LastCheckAt time.Time
	UpdatedAt   time.Time
}

func (LicenseInfo) TableName() string {
	return "license_info"
}

func NewLicenseInfo(
	method string,
	status string,
	fingerprint string,
	errorCount uint,
	expiresAt time.Time,
	lastCheckAt time.Time,
) LicenseInfo {
	licenseInfoModel := LicenseInfo{
		ID:          1,
		Method:      method,
		Status:      status,
		Fingerprint: fingerprint,
		ErrorCount:  errorCount,
		ExpiresAt:   expiresAt,
		LastCheckAt: lastCheckAt,
	}

	return licenseInfoModel
}

func (LicenseInfo) ToModel(entity entity.LicenseInfo) LicenseInfo {
	expiresAt := time.Unix(entity.ExpiresAt.Get(), 0)
	lastCheckAt := time.Unix(entity.LastCheckAt.Get(), 0)

	return NewLicenseInfo(
		entity.Method.String(),
		entity.Status.String(),
		entity.Fingerprint.String(),
		entity.ErrorCount,
		expiresAt,
		lastCheckAt,
	)
}

func (model LicenseInfo) ToEntity() (entity.LicenseInfo, error) {
	var licenseInfo entity.LicenseInfo

	licenseMethod, err := valueObject.NewLicenseMethod(model.Method)
	if err != nil {
		return licenseInfo, err
	}

	licenseStatus, err := valueObject.NewLicenseStatus(model.Status)
	if err != nil {
		return licenseInfo, err
	}

	licenseFingerprint, err := valueObject.NewLicenseFingerprint(model.Fingerprint)
	if err != nil {
		return licenseInfo, err
	}

	expiresAt := valueObject.UnixTime(model.ExpiresAt.Unix())
	lastCheckAt := valueObject.UnixTime(model.LastCheckAt.Unix())
	updatedAt := valueObject.UnixTime(model.UpdatedAt.Unix())

	licenseInfo = entity.NewLicenseInfo(
		licenseMethod,
		licenseStatus,
		licenseFingerprint,
		model.ErrorCount,
		expiresAt,
		lastCheckAt,
		updatedAt,
	)

	return licenseInfo, nil
}
