package tmdb

type Genre struct {
    ID                      uint64              `json:"id"`
    Name                    string              `json:"name"`
}

type ProductionCompany struct {
    ID                      uint64              `json:"id"`
    LogoPath                string              `json:"logo_path"`
    Name                    string              `json:"name"`
    OriginCountry           string              `json:"origin_country"`
}

type ProductionCountry struct {
    ISO31661                string              `json:"iso_3166_1"`
    Name                    string              `json:"name"`
}

type CommonImage struct {
    AspectRatio                         float64         `json:"aspect_ratio"`
    FilePath                            string          `json:"file_path"`
    Height                              uint64          `json:"height"`
    ISO6391                             string          `json:"iso_639_1"`
    VoteAverage                         float64         `json:"vote_average"`
    VoteCount                           uint64          `json:"vote_count"`
    Width                               uint64          `json:"width"`
}
