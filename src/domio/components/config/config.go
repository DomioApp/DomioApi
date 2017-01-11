package config

import (
    "encoding/json"
    "os"
    "log"
    "path/filepath"
    "domio/components/logger"
)

type Configuration struct {
    AWS_ACCESS_KEY_ID     string `json:"AWS_ACCESS_KEY_ID"`
    AWS_SECRET_ACCESS_KEY string `json:"AWS_SECRET_ACCESS_KEY"`
    DOMIO_DB_USER         string `json:"DOMIO_DB_USER"`
    DOMIO_DB_NAME         string `json:"DOMIO_DB_NAME"`
    DOMIO_DB_PASSWORD     string `json:"DOMIO_DB_PASSWORD"`
    PORT                  int    `json:"PORT"`
    ENV                   string `json:"ENV"`
}

var Config Configuration

func LoadConfig() Configuration {
    dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

    log.Println("Loading config...");
    log.Println(dir);

    defaultConfigFilePath := "config.dev.json"
    configFile := dir + "/" + defaultConfigFilePath
    //configFile := "C:\\Users\\sbasharov\\WebstormProjects\\domio\\DomioApi\\bin\\" + defaultConfigFilePath
    //configFile := "/Users/sergeibasharov/WebstormProjects/DomioApiGo/deploy/config/config.json"
    if _, err := os.Stat(configFile); os.IsNotExist(err) {
        logger.Logger.Crit("Config file doesn't exist, exitting...")
        log.Fatalln("error:", err)
    }

    file, _ := os.Open(configFile)

    decoder := json.NewDecoder(file)
    config := Configuration{}
    err := decoder.Decode(&config)
    if err != nil {
        logger.Logger.Crit("Config file couldn't be loaded, exitting...")
        //logger.Logger.Crit(file)
        log.Fatalln("error:", err)
    }
    logger.Logger.Info("Config loaded")
    return config
}