package loading

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func LoadVFC(source string){

}

func LoadSqlite(source string){
    db, err := sql.Open("sqlite3", source)
    if err != nil{
        fmt.Println(err)
    }
    // rows, err := db.Query("Select name from sqlite_master WHERE type='table';")
    rows, err := db.Query("Select * from folder_id_phone_list;")
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
