package tests

import (
	"bytes"
	"crypto/sha256"
	"io"
	"log"
	"os"
	"testing"
	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/vcf"
)



// Files won't be identical, 
// because go-vcard is structuring the VCARDs
// different from gnome-contacts
func TestExport(t *testing.T){
    var path1 = "../testdata/contacts.vcf"
    var path2 = "../testdata/contacts_encoded.vcf"
    cards := vcf.Import(path1)
    var writer io.Writer = &bytes.Buffer{}
    vcf.Export(cards, writer)
    buf := writer.(*bytes.Buffer)
    data := buf.Bytes()
    os.WriteFile(path2, data, 0644)

	fileOrig, err := os.Open(path1)
	if err != nil {
        log.Fatalf("Could open file: %s \n", path1)
	}
	defer fileOrig.Close()

	fileEnc, err := os.Open(path1)
	if err != nil {
        log.Fatalf("Could open file: %s \n", path2)
	}
	defer fileEnc.Close()

	hashOrig := sha256.New()
	if _, err := io.Copy(hashOrig, fileOrig); err != nil {
	}
	hashEnc := sha256.New()
	if _, err := io.Copy(hashEnc, fileEnc); err != nil {
	}

    if hashOrig != hashEnc{
        log.Fatalf("Files are not identical.\n HashOrig: %s \n HashEnc %s \n ", hashOrig, hashEnc)
    }


}

