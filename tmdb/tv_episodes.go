package tmdb

import format "fmt"

type TVEpisode struct {
    AirDate                         string          `json:"air_date"`
    Crew                            []struct        {
        ID                          uint64          `json:"id"`
        CreditID                    string          `json:"credit_id"`
        Name                        string          `json:"name"`
        Department                  string          `json:"department"`
        Job                         string          `json:"job"`
        ProfilePath                 string          `json:"profile_path"`
    }                                               `json:"crew"`
    EpisodeNumber                   uint64          `json:"episode_number"`
    GuestStars                      []struct        {
        ID                          uint64          `json:"id"`
        Name                        string          `json:"name"`
        CreditID                    string          `json:"credit_id"`
        Character                   string          `json:"character"`
        Order                       uint64          `json:"order"`
        ProfilePath                 string          `json:"profile_path"`
    }                                               `json:"guest_stars"`
    Name                            string          `json:"name"`
    Overview                        string          `json:"overview"`
    ID                              uint64          `json:"id"`
    ShowID                          uint64          `json:"show_id"`
    ProductionCode                  string          `json:"production_code"`
    SeasonNumber                    uint64          `json:"season_number"`
    StillPath                       string          `json:"still_path"`
    VoteAverage                     float64         `json:"vote_average"`
    VoteCount                       uint64          `json:"vote_count"`
}

type TVEpisodeGroups struct {
    Results                         []struct        {
        Description                 string          `json:"description"`
        EpisodeCount                uint64          `json:"episode_count"`
        GroupCount                  uint64          `json:"group_count"`
        ID                          string          `json:"id"`
        Name                        string          `json:"name"`
        Type                        uint64          `json:"type"`
        Network                     struct          {
            ID                      uint64          `json:"id"`
            LogoPath                string          `json:"logo_path"`
            Name                    string          `json:"name"`
            OriginCountry           string          `json:"origin_country"`
        }                                           `json:"network"`
    }                                               `json:"results"`
    ID                              uint64          `json:"id"`
}

type TVEpisodeChanges struct {
    Changes                         []struct        {
        Key                         string          `json:"key"`
        Items                       []struct        {
            ID                      string          `json:"id"`
            Action                  string          `json:"action"`
            Time                    string          `json:"time"`
            Value                   struct          {
                EpisodeID           uint64          `json:"episode_id"`
                EpisodeNumber       uint64          `json:"episode_number"`
            }                                       `json:"value"`
            ISO6391                 string          `json:"iso_639_1"`
            OriginalValue           string          `json:"original_value"`
        }                                           `json:"items"`
    }                                               `json:"changes"`
}

type TVEpisodeCredits struct {
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
    GuestStars                      []struct            {
        ID                          uint64              `json:"id"`
        Name                        string              `json:"name"`
        CreditID                    string              `json:"credit_id"`
        Character                   string              `json:"character"`
        Order                       uint64              `json:"order"`
        ProfilePath                 string              `json:"profile_path"`
    }                                                   `json:"guest_stars"`
}

type TVEpisodeExternalIDs struct {
    Freebase                        string              `json:"freebase_mid"`
    FreebaseID                      string              `json:"freebase_id"`
    TVDB                            uint64              `json:"tvdb_id"`
    TVRage                          uint64              `json:"tvrage_id"`
    ID                              uint64              `json:"id"`
}

type TVEpisodeImages struct {
    ID                              uint64              `json:"id"`
    Stills                          []CommonImage       `json:"stills"`
}

type TVEpisodeTranslations struct {
    ID                              uint64              `json:"id"`
    Translations                    []struct            {
        ISO31661                    string              `json:"iso_3166_1"`
        ISO6391                     string              `json:"iso_639_1"`
        Name                        string              `json:"name"`
        EnglishName                 string              `json:"english_name"`
        Data                        struct              {
            Name                    string              `json:"name"`
            Overview                string              `json:"overview"`
        }                                               `json:"data"`
    }                                                   `json:"translations"`
}

type TVEpisodeVideos struct {
    ID                              uint64              `json:"id"`
    Results                         []struct            {
        ID                          string              `json:"id"`
        ISO6391                     string              `json:"iso_639_1"`
        ISO31661                    string              `json:"iso_3166_1"`
        Key                         string              `json:"key"`
        Name                        string              `json:"name"`
        Site                        string              `json:"site"`
        Size                        uint64              `json:"size"`
        Type                        string              `json:"type"`
    }                                                   `json:"results"`
}

// Get the TV episode details by id.
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-details
func (client *TheMovieDatabase) GetTVEpisodeInformation(id uint64, seasonNumber uint64, episodeNumber uint64, options map[string]string) (*TVEpisode, error) {
    availableOptions := map[string]struct{} {
        "language":             {},
        "append_to_response":   {},
    }
    var structure TVEpisode
    result, err := client.sendGetRequest(
        client.buildRequestUrl("tv", id, format.Sprintf("season/%d/episode/%d", seasonNumber, episodeNumber), options, availableOptions),
        &structure,
    )
    return result.(*TVEpisode), err
}

// Get the changes for a TV episode.
// By default only the last 24 hours are returned.
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-changes
func (client *TheMovieDatabase) GetTVEpisodeChanges(id uint64, options map[string]string) (*TVEpisodeChanges, error) {
    availableOptions := map[string]struct{} {
        "language":             {},
        "start_date":           {},
        "end_date":             {},
        "page":                 {},
    }
    var structure TVEpisodeChanges
    result, err := client.sendGetRequest(
        client.buildRequestUrl("tv/episode", id, "changes", options, availableOptions),
        &structure,
    )
    return result.(*TVEpisodeChanges), err
}

// Get the credits (cast, crew and guest stars) for a TV episode.
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-credits
func (client *TheMovieDatabase) GetTVEpisodeCredits(id uint64, seasonNumber uint64, episodeNumber uint64, options map[string]string) (*TVEpisodeCredits, error) {
    availableOptions := map[string]struct{} {
        "language":             {},
    }
    var structure TVEpisodeCredits
    result, err := client.sendGetRequest(
        client.buildRequestUrl("tv", id, format.Sprintf("season/%d/episode/%d/credits", seasonNumber, episodeNumber), options, availableOptions),
        &structure,
    )
    return result.(*TVEpisodeCredits), err
}

// Get the external ids for a TV episode.
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-external-ids
func (client *TheMovieDatabase) GetTVEpisodeExternalIDs(id uint64, seasonNumber uint64, episodeNumber uint64, options map[string]string) (*TVEpisodeExternalIDs, error) {
    availableOptions := map[string]struct{} {
        "language":             {},
    }
    var structure TVEpisodeExternalIDs
    result, err := client.sendGetRequest(
        client.buildRequestUrl("tv", id, format.Sprintf("season/%d/episode/%d/external_ids", seasonNumber, episodeNumber), options, availableOptions),
        &structure,
    )
    return result.(*TVEpisodeExternalIDs), err
}

// Get the images that belong to a TV episode.
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-images
func (client *TheMovieDatabase) GetTVEpisodeImages(id uint64, seasonNumber uint64, episodeNumber uint64, options map[string]string) (*TVEpisodeImages, error) {
    availableOptions := map[string]struct{} {
        "language":             {},
    }
    var structure TVEpisodeImages
    result, err := client.sendGetRequest(
        client.buildRequestUrl("tv", id, format.Sprintf("season/%d/episode/%d/images", seasonNumber, episodeNumber), options, availableOptions),
        &structure,
    )
    return result.(*TVEpisodeImages), err
}

// Get the images that belong to a TV episode.
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-images
func (client *TheMovieDatabase) GetTVEpisodeTranslations(id uint64, seasonNumber uint64, episodeNumber uint64, options map[string]string) (*TVEpisodeTranslations, error) {
    availableOptions := map[string]struct{} {
        "language":             {},
    }
    var structure TVEpisodeTranslations
    result, err := client.sendGetRequest(
        client.buildRequestUrl("tv", id, format.Sprintf("season/%d/episode/%d/translations", seasonNumber, episodeNumber), options, availableOptions),
        &structure,
    )
    return result.(*TVEpisodeTranslations), err
}

// Get the videos that have been added to a TV episode.
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-videos
func (client *TheMovieDatabase) GetTVEpisodeVideos(id uint64, seasonNumber uint64, episodeNumber uint64, options map[string]string) (*TVEpisodeVideos, error) {
    availableOptions := map[string]struct{} {
        "language":             {},
    }
    var structure TVEpisodeVideos
    result, err := client.sendGetRequest(
        client.buildRequestUrl("tv", id, format.Sprintf("season/%d/episode/%d/videos", seasonNumber, episodeNumber), options, availableOptions),
        &structure,
    )
    return result.(*TVEpisodeVideos), err
}