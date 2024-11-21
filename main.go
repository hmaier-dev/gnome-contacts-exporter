package main

import (
	"fmt"
	"os"
    "path/filepath"
	args "github.com/hmaier-dev/gnome-contacts-exporter/pkg/arguments"
	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/daemon"
	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/db"
	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/vcf"
)

func main(){
    err := args.Define()
    if err != nil {
        for _, e := range err{
            fmt.Printf("Error: %s \n", e)
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
        fmt.Printf("Exporting from %s to %s is not supported.", args.Source, args.Destination)
    }
    

    
    if args.Permit != false {
        daemon.Run()
    }
}
