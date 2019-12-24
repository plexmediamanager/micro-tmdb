package tmdb

import format "fmt"

type TVSeason struct {
    AirDate                     string          `json:"air_date"`
    Episodes                    []TVEpisode     `json:"episodes"`
    Name                        string          `json:"name"`
    Overview                    string          `json:"overview"`
    ID                          uint64          `json:"id"`
    PosterPath                  string          `json:"poster_path"`
    SeasonNumber                uint64          `json:"season_number"`
}

type TVSeasonChanges struct {
    Changes                     []struct        {
        Key                     string          `json:"key"`
        Items                   []struct        {
            ID                  string          `json:"id"`
            Action              string          `json:"action"`
            Time                string          `json:"time"`
            Value               struct          {
                EpisodeID       uint64          `json:"episode_id"`
                EpisodeNumber   uint64          `json:"episode_number"`
            }                                   `json:"value"`
            ISO6391             string          `json:"iso_639_1"`
            OriginalValue       string          `json:"original_value"`
        }                                       `json:"items"`
    }                                           `json:"changes"`
}

type TVSeasonCredits struct {
    ID                              uint64              `json:"id"`
    Cast                            []struct            {
        Character                   string              `json:"character"`
        CreditID                    string              `json:"credit_id"`
        Gender                      uint64              `json:"gender"`
        ID                          uint64              `json:"id"`
        Name                        string              `json:"name"`
        Order                       uint64              `json:"order"`
        ProfilePath                 string              `json:"profile_path"`
    }                                                   `json:"cast"`
    Crew                            []struct            {
        CreditID                    string              `json:"credit_id"`
        Department                  string              `json:"department"`
        Gender                      uint64              `json:"gender"`
        ID                          uint64              `json:"id"`
        Job                         string              `json:"job"`
        Name                        string              `json:"name"`
        ProfilePath                 string              `json:"profile_path"`
    }                                                   `json:"crew"`
}

type TVSeasonExternalIDs struct {
    Freebase                        string              `json:"freebase_mid"`
    FreebaseID                      string              `json:"freebase_id"`
    TVDB                            uint64              `json:"tvdb_id"`
    TVRage                          uint64              `json:"tvrage_id"`
    ID                              uint64              `json:"id"`
}

type TVSeasonImages struct {
    ID                              uint64              `json:"id"`
    Posters                         []CommonImage       `json:"posters"`
}

type TVSeasonVideos struct {
    ID                          uint64              `json:"id"`
    Results                     []struct            {
        ID                      string              `json:"id"`
        ISO6391                 string              `json:"iso_639_1"`
        ISO31661                string              `json:"iso_3166_1"`
        Key                     string              `json:"key"`
        Name                    string              `json:"name"`
        Site                    string              `json:"site"`
        Size                    uint64              `json:"size"`
        Type                    string              `json:"type"`
    }                                               `json:"results"`
}

// Get the TV season details by id.
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-details
func (client *TheMovieDatabase) GetTVSeasonInformation(id uint64, number uint64, options map[string]string) (*TVSeason, error) {
    availableOptions := map[string]struct{} {
        "language":             {},
        "append_to_response":   {},
    }
    var structure TVSeason
    result, err := client.sendGetRequest(
        client.buildRequestUrl("tv", id, format.Sprintf("season/%d", number), options, availableOptions),
        &structure,
    )
    return result.(*TVSeason), err
}

// Get the changes for a TV season.
// By default only the last 24 hours are returned.
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-changes
func (client *TheMovieDatabase) GetTVSeasonChanges(id uint64, options map[string]string) (*TVSeasonChanges, error) {
    availableOptions := map[string]struct{} {
        "language":             {},
        "start_date":           {},
        "end_date":             {},
        "page":                 {},
    }
    var structure TVSeasonChanges
    result, err := client.sendGetRequest(
        client.buildRequestUrl("tv/season", id, "changes", options, availableOptions),
        &structure,
    )
    return result.(*TVSeasonChanges), err
}

// Get the credits for TV season.
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-credits
func (client *TheMovieDatabase) GetTVSeasonCredits(id uint64, number uint64, options map[string]string) (*TVSeasonCredits, error) {
    availableOptions := map[string]struct{} {
        "language":             {},
    }
    var structure TVSeasonCredits
    result, err := client.sendGetRequest(
        client.buildRequestUrl("tv", id, format.Sprintf("season/%d/credits", number), options, availableOptions),
        &structure,
    )
    return result.(*TVSeasonCredits), err
}

// Get the external ids for a TV season.
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-external-ids
func (client *TheMovieDatabase) GetTVSeasonExternalIDs(id uint64, number uint64, options map[string]string) (*TVSeasonExternalIDs, error) {
    availableOptions := map[string]struct{} {
        "language":             {},
    }
    var structure TVSeasonExternalIDs
    result, err := client.sendGetRequest(
        client.buildRequestUrl("tv", id, format.Sprintf("season/%d/external_ids", number), options, availableOptions),
        &structure,
    )
    return result.(*TVSeasonExternalIDs), err
}

// Get the images that belong to a TV season.
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-images
func (client *TheMovieDatabase) GetTVSeasonImages(id uint64, number uint64, options map[string]string) (*TVSeasonImages, error) {
    availableOptions := map[string]struct{} {
        "language":             {},
    }
    var structure TVSeasonImages
    result, err := client.sendGetRequest(
        client.buildRequestUrl("tv", id, format.Sprintf("season/%d/images", number), options, availableOptions),
        &structure,
    )
    return result.(*TVSeasonImages), err
}

// Get the videos that have been added to a TV season.
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-videos
func (client *TheMovieDatabase) GetTVSeasonVideos(id uint64, number uint64, options map[string]string) (*TVSeasonVideos, error) {
    availableOptions := map[string]struct{} {
        "language":             {},
    }
    var structure TVSeasonVideos
    result, err := client.sendGetRequest(
        client.buildRequestUrl("tv", id, format.Sprintf("season/%d/videos", number), options, availableOptions),
        &structure,
    )
    return result.(*TVSeasonVideos), err
}