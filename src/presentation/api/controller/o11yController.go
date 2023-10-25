package apiController

import (
	"net/http"

	"github.com/goinfinite/fleet/src/domain/useCase"
	"github.com/goinfinite/fleet/src/infra"
	apiHelper "github.com/goinfinite/fleet/src/presentation/api/helper"
	"github.com/labstack/echo/v4"
)

// O11yOverview  godoc
// @Summary      O11yOverview
// @Description  Show system information and resource usage.
// @Tags         o11y
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Success      200 {object} entity.O11yOverview
// @Router       /o11y/overview/ [get]
func O11yOverviewController(c echo.Context) error {
	o11yQueryRepo := infra.O11yQueryRepo{}
	o11yOverview, err := useCase.GetO11yOverview(o11yQueryRepo)
	if err != nil {
		return apiHelper.ResponseWrapper(c, http.StatusInternalServerError, err.Error())
	}

	return apiHelper.ResponseWrapper(c, http.StatusOK, o11yOverview)
}
