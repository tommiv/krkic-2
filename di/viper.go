package di

import (
    "fmt"
    "gopkg.in/spf13/viper.v0"
)

func SetupViper(appDataFolder string) {
    viper.SetConfigName("config")
    viper.AddConfigPath(appDataFolder)

    viper.SetDefault("log_level", "debug")
    viper.SetDefault("app_data_folder", appDataFolder)

    err := viper.ReadInConfig()
    if err != nil {
        panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }
}
