package router

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	_ "gin-vue/docs"
	"gin-vue/global"
	"gin-vue/middleware"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type IFnRegistRoute = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

var (
	gfnRoutes []IFnRegistRoute
)

func RegistRoute(fn IFnRegistRoute) {
	if fn == nil {
		return
	}
	gfnRoutes = append(gfnRoutes, fn)
}

func InitRounter() {
	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	r := gin.Default()
	r.Use(middleware.Cors())
	r.Use(middleware.NoCacheMiddleware())

	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1")

	initBasePlatformRoutes()

	// register custom validator
	registCustValidator()

	for _, fnRegistRoute := range gfnRoutes {
		fnRegistRoute(rgPublic, rgAuth)
	}

	// integrate swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	stPort := viper.GetString("server.port")
	if stPort == "" {
		stPort = "8999"
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", stPort),
		Handler: r,
	}

	// start server
	go func() {
		global.Logger.Infof("Start Server: http://127.0.0.1:%s", stPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Errorf("Start Server Error: %s", err.Error())
			return
		}
	}()

	<-ctx.Done()

	ctx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Errorf("Shutdown Server Error: %s", err.Error())
		return
	}
	global.Logger.Info("Shutdown Server Success")
	fmt.Println("Shutdown Server Success")

}

func initBasePlatformRoutes() {
	InitUserRoutes()
	InitHostRoutes()
}

// ! register custom validator
func registCustValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("first_is_a", func(fl validator.FieldLevel) bool {
			if val, ok := fl.Field().Interface().(string); ok {
				// first letter must be a
				return val != "" && val[0] == 'a'
			}
			return false
		})
	}
}
