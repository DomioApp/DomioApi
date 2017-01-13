package config

import (
    "encoding/json"
    "os"
    "log"
    "domio/components/logger"
)

type Configuration struct {
    AWS_ACCESS_KEY_ID     string `json:"AWS_ACCESS_KEY_ID"`
    AWS_SECRET_ACCESS_KEY string `json:"AWS_SECRET_ACCESS_KEY"`
    DOMIO_DB_NAME         string `json:"DOMIO_DB_NAME"`
    DOMIO_DB_USER         string `json:"DOMIO_DB_USER"`
    DOMIO_DB_PASSWORD     string `json:"DOMIO_DB_PASSWORD"`
    PORT                  uint    `json:"PORT"`
    ENV                   string `json:"ENV"`
}

type AppStatus struct {
    Buildstamp string `json:"app_buildstamp"`
    BuildAgo   string `json:"app_buildago"`
    Hash       string `json:"app_hash"`
    Version    string `json:"app_version"`
}

var AppStatusInfo AppStatus
var Config Configuration
var ConfigPath = "/usr/local/domio"

func LoadConfig() error {
    //dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

    configFile := "/usr/local/domio/config.json"
    log.Print(configFile)
    //configFile := "C:\\Users\\sbasharov\\WebstormProjects\\domio\\DomioApi\\bin\\" + defaultConfigFilePath
    //configFile := "/Users/sergeibasharov/WebstormProjects/DomioApiGo/deploy/config/config.json"
    if _, err := os.Stat(configFile); os.IsNotExist(err) {
        logger.Logger.Crit("Config file doesn't exist, exitting...")
        log.Fatalln("error:", err)
    }

    file, _ := os.Open(configFile)

    decoder := json.NewDecoder(file)
    err := decoder.Decode(&Config)
    if err != nil {
        logger.Logger.Crit("Config file couldn't be loaded, exitting...")
        //logger.Logger.Crit(file)
        log.Fatalln("error:", err)
    }

    return nil
}