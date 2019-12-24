package resolver

import (
    "context"
    "github.com/plexmediamanager/micro-tmdb/proto"
)

func (service TMDBService) GetTVInformation (_ context.Context, parameters *proto.TMDBIDOptions, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetTVInformation(parameters.Id, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetTVChanges (_ context.Context, parameters *proto.TMDBIDOptions, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetTVChanges(parameters.Id, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetTVCredits (_ context.Context, parameters *proto.TMDBIDOptions, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetTVCredits(parameters.Id, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetTVExternalIDs (_ context.Context, parameters *proto.TMDBIDOptions, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetTVExternalIDs(parameters.Id, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetTVImages (_ context.Context, parameters *proto.TMDBIDOptions, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetTVImages(parameters.Id, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetTVKeywords (_ context.Context, parameters *proto.TMDBIDOptions, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetTVKeywords(parameters.Id, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetTVVideos (_ context.Context, parameters *proto.TMDBIDOptions, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetTVVideos(parameters.Id, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetTVTranslations (_ context.Context, parameters *proto.TMDBIDOptions, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetTVTranslations(parameters.Id, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetTVSeasonInformation (_ context.Context, parameters *proto.TMDBSeason, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetTVSeasonInformation(parameters.Id, parameters.Season, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetTVSeasonChanges (_ context.Context, parameters *proto.TMDBIDOptions, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetTVSeasonChanges(parameters.Id, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetTVSeasonCredits (_ context.Context, parameters *proto.TMDBSeason, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetTVSeasonCredits(parameters.Id, parameters.Season, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetTVSeasonExternalIDs (_ context.Context, parameters *proto.TMDBSeason, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetTVSeasonExternalIDs(parameters.Id, parameters.Season, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetTVSeasonImages (_ context.Context, parameters *proto.TMDBSeason, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetTVSeasonImages(parameters.Id, parameters.Season, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetTVSeasonVideos (_ context.Context, parameters *proto.TMDBSeason, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetTVSeasonVideos(parameters.Id, parameters.Season, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetTVEpisodeInformation (_ context.Context, parameters *proto.TMDBSeasonEpisode, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetTVEpisodeInformation(parameters.Id, parameters.Season, parameters.Episode, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetTVEpisodeChanges (_ context.Context, parameters *proto.TMDBIDOptions, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetTVEpisodeChanges(parameters.Id, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetTVEpisodeCredits (_ context.Context, parameters *proto.TMDBSeasonEpisode, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetTVEpisodeCredits(parameters.Id, parameters.Season, parameters.Episode, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetTVEpisodeExternalIDs (_ context.Context, parameters *proto.TMDBSeasonEpisode, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetTVEpisodeExternalIDs(parameters.Id, parameters.Season, parameters.Episode, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetTVEpisodeImages (_ context.Context, parameters *proto.TMDBSeasonEpisode, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetTVEpisodeImages(parameters.Id, parameters.Season, parameters.Episode, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetTVEpisodeTranslations (_ context.Context, parameters *proto.TMDBSeasonEpisode, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetTVEpisodeTranslations(parameters.Id, parameters.Season, parameters.Episode, parameters.Options)
    return structureToBytesWithError(result, err, response)
}

func (service TMDBService) GetTVEpisodeVideos (_ context.Context, parameters *proto.TMDBSeasonEpisode, response *proto.TMDBResponse) error {
    result, err := service.TMDB.GetTVEpisodeVideos(parameters.Id, parameters.Season, parameters.Episode, parameters.Options)
    return structureToBytesWithError(result, err, response)
}
