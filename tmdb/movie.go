package tmdb

import (
    "github.com/plexmediamanager/micro-database/models"
)

type Movie struct {
    Adult                       bool                `json:"adult"`
    BackdropPath                string              `json:"backdrop_path"`
    Budget                      uint64              `json:"budget"`
    Genres                      []Genre             `json:"genres"`
    Homepage                    string              `json:"homepage"`
    ID                          uint64              `json:"id"`
    ImdbID                      string              `json:"imdb_id"`
    OriginalLanguage            string              `json:"original_language"`
    OriginalTitle               string              `json:"original_title"`
    Overview                    string              `json:"overview"`
    Popularity                  float64             `json:"popularity"`
    PosterPath                  string              `json:"poster_path"`
    ProductionCompanies         []ProductionCompany `json:"production_companies"`
    ProductionCountries         []ProductionCountry `json:"production_countries"`
    ReleaseDate                 string              `json:"release_date"`
    Revenue                     uint64              `json:"revenue"`
    Runtime                     uint64              `json:"runtime"`
    SpokenLanguages             []Language          `json:"spoken_languages"`
    Status                      string              `json:"status"`
    Tagline                     string              `json:"tagline"`
    Title                       string              `json:"title"`
    Video                       bool                `json:"video"`
    VoteAverage                 float64             `json:"vote_average"`
    VoteCount                   uint64              `json:"vote_count"`
}

type MovieAlternativeTitles struct {
    ID                          uint64              `json:"id"`
    Titles                      []struct            {
        ISO31661                string              `json:"iso_3166_1"`
        Title                   string              `json:"title"`
        Type                    string              `json:"type"`
    }                                               `json:"titles"`
}

type MovieChanges struct {
    Changes                     []struct            {
        Key                     string              `json:"key"`
        Items                   []struct            {
            ID                  string              `json:"id"`
            Action              string              `json:"action"`
            Time                string              `json:"time"`
            ISO6391             string              `json:"iso_639_1"`
            Value               string              `json:"value"`
            OriginalValue       string              `json:"original_value"`
        }                                           `json:"items"`
    }                                               `json:"changes"`
}

type MovieCredits struct {
    ID                          uint64              `json:"id"`
    Cast                        []struct            {
        CastID                  uint64              `json:"cast_id"`
        Character               string              `json:"string"`
        CreditID                string              `json:"credit_id"`
        Gender                  uint64              `json:"gender"`
        ID                      uint64              `json:"id"`
        Name                    string              `json:"name"`
        Order                   uint64              `json:"order"`
        ProfilePath             string              `json:"profile_path"`
    }                                               `json:"cast"`
    Crew                        []struct            {
        CreditID                string              `json:"credit_id"`
        Department              string              `json:"department"`
        Gender                  uint64              `json:"gender"`
        ID                      uint64              `json:"id"`
        Job                     string              `json:"job"`
        Name                    string              `json:"name"`
        ProfilePath             string              `json:"profile_path"`
    }                                               `json:"crew"`
}

type MovieExternalIDs struct {
    ID                          uint64              `json:"id"`
    IMDB                        string              `json:"imdb_id"`
    Facebook                    string              `json:"facebook_id"`
    Instagram                   string              `json:"instagram_id"`
    Twitter                     string              `json:"twitter_id"`
}

type MovieImages struct {
    ID                          uint64              `json:"id"`
    Backdrops                   []CommonImage       `json:"backdrops"`
    Posters                     []CommonImage       `json:"posters"`
}

type MovieKeywords struct {
    ID                          uint64              `json:"id"`
    Keywords                    []struct            {
        ID                      uint64              `json:"id"`
        Name                    string              `json:"name"`
    }                                               `json:"keywords"`
}

type MovieReleaseDates struct {
    ID                          uint64              `json:"id"`
    Results                     []struct            {
        ISO31661                string              `json:"iso_3166_1"`
        ReleaseDates            []struct            {
            Certification       string              `json:"certification"`
            ISO6391             string              `json:"iso_639_1"`
            ReleaseDate         string              `json:"release_date"`
            Type                uint64              `json:"type"`
            Note                string              `json:"note"`
        }                                           `json:"release_dates"`
    }                                               `json:"results"`
}

type MovieVideos struct {
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

type MovieTranslations struct {
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

// Get information for specific movie by its ID
// https://developers.themoviedb.org/3/movies/get-movie-details
func (client *TheMovieDatabase) GetMovieInformation(id uint64, options map[string]string) (*Movie, error) {
    availableOptions := map[string]struct{} {
        "language":             {},
        "append_to_response":   {},
    }
    var movie Movie
    result, err := client.sendGetRequest(
        client.buildRequestUrl("movie", id, nil, options, availableOptions),
        &movie,
    )
    return result.(*Movie), err
}

// Get information for specific movie by its ID (Ready for the database)
func (client *TheMovieDatabase) GetMovieInformationForDatabase(id uint64, options map[string]string) (*models.Movie, error) {
    result, err := client.GetMovieInformation(id, options)
    if err != nil {
        return nil, err
    }
    return ProcessMovie(result), nil
}

// Get alternative titles for the movie by its ID
// https://developers.themoviedb.org/3/movies/get-movie-alternative-titles
func (client *TheMovieDatabase) GetMovieAlternativeTitles(id uint64, options map[string]string) (*MovieAlternativeTitles, error) {
    availableOptions := map[string]struct{} {
        "country":              {},
        "append_to_response":   {},
    }
    var titles MovieAlternativeTitles
    result, err := client.sendGetRequest(
        client.buildRequestUrl("movie", id, "alternative_titles", options, availableOptions),
        &titles,
    )
    return result.(*MovieAlternativeTitles), err
}

// Get the changes for a movie. By default only the last 24 hours are returned.
// https://developers.themoviedb.org/3/movies/get-movie-changes
func (client *TheMovieDatabase) GetMovieChanges(id uint64, options map[string]string) (*MovieChanges, error) {
    availableOptions := map[string]struct{} {
        "start_date":   {},
        "end_date":     {},
    }
    var changes MovieChanges
    result, err := client.sendGetRequest(
        client.buildRequestUrl("movie", id, "changes", options, availableOptions),
        &changes,
    )
    return result.(*MovieChanges), err
}

// Get the cast and crew for a movie.
// https://developers.themoviedb.org/3/movies/get-movie-credits
func (client *TheMovieDatabase) GetMovieCredits(id uint64, options map[string]string) (*MovieCredits, error) {
    availableOptions := map[string]struct{} {
        "append_to_response":   {},
    }
    var credits MovieCredits
    result, err := client.sendGetRequest(
        client.buildRequestUrl("movie", id, "credits", options, availableOptions),
        &credits,
    )
    return result.(*MovieCredits), err
}

// Get the external ids for a movie.
// https://developers.themoviedb.org/3/movies/get-movie-external-ids
func (client *TheMovieDatabase) GetMovieExternalIDs(id uint64, options map[string]string) (*MovieExternalIDs, error) {
    var externalIDs MovieExternalIDs
    result, err := client.sendGetRequest(
        client.buildRequestUrl("movie", id, "external_ids", options, nil),
        &externalIDs,
    )
    return result.(*MovieExternalIDs), err
}

// Get the images that belong to a movie.
// Querying images with a language parameter will filter the results. If you want to include a fallback language (especially useful for backdrops) you can use the include_image_language parameter.
// https://developers.themoviedb.org/3/movies/get-movie-images
func (client *TheMovieDatabase) GetMovieImages(id uint64, options map[string]string) (*MovieImages, error) {
    availableOptions := map[string]struct{} {
        "language":                     {},
        "append_to_response":           {},
        "include_image_language":       {},
    }
    var movieImages MovieImages
    result, err := client.sendGetRequest(
        client.buildRequestUrl("movie", id, "images", options, availableOptions),
        &movieImages,
    )
    return result.(*MovieImages), err
}

// Get the keywords that have been added to a movie.
// https://developers.themoviedb.org/3/movies/get-movie-keywords
func (client *TheMovieDatabase) GetMovieKeywords(id uint64, options map[string]string) (*MovieKeywords, error) {
    availableOptions := map[string]struct{} {
        "append_to_response":           {},
    }
    var movieKeywords MovieKeywords
    result, err := client.sendGetRequest(
        client.buildRequestUrl("movie", id, "keywords", options, availableOptions),
        &movieKeywords,
    )
    return result.(*MovieKeywords), err
}

// Get the release date along with the certification for a movie.
// https://developers.themoviedb.org/3/movies/get-movie-release-dates
func (client *TheMovieDatabase) GetMovieReleaseDates(id uint64, options map[string]string) (*MovieReleaseDates, error) {
    availableOptions := map[string]struct{} {
        "append_to_response":           {},
    }
    var releaseDates MovieReleaseDates
    result, err := client.sendGetRequest(
        client.buildRequestUrl("movie", id, "release_dates", options, availableOptions),
        &releaseDates,
    )
    return result.(*MovieReleaseDates), err
}

// Get the videos that have been added to a movie.
// https://developers.themoviedb.org/3/movies/get-movie-videos
func (client *TheMovieDatabase) GetMovieVideos(id uint64, options map[string]string) (*MovieVideos, error) {
    availableOptions := map[string]struct{} {
        "language":                     {},
        "append_to_response":           {},
    }
    var structure MovieVideos
    result, err := client.sendGetRequest(
        client.buildRequestUrl("movie", id, "videos", options, availableOptions),
        &structure,
    )
    return result.(*MovieVideos), err
}

// Get a list of translations that have been created for a movie.
// https://developers.themoviedb.org/3/movies/get-movie-translations
func (client *TheMovieDatabase) GetMovieTranslations(id uint64, options map[string]string) (*MovieTranslations, error) {
    availableOptions := map[string]struct{} {
        "append_to_response":           {},
    }
    var structure MovieTranslations
    result, err := client.sendGetRequest(
        client.buildRequestUrl("movie", id, "translations", options, availableOptions),
        &structure,
    )
    return result.(*MovieTranslations), err
}