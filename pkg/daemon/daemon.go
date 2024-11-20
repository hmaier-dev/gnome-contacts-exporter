package daemon

import "fmt"

var Permit bool

func Run(){
    if Permit != false {
        fmt.Println("Daemon is running!")
    }

}
