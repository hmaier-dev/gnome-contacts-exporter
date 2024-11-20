package arguments

import (
	"errors"
	"flag"
	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/db"
	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/daemon"
)

// Sets arguments for the needed variables
func Define() []error{
    ret := []error{}

    source := flag.String("source", "", "Sets the location from where to read the database.")
	destination := flag.String("destination", "", "Sets the location where to write the vCard-file.")
    permit := flag.Bool("daemon", false, "Runs the exporter as daemon, which listens for changes in the database.")
    flag.Parse()
    
    // Checking mandatory flags
    if *source == "" {
        ret = append(ret, errors.New("Source hasn't been set."))
    }

    if *destination == ""{
        ret = append(ret, errors.New("Destination hasn't been set."))
    }

    if len(ret) > 0{
        return ret

    }else{
        db.Destination = *destination
        db.Source = *source
        daemon.Permit = *permit
        return nil
    }


}
