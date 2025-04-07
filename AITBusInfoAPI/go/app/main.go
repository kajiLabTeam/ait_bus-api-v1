package main

import (
	controller "AITBusInfo/controller"
	"AITBusInfo/docs"
	"os"
	"time"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	SetUpServer().Run(":" + port)
}

func SetUpServer() *gin.Engine {

	engine := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	// ミドルウェア
	// engine.Use(middleware.RecordUaAndTime)
	// CRUD 書籍
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*", "http://localhost:8080", "https://localhost:8080", "https://bus-api.bigbell.dev"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	versionEngine := engine.Group("api/v1")
	{
		// @title APIドキュメントのタイトル
		// @version バージョン(1.0)
		// @description 仕様書に関する内容説明
		// @termsOfService 仕様書使用する際の注意事項

		// @contact.name APIサポーター
		// @contact.url http://www.swagger.io/support
		// @contact.email support@swagger.io

		// @license.name ライセンス(必須)
		// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

		// @host bus-api.bigbell.dev
		// @BasePath /
		// http://localhost:8080/api/v1/swagger/index.html
		versionEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		versionEngine.GET("/nextbus", controller.GetNextTime)
		versionEngine.GET("/mode", controller.GetOperationMode)
		versionEngine.GET("/modeall", controller.GetOperationModeAll)
		versionEngine.GET("/timetable", controller.GetTimeTable)
		versionEngine.GET("/timetableall", controller.GetTimeTableAll)
		// versionEngine.GET("/test/nextbus", controller.TestGetNextTime)
		// versionEngine.POST("/stayers", controller.Beacon)
	}

	versionEngineV2 := engine.Group("api/v2")
	{
		versionEngineV2.GET("/nextbus", controller.GetNextTimeV2)
		// versionEngineV2.GET("/test/nextbus", controller.TestGetNextTime2)
	}
	return engine
}
