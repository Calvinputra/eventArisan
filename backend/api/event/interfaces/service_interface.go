package interfaces

import (
    "event/backend/api/event/model"
	basemodel "event/backend/model"
)

type EventInterface interface {
    CreateEvent(input *model.CreateEventRequest, inputter string) basemodel.Response
    UpdateEvent(input *model.UpdateEventRequest, inputter string) basemodel.Response
    DeleteEvent(recid string, inputter string) basemodel.Response
	ListEvent(typeEvent string) basemodel.Response
    GetEvent(recid string, inputter string) basemodel.Response
}
