package arguments

import (
    "gopkg.in/alecthomas/kingpin.v2"
    "os"
    "strings"
    "log"
)

func ProcessArguments() string {
    var (
        app = kingpin.New("domio", "Domio domains rental server.")
        _ = app.Flag("debug", "Enable debug mode.").Bool()
        _ = app.Flag("server", "Server address.").Default("127.0.0.1").IP()

        init = app.Command("init", "Init a new config file.")
        //registerNick = register.Arg("nick", "Nickname for user.").String()
        //registerName = register.Arg("name", "Name of user.").String()

        start = app.Command("start", "Start server.")
        check = app.Command("check", "Check config.")
        postImage = check.Flag("image", "Image to post.").File()
        configFilePath = start.Arg("config", "Config file to use.").Default("").Strings()
    )

    log.Print("*****************************************************")
    log.Print(init)
    //log.Print(*debug)
    //log.Print(*serverIP)
    //log.Print(*registerName)
    //log.Print(*check)
    //log.Print(*start)
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

    return "success"
}