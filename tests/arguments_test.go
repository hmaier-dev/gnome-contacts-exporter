package tests

import (
	"flag"
	"testing"

	"github.com/hmaier-dev/gnome-contacts-exporter/pkg/arguments"
)

func TestArgumentsWithoutArguments(t *testing.T) {

	test := struct {
		name string
		args []string
		want []string
	}{
		name: "Missing source and destination",
		args: []string{""},
		want: []string{"Source hasn't been set.", "Destination hasn't been set."},
	}

	flag.CommandLine = flag.NewFlagSet(test.name, flag.ContinueOnError)

	errs := arguments.Define(test.args)
	// Validate errors
	if len(errs) != len(test.want) {
		t.Fatalf("expected %d errors, got %d", len(test.want), len(errs))
	}
	for i, err := range errs {
		if err.Error() != test.want[i] {
			t.Errorf("expected error %q, got %q", test.want[i], err.Error())
		}
	}
}
func TestWithoutSource(t *testing.T) {

	test := struct {
		name string
		args []string
		want []string
	}{
		name: "Missing source",
		args: []string{"--destination", "contacts.vcf"},
		want: []string{"Source hasn't been set."},
	}

	flag.CommandLine = flag.NewFlagSet(test.name, flag.ContinueOnError)

	errs := arguments.Define(test.args)
	// Validate errors
	if len(errs) != len(test.want) {
		t.Fatalf("expected %d errors, got %d", len(test.want), len(errs))
	}
	for i, err := range errs {
		if err.Error() != test.want[i] {
			t.Errorf("expected error %q, got %q", test.want[i], err.Error())
		}
	}
}


func TestSameFileExtension(t *testing.T) {

	test := struct {
		name string
		args []string
		want []string
	}{
		name: "Same filetype",
		args: []string{"--destination", "contacts.db", "--source", "contacts.db"},
		want: []string{"Cannot export into the same filetype."},
	}

	flag.CommandLine = flag.NewFlagSet(test.name, flag.ContinueOnError)

	errs := arguments.Define(test.args)
	// Validate errors
	if len(errs) != len(test.want) {
		t.Fatalf("expected %d errors, got %d", len(test.want), len(errs))
	}
	for i, err := range errs {
		if err.Error() != test.want[i] {
			t.Errorf("expected error %q, got %q", test.want[i], err.Error())
		}
	}
}
