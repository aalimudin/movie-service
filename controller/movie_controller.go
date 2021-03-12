package controller

import (
	"movie-services/service"
	"net/http"

	"github.com/labstack/echo"
)

type (
	MovieController interface {
		GetMovieListAction(c echo.Context) error
		GetMovieDetailAction(c echo.Context) error
	}

	movieController struct {
		movieService service.Movie
	}

	errorResponse struct {
		Message string `json:"message"`
	}
)

func NewMovieController(movieService service.Movie) MovieController {
	mc := movieController{}
	mc.movieService = movieService

	return mc
}

func (m movieController) GetMovieListAction(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	page := c.QueryParam("page")

	response, err := m.movieService.GetMovieList(keyword, page)

	if err != nil {
		errRes := errorResponse{}
		errRes.Message = "Data not available"
		return c.JSON(http.StatusBadRequest, errRes)
	}

	return c.JSON(http.StatusOK, response)
}

func (m movieController) GetMovieDetailAction(c echo.Context) error {
	title := c.QueryParam("title")
	response, err := m.movieService.GetMovieDetail(title)
	if err != nil {
		errRes := errorResponse{}
		errRes.Message = "Data not available"
		return c.JSON(http.StatusBadRequest, errRes)
	}

	return c.JSON(http.StatusOK, response)
}
