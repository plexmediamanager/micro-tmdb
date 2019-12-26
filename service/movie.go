package service

import (
    "context"
    microClient "github.com/micro/go-micro/client"
    "github.com/plexmediamanager/micro-database/models"
    "github.com/plexmediamanager/micro-tmdb/proto"
    "github.com/plexmediamanager/micro-tmdb/tmdb"
)

func TMDBGetMovieInformation (client microClient.Client, id uint64, options map[string]string) (*tmdb.Movie, error) {
    var result *tmdb.Movie
    service := GetTMDBService(client)
    parameters := &proto.TMDBIDOptions{ Id: id, Options: options }
    response, err := service.GetMovieInformation(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func GetMovieInformationForDatabase (client microClient.Client, id uint64, options map[string]string) (*models.Movie, error) {
    var result *models.Movie
    service := GetTMDBService(client)
    parameters := &proto.TMDBIDOptions{ Id: id, Options: options }
    response, err := service.GetMovieInformationForDatabase(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetMovieAlternativeTitles (client microClient.Client, id uint64, options map[string]string) (*tmdb.MovieAlternativeTitles, error) {
    var result *tmdb.MovieAlternativeTitles
    service := GetTMDBService(client)
    parameters := &proto.TMDBIDOptions{ Id: id, Options: options }
    response, err := service.GetMovieAlternativeTitles(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetMovieChanges (client microClient.Client, id uint64, options map[string]string) (*tmdb.MovieChanges, error) {
    var result *tmdb.MovieChanges
    service := GetTMDBService(client)
    parameters := &proto.TMDBIDOptions{ Id: id, Options: options }
    response, err := service.GetMovieChanges(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetMovieCredits (client microClient.Client, id uint64, options map[string]string) (*tmdb.MovieCredits, error) {
    var result *tmdb.MovieCredits
    service := GetTMDBService(client)
    parameters := &proto.TMDBIDOptions{ Id: id, Options: options }
    response, err := service.GetMovieCredits(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetMovieExternalIDs (client microClient.Client, id uint64, options map[string]string) (*tmdb.MovieExternalIDs, error) {
    var result *tmdb.MovieExternalIDs
    service := GetTMDBService(client)
    parameters := &proto.TMDBIDOptions{ Id: id, Options: options }
    response, err := service.GetMovieExternalIDs(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetMovieImages (client microClient.Client, id uint64, options map[string]string) (*tmdb.MovieImages, error) {
    var result *tmdb.MovieImages
    service := GetTMDBService(client)
    parameters := &proto.TMDBIDOptions{ Id: id, Options: options }
    response, err := service.GetMovieImages(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetMovieKeywords (client microClient.Client, id uint64, options map[string]string) (*tmdb.MovieKeywords, error) {
    var result *tmdb.MovieKeywords
    service := GetTMDBService(client)
    parameters := &proto.TMDBIDOptions{ Id: id, Options: options }
    response, err := service.GetMovieKeywords(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetMovieReleaseDates (client microClient.Client, id uint64, options map[string]string) (*tmdb.MovieReleaseDates, error) {
    var result *tmdb.MovieReleaseDates
    service := GetTMDBService(client)
    parameters := &proto.TMDBIDOptions{ Id: id, Options: options }
    response, err := service.GetMovieReleaseDates(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetMovieVideos (client microClient.Client, id uint64, options map[string]string) (*tmdb.MovieVideos, error) {
    var result *tmdb.MovieVideos
    service := GetTMDBService(client)
    parameters := &proto.TMDBIDOptions{ Id: id, Options: options }
    response, err := service.GetMovieVideos(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBGetMovieTranslations (client microClient.Client, id uint64, options map[string]string) (*tmdb.MovieTranslations, error) {
    var result *tmdb.MovieTranslations
    service := GetTMDBService(client)
    parameters := &proto.TMDBIDOptions{ Id: id, Options: options }
    response, err := service.GetMovieTranslations(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

