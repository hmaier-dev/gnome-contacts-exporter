package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var testdb string = "testdata/contacts.db"
var db *sql.DB
var err error

type categories struct {
	category    string
	unusedvalue string
	refs        int
}

type folder_id struct {
	uid                   string
	Rev                   string
	file_as               string
	file_as_localized     string
	nickname              string
	full_name             string
	given_name            string
	given_name_localized  string
	family_name           string
	family_name_localized string
	is_list               string
	list_show_addresses   int
	wants_html            int
	x509Cert              int
	pgpCert               int
	vcard                 string
	bdata                 string
}

type folder_id_email_list struct {
	uid   string
	value string
}

type folder_id_phone_list struct {
	uid   string
	value string
}

type folders struct {
	folder_id   string
	version     int
	multivalues string
	lc_collate  string
	countrycode string
}

type keys struct {
	key       string
	value     string
	folder_id string
}

func init() {
	source := testdb
	db, err = sql.Open("sqlite3", source)
	if err != nil {
		log.Fatalf("%s when loading %s", err, source)
	}
}

// Export exports all vCards from the database
func Export() {

	rows := queryCommand("select vcard from folder_id;")
    recv := []folder_id{}
	for rows.Next() {
		var row folder_id
		err := rows.Scan(&row.vcard)
		if err != nil {
			log.Fatal("problems scanning the database", err)
		}
        recv = append(recv, row)
	}

}

func LoadSqlite(source string) {
	rows := queryCommand("select vcard from folder_id;")
	cols, err := rows.Columns()

	if err != nil {
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

func queryCommand(q string) *sql.Rows {
	rows, err := db.Query(q)
	if err != nil {
		log.Fatalf("%#v\n", err)
	}
	return rows

}
