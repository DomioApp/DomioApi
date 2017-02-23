package utils

import (
    "log"
    "github.com/fatih/color"
)

func ShowError(err interface{}) {
    color.Set(color.FgHiRed)
    log.Print(err)
    color.Unset()

}