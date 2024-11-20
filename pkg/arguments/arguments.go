package arguments

import (
	"flag"
	"fmt"
	"os"
	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/db"
)

// Sets arguments for the needed variables
func Define(){
    source := flag.String("source", "", "Sets the location from where to read the database.")
	destination := flag.String("destination", "", "Sets the location where to write the vCard-file.")
    flag.Parse()

    if *source == "" {
        fmt.Println("Error: Both source and destination are mandatory.")
        fmt.Println("Usage:")
        flag.PrintDefaults()
        os.Exit(1)
    }

    if *destination == ""{
        fmt.Println("Error: Both source and destination are mandatory.")
        fmt.Println("Usage:")
        flag.PrintDefaults()
        os.Exit(1)
    }

    db.Destination = *destination
    db.Source = *source

}
