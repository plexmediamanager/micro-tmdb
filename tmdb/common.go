package tmdb

type CommonImage struct {
    AspectRatio                         float64         `json:"aspect_ratio"`
    FilePath                            string          `json:"file_path"`
    Height                              uint64          `json:"height"`
    ISO6391                             string          `json:"iso_639_1"`
    VoteAverage                         float64         `json:"vote_average"`
    VoteCount                           uint64          `json:"vote_count"`
    Width                               uint64          `json:"width"`
}
