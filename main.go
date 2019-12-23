package main

import (
    "github.com/micro/go-micro/client"
    "github.com/plexmediamanager/service"
    "github.com/plexmediamanager/service/log"
    "time"
)

func main() {
    application := service.CreateApplication()

    err := application.InitializeConfiguration()
    if err != nil {
        log.Panic(err)
    }

    err = application.InitializeMicroService()
    if err != nil {
        log.Panic(err)
    }

    err = application.Service().Client().Init(
        client.PoolSize(10),
        client.Retries(30),
        client.RequestTimeout(1 * time.Second),
    )
    if err != nil {
        log.Panic(err)
    }

    //client := tmdb.Initialize(helpers.GetEnvironmentVariableAsString("TMDB_API_KEY", ""))
    //
    //result, err := client.GetMovieTranslations(475557, nil)
    //if err != nil {
    //    format.Println(err)
    //} else {
    //    result, _ := json.Marshal(result)
    //
    //    format.Println(string(result))
    //}

    go application.StartMicroService()

    service.WaitForOSSignal(1)
}
