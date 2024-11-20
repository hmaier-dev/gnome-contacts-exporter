package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/arguments"
	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/daemon"
	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/db"
	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/vcf"
	// "github.com/hmaier-dev/gnome-contacts-exporter/pkg/db"
)

func main(){
    err := arguments.Define()
    if err != nil {
        for _, e := range err{
            fmt.Printf("Error: %s \n", e)
        }
        os.Exit(1)
    }
    sourceExt := filepath.Ext(arguments.Source)
    destExt := filepath.Ext(arguments.Destination)
    
    // Export from sqlite into vcf
    if sourceExt == ".db"{
        db.Export(arguments.Source, arguments.Destination)
    }
    // Import vcf into a sqlite-database
    if destExt == ".db"{
        vcf.Import(arguments.Source)
    }
    
    if arguments.Permit != false {
        daemon.Run()
    }
}
