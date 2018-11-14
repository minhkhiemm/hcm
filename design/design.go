package design

import (
	. "github.com/goadesign/goa/design"
	. "gopkg.in/goadesign/goa.v1/design/apidsl"
)

var _ = API("hcm", func() {
	Title("The health care management API")
	Description("web hook for collect data for health care management project")
	Host("localhost:8080")
	Scheme("http")
	BasePath("/v1")
})

var Patient = MediaType("application/vnd.patient+json", func() {
	Description("patient model")
	Attributes(func() {
		Attribute("id",UUID,"id for patient is unique uuid")
		Attribute("name",String,"name of patient")
		Attribute("date_of_birth",DateTime,"date of birth")
		Required("name")
	})
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("date_of_birth")
	})
})

var _ = Resource("partients", func() {
	BasePath("/patients")
	Description("patients resources")
	Action("create", func() {
		Routing(POST("/"))
		Payload(Patient)
		Response(Created)
	})
	Action("list", func() {
		Routing(GET("/"))
		Response(OK,CollectionOf(Patient))
	})
})