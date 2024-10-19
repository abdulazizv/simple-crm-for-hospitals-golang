package api

import (
	h "net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "gitlab.com/backend/api/docs"
	v1 "gitlab.com/backend/api/handler/v1"
	"gitlab.com/backend/api/middleware"
	"gitlab.com/backend/api/tokens"
	"gitlab.com/backend/config"
	"gitlab.com/backend/pkg/logger"
	"gitlab.com/backend/storage"
)

type Options struct {
	Cfg            config.Config
	Storage        storage.StorageI
	Log            logger.Logger
	CasbinEnforcer *casbin.Enforcer
}

// New ...
// @title           Clinic
// @version         1.0
// @description     This is Clinic server api
// @termsOfService  1 term Traffic Light
// host      		localhost:5000
// host      		clinic.addscontrol.uz
// @BasePath  		/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func New(opt *Options) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	corConfig := cors.DefaultConfig()
	corConfig.AllowAllOrigins = true
	corConfig.AllowCredentials = true
	corConfig.AllowHeaders = []string{"*"}
	corConfig.AllowBrowserExtensions = true
	corConfig.AllowMethods = []string{"*"}
	router.Use(cors.New(corConfig))

	jwtHandler := tokens.JWTHandler{
		SigninKey: opt.Cfg.SigningKey,
		Log:       opt.Log,
	}

	handlerV1 := v1.New(&v1.HandlerV1Option{
		Cfg:        &opt.Cfg,
		Storage:    opt.Storage,
		Log:        opt.Log,
		JwtHandler: jwtHandler,
	})

	router.Use(middleware.NewAuth(opt.CasbinEnforcer, jwtHandler, config.Load()))

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(h.StatusOK, gin.H{
			"message": "Server is running!!!",
		})
	})

	router.MaxMultipartMemory = 8 << 20 // 8 Mib

	api := router.Group("/v1")

	// file-upload api
	api.POST("/media/file-upload", handlerV1.FileUpload)
	// admin apis
	api.POST("/admin", handlerV1.RegisterAdmin)
	api.POST("/admin/login", handlerV1.LoginAdmin)
	api.GET("/admin/:id", handlerV1.GetAdmin)
	api.GET("/admins", handlerV1.GetAllAdmin)

	//doctor apis
	api.POST("/doctor", handlerV1.RegisterDoctor)
	api.POST("/doctor/login", handlerV1.LoginDoctor)
	api.GET("/doctor/:id", handlerV1.GetDoctor)
	api.GET("/doctors/search/:clinic_id", handlerV1.GetAllDoctor)
	api.GET("/doctors/:clinic_id", handlerV1.GetDoctorsList)
	api.PUT("/doctors", handlerV1.UpdateDoctor)
	api.DELETE("/doctors/delete/:id", handlerV1.DeleteDoctor)
	api.GET("/doctor/service/:clinic_id", handlerV1.GetDoctorsByService)
	api.GET("/doctors/service/:service_id", handlerV1.GetDoctorsByServiceId)
	api.GET("/doctor/customers/:doctor_id", handlerV1.GetCustomersByDoctor)

	// korik apis
	api.POST("/korik", handlerV1.CreateKorik)
	api.GET("/korik/:id", handlerV1.GetKorik)
	api.GET("/korik", handlerV1.GetAllKoriks)
	api.GET("/korik/user/:id", handlerV1.GetKorikByUserId)
	api.PUT("/korik", handlerV1.UpdateKorik)
	api.DELETE("/korik/:id", handlerV1.DeleteKorik)

	// clinic apis
	api.POST("/clinic", handlerV1.CreateClinic)
	api.GET("/clinic/:id", handlerV1.GetClinic)
	api.GET("/clinics", handlerV1.GetClinicsList)
	api.PUT("/clinics", handlerV1.UpdateClinics)
	api.DELETE("/clinics/delete/:id", handlerV1.DeleteClinic)
	//services apis
	api.POST("/service", handlerV1.CreateServices)
	api.GET("/service/:id", handlerV1.GetServices)
	api.GET("/services", handlerV1.GetAllServices)
	api.PUT("/service", handlerV1.UpdateServices)
	api.DELETE("/service/delete/:id", handlerV1.DeleteServices)

	// client apis
	api.POST("/client/login", handlerV1.LoginClient)
	api.POST("/client", handlerV1.CreateClient)
	api.GET("/client/:id", handlerV1.GetClient)
	api.GET("/clients", handlerV1.GetClients)
	api.PUT("/client", handlerV1.UpdateClient)
	api.DELETE("/client/delete/:id", handlerV1.DeleteClient)

	// queue
	api.POST("/queue", handlerV1.CreateQueue)
	api.DELETE("/queue/cancel", handlerV1.CancelQueue)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
