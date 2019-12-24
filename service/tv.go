package service

import (
    "context"
    microClient "github.com/micro/go-micro/client"
    "github.com/plexmediamanager/micro-tmdb/proto"
    "github.com/plexmediamanager/micro-tmdb/tmdb"
)

func TMDBGetTVInformation (client microClient.Client, id uint64, options map[string]string) (*tmdb.TV, error) {
    var result *tmdb.TV
    service := GetTMDBService(client)
    parameters := &proto.TMDBIDOptions{ Id: id, Options: options }
    response, err := service.GetTVInformation(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetTVChanges (client microClient.Client, id uint64, options map[string]string) (*tmdb.TVChanges, error) {
    var result *tmdb.TVChanges
    service := GetTMDBService(client)
    parameters := &proto.TMDBIDOptions{ Id: id, Options: options }
    response, err := service.GetTVChanges(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetTVCredits (client microClient.Client, id uint64, options map[string]string) (*tmdb.TVCredits, error) {
    var result *tmdb.TVCredits
    service := GetTMDBService(client)
    parameters := &proto.TMDBIDOptions{ Id: id, Options: options }
    response, err := service.GetTVCredits(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetTVExternalIDs (client microClient.Client, id uint64, options map[string]string) (*tmdb.TVExternalIDs, error) {
    var result *tmdb.TVExternalIDs
    service := GetTMDBService(client)
    parameters := &proto.TMDBIDOptions{ Id: id, Options: options }
    response, err := service.GetTVExternalIDs(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetTVImages (client microClient.Client, id uint64, options map[string]string) (*tmdb.TVImages, error) {
    var result *tmdb.TVImages
    service := GetTMDBService(client)
    parameters := &proto.TMDBIDOptions{ Id: id, Options: options }
    response, err := service.GetTVImages(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetTVKeywords (client microClient.Client, id uint64, options map[string]string) (*tmdb.TVKeywords, error) {
    var result *tmdb.TVKeywords
    service := GetTMDBService(client)
    parameters := &proto.TMDBIDOptions{ Id: id, Options: options }
    response, err := service.GetTVKeywords(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetTVVideos (client microClient.Client, id uint64, options map[string]string) (*tmdb.TVVideos, error) {
    var result *tmdb.TVVideos
    service := GetTMDBService(client)
    parameters := &proto.TMDBIDOptions{ Id: id, Options: options }
    response, err := service.GetTVVideos(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetTVTranslations (client microClient.Client, id uint64, options map[string]string) (*tmdb.TVTranslations, error) {
    var result *tmdb.TVTranslations
    service := GetTMDBService(client)
    parameters := &proto.TMDBIDOptions{ Id: id, Options: options }
    response, err := service.GetTVTranslations(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetTVSeasonInformation (client microClient.Client, seriesID uint64, seasonNumber uint64, options map[string]string) (*tmdb.TVSeason, error) {
    var result *tmdb.TVSeason
    service := GetTMDBService(client)
    parameters := &proto.TMDBSeason{ Id: seriesID, Season: seasonNumber, Options: options }
    response, err := service.GetTVSeasonInformation(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetTVSeasonChanges (client microClient.Client, seasonID uint64, options map[string]string) (*tmdb.TVSeasonChanges, error) {
    var result *tmdb.TVSeasonChanges
    service := GetTMDBService(client)
    parameters := &proto.TMDBIDOptions{ Id: seasonID, Options: options }
    response, err := service.GetTVSeasonChanges(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetTVSeasonCredits (client microClient.Client, seriesID uint64, seasonNumber uint64, options map[string]string) (*tmdb.TVSeasonCredits, error) {
    var result *tmdb.TVSeasonCredits
    service := GetTMDBService(client)
    parameters := &proto.TMDBSeason{ Id: seriesID, Season: seasonNumber, Options: options }
    response, err := service.GetTVSeasonCredits(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetTVSeasonExternalIDs (client microClient.Client, seriesID uint64, seasonNumber uint64, options map[string]string) (*tmdb.TVSeasonExternalIDs, error) {
    var result *tmdb.TVSeasonExternalIDs
    service := GetTMDBService(client)
    parameters := &proto.TMDBSeason{ Id: seriesID, Season: seasonNumber, Options: options }
    response, err := service.GetTVSeasonExternalIDs(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetTVSeasonImages (client microClient.Client, seriesID uint64, seasonNumber uint64, options map[string]string) (*tmdb.TVSeasonImages, error) {
    var result *tmdb.TVSeasonImages
    service := GetTMDBService(client)
    parameters := &proto.TMDBSeason{ Id: seriesID, Season: seasonNumber, Options: options }
    response, err := service.GetTVSeasonImages(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetTVSeasonVideos (client microClient.Client, seriesID uint64, seasonNumber uint64, options map[string]string) (*tmdb.TVSeasonVideos, error) {
    var result *tmdb.TVSeasonVideos
    service := GetTMDBService(client)
    parameters := &proto.TMDBSeason{ Id: seriesID, Season: seasonNumber, Options: options }
    response, err := service.GetTVSeasonVideos(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetTVEpisodeInformation (client microClient.Client, seriesID uint64, seasonNumber uint64, episodeNumber uint64, options map[string]string) (*tmdb.TVEpisode, error) {
    var result *tmdb.TVEpisode
    service := GetTMDBService(client)
    parameters := &proto.TMDBSeasonEpisode{ Id: seriesID, Season: seasonNumber, Episode: episodeNumber, Options: options }
    response, err := service.GetTVEpisodeInformation(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetTVEpisodeChanges (client microClient.Client, episodeID uint64, options map[string]string) (*tmdb.TVEpisodeChanges, error) {
    var result *tmdb.TVEpisodeChanges
    service := GetTMDBService(client)
    parameters := &proto.TMDBIDOptions{ Id: episodeID, Options: options }
    response, err := service.GetTVEpisodeChanges(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetTVEpisodeCredits (client microClient.Client, seriesID uint64, seasonNumber uint64, episodeNumber uint64, options map[string]string) (*tmdb.TVEpisodeCredits, error) {
    var result *tmdb.TVEpisodeCredits
    service := GetTMDBService(client)
    parameters := &proto.TMDBSeasonEpisode{ Id: seriesID, Season: seasonNumber, Episode: episodeNumber, Options: options }
    response, err := service.GetTVEpisodeCredits(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetTVEpisodeExternalIDs (client microClient.Client, seriesID uint64, seasonNumber uint64, episodeNumber uint64, options map[string]string) (*tmdb.TVEpisodeExternalIDs, error) {
    var result *tmdb.TVEpisodeExternalIDs
    service := GetTMDBService(client)
    parameters := &proto.TMDBSeasonEpisode{ Id: seriesID, Season: seasonNumber, Episode: episodeNumber, Options: options }
    response, err := service.GetTVEpisodeExternalIDs(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetTVEpisodeImages (client microClient.Client, seriesID uint64, seasonNumber uint64, episodeNumber uint64, options map[string]string) (*tmdb.TVEpisodeImages, error) {
    var result *tmdb.TVEpisodeImages
    service := GetTMDBService(client)
    parameters := &proto.TMDBSeasonEpisode{ Id: seriesID, Season: seasonNumber, Episode: episodeNumber, Options: options }
    response, err := service.GetTVEpisodeImages(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetTVEpisodeTranslations (client microClient.Client, seriesID uint64, seasonNumber uint64, episodeNumber uint64, options map[string]string) (*tmdb.TVEpisodeTranslations, error) {
    var result *tmdb.TVEpisodeTranslations
    service := GetTMDBService(client)
    parameters := &proto.TMDBSeasonEpisode{ Id: seriesID, Season: seasonNumber, Episode: episodeNumber, Options: options }
    response, err := service.GetTVEpisodeTranslations(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetTVEpisodeVideos (client microClient.Client, seriesID uint64, seasonNumber uint64, episodeNumber uint64, options map[string]string) (*tmdb.TVEpisodeVideos, error) {
    var result *tmdb.TVEpisodeVideos
    service := GetTMDBService(client)
    parameters := &proto.TMDBSeasonEpisode{ Id: seriesID, Season: seasonNumber, Episode: episodeNumber, Options: options }
    response, err := service.GetTVEpisodeVideos(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}