package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	// When no one is giving for a variable it is anonymous
	// It is now embedded into the upper struct
	Date
}
func (e *Event) Title() string {
	return e.title
}
func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("Invalid Title - Too long")
	}

	e.title = title
	return nil
}