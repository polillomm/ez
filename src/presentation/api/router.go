package api

import (
	_ "embed"

	"github.com/labstack/echo/v4"
	"github.com/speedianet/control/src/infra/db"
	apiController "github.com/speedianet/control/src/presentation/api/controller"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/speedianet/control/src/presentation/api/docs"
)

type Router struct {
	baseRoute       *echo.Group
	persistentDbSvc *db.PersistentDatabaseService
	transientDbSvc  *db.TransientDatabaseService
}

func NewRouter(
	baseRoute *echo.Group,
	persistentDbSvc *db.PersistentDatabaseService,
	transientDbSvc *db.TransientDatabaseService,
) *Router {
	return &Router{
		baseRoute:       baseRoute,
		persistentDbSvc: persistentDbSvc,
		transientDbSvc:  transientDbSvc,
	}
}

func (router *Router) swaggerRoute() {
	swaggerGroup := router.baseRoute.Group("/swagger")
	swaggerGroup.GET("/*", echoSwagger.WrapHandler)
}

func (router *Router) authRoutes() {
	authGroup := router.baseRoute.Group("/v1/auth")

	authController := apiController.NewAuthController(router.persistentDbSvc)
	authGroup.POST("/login/", authController.Login)
}

func (router *Router) accountRoutes() {
	accountGroup := router.baseRoute.Group("/v1/account")

	accountController := apiController.NewAccountController(router.persistentDbSvc)
	accountGroup.GET("/", accountController.Read)
	accountGroup.POST("/", accountController.Create)
	accountGroup.PUT("/", accountController.Update)
	accountGroup.DELETE("/:accountId/", accountController.Delete)
	go accountController.AutoUpdateAccountsQuotaUsage()
}

func (router *Router) containerRoutes() {
	containerGroup := router.baseRoute.Group("/v1/container")

	containerController := apiController.NewContainerController(router.persistentDbSvc)
	containerGroup.GET("/", containerController.Read)
	containerGroup.GET("/metrics/", containerController.ReadWithMetrics)
	containerGroup.GET("/auto-login/:containerId/", containerController.AutoLogin)
	containerGroup.POST("/", containerController.Create)
	containerGroup.PUT("/", containerController.Update)
	containerGroup.DELETE(
		"/:accountId/:containerId/",
		containerController.Delete,
	)

	containerProfileGroup := containerGroup.Group("/profile")
	containerProfileGroup.GET("/", apiController.GetContainerProfilesController)
	containerProfileGroup.POST("/", apiController.CreateContainerProfileController)
	containerProfileGroup.PUT("/", apiController.UpdateContainerProfileController)
	containerProfileGroup.DELETE(
		"/:profileId/",
		apiController.DeleteContainerProfileController,
	)

	containerRegistryGroup := containerGroup.Group("/registry")
	containerRegistryGroup.GET("/image/", apiController.GetContainerRegistryImagesController)
	containerRegistryGroup.GET(
		"/image/tagged/",
		apiController.GetContainerRegistryTaggedImageController,
	)
}

func (router *Router) licenseRoutes() {
	licenseGroup := router.baseRoute.Group("/v1/license")
	licenseGroup.GET("/", apiController.GetLicenseInfoController)
	go apiController.AutoLicenseValidationController(
		router.persistentDbSvc,
		router.transientDbSvc,
	)
}

func (router *Router) mappingRoutes() {
	mappingGroup := router.baseRoute.Group("/v1/mapping")

	mappingController := apiController.NewMappingController(router.persistentDbSvc)
	mappingGroup.GET("/", mappingController.Read)
	mappingGroup.POST("/", mappingController.Create)
	mappingGroup.DELETE("/:mappingId/", mappingController.Delete)
	mappingGroup.POST("/target/", mappingController.CreateTarget)
	mappingGroup.DELETE(
		"/:mappingId/target/:targetId/",
		mappingController.DeleteTarget,
	)
}

func (router *Router) o11yRoutes() {
	o11yGroup := router.baseRoute.Group("/v1/o11y")
	o11yGroup.GET("/overview/", apiController.O11yOverviewController)
}

func (router *Router) scheduledTaskRoutes() {
	scheduledTaskGroup := router.baseRoute.Group("/v1/scheduled-task")

	scheduledTaskController := apiController.NewScheduledTaskController(router.persistentDbSvc)
	scheduledTaskGroup.GET("/", scheduledTaskController.Read)
	scheduledTaskGroup.PUT("/", scheduledTaskController.Update)
	go scheduledTaskController.Run()
}

func (router *Router) RegisterRoutes() {
	router.swaggerRoute()
	router.authRoutes()
	router.accountRoutes()
	router.containerRoutes()
	router.licenseRoutes()
	router.mappingRoutes()
	router.o11yRoutes()
	router.scheduledTaskRoutes()
}
