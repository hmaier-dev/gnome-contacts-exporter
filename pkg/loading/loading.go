package loading

import (
	"log"
    "os"
    "io"
    "github.com/hmaier-dev/go-vcard"

)


// LoadVFC opens a vcf-file from the given path and stores it in a Card.
func LoadVFC(source string) []vcard.Card {
    f, err := os.Open(source)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    var out []vcard.Card 
    dec := vcard.NewDecoder(f)
    for {
        card, err := dec.Decode()
        if err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        }
        out = append(out, card)
    }
    return out
}



func WriteVCF(cards []vcard.Card, w io.Writer){
   
    enc := vcard.NewEncoder(w)
    for _, c := range cards{
        enc.Encode(c)
        w.Write([]byte("\n"))
    }

}


