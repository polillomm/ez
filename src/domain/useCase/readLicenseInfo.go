package useCase

import (
	"errors"
	"log/slog"

	"github.com/goinfinite/ez/src/domain/entity"
	"github.com/goinfinite/ez/src/domain/repository"
)

func ReadLicenseInfo(
	licenseQueryRepo repository.LicenseQueryRepo,
) (entity.LicenseInfo, error) {
	licenseInfo, err := licenseQueryRepo.Read()
	if err != nil {
		slog.Error("ReadLicenseInfoInfraError", slog.Any("error", err))
		return licenseInfo, errors.New("ReadLicenseInfoInfraError")
	}

	return licenseInfo, nil
}
