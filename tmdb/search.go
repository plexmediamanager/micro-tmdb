package tmdb

import format "fmt"

type SearchCompanies struct {
    Page                        uint64          `json:"page"`
    Results                     []struct        {
        ID                      uint64          `json:"id"`
        LogoPath                string          `json:"logo_path"`
        Name                    string          `json:"name"`
    }                                           `json:"results"`
    TotalPages                  uint64          `json:"total_pages"`
    TotalResults                uint64          `json:"total_results"`
}

type SearchMovies struct {
    Page                        uint64          `json:"page"`
    Results                     []struct        {
        PosterPath              string          `json:"poster_path"`
        Adult                   bool            `json:"adult"`
        Overview                string          `json:"overview"`
        ReleaseDate             string          `json:"release_date"`
        GenreIDs                []uint64        `json:"genre_ids"`
        ID                      uint64          `json:"id"`
        OriginalTitle           string          `json:"original_title"`
        OriginalLanguage        string          `json:"original_language"`
        Title                   string          `json:"title"`
        BackdropPath            string          `json:"backdrop_path"`
        Popularity              float64         `json:"popularity"`
        VoteCount               uint64          `json:"vote_count"`
        Video                   bool            `json:"video"`
        VoteAverage             float64         `json:"vote_average"`
    }                                           `json:"results"`
    TotalPages                  uint64          `json:"total_pages"`
    TotalResults                uint64          `json:"total_results"`
}

type SearchTV struct {
    Page                        uint64          `json:"page"`
    Results                     []struct        {
        PosterPath              string          `json:"poster_path"`
        Popularity              float64         `json:"popularity"`
        ID                      uint64          `json:"id"`
        BackdropPath            string          `json:"backdrop_path"`
        VoteAverage             float64         `json:"vote_average"`
        Overview                string          `json:"overview"`
        FirstAirDate            string          `json:"first_air_date"`
        OriginCountry           []string        `json:"origin_country"`
        GenreIDs                []uint64        `json:"genre_ids"`
        OriginalLanguage        string          `json:"original_language"`
        VoteCount               uint64          `json:"vote_count"`
        Name                    string          `json:"name"`
        OriginalName            string          `json:"original_name"`
    }                                           `json:"results"`
    TotalPages                  uint64          `json:"total_pages"`
    TotalResults                uint64          `json:"total_results"`
}

// Search for companies.
// https://developers.themoviedb.org/3/search/search-companies
func (client *TheMovieDatabase) SearchCompanies(query string, page uint64, language string) (*SearchCompanies, error) {
    var structure SearchCompanies
    options := map[string]string {
        "query":    query,
        "page":     format.Sprint(page),
        "language": language,
    }

    availableOptions := map[string]struct{} {
        "language":             {},
        "query":                {},
        "page":                 {},
    }

    result, err := client.sendGetRequest(
        client.buildRequestUrl("search/company", nil, nil, options, availableOptions),
        &structure,
    )
    return result.(*SearchCompanies), err
}

// Search for movies.
// https://developers.themoviedb.org/3/search/search-movies
func (client *TheMovieDatabase) SearchMovies(query string, year uint64, options map[string]string) (*SearchMovies, error) {
    var structure SearchMovies
    mergedOptions := make(map[string]string)
    if options != nil {
        for key, value := range options {
            mergedOptions[key] = value
        }
    }

    mergedOptions["include_adult"] = "true"
    mergedOptions["query"] = query
    mergedOptions["year"] = format.Sprint(year)

    availableOptions := map[string]struct{} {
        "language":             {},
        "query":                {},
        "page":                 {},
        "include_adult":        {},
        "region":               {},
        "year":                 {},
        "primary_release_year": {},
    }

    result, err := client.sendGetRequest(
        client.buildRequestUrl("search/movie", nil, nil, mergedOptions, availableOptions),
        &structure,
    )
    return result.(*SearchMovies), err
}

// Search for a TV show.
// https://developers.themoviedb.org/3/search/search-tv-shows
func (client *TheMovieDatabase) SearchTV(query string, year uint64, options map[string]string) (*SearchTV, error) {
    var structure SearchTV
    mergedOptions := make(map[string]string)
    if options != nil {
        for key, value := range options {
            mergedOptions[key] = value
        }
    }

    mergedOptions["query"] = query
    mergedOptions["first_air_date_year"] = format.Sprint(year)

    availableOptions := map[string]struct{} {
        "language":             {},
        "query":                {},
        "page":                 {},
        "first_air_date_year":  {},
    }

    result, err := client.sendGetRequest(
        client.buildRequestUrl("search/tv", nil, nil, mergedOptions, availableOptions),
        &structure,
    )
    return result.(*SearchTV), err
}
