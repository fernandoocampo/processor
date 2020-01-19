package domain_test

import (
	"errors"
	"testing"

	"github.com/fernandoocampo/processor/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewEmployeeWithRecord(t *testing.T) {
	cases := []struct {
		name           string
		wantedError    error
		wantedEmployee *domain.Employee
		given          []string
	}{
		{
			name:           "good_record",
			wantedError:    nil,
			wantedEmployee: &domain.Employee{"123123123A", "Leopoldo Enriquez", "Gomez Ruiz", "12/11/1976", "sales"},
			given:          []string{"123123123A", "Gomez Ruiz", "Leopoldo Enriquez", "12/11/1976", "sales"},
		},
		{
			name:           "incomplete_record",
			wantedError:    errors.New("record does not have the expected number of fields"),
			wantedEmployee: nil,
			given:          []string{"123123123A", "Gomez Ruiz", "Leopoldo Enriquez", "12/11/1976"},
		},
		{
			name:           "more_fields_record",
			wantedError:    errors.New("record does not have the expected number of fields"),
			wantedEmployee: nil,
			given:          []string{"123123123A", "Gomez Ruiz", "Leopoldo Enriquez", "12/11/1976", "sales", "Berlin", "Rome"},
		},
		{
			name:           "empty_record",
			wantedError:    errors.New("record does not have the expected number of fields"),
			wantedEmployee: nil,
			given:          []string{},
		},
		{
			name:           "nil_record",
			wantedError:    errors.New("record does not have the expected number of fields"),
			wantedEmployee: nil,
			given:          []string{},
		},
	}

	for _, acase := range cases {
		got, err := domain.NewEmployeeWithRecord(acase.given)
		assert.Equal(t, acase.wantedError, err, acase.name)
		assert.Equal(t, acase.wantedEmployee, got, acase.name)
	}
}
