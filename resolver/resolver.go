package resolver

import (
    "encoding/json"
    "github.com/plexmediamanager/micro-tmdb/errors"
    "github.com/plexmediamanager/micro-tmdb/proto"
    "github.com/plexmediamanager/micro-tmdb/tmdb"
)

type TMDBService struct {
    TMDB           *tmdb.TheMovieDatabase
}

// Convert structure to bytes and return error if needed
func structureToBytesWithError(structure interface{}, err error, response *proto.TMDBResponse) error {
    if err != nil {
        return err
    }
    bytes, err := json.Marshal(structure)
    if err != nil {
        return errors.TMDBUnmarshalError.ToError(err)
    }
    response.Result = bytes
    return nil
}

