package config

import (
    "encoding/json"
    "os"
    "log"
    "path/filepath"
    "domio/components/logger"
    "gopkg.in/alecthomas/kingpin.v2"
    "strings"
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

func init() {
    Config = LoadConfig()
}
func LoadConfig() Configuration {
    defaultConfigFilePath := "./config.dev.json"
    var (
        app = kingpin.New("domio", "Domio domains rental server.")
        debug = app.Flag("debug", "Enable debug mode.").Bool()
        serverIP = app.Flag("server", "Server address.").Default("127.0.0.1").IP()

        init = app.Command("init", "Init a new config file.")
        //registerNick = register.Arg("nick", "Nickname for user.").String()
        //registerName = register.Arg("name", "Name of user.").String()

        start = app.Command("start", "Start server.")
        check = app.Command("check", "Check config.")
        postImage = check.Flag("image", "Image to post.").File()
        //        postChannel = post.Arg("channel", "Channel to post to.").Required().String()
        configFilePath = start.Arg("config", "Config file to use.").Default(defaultConfigFilePath).Strings()
    )

    log.Print("*****************************************************")
    log.Print(init)
    log.Print(*debug)
    log.Print(*serverIP)
    //log.Print(*registerName)
    log.Print(*check)
    log.Print(*start)
    log.Print("*****************************************************")

    switch kingpin.MustParse(app.Parse(os.Args[1:])) {
    // Register user
    /*
    case register.FullCommand():
        println(*registerNick)
    */

    case start.FullCommand():
        if *postImage != nil {
        }
        text := strings.Join(*configFilePath, " ")
        println("Post:", text)
    }

    dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

    log.Println("Loading config...");
    log.Println(dir);

    configFile := dir + defaultConfigFilePath
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