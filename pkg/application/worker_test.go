package application_test

import (
	"context"
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

	ctx := context.TODO()
	got, err := application.Process(ctx, filepath)

	assert.NoError(t, err)

	assert.Equal(t, wanted, got)
}

func TestProcessFileWithOneInvalidRecord(t *testing.T) {
	filepath := buildFileTestPath(t, "oneinvalidrecord.csv")

	ctx := context.TODO()
	got, err := application.Process(ctx, filepath)

	assert.Nil(t, got)
	assert.Error(t, err)
}

func TestProcessNonexistingFile(t *testing.T) {
	filepath := buildFileTestPath(t, "nonexisting.csv")
	ctx := context.TODO()
	got, err := application.Process(ctx, filepath)
	assert.Nil(t, got)
	assert.Error(t, err)
}

func BenchmarkProcessBigFile(b *testing.B) {
	ctx := context.TODO()
	filepath := buildFilePath("bulk.csv")
	application.Process(ctx, filepath) //0.139s
}

// buildFileTestPath helps to build the file path of the given file name.
func buildFileTestPath(t *testing.T, name string) string {
	t.Helper()
	return buildFilePath(name)
}

func buildFilePath(name string) string {
	dir, _ := os.Getwd()
	return filepath.Join(dir+"/testdata", name)
}
