package tmdb

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