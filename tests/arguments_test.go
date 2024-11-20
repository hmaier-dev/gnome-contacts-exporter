package tests

import (
	"log"
	"os"
	"testing"

	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/arguments"
)


func TestArgumentsWithoutArguments(t *testing.T){

    err := arguments.Define() 
    if err == nil{
        log.Fatalf("Error: %s", err)
    }
}

func TestArgumentsWithArguments(t *testing.T){

    os.Args = append(os.Args, "--destination")
    os.Args = append(os.Args, "contacts.vcf")
    os.Args = append(os.Args, "--source")
    os.Args = append(os.Args, "contacts.db")
    err := arguments.Define() 
    if err != nil{
        log.Fatalf("Error: %s", err)
    }
}
