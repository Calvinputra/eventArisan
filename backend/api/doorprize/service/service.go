package service

import (
	"event/backend/api/doorprize/entity"
	"event/backend/api/doorprize/model"
	doorprizeRepository "event/backend/api/doorprize/repository"
	"event/backend/config"
	"event/backend/constants"
	"event/backend/helper"
	basemodel "event/backend/model"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gitlab.universedigital.my.id/library/golang/crud/crud"
	"gorm.io/gorm"
)

type DoorprizeService struct {
    DB                        *gorm.DB
	Log                       *logrus.Logger
	Validate                  *validator.Validate
	CustomValidation          *config.CustomValidation
	ResponseParameter         *config.ResponseParameter
    DoorprizeRepository           *doorprizeRepository.DoorprizeRepository
	crudLib                   crud.LiveHisNau[entity.Doorprize]

	CurrentEventRecid         string
}

func NewDoorprizeService(
	db *gorm.DB,
	log *logrus.Logger,
	validate *validator.Validate,
	customValidation *config.CustomValidation,
	responseParameter *config.ResponseParameter,
	doorprizeRepository *doorprizeRepository.DoorprizeRepository,
) *DoorprizeService {
	crudLib := crud.LiveHisNau[entity.Doorprize]{
		LiveTable: entity.Doorprize{}.TableName(),
		HisTable:  entity.DoorprizeHis{}.TableName(),
		NauTable: "",
		NauCount:  0,
		DB:        db,
	}

	return &DoorprizeService{
		DB:                        db,
		Log:                       log,
		Validate:                  validate,
		CustomValidation:          customValidation,
		ResponseParameter:         responseParameter,
		DoorprizeRepository:        doorprizeRepository,
		crudLib:                   crudLib,
	}
}


func (s *DoorprizeService) CreateDoorprize (req *model.CreateDoorprizeRequest, inputter string) basemodel.Response {
	doorprizeTemp := req.ToEntity(inputter, constants.INSERT)

	tx := s.DB.Begin()
	defer helper.RollbackHelper(tx)

	crudResponse, doorprize := s.crudLib.WithUuidAsRecid().Create(tx, doorprizeTemp, inputter)
	helper.ThrowWithoutMessage(crudResponse.Err)

	tx.Commit()

	response := model.DoorprizeResponse{}.ToResponse(&doorprize)
	return s.ResponseParameter.SetResponse(constants.ResponseSuccessInsertRecord, nil, response, nil)
}


func (s *DoorprizeService) ListDoorprize(
    typeDoorprize string,
) basemodel.Response {

    if typeDoorprize == "" || typeDoorprize != "LIVE" {
        return s.ResponseParameter.SetResponse(constants.ResponseErrorInvalidTypeParameter, nil, nil, nil)
    }

    if s.CurrentEventRecid == "" {
        return s.ResponseParameter.SetResponse(constants.ResponseErrorBadRequest, nil, nil, nil)
    }

    tx := s.DB.Begin()
    defer helper.RollbackHelper(tx)

    var doorprize []entity.Doorprize
    err := tx.
        Preload("Attendance").
        Where("event_recid = ?", s.CurrentEventRecid).
        Find(&doorprize).
        Error

    helper.ThrowWithoutMessage(err)

    tx.Commit()

    response := model.DoorprizeResponse{}.ToResponseList(doorprize)
    return s.ResponseParameter.SetResponse(constants.ResponseSuccessListRecord, nil, response, nil)
}
