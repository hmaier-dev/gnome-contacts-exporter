package arguments

import (
    "flag"
	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/db"
)

func Define(){
    
    flag.Func("from", "Sets the location from where to read the database.", func(s string) error {
        err := db.SetSource(s)
        if err != nil {
            return err
        }
        return nil
    })
    flag.Func("to", "Sets the location where to write the vCard-file.", func(s string) error {
        err := db.SetDestionation(s)
        if err != nil {
            return err
        }
        return nil
    })

    flag.Parse()
}
