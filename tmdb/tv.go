package tmdb

type TV struct {
    BackdropPath                    string          `json:"backdrop_path"`
    CreatedBy                       []struct        {
        ID                          uint64          `json:"id"`
        CreditID                    string          `json:"credit_it"`
        Name                        string          `json:"name"`
        Gender                      uint64          `json:"gender"`
        ProfilePath                 string          `json:"profile_path"`
    }                                               `json:"created_by"`
    EpisodeRuntime                  []uint64        `json:"episode_run_time"`
    FirstAirDate                    string          `json:"first_air_date"`
    Genres                          []struct        {
        ID                          uint64          `json:"id"`
        Name                        string          `json:"name"`
    }                                               `json:"genres"`
    Homepage                        string          `json:"homepage"`
    ID                              uint64          `json:"id"`
    InProduction                    bool            `json:"in_production"`
    LastAirDate                     string          `json:"last_air_date"`
    LastEpisodeToAir                TVEpisode       `json:"last_episode_to_air"`
    Name                            string          `json:"name"`
    Networks                        []struct        {
        Name                        string          `json:"name"`
        ID                          uint64          `json:"id"`
        LogoPath                    string          `json:"logo_path"`
        OriginCountry               string          `json:"origin_country"`
    }                                               `json:"networks"`
    NumberOfEpisodes                uint64          `json:"number_of_episodes"`
    NumberOfSeasons                 uint64          `json:"number_of_seasons"`
    OriginCountry                   []string        `json:"origin_country"`
    OriginalLanguage                string          `json:"original_language"`
    OriginalName                    string          `json:"original_name"`
    Overview                        string          `json:"overview"`
    Popularity                      float64         `json:"popularity"`
    PosterPath                      string          `json:"poster_path"`
    ProductionCompanies             []struct        {
        ID                          uint64          `json:"id"`
        LogoPath                    string          `json:"logo_path"`
        Name                        string          `json:"name"`
        OriginCountry               string          `json:"origin_country"`
    }                                               `json:"production_companies"`
    Seasons                         []struct        {
        AirDate                     string          `json:"air_date"`
        EpisodeCount                uint64          `json:"episode_count"`
        ID                          uint64          `json:"id"`
        Name                        string          `json:"name"`
        Overview                    string          `json:"overview"`
        PosterPath                  string          `json:"poster_path"`
        SeasonNumber                uint64          `json:"season_number"`
    }                                               `json:"seasons"`
    Status                          string          `json:"status"`
    Type                            string          `json:"type"`
    VoteAverage                     float64         `json:"vote_average"`
    VoteCount                       uint64          `json:"vote_count"`
}

type TVAlternativeTitles struct {
    ID                              uint64          `json:"id"`
    Results                         []struct        {
        Title                       string          `json:"title"`
        ISO31661                    string          `json:"iso_3166_1"`
        Type                        string          `json:"type"`
    }                                               `json:"results"`
}

type TVChanges struct {
    Changes                         []struct            {
        Key                         string              `json:"key"`
        Items                       []struct            {
            ID                      string              `json:"id"`
            Action                  string              `json:"action"`
            Time                    string              `json:"time"`
        }                                               `json:"items"`
    }                                                   `json:"changes"`
}

type TVContentRatings struct {
    Results                         []struct            {
        ISO31661                    string              `json:"iso_3166_1"`
        Rating                      string              `json:"rating"`
    }                                                   `json:"results"`
    ID                              uint64              `json:"id"`
}

type TVCredits struct {
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

type TVExternalIDs struct {
    IMDB                            string              `json:"imdb_id"`
    Freebase                        string              `json:"freebase_mid"`
    FreebaseID                      string              `json:"freebase_id"`
    TVDB                            uint64              `json:"tvdb_id"`
    TVRage                          uint64              `json:"tvrage_id"`
    Facebook                        string              `json:"facebook_id"`
    Instagram                       string              `json:"instagram_id"`
    Twitter                         string              `json:"twitter_id"`
    ID                              uint64              `json:"id"`
}

type TVImages struct {
    ID                          uint64              `json:"id"`
    Backdrops                   []CommonImage       `json:"backdrops"`
    Posters                     []CommonImage       `json:"posters"`
}

type TVKeywords struct {
    ID                          uint64              `json:"id"`
    Result                      []struct            {
        ID                      uint64              `json:"id"`
        Name                    string              `json:"name"`
    }                                               `json:"results"`
}

type TVTranslations struct {
    ID                          uint64              `json:"id"`
    Translations                []struct            {
        ISO31661                string              `json:"iso_3166_1"`
        ISO6391                 string              `json:"iso_639_1"`
        Name                    string              `json:"name"`
        EnglishName             string              `json:"english_name"`
        Data                    struct              {
            Title               string              `json:"title"`
            Overview            string              `json:"overview"`
            Homepage            string              `json:"homepage"`
        }                                           `json:"data"`
    }                                               `json:"translations"`
}

type TVVideos struct {
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

// Get the primary TV show details by id.
// https://developers.themoviedb.org/3/tv/get-tv-details
func (client *TheMovieDatabase) GetTVInformation(id uint64, options map[string]string) (*TV, error) {
    availableOptions := map[string]struct{} {
        "language":             {},
        "append_to_response":   {},
    }
    var structure TV
    result, err := client.sendGetRequest(
        client.buildRequestUrl("tv", id, nil, options, availableOptions),
        &structure,
    )
    return result.(*TV), err
}

// Returns all of the alternative titles for a TV show.
// https://developers.themoviedb.org/3/tv/get-tv-alternative-titles
func (client *TheMovieDatabase) GetTVAlternativeTitle(id uint64, options map[string]string) (*TVAlternativeTitles, error) {
    availableOptions := map[string]struct{} {
        "language":             {},
    }
    var structure TVAlternativeTitles
    result, err := client.sendGetRequest(
        client.buildRequestUrl("tv", id, "alternative_titles", options, availableOptions),
        &structure,
    )
    return result.(*TVAlternativeTitles), err
}

// Get the changes for a TV show. By default only the last 24 hours are returned.
// https://developers.themoviedb.org/3/tv/get-tv-changes
func (client *TheMovieDatabase) GetTVChanges(id uint64, options map[string]string) (*TVChanges, error) {
    availableOptions := map[string]struct{} {
        "start_date":           {},
        "end_date":             {},
        "page":                 {},
    }
    var structure TVChanges
    result, err := client.sendGetRequest(
        client.buildRequestUrl("tv", id, "changes", options, availableOptions),
        &structure,
    )
    return result.(*TVChanges), err
}

// Get the list of content ratings (certifications) that have been added to a TV show.
// https://developers.themoviedb.org/3/tv/get-tv-content-ratings
func (client *TheMovieDatabase) GetTVContentRatings(id uint64, options map[string]string) (*TVContentRatings, error) {
    availableOptions := map[string]struct{} {
        "language":             {},
    }
    var structure TVContentRatings
    result, err := client.sendGetRequest(
        client.buildRequestUrl("tv", id, "content_ratings", options, availableOptions),
        &structure,
    )
    return result.(*TVContentRatings), err
}

// Get the credits (cast and crew) that have been added to a TV show.
// https://developers.themoviedb.org/3/tv/get-tv-credits
func (client *TheMovieDatabase) GetTVCredits(id uint64, options map[string]string) (*TVCredits, error) {
    availableOptions := map[string]struct{} {
        "language":             {},
    }
    var structure TVCredits
    result, err := client.sendGetRequest(
        client.buildRequestUrl("tv", id, "credits", options, availableOptions),
        &structure,
    )
    return result.(*TVCredits), err
}

// Get all of the episode groups that have been created for a TV show.
// https://developers.themoviedb.org/3/tv/get-tv-episode-groups
func (client *TheMovieDatabase) GetTVEpisodeGroups(id uint64, options map[string]string) (*TVEpisodeGroups, error) {
    availableOptions := map[string]struct{} {
        "language":             {},
    }
    var structure TVEpisodeGroups
    result, err := client.sendGetRequest(
        client.buildRequestUrl("tv", id, "episode_groups", options, availableOptions),
        &structure,
    )
    return result.(*TVEpisodeGroups), err
}

// Get the external ids for a TV show. We currently support the following external sources.
// https://developers.themoviedb.org/3/tv/get-tv-external-ids
func (client *TheMovieDatabase) GetTVExternalIDs(id uint64, options map[string]string) (*TVExternalIDs, error) {
    var structure TVExternalIDs
    result, err := client.sendGetRequest(
        client.buildRequestUrl("tv", id, "external_ids", options, nil),
        &structure,
    )
    return result.(*TVExternalIDs), err
}

// Querying images with a language parameter will filter the results.
// If you want to include a fallback language (especially useful for backdrops) you can use the include_image_language parameter.
// https://developers.themoviedb.org/3/tv/get-tv-images
func (client *TheMovieDatabase) GetTVImages(id uint64, options map[string]string) (*TVImages, error) {
    availableOptions := map[string]struct{} {
        "language":                     {},
        "append_to_response":           {},
        "include_image_language":       {},
    }
    var structure TVImages
    result, err := client.sendGetRequest(
        client.buildRequestUrl("tv", id, "images", options, availableOptions),
        &structure,
    )
    return result.(*TVImages), err
}

// Get the keywords that have been added to a TV show.
// https://developers.themoviedb.org/3/tv/get-tv-keywords
func (client *TheMovieDatabase) GetTVKeywords(id uint64, options map[string]string) (*TVKeywords, error) {
    availableOptions := map[string]struct{} {
        "append_to_response":           {},
    }
    var structure TVKeywords
    result, err := client.sendGetRequest(
        client.buildRequestUrl("tv", id, "keywords", options, availableOptions),
        &structure,
    )
    return result.(*TVKeywords), err
}

// Get a list of the translations that exist for a TV show.
// https://developers.themoviedb.org/3/tv/get-tv-translations
func (client *TheMovieDatabase) GetTVTranslations(id uint64, options map[string]string) (*TVTranslations, error) {
    availableOptions := map[string]struct{} {
        "append_to_response":           {},
    }
    var structure TVTranslations
    result, err := client.sendGetRequest(
        client.buildRequestUrl("tv", id, "translations", options, availableOptions),
        &structure,
    )
    return result.(*TVTranslations), err
}

// Get the videos that have been added to a TV show.
// https://developers.themoviedb.org/3/tv/get-tv-videos
func (client *TheMovieDatabase) GetTVVideos(id uint64, options map[string]string) (*TVVideos, error) {
    availableOptions := map[string]struct{} {
        "language":                     {},
        "append_to_response":           {},
    }
    var structure TVVideos
    result, err := client.sendGetRequest(
        client.buildRequestUrl("tv", id, "videos", options, availableOptions),
        &structure,
    )
    return result.(*TVVideos), err
}
