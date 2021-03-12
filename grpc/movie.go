package movie

import (
	"movie-services/service"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) GetMovieList(ctx context.Context, in *MovieListParam) (*MovieSearchList, error) {
	result := service.NewMovieService().GetMovieList(in.Keyword, in.Page)
	return &MovieSearchList{
		Search:       nil,
		TotalResults: result.TotalResults,
		Response:     result.Response,
	}, nil
}

func (s *Server) GetMovieDetail(ctx context.Context, in *MovieDetailParam) (*MovieDetail, error) {
	result := service.NewMovieService().GetMovieDetaill(in.Title)
	return &MovieDetail{
		Title:      result.Title,
		Year:       result.Year,
		Rated:      result.Rated,
		Released:   result.Released,
		Runtime:    result.Runtime,
		Genre:      result.Genre,
		Director:   result.Director,
		Writer:     result.Writer,
		Actors:     result.Actors,
		Plot:       result.Plot,
		Language:   result.Language,
		Country:    result.Country,
		Awards:     result.Awards,
		Poster:     result.Poster,
		Ratings:    nil,
		Metascore:  result.Metascore,
		ImdbRating: result.ImdbRating,
		ImdbVotes:  result.ImdbVotes,
		ImdbID:     result.ImdbID,
		Type:       result.Type,
		DVD:        result.DVD,
		BoxOffice:  result.BoxOffice,
		Production: result.Production,
		Website:    result.Website,
		Response:   result.Response,
	}, nil
}
