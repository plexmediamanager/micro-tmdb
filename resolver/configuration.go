package resolver

import (
    "context"
    "github.com/plexmediamanager/micro-tmdb/proto"
)

func (service TMDBService) GetAPIConfiguration (_ context.Context, parameters *proto.TMDBEmpty, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetAPIConfiguration()
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetCountries (_ context.Context, parameters *proto.TMDBEmpty, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetCountries()
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetLanguages (_ context.Context, parameters *proto.TMDBEmpty, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetLanguages()
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetPrimaryTranslations (_ context.Context, parameters *proto.TMDBEmpty, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetPrimaryTranslations()
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetTimeZones (_ context.Context, parameters *proto.TMDBEmpty, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetTimeZones()
    return structureToBytesWithError(result, err, response)
}

