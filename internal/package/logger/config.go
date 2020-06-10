package logger

import (
    "github.com/spf13/viper"
    "fmt"
)


func ConfigListing(){

    fmt.Println("App Config")
    appmode := viper.GetString("app.mode")
    fmt.Println(appmode)
    port := viper.GetString("app.port")
    fmt.Println(port)
    domain := viper.GetString("app.domain")
    fmt.Println(domain)
    fmt.Println("..........")

    fmt.Println("App Log")
    logLevel := viper.GetString("log.level")
    fmt.Println(logLevel)
    logPath := viper.GetString("log.path")
    fmt.Println(logPath)
    fmt.Println("..........")

    fmt.Println("Redis")
    redis := viper.GetString("redis.addr")
    fmt.Println(redis)
    fmt.Println("..........")
}