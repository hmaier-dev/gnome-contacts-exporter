package arguments

import (
	"errors"
	"flag"
)

var Source string
var Destination string
var Permit bool

// Sets arguments for the needed variables
func Define(args []string) []error{
    ret := []error{}

    permit := flag.Bool("daemon", false, "Runs the exporter as daemon, which listens for changes in the database.")
    source := flag.String("source", "", "Sets the location from where to read the database.")
	destination := flag.String("destination", "", "Sets the location where to write the vCard-file.")
    flag.CommandLine.Parse(args)

    // Checking mandatory flags
    if *source == "" {
        ret = append(ret, errors.New("Source hasn't been set."))
    }

    if *destination == ""{
        ret = append(ret, errors.New("Destination hasn't been set."))
    }

    if len(ret) > 0{
        return ret

    }
    
    Permit = *permit
    Source = *source
    Destination = *destination
    return nil

}
