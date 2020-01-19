package application_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/fernandoocampo/processor/pkg/application"
	"github.com/fernandoocampo/processor/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func TestProcessFile(t *testing.T) {
	filepath := buildFileTestPath(t, "tworecords.csv")
	wanted := []*domain.Employee{
		&domain.Employee{"123123123A", "Leopoldo Enriquez", "Gomez Ruiz", "12/11/1976", "sales"},
		&domain.Employee{"158123123A", "Vladimir Lucilo", "Garcia Caicedo", "12/11/1966", "it"},
	}

	got, err := application.Process(filepath)

	assert.NoError(t, err)

	assert.Equal(t, wanted, got)
}

func TestProcessFileWithOneInvalidRecord(t *testing.T) {
	filepath := buildFileTestPath(t, "oneinvalidrecord.csv")

	got, err := application.Process(filepath)

	assert.Nil(t, got)
	assert.Error(t, err)
}

func TestProcessNonexistingFile(t *testing.T) {
	filepath := buildFileTestPath(t, "nonexisting.csv")
	got, err := application.Process(filepath)
	assert.Nil(t, got)
	assert.Error(t, err)
}

// buildFileTestPath helps to build the file path of the given file name.
func buildFileTestPath(t *testing.T, name string) string {
	t.Helper()
	dir, _ := os.Getwd()
	return filepath.Join(dir+"/testdata", name)
}
