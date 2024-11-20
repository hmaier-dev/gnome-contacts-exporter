package main

import (
	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/arguments"
	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/db"
)

func main(){
    arguments.Define()
    // Fails if --source and --destination haven't been defined
    db.Export()
}
