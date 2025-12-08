package service

import (
	"event/backend/api/attendance/entity"
	"event/backend/api/attendance/model"
	attendanceRepository "event/backend/api/attendance/repository"
	"event/backend/config"
	"event/backend/constants"
	"event/backend/helper"
	basemodel "event/backend/model"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
	"gitlab.universedigital.my.id/library/golang/crud/crud"
	"gorm.io/gorm"
)

type AttendanceService struct {
    DB                        *gorm.DB
	Log                       *logrus.Logger
	Validate                  *validator.Validate
	CustomValidation          *config.CustomValidation
	ResponseParameter         *config.ResponseParameter
    AttendanceRepository           *attendanceRepository.AttendanceRepository
	crudLib                   crud.LiveHisNau[entity.Attendance]
}

func NewAttendanceService(
	db *gorm.DB,
	log *logrus.Logger,
	validate *validator.Validate,
	customValidation *config.CustomValidation,
	responseParameter *config.ResponseParameter,
	attendanceRepository *attendanceRepository.AttendanceRepository,
) *AttendanceService {
	crudLib := crud.LiveHisNau[entity.Attendance]{
		LiveTable: entity.Attendance{}.TableName(),
		HisTable:  entity.AttendanceHis{}.TableName(),
		NauTable:  "",
		NauCount:  0,
		DB:        db,
	}

	return &AttendanceService{
		DB:                        db,
		Log:                       log,
		Validate:                  validate,
		CustomValidation:          customValidation,
		ResponseParameter:         responseParameter,
		AttendanceRepository:          attendanceRepository,
		crudLib:                   crudLib,
	}
}

func (s *AttendanceService) CreateAttendance (req *model.CreateAttendanceRequest, inputter string) basemodel.Response {
	attendanceTemp := req.ToEntity(inputter, constants.INSERT)

	tx := s.DB.Begin()
	defer helper.RollbackHelper(tx)

	crudResponse, attendance := s.crudLib.WithUuidAsRecid().Create(tx, attendanceTemp, inputter)
	helper.ThrowWithoutMessage(crudResponse.Err)

	tx.Commit()

	response := model.AttendanceResponse{}.ToResponse(&attendance)
	return s.ResponseParameter.SetResponse(constants.ResponseSuccessInsertRecord, nil, response, nil)
}

func (s *AttendanceService) UpdateAttendance(req *model.UpdateAttendanceRequest, inputter string) basemodel.Response {
	crudResponse, currentAttendance := s.crudLib.Get(constants.RecordStatusParameterLIVE, req.Recid)
	if crudResponse.Err != nil {
		return s.ResponseParameter.SetResponse(crudResponse.Message, nil, nil, nil)
	}

	attendanceTemp := req.ToEntity(currentAttendance, constants.UPDATE)

	tx := s.DB.Begin()
	defer helper.RollbackHelper(tx)

	crudResponse, attendanceLive := s.crudLib.Update(tx, attendanceTemp, inputter)
	helper.ThrowWithoutMessage(crudResponse.Err)

	tx.Commit()

	response := model.AttendanceResponse{}.ToResponse(&attendanceLive)
	return s.ResponseParameter.SetResponse(constants.ResponseSuccessUpdateRecord, nil, response, nil)
}

func (s *AttendanceService) DeleteAttendance(recid string, inputter string) basemodel.Response {
	// get current live
	crudResponse, currentAttendance := s.crudLib.Get(constants.RecordStatusParameterLIVE, recid)
	if crudResponse.Err != nil {
		return s.ResponseParameter.SetResponse(crudResponse.Message, nil, nil, nil)
	}

	// not found
	if currentAttendance == nil {
		return s.ResponseParameter.SetResponse(constants.ResponseErrorNotFound, nil, nil, nil)
	}

	tx := s.DB.Begin()
	defer helper.RollbackHelper(tx)

	// execute delete
	crudResponse, attendanceDeleted := s.crudLib.DeleteLive(tx, recid, inputter)
	helper.ThrowWithoutMessage(crudResponse.Err)

	tx.Commit()
	// return
	response := model.AttendanceResponse{}.ToResponse(&attendanceDeleted)

	return s.ResponseParameter.SetResponse(constants.ResponseSuccessDeleteRecord, nil, response, nil)

}

func (s *AttendanceService) ImportAttendance(
    eventRecid string,
    fileHeader *multipart.FileHeader,
    inputter string,
) basemodel.Response {

    if eventRecid == "" || fileHeader == nil {
        return s.ResponseParameter.SetResponse(constants.ResponseErrorInvalidTypeParameter, nil, nil, nil)
    }

    file, err := fileHeader.Open()
    if err != nil {
        s.Log.WithError(err).Error("open upload file gagal")
        return s.ResponseParameter.SetResponse("open upload file gagal", nil, nil, nil)
    }
    defer file.Close()

    excelFile, err := excelize.OpenReader(file)
    if err != nil {
        s.Log.WithError(err).Error("buka excel gagal")
        return s.ResponseParameter.SetResponse("buka excel gagal", nil, nil, nil)
    }

    sheetName := excelFile.GetSheetName(0)
    if sheetName == "" {
        s.Log.Error("sheet pertama tidak ditemukan")
        return s.ResponseParameter.SetResponse("sheet pertama tidak ditemukan", nil, nil, nil)
    }

    rows, err := excelFile.GetRows(sheetName)
    if err != nil {
        s.Log.WithError(err).Error("get rows excel gagal")
        return s.ResponseParameter.SetResponse("get rows excel gagal", nil, nil, nil)
    }

    var list []entity.Attendance

    for i, row := range rows {
        if i == 0 { 
            continue
        }
        if len(row) == 0 {
            continue
        }

        name := ""
        code := ""
        var noTable int64

        if len(row) > 0 {
            name = row[0]
        }
        if len(row) > 1 {
            code = row[1]
        }
        if len(row) > 2 {
            if v, err := strconv.ParseInt(row[2], 10, 64); err == nil {
                noTable = v
            }
        }

        if name == "" && code == "" {
            continue
        }

        list = append(list, entity.Attendance{
            EventRecid:     eventRecid,
            Name:           name,
            Code:           code,
            NoTable:        noTable,
            StatusCheckin:  0,
            StatusSouvenir: 0,
            CheckinTime:    0,
            SouvenirTime:    0,
        })
    }

    if len(list) == 0 {
        s.Log.Warn("import attendance: list kosong")
        return s.ResponseParameter.SetResponse("import attendance: list kosong", nil, nil, nil)
    }

    tx := s.DB.Begin()
    defer helper.RollbackHelper(tx)
	for i := range list {
		list[i].Recid = uuid.NewString()
		list[i].CurrNo = 1
		list[i].RecordStatus = "LIVE"
		list[i].Inputter = inputter
		list[i].InputDateTime = time.Now().Unix()
		list[i].CreatedBy = inputter
		list[i].CreatedDateTime = time.Now().Unix()
	}

    if err := s.AttendanceRepository.BulkCreate(tx, list); err != nil {
        s.Log.WithError(err).Error("bulk insert attendance gagal")
        return s.ResponseParameter.SetResponse("bulk insert attendance gagal", nil, nil, nil)
    }

    tx.Commit()

    return s.ResponseParameter.SetResponse(constants.ResponseSuccessInsertRecord, nil, list, nil)
}

func (s *AttendanceService) ScanAttendance(
    req *model.ScanAttendanceRequest,
    inputter string,
) basemodel.Response {

    if req.EventRecid == "" || req.Code == "" {
        return s.ResponseParameter.SetResponse(
            constants.ResponseErrorInvalidTypeParameter,
            nil, "event_recid dan code wajib diisi", nil,
        )
    }

    tx := s.DB.Begin()
    defer helper.RollbackHelper(tx)

    attendanceData, err := s.AttendanceRepository.GetByEventRecidAndCode(req.EventRecid, req.Code)
    if err != nil {
        return s.ResponseParameter.SetResponse("ResponseErrorInternalServer", nil, nil, nil)
    }
    if attendanceData == nil {
        return s.ResponseParameter.SetResponse(
            "Kode tidak ditemukan / tidak valid",
            nil, "Kode tidak ditemukan / tidak valid", nil,
        )
    }

    if attendanceData.StatusCheckin == 1 && attendanceData.StatusSouvenir == 1 {
        return s.ResponseParameter.SetResponse(
            "Kode sudah digunakan",
            nil, "Peserta sudah check-in dan ambil souvenir, tidak dapat diproses lagi", nil,
        )
    }

    now := time.Now().Unix()

    // Checkin hanya sekali
    if attendanceData.StatusCheckin == 0 {
        attendanceData.StatusCheckin = 1
        attendanceData.CheckinTime = now
    }

    // souvenir jika di centang
    if req.Souvenir {
        if attendanceData.StatusSouvenir == 0 {
            attendanceData.StatusSouvenir = 1
            attendanceData.SouvenirTime = now
        }
    }

    crudResponse, attendanceLive := s.crudLib.Update(tx, attendanceData, inputter)
    if crudResponse.Err != nil {
        return s.ResponseParameter.SetResponse(crudResponse.Message, nil, nil, nil)
    }

    tx.Commit()

    response := model.AttendanceResponse{}.ToResponse(&attendanceLive)
    return s.ResponseParameter.SetResponse(crudResponse.Message, nil, response, nil)
}

func (s *AttendanceService) ListAttendance(
    userType string,
) basemodel.Response {
    users, err := s.AttendanceRepository.FindAllByType(userType)
    if err != nil {
        return s.ResponseParameter.SetResponse(constants.ResponseErrorNotFound, nil, nil, nil)
    }

    attendances, ok := users.([]entity.Attendance)
    if !ok {
        return s.ResponseParameter.SetResponse(constants.ResponseErrorInvalidTypeParameter, nil, nil, nil)
    }

    response := model.AttendanceResponse{}.ToResponseList(attendances)
    return s.ResponseParameter.SetResponse(constants.ResponseSuccessListRecord, nil, response, nil)
}   

func (s *AttendanceService) GetAttendance(userType string, recid string) basemodel.Response {
    attendance, err := s.AttendanceRepository.FindByIdAndType(s.DB, recid, userType)
    if err != nil {
        return s.ResponseParameter.SetResponse(constants.ResponseErrorNotFound, nil, nil, nil)
    }


    response := model.AttendanceResponse{}.ToResponse(attendance)
    return s.ResponseParameter.SetResponse(constants.ResponseSuccessGetRecord, nil, response, nil)
}