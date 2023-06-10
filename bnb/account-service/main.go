package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/account-service/controllers"
	grpcserver "github.com/rruzicic/globetrotter/bnb/account-service/grpc_server"
	"github.com/rruzicic/globetrotter/bnb/account-service/jwt"
	"github.com/rruzicic/globetrotter/bnb/account-service/middlewares"
	"github.com/rruzicic/globetrotter/bnb/account-service/repos"
)

const serviceName = "account-service"

func main() {
	repos.Connect()
	go ginSetup()
	grpcserver.InitServer()
	repos.Disconnect()
}

func ginSetup() {
	gin.DisableConsoleColor()
	f, _ := os.Create("log/" + serviceName + ".log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.CORSMiddleware())
	r.NoRoute()

	public := r.Group("/user")

	protected := r.Group("/user")
	protected.Use(jwt.AnyUserAuthMiddleware())

	public.GET("/health", controllers.HealthCheck)
	public.GET("/all", controllers.GetAll)
	public.GET("/id/:id", controllers.GetById)
	public.GET("/email/:email", controllers.GetByEmail)

	public.POST("/register/host", controllers.RegisterHost)
	public.POST("/register/guest", controllers.RegisterGuest)
	protected.POST("/update", controllers.UpdateUser)
	public.POST("/login", controllers.Login)

	protected.DELETE("/delete/:id", controllers.DeleteUser)
	r.Run(":8080")
	log.Println("HTTP server running on port 8080")
}

func logOutput() func() {
	logfile := `app.log`
	// open file read/write | create if not exist | clear file at open if exists
	f, _ := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)

	// save existing stdout | MultiWriter writes to saved stdout and file
	out := os.Stdout
	mw := io.MultiWriter(out, f)

	// get pipe reader and writer | writes to pipe writer come out pipe reader
	r, w, _ := os.Pipe()

	// replace stdout,stderr with pipe writer | all writes to stdout, stderr will go through pipe instead (fmt.print, log)
	os.Stdout = w
	os.Stderr = w

	// writes with log.Print should also write to mw
	log.SetOutput(mw)

	//create channel to control exit | will block until all copies are finished
	exit := make(chan bool)

	go func() {
		// copy all reads from pipe to multiwriter, which writes to stdout and file
		_, _ = io.Copy(mw, r)
		// when r or w is closed copy will finish and true will be sent to channel
		exit <- true
	}()

	// function to be deferred in main until program exits
	return func() {
		// close writer then block on exit channel | this will let mw finish writing before the program exits
		_ = w.Close()
		<-exit
		// close file after all writes have finished
		_ = f.Close()
	}

}
