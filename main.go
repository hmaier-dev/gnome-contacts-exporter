package main

import (
	"fmt"
	"github.com/hmaier-dev/contacts_converter/pkg/loading"
	"github.com/hmaier-dev/go-vcard"
)

func main(){

    p := "testdata/contacts.vcf"
    c := loading.LoadVFC(p)
    for _, vc := range c{
        fmt.Println(vc.AllValues(vcard.FieldFormattedName))
        fmt.Println(vc.AllValues(vcard.FieldAddress))
        fmt.Println(vc.AllValues(vcard.FieldEmail))
    }
    // buf := writer.(*bytes.Buffer)
    // data := buf.Bytes()
    // os.WriteFile("../testdata/contacts2.vcf", data, 0744)

}
