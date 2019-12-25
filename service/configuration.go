package service

import (
    "context"
    microClient "github.com/micro/go-micro/client"
    "github.com/plexmediamanager/micro-tmdb/proto"
    "github.com/plexmediamanager/micro-tmdb/tmdb"
)

func TMDBServiceGetAPIConfiguration (client microClient.Client) (*tmdb.APIConfiguration, error) {
    var result *tmdb.APIConfiguration
    service := GetTMDBService(client)
    parameters := &proto.TMDBEmpty{}
    response, err := service.GetAPIConfiguration(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBServiceGetCountries (client microClient.Client) (*tmdb.Countries, error) {
    var result *tmdb.Countries
    service := GetTMDBService(client)
    parameters := &proto.TMDBEmpty{}
    response, err := service.GetCountries(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBServiceGetLanguages (client microClient.Client) ([]tmdb.Language, error) {
    var result []tmdb.Language
    service := GetTMDBService(client)
    parameters := &proto.TMDBEmpty{}
    response, err := service.GetLanguages(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBServiceGetPrimaryTranslations (client microClient.Client) (*tmdb.PrimaryTranslations, error) {
    var result *tmdb.PrimaryTranslations
    service := GetTMDBService(client)
    parameters := &proto.TMDBEmpty{}
    response, err := service.GetPrimaryTranslations(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBServiceGetTimeZones (client microClient.Client) (*tmdb.TimeZones, error) {
    var result *tmdb.TimeZones
    service := GetTMDBService(client)
    parameters := &proto.TMDBEmpty{}
    response, err := service.GetTimeZones(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

