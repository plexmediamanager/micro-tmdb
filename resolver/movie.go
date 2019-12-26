package resolver

import (
    "context"
    "github.com/plexmediamanager/micro-tmdb/proto"
)

func (service TMDBService) GetMovieInformation (_ context.Context, parameters *proto.TMDBIDOptions, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetMovieInformation(parameters.Id, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetMovieInformationForDatabase (_ context.Context, parameters *proto.TMDBIDOptions, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetMovieInformationForDatabase(parameters.Id, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetMovieAlternativeTitles (_ context.Context, parameters *proto.TMDBIDOptions, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetMovieAlternativeTitles(parameters.Id, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetMovieChanges (_ context.Context, parameters *proto.TMDBIDOptions, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetMovieChanges(parameters.Id, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetMovieCredits (_ context.Context, parameters *proto.TMDBIDOptions, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetMovieCredits(parameters.Id, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetMovieExternalIDs (_ context.Context, parameters *proto.TMDBIDOptions, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetMovieExternalIDs(parameters.Id, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetMovieImages (_ context.Context, parameters *proto.TMDBIDOptions, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetMovieImages(parameters.Id, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetMovieKeywords (_ context.Context, parameters *proto.TMDBIDOptions, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetMovieKeywords(parameters.Id, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetMovieReleaseDates (_ context.Context, parameters *proto.TMDBIDOptions, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetMovieReleaseDates(parameters.Id, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetMovieVideos (_ context.Context, parameters *proto.TMDBIDOptions, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetMovieVideos(parameters.Id, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetMovieTranslations (_ context.Context, parameters *proto.TMDBIDOptions, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetMovieTranslations(parameters.Id, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

