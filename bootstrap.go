package main

import (
	"movie-services/controller"
	"movie-services/repository"
	"movie-services/service"
)

func (a *application) registerRoutes() {
	//Repository
	apiLogRepository := repository.NewApiLogRepository(a.movieDB)

	//Service
	apiLogService := service.NewApiLogService(apiLogRepository)
	movieService := service.NewMovieService(apiLogService)

	//Controller
	movieController := controller.NewMovieController(movieService)

	//Route
	a.echo.GET("/movies", movieController.GetMovieListAction)
	a.echo.GET("/movie", movieController.GetMovieDetailAction)

}
