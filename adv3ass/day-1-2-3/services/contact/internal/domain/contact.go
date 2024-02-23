package domain

import (
	"fmt"
)

type Contact struct {
	ID          string
	firstName   string
	lastName    string
	patronymic  string
	PhoneNumber string
}

func (c *Contact) FullName() string {
	return fmt.Sprintf("%s %s %s", c.lastName, c.firstName, c.patronymic)
}

func (c *Contact) SetFullName(firstName, lastName, patronymic string) {
	c.firstName = firstName
	c.lastName = lastName
	c.patronymic = patronymic
}

type Group struct {
	ID   string
	Name string
}
