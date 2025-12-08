package interfaces

import (
    "event/backend/api/doorprize/model"
	basemodel "event/backend/model"

)

type DoorprizeInterface interface {
	CreateDoorprize(req *model.CreateDoorprizeRequest, inputter string) basemodel.Response
	ListDoorprize(typeDoorprize string) basemodel.Response
}