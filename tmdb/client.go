package tmdb

import (
    "encoding/json"
    format "fmt"
    "github.com/plexmediamanager/micro-tmdb/errors"
    "github.com/plexmediamanager/service/helpers"
    "io/ioutil"
    "net/http"
    "strings"
)

type apiErrorResponse struct {
    Code            int         `json:"status_code"`
    Message         string      `json:"status_message"`
}

type TheMovieDatabase struct {
    apiKey              string
    apiVersion          uint64
    apiEndpoint         string
}

// Initialize TheMovieDatabase
func Initialize() *TheMovieDatabase {
    return &TheMovieDatabase {
        apiKey:         helpers.GetEnvironmentVariableAsString("TMDB_API_KEY", ""),
        apiVersion:     3,
        apiEndpoint:    "https://api.themoviedb.org",
    }
}

// Send Get Request
func (client *TheMovieDatabase) sendGetRequest(url string, payload interface{}) (interface{}, error) {
    // TODO: Create proper errors instead of just plain Go errors
    var httpClient http.Client

    response, err := httpClient.Get(url)
    if err != nil {
        return payload, err
    }
    defer response.Body.Close()

    responseBody, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return payload, err
    }

    if response.StatusCode >= 200 && response.StatusCode < 300 {
        err = json.Unmarshal(responseBody, &payload)
        if err != nil {
            return payload, err
        }
        return payload, nil
    }

    var apiError apiErrorResponse
    err = json.Unmarshal(responseBody, &apiError)
    if err != nil {
        return payload, err
    }
    return payload, errors.NetworkAPIResponseError.ToErrorWithArguments(nil, apiError.Code, apiError.Message)
}

// Build request url
func (client *TheMovieDatabase) buildRequestUrl(
    endpoint string,
    item interface{},
    subQuery interface{},
    options map[string]string,
    availableOptions map[string]struct{},
) string {
    requestUrl := format.Sprintf(
        "%s/%d/%s",
        client.apiEndpoint,
        client.apiVersion,
        endpoint,
    )
    if item != nil {
        requestUrl += format.Sprintf("/%v", item)
    }
    if subQuery != nil {
        requestUrl += format.Sprintf("/%v", subQuery)
    }
    return requestUrl + client.buildOptionsString(options, availableOptions)
}

// Build request options string
func (client *TheMovieDatabase) buildOptionsString(options map[string]string, availableOptions map[string]struct{}) string {
    var optionsString strings.Builder
    optionsString.WriteString(format.Sprintf("?api_key=%s", client.apiKey))
    for key, value := range options {
        if _, exists := availableOptions[key]; exists {
            optionsString.WriteString(format.Sprintf("&%s=%s", key, strings.Replace(value, " ", "%20", -1)))
        }
    }
    return optionsString.String()
}