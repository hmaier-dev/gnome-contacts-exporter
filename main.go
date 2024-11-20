package main

import (
	"fmt"
	"os"

	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/arguments"
	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/daemon"
	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/db"
)

func main(){
    // Fails if --source and --destination haven't been defined
    db.Export()
    daemon.Run()
}
