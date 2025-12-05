package config

import (
	"event/backend/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"github.com/sirupsen/logrus"
)

type ResponseData struct {
	Recid    string `json:"recid"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
	HttpCode int    `json:"httpCode"`
}

type ResponseParameter struct {
	Responses             map[string]ResponseData
	ValidatorErrorMapping map[string]string
}

// fallback kalau gagal ambil dari HTTP
func defaultResponseParameter() *ResponseParameter {
	return &ResponseParameter{
		Responses:             make(map[string]ResponseData),
		ValidatorErrorMapping: getValidateMessages(),
	}
}

func NewResponseParameter(url string, log *logrus.Logger) *ResponseParameter {
	log.Infoln("Initialize ResponseData PARAMETER")

	// Kalau URL kosong, langsung pakai default, jangan panic
	if url == "" {
		log.Errorln("[INIT RESPONSE PARAMETER] Empty URL (RESPONSE_PARAMETER_URL kosong), using default")
		return defaultResponseParameter()
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorln("[INIT RESPONSE PARAMETER] error when creating new HTTP Request, using default:", err)
		return defaultResponseParameter()
	}

	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Errorln("[INIT RESPONSE PARAMETER] error request to client, using default:", err)
		return defaultResponseParameter()
	}
	defer response.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Errorln("[INIT RESPONSE PARAMETER] error processing response body from client, using default:", err)
		return defaultResponseParameter()
	}

	// Save the response from API
	var dataJson map[string]interface{}
	if err := json.Unmarshal(respBody, &dataJson); err != nil {
		log.Errorln("[INIT RESPONSE PARAMETER] error json unmarshal response body from client, using default:", err)
		return defaultResponseParameter()
	}

	formattedResponseParameter := make(map[string]ResponseData)

	// pastikan key "data" ada dan bentuknya map
	rawData, ok := dataJson["data"].(map[string]interface{})
	if !ok {
		log.Errorln("[INIT RESPONSE PARAMETER] response body has no valid 'data' field, using default")
		return defaultResponseParameter()
	}

	for key, value := range rawData {
		vMap, ok := value.(map[string]interface{})
		if !ok {
			continue
		}
		code, _ := strconv.Atoi(fmt.Sprint(vMap["code"]))
		httpCode, _ := strconv.Atoi(fmt.Sprint(vMap["httpCode"]))

		formattedResponseParameter[key] = ResponseData{
			Recid:    key,
			Code:     code,
			Message:  fmt.Sprint(vMap["message"]),
			HttpCode: httpCode,
		}
	}

	return &ResponseParameter{
		Responses:             formattedResponseParameter,
		ValidatorErrorMapping: getValidateMessages(),
	}
}

func (c *ResponseParameter) GetResponse(recId string) ResponseData {
	if c.Responses[recId].Message != "" {
		return ResponseData{
			Recid:    c.Responses[recId].Recid,
			Code:     c.Responses[recId].Code,
			Message:  c.Responses[recId].Message,
			HttpCode: c.Responses[recId].HttpCode,
		}
	}

	return ResponseData{
		Recid:    recId,
		Code:     9999,
		Message:  "",
		HttpCode: http.StatusBadRequest,
	}
}

func (c *ResponseParameter) GetResponseMessageOnly(recid string) string {
	if c.Responses[recid].Message != "" {
		return c.Responses[recid].Message
	}
	return ""
}

func (c *ResponseParameter) SetResponse(recId string, err any, data any, nextPage *string) model.Response {
	if resp, ok := c.Responses[recId]; ok && resp.Message != "" {
		return model.Response{
			HttpCode:    resp.HttpCode,
			Message:     resp.Message,
			MessageId:   resp.Recid,
			MessageCode: resp.Code,
			Errors:      err,
			Data:        data,
			NextPage:    nextPage,
		}
	}

	if strings.HasPrefix(recId, "SUCCESS_") {
		return model.Response{
			HttpCode:    http.StatusOK,
			Message:     "success",
			MessageId:   recId,
			MessageCode: 0,
			Errors:      err,
			Data:        data,
			NextPage:    nil,
		}
	}

	// fallback error
	return model.Response{
		HttpCode:    http.StatusBadRequest,
		Message:     "GAGAL",
		MessageId:   recId,
		MessageCode: 9999,
		Errors:      err,
		Data:        nil,
		NextPage:    nil,
	}
}


func (c *ResponseParameter) SetResponseWithDataRecid(recId string, err any, data any, nextPage *string, dataRecid string) model.Response {
	responseParameter := c.SetResponse(recId, err, data, nextPage)
	responseParameter.Message = fmt.Sprintf(responseParameter.Message, dataRecid)
	return responseParameter
}
