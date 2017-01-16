package logger

import (
    "log"
    syslog "github.com/RackSec/srslog"
)

var Logger *syslog.Writer

func init() {
    var err error
    Logger, err = syslog.Dial("udp", "logs5.papertrailapp.com:18422", syslog.LOG_EMERG | syslog.LOG_KERN, "domio-api")
    if err != nil {
        log.Fatal("failed to dial syslog")
    }
}

var Crit = Logger.Crit
var Info = Logger.Info
var Err = Logger.Err