package main

import "fmt"

// Element is the type for the ListView elements
type Element struct {
	ID   int
	Text string
}

// Elementer interface for producing the string
type Elementer interface {
	FormatterForDB() string
}

// FormatterForDB function to generate string that can be inserted into database
func (e *Element) FormatterForDB() string {
	return fmt.Sprintf("Text: %s", e.Text)
}

func getAllElements() ([]Element, error) {
	db := openDBConnection()
	selectStatement := `SELECT * FROM elements ORDER BY id DESC`
	rows, err := db.Query(selectStatement)
	if err != nil {
		fmt.Println(err.Error())
		return []Element{}, err
	}
	elements := []Element{}
	elementObj := Element{}

	for rows.Next() {
		var (
			id      int
			element string
		)
		err = rows.Scan(&id, &element)
		if err != nil {
			fmt.Println(err.Error())
		}
		elementObj.ID = id
		elementObj.Text = element
		elements = append(elements, elementObj)
	}
	defer db.Close()
	return elements, err
}

func inserElement(ele string) bool {
	db := openDBConnection()
	insertStatement := `INSERT INTO elements ( element ) VALUES ($1)`
	elementObj := Element{Text: ele}
	_, err := db.Exec(insertStatement, elementObj.FormatterForDB())
	if err != nil {
		return false
	}
	return true
}
