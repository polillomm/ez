package apiController

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/speedianet/control/src/domain/useCase"
	"github.com/speedianet/control/src/domain/valueObject"
	"github.com/speedianet/control/src/infra"
	"github.com/speedianet/control/src/infra/db"
	apiHelper "github.com/speedianet/control/src/presentation/api/helper"
)

// GetContainerRegistryImages	 godoc
// @Summary      GetContainerRegistryImages
// @Description  Get container registry images.
// @Tags         container
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        name    query     string  false  "ImageName"
// @Success      200 {array} entity.RegistryImage
// @Router       /container/registry/image/ [get]
func GetContainerRegistryImagesController(c echo.Context) error {
	dbSvc := c.Get("dbSvc").(*db.DatabaseService)

	imageNameStr := c.QueryParam("name")
	var imageNamePtr *valueObject.RegistryImageName
	if imageNameStr != "" {
		imageName := valueObject.NewRegistryImageNamePanic(imageNameStr)
		imageNamePtr = &imageName
	}

	containerRegistryQueryRepo := infra.NewContainerRegistryQueryRepo(dbSvc)
	imagesList, err := useCase.GetRegistryImages(
		containerRegistryQueryRepo,
		imageNamePtr,
	)
	if err != nil {
		return apiHelper.ResponseWrapper(c, http.StatusInternalServerError, err.Error())
	}

	return apiHelper.ResponseWrapper(c, http.StatusOK, imagesList)
}
