package tmdb

import (
    "encoding/json"
    "github.com/plexmediamanager/micro-database/models"
    "strings"
)

// Process movie entry and convert it to the database ready structure
func ProcessMovie(movie *Movie) *models.Movie {
    return &models.Movie {
        ID:                             movie.ID,
        Title:                          movie.Title,
        OriginalTitle:                  movie.OriginalTitle,
        OriginalLanguage:               movie.OriginalLanguage,
        Languages:                      extractLanguages(movie.SpokenLanguages),
        Overview:                       movie.Overview,
        Tagline:                        movie.Tagline,
        Genres:                         extractGenres(movie.Genres),
        Homepage:                       movie.Homepage,
        Runtime:                        movie.Runtime,
        Status:                         extractStatus(movie.Status),
        Adult:                          movie.Adult,
        ImdbId:                         movie.ImdbID,
        ReleaseDate:                    movie.ReleaseDate,
        ProductionCompanies:            extractProductionCompanies(movie.ProductionCompanies),
        ProductionCountries:            extractProductionCountries(movie.ProductionCountries),
        VoteAverage:                    movie.VoteAverage,
        VoteCount:                      movie.VoteCount,
        Popularity:                     movie.Popularity,
        Budget:                         movie.Budget,
        Revenue:                        movie.Revenue,
        Backdrop:                       clearImagePath(movie.BackdropPath),
        Poster:                         clearImagePath(movie.PosterPath),
    }
}

// Extract languages
func extractLanguages (languages []Language) json.RawMessage {
    temp := make([]string, 0)
    for _, language := range languages {
        temp = append(temp, language.ISO6391)
    }
    result, err := json.Marshal(temp)
    if err != nil {
        return []byte{}
    }
    return result
}

// Extract genres
func extractGenres(genres []Genre) json.RawMessage {
    temp := make([]uint64, 0)
    for _, genre := range genres {
        temp = append(temp, genre.ID)
    }
    result, err := json.Marshal(temp)
    if err != nil {
        return []byte{}
    }
    return result
}

// Extract production companies
func extractProductionCompanies(companies []ProductionCompany) json.RawMessage {
    temp := make([]uint64, 0)
    for _, company := range companies {
        temp = append(temp, company.ID)
    }
    result, err := json.Marshal(temp)
    if err != nil {
        return []byte{}
    }
    return result
}

// Extract production countries
func extractProductionCountries(countries []ProductionCountry) json.RawMessage {
    temp := make([]string, 0)
    for _, country := range countries {
        temp = append(temp, country.ISO31661)
    }
    result, err := json.Marshal(temp)
    if err != nil {
        return []byte{}
    }
    return result
}

// Extract status
func extractStatus(statusString string) uint64 {
    var result uint64
    switch statusString {
        case "Released":
            result = 1
        case "Running":
            result = 2
        case "Ended":
            result = 3
        case "Canceled":
            result = 4
        case "Unknown":
            result = 5
        default:
            result = 0
    }
    return result
}

// Clear image path
func clearImagePath(path string) string {
    return strings.TrimPrefix(path, "/")
}
