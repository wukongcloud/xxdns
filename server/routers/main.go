package routers

import (
	// import need sorted
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/wukongcloud/xxdns/server/controllers"
	"github.com/wukongcloud/xxdns/server/docs"
	//"github.com/wukongcloud/xxdns/middleware"
	"time"
)

func init() {
	docs.SwaggerInfo.Title = "DNS服务"
	docs.SwaggerInfo.Description = "DNS服务后端API接口文档"
	docs.SwaggerInfo.Version = "2.0"
	docs.SwaggerInfo.Host = "127.0.0.1:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}
}

func SetupRouter() *gin.Engine {
	router := gin.New()
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf(`{"time":"%s","dest_ip":"%s","http_method":"%s","uri_path":"%s","proto":"%s","status":%d,"response_time":"%s","http_user_agent":"%s","bytes_in":%d,"errmsg":"%s"}%v`,
			param.TimeStamp.Format(time.UnixDate),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.BodySize,
			param.ErrorMessage,
			"\r\n",
		)
	}))
	router.Use(gin.Recovery())

	// health check url
	router.GET("/healthz", controllers.Health)

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// api v1 group
	v1 := router.Group("/api/v1")
	{
		domains := v1.Group("domains")
		//domain.Use(middleware.AuthRequired)
		{
			domains.GET("", controllers.GetDomains)
			domains.GET(":id", controllers.GetDomainById)
			domains.POST("", controllers.AddDomain)
			domains.PUT(":id", controllers.UpdateDomainById)
			domains.DELETE(":id", controllers.DeleteDomain)
		}
		views := v1.Group("views")
		{
			views.GET("", controllers.GetViews)
			views.GET(":id", controllers.GetViewById)
			views.POST("", controllers.AddView)
			views.PUT(":id", controllers.UpdateViewById)
			views.DELETE(":id", controllers.DeleteView)
		}
		acls := v1.Group("acls")
		{
			acls.GET("", controllers.GetAcls)
			acls.GET(":id", controllers.GetAclById)
			acls.POST("", controllers.AddAcl)
			acls.PUT(":id", controllers.UpdateAclById)
			acls.DELETE(":id", controllers.DeleteAcl)
		}
		records := v1.Group("records")
		{
			records.GET("", controllers.GetRecords)
			records.GET(":id", controllers.GetRecordById)
			records.POST("", controllers.AddRecord)
			records.PUT(":id", controllers.UpdateRecordById)
			records.DELETE(":id", controllers.DeleteRecord)
		}
		ipdb := v1.Group("ipdb")
		{
			ipdb.POST("", controllers.AddIPDB)
			ipdb.GET("", controllers.GetIPDB)
		}
		//auth := v1.Group("auth")
		//{
		//	auth.POST("login", controllers.Login)
		//	auth.GET("logout", controllers.Login) //todo
		//}
	}

	return router
}
