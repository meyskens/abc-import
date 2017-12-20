package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	valid "gopkg.in/asaskevich/govalidator.v4"
)

// Person contains the info of one person
type Person struct {
	Name           string
	Address        string
	PostalCode     string
	City           string
	ComminucateVia string
	Emails         []string
}

func main() {
	emails := collectEmails(getPeople())
	fmt.Println(strings.Join(emails, "\n"))
}

func getPeople() []Person {
	csvFile, _ := os.Open("out.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var people []Person

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		emails := []string{}
		for i := 4; i < len(line); i++ {
			if valid.IsEmail(line[i]) {
				emails = append(emails, line[i])
			}
		}
		people = append(people, Person{
			Name:           line[0],
			Address:        line[1],
			PostalCode:     line[2],
			City:           line[3],
			ComminucateVia: line[4],
			Emails:         emails,
		})
	}

	return people
}

func collectEmails(people []Person) []string {
	out := []string{}
	for _, person := range people {
		if person.ComminucateVia == "" || person.ComminucateVia == "e-mail" {
			out = append(out, person.Emails...)
		}
	}

	return out
}
