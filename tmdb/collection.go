package tmdb

type Collection struct {
    ID                          uint64      `json:"id"`
    Name                        string      `json:"name"`
    Overview                    string      `json:"overview"`
    PosterPath                  string      `json:"poster_path"`
    BackdropPath                string      `json:"backdrop_path"`
    Parts                       struct      {
        Adult                   bool        `json:"adult"`
        BackdropPath            string      `json:"backdrop_path"`
        GenreIDs                []uint64    `json:"genre_ids"`
        ID                      uint64      `json:"id"`
        OriginalLanguage        string      `json:"original_language"`
        OriginalTitle           string      `json:"original_title"`
        Overview                string      `json:"overview"`
        ReleaseDate             string      `json:"release_date"`
        PosterPath              string      `json:"poster_path"`
        Popularity              float64     `json:"popularity"`
        Title                   string      `json:"title"`
        Video                   bool        `json:"video"`
        VoteAverage             float64     `json:"vote_average"`
        VoteCount               uint64      `json:"vote_count"`
    }
}

type CollectionShort struct {
    ID              uint64      `json:"id"`
    Name            string      `json:"name"`
    PosterPath      string      `json:"poster_path"`
    BackdropPath    string      `json:"backdrop_path"`
}