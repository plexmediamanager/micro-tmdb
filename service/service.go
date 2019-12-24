package service

import (
    "encoding/json"
    microClient "github.com/micro/go-micro/client"
    "github.com/plexmediamanager/micro-tmdb/errors"
    "github.com/plexmediamanager/micro-tmdb/proto"
)

// Convert response to structure
func protoToStructure(output interface{}, result *proto.TMDBResponse, err error) error {
    if err != nil {
        return err
    }
    err = json.Unmarshal(result.Result, &output)
    if err != nil {
        return errors.TMDBUnmarshalError.ToError(err)
    }
    return nil
}

func GetTMDBService(client microClient.Client) proto.TMDBService {
    return proto.NewTMDBService("micro.tmdb", client)
}
