package resolver

import (
    "context"
    "github.com/plexmediamanager/micro-tmdb/proto"
)

func (service TMDBService) SearchCompanies (_ context.Context, parameters *proto.TMDBQueryOptions, response *proto.TMDBResponse) error {
    result, err := service.TMDB.SearchCompanies(parameters.Query, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) SearchMovies (_ context.Context, parameters *proto.TMDBQueryYearOptions, response *proto.TMDBResponse) error {
    result, err := service.TMDB.SearchMovies(parameters.Query, parameters.Year, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) SearchTV (_ context.Context, parameters *proto.TMDBQueryYearOptions, response *proto.TMDBResponse) error {
    result, err := service.TMDB.SearchTV(parameters.Query, parameters.Year, parameters.Options)
    return structureToBytesWithError(result, err, response)
}
