// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "hcm": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/minhkhiemm/hcm/design
// --out=$(GOPATH)/src/github.com/minhkhiemm/hcm
// --version=v1.3.1

package app

import (
	"github.com/goadesign/goa"
	uuid "github.com/satori/go.uuid"
	"time"
)

// patient model (default view)
//
// Identifier: application/vnd.patient+json; view=default
type Patient struct {
	// date of birth
	DateOfBirth *time.Time `form:"date_of_birth,omitempty" json:"date_of_birth,omitempty" yaml:"date_of_birth,omitempty" xml:"date_of_birth,omitempty"`
	// id for patient is unique uuid
	ID *uuid.UUID `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	// name of patient
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
}

// Validate validates the Patient media type instance.
func (mt *Patient) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// PatientCollection is the media type for an array of Patient (default view)
//
// Identifier: application/vnd.patient+json; type=collection; view=default
type PatientCollection []*Patient

// Validate validates the PatientCollection media type instance.
func (mt PatientCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}