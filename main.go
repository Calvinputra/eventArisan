package main

import(
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"event/backend/helper"
	"event/backend/config"
	// event
	eventRepository "event/backend/api/event/repository"
	eventService "event/backend/api/event/service"
	eventController "event/backend/api/event/controller"
	// attendance
	attendanceRepository "event/backend/api/attendance/repository"
	attendanceService "event/backend/api/attendance/service"
	attendanceController "event/backend/api/attendance/controller"
	
	"event/backend/migrate"
)


func main() {
	// router
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"POST", "PUT", "DELETE", "GET", "PATCH", "OPTION"},
		AllowHeaders:  []string{"Origin", "Content-Length", "Content-TypeRecid", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders: []string{"Content-Disposition", "Pagination-Limit", "Pagination-Max-Page", "Pagination-Page", "Pagination-Total-Data", "Access-Control-Allow-Origin"},
	}))

	// config
	helper.InitEnv()
	db := config.NewConnection()
	log := config.NewLog()
	validate := config.NewValidator()
	rp := config.NewResponseParameter(helper.GetStringEnv("RESPONSE_PARAMETER_URL"), log)
	cv := config.NewCustomValidation(rp)
	// migrate
	if err := migrate.AutoMigrate(db); err != nil {
		log.WithError(err).Fatal("failed to migrate database")
		return
	}
	// routing
	apiRouter := router.Group("/api")
	appRouter := apiRouter.Group("/app")

	// Repository
	eventRepository := eventRepository.NewEventRepository(db,log)
	attendanceRepository := attendanceRepository.NewAttendanceRepository(log, db)

	// Service
	eventService := eventService.NewEventService(db, log, validate, cv, rp, eventRepository)
	attendanceService := attendanceService.NewAttendanceService(db, log, validate, cv, rp, attendanceRepository)

	// Controller
	eventController.NewEventController(appRouter, log, eventService, rp)
	attendanceController.NewAttendanceController(appRouter, log, attendanceService, rp)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}