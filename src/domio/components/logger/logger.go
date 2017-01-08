package logger

import (
    "log"
    "log/syslog"
)

var Logger *syslog.Writer

func init() {
    var err error
    Logger, err = syslog.Dial("udp", "logs5.papertrailapp.com:18422", syslog.LOG_EMERG | syslog.LOG_KERN, "domio")
    if err != nil {
        log.Fatal("failed to dial syslog")
    }
}