package controller

import (
	AttendanceInterface "event/backend/api/attendance/interfaces"
	AttendanceModel "event/backend/api/attendance/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"event/backend/config"
	"event/backend/constants"
	"net/http"
)

type AttendanceController struct {
	Log     *logrus.Logger
	ServiceAttendance AttendanceInterface.AttendanceInterface
	ResponseParameter *config.ResponseParameter
}


func NewAttendanceController(
	appRouter *gin.RouterGroup,
	log *logrus.Logger,
	service AttendanceInterface.AttendanceInterface,
	responseParameter *config.ResponseParameter,
) {
	controller := AttendanceController{log, service, responseParameter}
	appRoute := appRouter.Group("/attendance")

	appRoute.POST("/register", controller.CreateAttendance)
	appRoute.PUT("/", controller.UpdateAttendance)
	appRoute.DELETE("/:recid", controller.DeleteAttendance)
	appRoute.GET("/list/:type", controller.ListAttendance)
	appRoute.GET("/:type/:recid", controller.GetAttendance)
	
	appRoute.POST("/import", controller.ImportAttendance)
	appRoute.POST("/scan", controller.ScanAttendance)
}

func (c *AttendanceController) CreateAttendance(ctx *gin.Context) {
	var req AttendanceModel.CreateAttendanceRequest
	if err := ctx.ShouldBind(&req); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		response := c.ResponseParameter.SetResponse(constants.ResponseErrorBadRequest, nil, nil, nil)
		ctx.JSON(response.HttpCode, response)
		return
	}
	inputter := "System"
	
	responseFromService := c.ServiceAttendance.CreateAttendance(&req, inputter)
	ctx.JSON(responseFromService.HttpCode, responseFromService)
}

func (c *AttendanceController) UpdateAttendance(ctx *gin.Context) {
	var req AttendanceModel.UpdateAttendanceRequest
	if err := ctx.ShouldBind(&req); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		response := c.ResponseParameter.SetResponse(constants.ResponseErrorBadRequest, nil, nil, nil)
		ctx.JSON(response.HttpCode, response)
		return
	}

	inputter := "System"
	responseFromService := c.ServiceAttendance.UpdateAttendance(&req, inputter)
	ctx.JSON(responseFromService.HttpCode, responseFromService)
}

func (c *AttendanceController) DeleteAttendance(ctx *gin.Context) {
	recid := ctx.Param("recid")

	inputter := "System"
	
	responseFromService := c.ServiceAttendance.DeleteAttendance(recid, inputter)
	ctx.JSON(responseFromService.HttpCode, responseFromService)
}

func (c *AttendanceController) ImportAttendance(ctx *gin.Context) {
    eventRecid := ctx.PostForm("event_recid")
    if eventRecid == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "message": "event_recid wajib diisi",
        })
        return
    }

    fileHeader, err := ctx.FormFile("file")
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "message": "file wajib diupload",
        })
        return
    }

    inputter := "System"

    resp := c.ServiceAttendance.ImportAttendance(eventRecid, fileHeader, inputter)

    ctx.JSON(resp.HttpCode, resp)
}

func (c *AttendanceController) ScanAttendance(ctx *gin.Context) {
    var req AttendanceModel.ScanAttendanceRequest
    if err := ctx.ShouldBind(&req); err != nil {
        c.Log.Warnf("Failed to parse request body : %+v", err)
        response := c.ResponseParameter.SetResponse(constants.ResponseErrorBadRequest, nil, nil, nil)
        ctx.JSON(response.HttpCode, response)
        return
    }

    inputter := "System"
    resp := c.ServiceAttendance.ScanAttendance(&req, inputter)
    ctx.JSON(resp.HttpCode, resp)
}

func (c *AttendanceController) ListAttendance(ctx *gin.Context) {
    typeAttendance:= ctx.Param("type")

    responseFromService := c.ServiceAttendance.ListAttendance(typeAttendance)
    ctx.JSON(responseFromService.HttpCode, responseFromService)
}	

func (c *AttendanceController) GetAttendance(ctx *gin.Context) {
    typeAttendance:= ctx.Param("type")
    recid := ctx.Param("recid")

    responseFromService := c.ServiceAttendance.GetAttendance(typeAttendance, recid)
    ctx.JSON(responseFromService.HttpCode, responseFromService)
}