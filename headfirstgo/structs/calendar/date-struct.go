package calendar

import "errors"

// Methods with Structs
type Date struct {
	// Make lowercase so they dont get exported
	// Similar to private in other languages
	year int
	month int
	day int
}

// Getter Functions
func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}

// Setter Functions
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("Invalid year")
	}

	d.year = year
	// Instead of *p.Age++
	// This is because Go automatically handles dereferencing for you when you use a pointer receiver.
	return nil
}

func (d *Date) SetMonth(month int) error {

	if month < 1 || month > 12 {
		return errors.New("Invalid month")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(month int) error {

	if month < 1 || month > 31 {
		return errors.New("Invalid day")
	}
	d.day = month
	return nil
}