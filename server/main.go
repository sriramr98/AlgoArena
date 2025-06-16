package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sriramr98/dsa_server/controllers"
)

// Not using router.Group because of https://github.com/gin-gonic/gin/issues/3546
func SetupProblemRoutes(router *gin.Engine) {
	pc := controllers.ProblemController{}
	router.GET("/api/problems", pc.GetProblems)
	router.GET("/api/problems/:id", pc.GetProblemDetails)
	router.GET("/api/problems/:id/stub/:language", pc.GetProblemStub)
	router.POST("/api/problems/:id/submit", pc.SubmitProblem)
	router.GET("/api/problems/:id/testcases", pc.GetProblemTestCases)
}

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	SetupProblemRoutes(router)

	return router
}

func main() {
	router := GetRouter()

	srv := &http.Server{
		Addr:    ":5000",
		Handler: router.Handler(),
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no params) by default sends syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
}
