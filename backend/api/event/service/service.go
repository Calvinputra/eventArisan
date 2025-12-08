package service

import (
	"event/backend/api/event/entity"
	"event/backend/api/event/model"
	basemodel "event/backend/model"
	eventRepository "event/backend/api/event/repository"
	"gorm.io/gorm"
    "event/backend/config"
	"gitlab.universedigital.my.id/library/golang/crud/crud"
    "github.com/sirupsen/logrus"
    "github.com/go-playground/validator/v10"
    "event/backend/constants"
    "event/backend/helper"
)

type EventService struct {
    DB                        *gorm.DB
	Log                       *logrus.Logger
	Validate                  *validator.Validate
	CustomValidation          *config.CustomValidation
	ResponseParameter         *config.ResponseParameter
    EventRepository           *eventRepository.EventRepository
	crudLib                   crud.LiveHisNau[entity.Event]
}

func NewEventService(
	db *gorm.DB,
	log *logrus.Logger,
	validate *validator.Validate,
	customValidation *config.CustomValidation,
	responseParameter *config.ResponseParameter,
	eventRepository *eventRepository.EventRepository,
) *EventService {
	crudLib := crud.LiveHisNau[entity.Event]{
		LiveTable: entity.Event{}.TableName(),
		HisTable:  entity.EventHis{}.TableName(),
		NauTable:  "",
		NauCount:  0,
		DB:        db,
	}

	return &EventService{
		DB:                        db,
		Log:                       log,
		Validate:                  validate,
		CustomValidation:          customValidation,
		ResponseParameter:         responseParameter,
		EventRepository:           eventRepository,
		crudLib:                   crudLib,
	}
}


func (s *EventService) CreateEvent (req *model.CreateEventRequest, inputter string) basemodel.Response {
	eventTemp := req.ToEntity(inputter, constants.INSERT)

	tx := s.DB.Begin()
	defer helper.RollbackHelper(tx)

	crudResponse, event := s.crudLib.WithUuidAsRecid().Create(tx, eventTemp, inputter)
	helper.ThrowWithoutMessage(crudResponse.Err)

	tx.Commit()

	response := model.EventResponse{}.ToResponse(&event)
	return s.ResponseParameter.SetResponse(constants.ResponseSuccessInsertRecord, nil, response, nil)
}

func (s *EventService) UpdateEvent(req *model.UpdateEventRequest, inputter string) basemodel.Response {
	crudResponse, currentEvent := s.crudLib.Get(constants.RecordStatusParameterLIVE, req.Recid)
	if crudResponse.Err != nil {
		return s.ResponseParameter.SetResponse(crudResponse.Message, nil, nil, nil)
	}

	eventTemp := req.ToEntity(currentEvent, constants.UPDATE)

	tx := s.DB.Begin()
	defer helper.RollbackHelper(tx)

	crudResponse, eventLive := s.crudLib.Update(tx, eventTemp, inputter)
	helper.ThrowWithoutMessage(crudResponse.Err)

	tx.Commit()

	response := model.EventResponse{}.ToResponse(&eventLive)
	return s.ResponseParameter.SetResponse(constants.ResponseSuccessUpdateRecord, nil, response, nil)
}

func (s *EventService) DeleteEvent(recid string, inputter string) basemodel.Response {
	// get current live
	crudResponse, currentEvent := s.crudLib.Get(constants.RecordStatusParameterLIVE, recid)
	if crudResponse.Err != nil {
		return s.ResponseParameter.SetResponse(crudResponse.Message, nil, nil, nil)
	}

	// not found
	if currentEvent == nil {
		return s.ResponseParameter.SetResponse(constants.ResponseErrorNotFound, nil, nil, nil)
	}

	tx := s.DB.Begin()
	defer helper.RollbackHelper(tx)

	// execute delete
	crudResponse, eventDeleted := s.crudLib.DeleteLive(tx, recid, inputter)
	helper.ThrowWithoutMessage(crudResponse.Err)

	// update status
	eventDeleted.Status = constants.DELETED

	tx.Commit()
	// return
	response := model.EventResponse{}.ToResponse(&eventDeleted)

	return s.ResponseParameter.SetResponse(constants.ResponseSuccessDeleteRecord, nil, response, nil)

}

func (s *EventService) ListEvent(userType string) basemodel.Response {
    users, err := s.EventRepository.FindAllByType(userType)
    if err != nil {
        return s.ResponseParameter.SetResponse(constants.ResponseErrorNotFound, nil, nil, nil)
    }

    events, ok := users.([]entity.Event)
    if !ok {
        return s.ResponseParameter.SetResponse(constants.ResponseErrorInvalidTypeParameter, nil, nil, nil)
    }

    response := model.EventResponse{}.ToResponseList(events)
    return s.ResponseParameter.SetResponse(constants.ResponseSuccessListRecord, nil, response, nil)
}

func (s *EventService) GetEvent(recid string, typeEvent string) basemodel.Response {
    event, err := s.EventRepository.FindByIdAndType(s.DB, recid, typeEvent)
    if err != nil {
        return s.ResponseParameter.SetResponse(constants.ResponseErrorNotFound, nil, nil, nil)
    }

    response := model.EventResponse{}.ToResponse(event)
    return s.ResponseParameter.SetResponse(constants.ResponseSuccessGetRecord, nil, response, nil)
}