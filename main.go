package main

import (
	"fmt"
	"os"

	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/arguments"
	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/daemon"
	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/db"
)

func main(){
    err := arguments.Define()
    if err != nil {
        for _, e := range err{
            fmt.Printf("Error: %s \n", e)
        }
        os.Exit(1)
    }
    // Fails if --source and --destination haven't been defined
    db.Export()
    daemon.Run()
}
