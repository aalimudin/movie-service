package service

import (
	"movie-services/model/db"
	"movie-services/repository"
)

type (
	ApiLog interface {
		SaveSearchRequestLog(RequestModel) error
	}

	apiLog struct {
		apiLogRepository repository.ApiLogRepository
	}

	RequestModel struct {
		RequestUrl       string
		RequestParameter string
		StatusCode       int
		Response         string
	}
)

func NewApiLogService(apiLogRepository repository.ApiLogRepository) ApiLog {
	al := &apiLog{}
	al.apiLogRepository = apiLogRepository

	return al
}

func (al *apiLog) SaveSearchRequestLog(requestModel RequestModel) error {
	apiLog := db.ApiLog{}
	apiLog.RequestURL = requestModel.RequestUrl
	apiLog.RequestParameter = requestModel.RequestParameter
	apiLog.StatusCode = requestModel.StatusCode
	apiLog.Response = requestModel.Response

	_, err := al.apiLogRepository.SaveLog(&apiLog)

	if err != nil {
		return err
	}

	return nil

}
