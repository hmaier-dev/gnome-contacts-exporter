package arguments

import(
    "flag"
    "github.com/hmaier-dev/gnome-contacts-exporter/pkg/loading"
)

func define(){
    flag.Func("exporter", "Export the contacts from sqlite")
    
}
