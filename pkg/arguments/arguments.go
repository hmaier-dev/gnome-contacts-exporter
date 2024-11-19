package arguments

import (
	"fmt"
	"os"
	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/db"
)

func Define(){
    if len(os.Args) > 1{
        for _, s := range os.Args[1:]{ // index 0 is the name of program, so slice it away
            switch s {
            case "--exporter":
                db.Export()
            default:
                fmt.Printf("Argument unknown: %s \n", s)
                os.Exit(0)
            }
        }

    }

}
