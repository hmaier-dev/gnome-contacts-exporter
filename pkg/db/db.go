package db

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)


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

var testdb string = "testdata/contacts.db"

var db *sql.DB
var err error

func init() {
	source := testdb
	db, err = sql.Open("sqlite3", source)
	if err != nil {
		log.Fatalf("%s when loading %s", err, source)
	}
}

// Export exports all vCards from the database
func Export(source string, dest string) {
    if source == ""{
        panic("Source is not defined.")
    }
    if dest == ""{
        panic("Destination is not defined.")
    }

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

func queryCommand(q string) *sql.Rows {
	rows, err := db.Query(q)
	if err != nil {
		log.Fatalf("%#v\n", err)
	}
	return rows

}
