package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) String() string {
	formatted := ""
	for k, t := range *l {
		prefix := "  "
		if t.Done {
			prefix = "X "
		}
		// Adjust the item number k to print numbers starting from 1 instead of 0
		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
	}
	return formatted
}
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)
}

func (l *List) Complete(i int) error {
	list := *l
	if i <= 0 || i > len(list) {
		return fmt.Errorf("item  %d does not exist", i)
	}

	list[i-1].Done = true
	list[i-1].CompletedAt = time.Now()

	return nil

}

func (l *List) Delete(i int) error {

	list := *l
	if i <= 0 || i > len(list) {
		return fmt.Errorf("item  %d does not exist", i)
	}

	*l = append(list[:i-1], list[i:]...)
	return nil
}

func (l *List) Save(filename string) error {
	log.Printf("Starting to save list to file: %s", filename)

	// Log the type and value of `l`
	log.Printf("Type of l: %T, Value of l: %+v", l, l)

	// Marshal the list into JSON
	js, err := json.Marshal(l)
	if err != nil {
		log.Printf("Error marshaling list to JSON: %v", err)
		return err
	}
	log.Printf("Successfully marshaled list to JSON.")

	// Log the type and value of `js`
	log.Printf("Type of js: %T, Value of js: %s", js, js)

	log.Printf("Completed saving list to file: %s", filename)
	return ioutil.WriteFile(filename, js, 0644)
}

func (l *List) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}
