package tmdb

type APIConfiguration struct {
    Images struct {
        BaseURL             string      `json:"base_url"`
        SecureBaseURL       string      `json:"secure_base_url"`
        BackdropSizes       []string    `json:"backdrop_sizes"`
        LogoSizes           []string    `json:"logo_sizes"`
        PosterSizes         []string    `json:"poster_sizes"`
        ProfileSizes        []string    `json:"profile_sizes"`
        StillSizes          []string    `json:"still_sizes"`
    }                                   `json:"images"`
    ChangeKeys              []string    `json:"change_keys"`
}

type Country struct {
    ISO31661                string      `json:"iso_3166_1"`
    EnglishName             string      `json:"english_name"`
}

type Countries []Country

type Language struct {
    ISO6391                 string      `json:"iso_639_1"`
    EnglishName             string      `json:"english_name"`
    Name                    string      `json:"name"`
}

type Languages []Language

type PrimaryTranslations []string

type TimeZone struct {
    ISO31661                string      `json:"iso_3166_1"`
    Zones                   []string    `json:"zones"`
}

type TimeZones []TimeZone

// Get the system wide configuration information.
// https://developers.themoviedb.org/3/configuration/get-api-configuration
func (client *TheMovieDatabase) GetAPIConfiguration() (*APIConfiguration, error) {
    var structure APIConfiguration
    result, err := client.sendGetRequest(
        client.buildRequestUrl("configuration", nil, nil, nil, nil),
        &structure,
    )
    return result.(*APIConfiguration), err
}

// Get the list of countries (ISO 3166-1 tags) used throughout TMDb.
// https://developers.themoviedb.org/3/configuration/get-countries
func (client *TheMovieDatabase) GetCountries() (*Countries, error) {
    var structure Countries
    result, err := client.sendGetRequest(
        client.buildRequestUrl("configuration", nil, "countries", nil, nil),
        &structure,
    )
    return result.(*Countries), err
}

// Get the list of languages (ISO 639-1 tags) used throughout TMDb.
// https://developers.themoviedb.org/3/configuration/get-languages
func (client *TheMovieDatabase) GetLanguages() ([]Language, error) {
    var structure Languages
    _, err := client.sendGetRequest(
        client.buildRequestUrl("configuration", nil, "languages", nil, nil),
        &structure,
    )
    var languages []Language
    for _, language := range structure {
        languages = append(languages, language)
    }
    return languages, err
}

// Get a list of the officially supported translations on TMDb.
// https://developers.themoviedb.org/3/configuration/get-primary-translations
func (client *TheMovieDatabase) GetPrimaryTranslations() (*PrimaryTranslations, error) {
    var structure PrimaryTranslations
    result, err := client.sendGetRequest(
        client.buildRequestUrl("configuration", nil, "primary_translations", nil, nil),
        &structure,
    )
    return result.(*PrimaryTranslations), err
}

// Get the list of timezones used throughout TMDb.
// https://developers.themoviedb.org/3/configuration/get-timezones
func (client *TheMovieDatabase) GetTimeZones() (*TimeZones, error) {
    var structure TimeZones
    result, err := client.sendGetRequest(
        client.buildRequestUrl("configuration", nil, "timezones", nil, nil),
        &structure,
    )
    return result.(*TimeZones), err
}