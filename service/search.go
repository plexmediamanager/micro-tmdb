package service

import (
    "context"
    microClient "github.com/micro/go-micro/client"
    "github.com/plexmediamanager/micro-tmdb/proto"
    "github.com/plexmediamanager/micro-tmdb/tmdb"
)

func TMDBSearchCompanies (client microClient.Client, query string, options map[string]string) (*tmdb.SearchCompanies, error) {
    var result *tmdb.SearchCompanies
    service := GetTMDBService(client)
    parameters := &proto.TMDBQueryOptions{ Query: query, Options: options }
    response, err := service.SearchCompanies(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBSearchMovies (client microClient.Client, query string, year uint64, options map[string]string) (*tmdb.SearchMovies, error) {
    var result *tmdb.SearchMovies
    service := GetTMDBService(client)
    parameters := &proto.TMDBQueryYearOptions{ Query: query, Year: year, Options: options }
    response, err := service.SearchMovies(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TMDBSearchTV (client microClient.Client, query string, year uint64, options map[string]string) (*tmdb.SearchTV, error) {
    var result *tmdb.SearchTV
    service := GetTMDBService(client)
    parameters := &proto.TMDBQueryYearOptions{ Query: query, Year: year, Options: options }
    response, err := service.SearchTV(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}
