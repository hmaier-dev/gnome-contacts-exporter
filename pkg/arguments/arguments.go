package arguments

import (
	"errors"
	"flag"
	"fmt"
	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/db"
)

// Sets arguments for the needed variables
func Define() error{
    source := flag.String("source", "", "Sets the location from where to read the database.")
	destination := flag.String("destination", "", "Sets the location where to write the vCard-file.")
    flag.Parse()

    if *source == "" {
        fmt.Println("Error: Both source and destination are mandatory.")
        fmt.Println("Usage:")
        flag.PrintDefaults()
        return errors.New("Source hasn't been set.")
    }

    if *destination == ""{
        fmt.Println("Error: Both source and destination are mandatory.")
        fmt.Println("Usage:")
        flag.PrintDefaults()
        return errors.New("Destination hasn't been set.")
    }

    db.Destination = *destination
    db.Source = *source
    return nil

}
