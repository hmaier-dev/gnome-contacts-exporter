package tests

import (
	"bytes"
	"os"
	"testing"
    "io"

	"github.com/hmaier-dev/contacts_converter/pkg/loading"
)


func TestLoadVCF(t *testing.T){
    var path = "../testdata/contacts.vcf"
    loading.LoadVFC(path)
}

func TestLoadSqlite(t *testing.T){
    var path = "../testdata/contacts.db"
    loading.LoadSqlite(path)
    
}

func TestWriteVCF(t *testing.T){
    var path = "../testdata/contacts.vcf"
    cards := loading.LoadVFC(path)
    var writer io.Writer = &bytes.Buffer{}
    loading.WriteVCF(cards, writer)
    buf := writer.(*bytes.Buffer)
    data := buf.Bytes()
    os.WriteFile("../testdata/contacts2.vcf", data, 0744)

}
