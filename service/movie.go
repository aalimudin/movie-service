package service

import (
	"encoding/json"
	"fmt"
	"log"
	"movie-services/model"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
)

type (
	Movie interface {
		GetMovieList(string, string) (model.MovieList, error)
		GetMovieDetail(string) (model.MovieDetail, error)
	}

	movieService struct {
		apiLogService ApiLog
	}
)

func NewMovieService(apiLogService ApiLog) Movie {
	ms := &movieService{}
	ms.apiLogService = apiLogService
	return ms
}

func (ms *movieService) GetMovieList(keyword string, page string) (model.MovieList, error) {
	baseUrl := viper.GetString("omdb.url")
	apiKey := viper.GetString("omdb.key")
	reqUrl := fmt.Sprintf("%v/?apikey=%v&s=%v&page=%v", baseUrl, apiKey, url.QueryEscape(keyword), page)
	movieList := model.MovieList{}

	req, err := http.NewRequest("GET", reqUrl, nil)

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return movieList, err
	}

	if err != nil {
		log.Println(err)
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&movieList)
	resBody, nil := json.Marshal(movieList)

	rm := RequestModel{}
	rm.RequestUrl = reqUrl
	rm.RequestParameter = "keyword: " + keyword + ", page: " + page
	rm.StatusCode = response.StatusCode
	rm.Response = string(resBody)

	err = ms.apiLogService.SaveSearchRequestLog(rm)

	if err != nil {
		log.Println(err)
		return movieList, err
	}

	return movieList, nil
}

func (ms *movieService) GetMovieDetail(title string) (model.MovieDetail, error) {
	baseUrl := viper.GetString("omdb.url")
	apiKey := viper.GetString("omdb.key")
	reqUrl := fmt.Sprintf("%v/?apikey=%v&t=%v", baseUrl, apiKey, url.QueryEscape(title))
	movieDetail := model.MovieDetail{}

	req, err := http.NewRequest("GET", reqUrl, nil)

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return movieDetail, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&movieDetail)

	if err != nil {
		log.Println(err)
		return movieDetail, err
	}

	return movieDetail, nil
}
