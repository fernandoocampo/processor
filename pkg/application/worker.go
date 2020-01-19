package application

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fernandoocampo/processor/pkg/domain"
)

// Process processes the file in the given file path and return
// an slice of employees. If something goes wrong while iterating
// the file an error will be returned.
func Process(filepath string) ([]*domain.Employee, error) {
	var result []*domain.Employee
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	linenumber := 1
	for scanner.Scan() {
		recordline := scanner.Text()
		employee, err := processLine(linenumber, recordline)
		if err != nil {
			return nil, fmt.Errorf("line %d has errors: %s", linenumber, err.Error())
		}
		result = append(result, employee)
		linenumber++
	}
	return result, nil
}

func processLine(lineNumber int, recordLine string) (*domain.Employee, error) {
	record := strings.Split(recordLine, domain.FieldSeparator)
	return domain.NewEmployeeWithRecord(record)
}
