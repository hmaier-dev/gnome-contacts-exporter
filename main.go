package main

import (
	"log"
	"os"
	"path/filepath"

	args "github.com/hmaier-dev/gnome-contacts-exporter/pkg/arguments"
	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/daemon"
	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/db"
	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/vcf"
)

func main(){
    // index 0 is always the name of the program
    err := args.Define(os.Args[1:])
    if err != nil {
        for _, e := range err{
            log.Printf("Error: %s \n", e)
        }
        os.Exit(1)
    }
    sourceExt := filepath.Ext(args.Source)
    destExt := filepath.Ext(args.Destination)
    
    if sourceExt == ".db" && destExt == ".vcf"{
        db.Export(args.Source, args.Destination)

        // db.Import -> []vcard
        // vcf.Export -> write to file
    } else if sourceExt == ".vcf" && destExt == ".db"{
        vcf.Import(args.Source)

        // vcf.Import -> []vcard
        // db.Export -> connect to db
    }else{
        log.Fatalf("Exporting from %s to %s is not supported.", args.Source, args.Destination)
    }
    
    if args.Permit != false {
        daemon.Run()
    }
}
