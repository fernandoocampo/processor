package domain

import "errors"

const (
	// FieldSeparator field separator.
	FieldSeparator = ","
	// ExpectedFields the expected number of fields.
	ExpectedFields = 5
)

const (
	// DocumentFieldPos position of the document field in the file.
	DocumentFieldPos int = iota
	// LastNameFieldPos position of the last name field in the file.
	LastNameFieldPos
	// FirstNameFieldPos position of the first name field in the file.
	FirstNameFieldPos
	// BirthDateFieldPos position of the birth date field in the file.
	BirthDateFieldPos
	// DeparmentFieldPos position of the department field in the file.
	DeparmentFieldPos
)

// Employee contains data for an employee
type Employee struct {
	Document   string
	FirstName  string
	LastName   string
	BirthDate  string
	Department string
}

// NewEmployeeWithRecord builds and returns an employee from the given record.
// return an error if something is invalid in the given record.
func NewEmployeeWithRecord(record []string) (*Employee, error) {
	if len(record) != ExpectedFields {
		return nil, errors.New("record does not have the expected number of fields")
	}
	employee := Employee{
		Document:   record[DocumentFieldPos],
		BirthDate:  record[BirthDateFieldPos],
		Department: record[DeparmentFieldPos],
		FirstName:  record[FirstNameFieldPos],
		LastName:   record[LastNameFieldPos],
	}
	return &employee, nil
}
