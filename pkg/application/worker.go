package application

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/fernandoocampo/processor/pkg/domain"
)

// input contains record data
type input struct {
	line    int
	record  string
	anerror error
}

// result contains the outcome after process an input.
type result struct {
	employee *domain.Employee
	anerror  error
}

// Process processes the file in the given file path and return
// an slice of employees. If something goes wrong while iterating
// the file an error will be returned.
func Process(ctx context.Context, filepath string) ([]*domain.Employee, error) {
	var finalResult []*domain.Employee

	done := make(chan interface{})
	defer close(done)

	recordsStream := getRecords(done, filepath)

	numProcessors := runtime.NumCPU()
	fmt.Printf("Spinning up %d record processors.\n", numProcessors)
	processors := make([]<-chan *result, numProcessors)
	for i := 0; i < numProcessors; i++ {
		processors[i] = processRecord(done, recordsStream)
	}

	for v := range fanIn(done, processors...) {
		if v.anerror != nil {
			log.Printf("error: %s", v.anerror.Error())
			return nil, v.anerror
		}
		finalResult = append(finalResult, v.employee)
	}
	return finalResult, nil
}

func fanIn(done <-chan interface{}, channels ...<-chan *result) <-chan *result {
	var wg sync.WaitGroup
	multiplexedStream := make(chan *result)

	multiplex := func(c <-chan *result) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}

	// Select from all the channels
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	// Wait for all the reads to complete
	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}

// getRecords create stream to get the records from file
func getRecords(done <-chan interface{}, filepath string) <-chan *input {
	inputStream := make(chan *input)
	go func() {
		defer close(inputStream)
		file, err := os.Open(filepath)
		if err != nil {
			inputStream <- &input{
				anerror: err,
			}
			return
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for {
			select {
			case <-done:
				return
			default:
				linenumber := 1
				if !scanner.Scan() {
					return
				}
				inputStream <- &input{
					line:   linenumber,
					record: scanner.Text(),
				}
				linenumber++
			}
		}
	}()
	return inputStream
}

// processRecord process input stream
func processRecord(done <-chan interface{}, inputStream <-chan *input) <-chan *result {
	resultStream := make(chan *result)
	go func() {
		defer close(resultStream)
		for {
			select {
			case <-done:
				return
			case v := <-inputStream:
				if v == nil {
					return
				}
				employee, err := processLine(v.line, v.record)
				if err != nil {
					resultStream <- &result{
						anerror: fmt.Errorf("line %d has errors: %s", v.line, err.Error()),
					}
					return
				}
				resultStream <- &result{
					employee: employee,
				}
			}
		}
	}()
	return resultStream
}

// processLine process a string line to convert it to a employee.
func processLine(lineNumber int, recordLine string) (*domain.Employee, error) {
	record := strings.Split(recordLine, domain.FieldSeparator)
	return domain.NewEmployeeWithRecord(record)
}
