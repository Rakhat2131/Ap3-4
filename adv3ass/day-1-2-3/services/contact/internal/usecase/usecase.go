package usecase

import (
	"fmt"
)

type Contact struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

type Group struct {
	ID       string
	Name     string
	Members  []*Contact
	Contacts []string
}

type ContactUseCase struct {
	contacts map[string]*Contact
}

func NewContactUseCase() *ContactUseCase {
	return &ContactUseCase{
		contacts: make(map[string]*Contact),
	}
}

func (uc *ContactUseCase) CreateContact(contact *Contact) error {
	// Ваша логика для создания контакта
	uc.contacts[contact.ID] = contact
	return nil
}

func (uc *ContactUseCase) ReadContact(contactID string) (*Contact, error) {
	contact, ok := uc.contacts[contactID]
	if !ok {
		return nil, fmt.Errorf("contact with ID %s not found", contactID)
	}
	return contact, nil
}

func (uc *ContactUseCase) UpdateContact(contact *Contact) error {

	uc.contacts[contact.ID] = contact
	return nil
}

func (uc *ContactUseCase) DeleteContact(contactID string) error {

	delete(uc.contacts, contactID)
	return nil
}

type GroupUseCase struct {
	groups   map[string]*Group
	Contacts []string
	contacts interface{}
}

func NewGroupUseCase() *GroupUseCase {
	return &GroupUseCase{
		groups: make(map[string]*Group),
	}
}

func (uc *GroupUseCase) CreateGroup(group *Group) error {

	uc.groups[group.ID] = group
	return nil
}

func (uc *GroupUseCase) ReadGroup(groupID string) (*Group, error) {
	group, ok := uc.groups[groupID]
	if !ok {
		return nil, fmt.Errorf("group with ID %s not found", groupID)
	}
	return group, nil
}

func (uc *GroupUseCase) AddContactToGroup(groupID, contactID string) error {
	group, ok := uc.groups[groupID]
	if !ok {
		return fmt.Errorf("group with ID %s not found", groupID)
	}

	group.Contacts = append(group.Contacts, contactID)
	return nil
}
