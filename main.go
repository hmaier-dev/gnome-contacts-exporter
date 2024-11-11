package main


import (
    "github.com/hmaier-dev/contacts_converter/pkg/loading"
)

func main(){

    p := "testdata/contacts.vcf"
    loading.LoadVFC(p)

}
