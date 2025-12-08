package controller

import (
	DoorprizeInterface "event/backend/api/doorprize/interfaces"
	DoorprizeModel "event/backend/api/doorprize/model"
	"event/backend/config"
	"event/backend/constants"
	"event/backend/api/doorprize/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type DoorprizeController struct {
	Log     *logrus.Logger
	ServiceDoorprize DoorprizeInterface.DoorprizeInterface
	ResponseParameter *config.ResponseParameter
}


func NewDoorprizeController(
	appRouter *gin.RouterGroup,
	log *logrus.Logger,
	service DoorprizeInterface.DoorprizeInterface,
	responseParameter *config.ResponseParameter,
) {
	controller := DoorprizeController{log, service, responseParameter}
	appRoute := appRouter.Group("/doorprize")

	appRoute.POST("/create", controller.CreateDoorprize)
	appRoute.GET("/list/:type", controller.ListDoorprize)
}

func (c *DoorprizeController) CreateDoorprize(ctx *gin.Context) {
	var req DoorprizeModel.CreateDoorprizeRequest
	if err := ctx.ShouldBind(&req); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		response := c.ResponseParameter.SetResponse(constants.ResponseErrorBadRequest, nil, nil, nil)
		ctx.JSON(response.HttpCode, response)
		return
	}
	inputter := "System"

	responseFromService := c.ServiceDoorprize.CreateDoorprize(&req, inputter)
	ctx.JSON(responseFromService.HttpCode, responseFromService)
}

func (c *DoorprizeController) ListDoorprize(ctx *gin.Context) {
    typeDoorprize := ctx.Param("type")
    eventRecid := ctx.Query("event") 

    c.ServiceDoorprize.(*service.DoorprizeService).CurrentEventRecid = eventRecid

    responseFromService := c.ServiceDoorprize.ListDoorprize(typeDoorprize)
    ctx.JSON(responseFromService.HttpCode, responseFromService)
}
