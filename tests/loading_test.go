package tests

import (
	"testing"
	"github.com/hmaier-dev/contacts_converter/pkg/loading"
)

func TestLoadSqlite(t *testing.T){
    var path = "testdata/contacts.db"
    loading.LoadSqlite(path)
    
}

func TestLoadVCF(t *testing.T){
    var path = "testdata/contacts.vcf"
    loading.LoadVFC(path)
}
