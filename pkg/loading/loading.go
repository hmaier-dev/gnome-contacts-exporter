package loading

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"log"
    "os"
    "io"
    "github.com/hmaier-dev/go-vcard"

)


func Export(){
    source := "testdata/contacts.db"
    LoadSqlite(source)
}

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

type Folder_ID struct {
	UID                  string 
	Rev                  string 
	FileAs               string 
	FileAsLocalized      string 
	Nickname             string 
	FullName             string 
	GivenName            string 
	GivenNameLocalized   string 
	FamilyName           string 
	FamilyNameLocalized  string 
	IsList               int    
	ListShowAddresses    int    
	WantsHTML            int    
	X509Cert             int    
	PgpCert              int    
	VCard                string 
	BData                string 
}

func LoadSqlite(source string){
    db, err := sql.Open("sqlite3", source)
    if err != nil{
        fmt.Println(err)
    }
    rows, err := db.Query("Select * from folder_id WHERE vcard != NULL;")
    if err != nil{
        log.Fatalf("%#v\n", err)
    } 
    cols, err := rows.Columns()

    if err != nil{
        log.Fatalf("%#v\n", err)
    } 
	rawResult := make([][]byte, len(cols)) // [row][values] -> e.g. row: [[value][value][value]]
	dest := make([]interface{}, len(cols)) // .Scan() needs []any as result type
	allRows := make([][]string, 0)
	for i := range cols {
		dest[i] = &rawResult[i] // mapping dest indices to byte slice
	}
	for rows.Next() {
		err := rows.Scan(dest...)
		if err != nil {
			log.Fatal("problems scanning the database", err)
		}
		singleRow := make([]string, len(cols))
		for i, raw := range rawResult {
			singleRow[i] = string(raw) // from byte to string
			//fmt.Printf("%v -> %v \n", i, singleRow)
		}
		allRows = append(allRows, singleRow)
	}
    
    fmt.Printf("%v\n", allRows)
}

func WriteVCF(cards []vcard.Card, w io.Writer){
   
    enc := vcard.NewEncoder(w)
    for _, c := range cards{
        enc.Encode(c)
        w.Write([]byte("\n"))
    }

}


