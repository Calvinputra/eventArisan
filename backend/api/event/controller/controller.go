package controller

import (
	EventInterface "event/backend/api/event/interfaces"
	EventModel "event/backend/api/event/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"event/backend/constants"
	"event/backend/config"
)

type EventController struct {
	Log     *logrus.Logger
	ServiceEvent EventInterface.EventInterface
	ResponseParameter *config.ResponseParameter
}


func NewEventController(
	appRouter *gin.RouterGroup,
	log *logrus.Logger,
	service EventInterface.EventInterface,
	responseParameter *config.ResponseParameter,
) {
	controller := EventController{log, service, responseParameter}

	appRoute := appRouter.Group("/event")
	appRoute.POST("/register", controller.RegisterEvent)
	appRoute.PUT("/", controller.UpdateEvent)
	appRoute.DELETE("/:recid", controller.DeleteEvent)
	appRoute.GET("/list/:type", controller.ListEvent)
}

func (c *EventController) RegisterEvent(ctx *gin.Context) {
	var req EventModel.CreateEventRequest
	if err := ctx.ShouldBind(&req); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		response := c.ResponseParameter.SetResponse(constants.ResponseErrorBadRequest, nil, nil, nil)
		ctx.JSON(response.HttpCode, response)
		return
	}

	inputter := "System"

	responseFromService := c.ServiceEvent.CreateEvent(&req, inputter)
	ctx.JSON(responseFromService.HttpCode, responseFromService)
}

func (c *EventController) UpdateEvent(ctx *gin.Context) {
	var req EventModel.UpdateEventRequest
	if err := ctx.ShouldBind(&req); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		response := c.ResponseParameter.SetResponse(constants.ResponseErrorBadRequest, nil, nil, nil)
		ctx.JSON(response.HttpCode, response)
		return
	}
	inputter := "System"

	responseFromService := c.ServiceEvent.UpdateEvent(&req, inputter)
	ctx.JSON(responseFromService.HttpCode, responseFromService)
}

func (c *EventController) DeleteEvent(ctx *gin.Context) {
	recid := ctx.Param("recid")

	inputter := "System"
	
	responseFromService := c.ServiceEvent.DeleteEvent(recid, inputter)
	ctx.JSON(responseFromService.HttpCode, responseFromService)
}

func (c *EventController) ListEvent(ctx *gin.Context) {
	typeEvent:= ctx.Param("type")

	responseFromService := c.ServiceEvent.ListEvent(typeEvent)
	ctx.JSON(responseFromService.HttpCode, responseFromService)

}