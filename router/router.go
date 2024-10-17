package router

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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

	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1")

	InitBasePlatformRoutes()

	for _, fnRegistRoute := range gfnRoutes {
		fnRegistRoute(rgPublic, rgAuth)
	}

	stPort := viper.GetString("server.port")
	if stPort == "" {
		stPort = "8999"
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", stPort),
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// todo: log error
			fmt.Println(fmt.Sprintf("Start Server Error: %s", err.Error()))
			return
		}
		fmt.Println(fmt.Sprintf("Start Server Listen: %s", stPort))
	}()

	<-ctx.Done()

	ctx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := server.Shutdown(ctx); err != nil {
		// todo: log error
		fmt.Println(fmt.Sprintf("Shutdown Server Error: %s", err.Error()))
		return
	}

	fmt.Println("Shutdown Server Success")

}

func InitBasePlatformRoutes() {
	InitUserRoutes()
}
