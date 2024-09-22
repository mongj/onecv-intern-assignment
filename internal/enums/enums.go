package enums

type Sex string
type EmploymentStatus string
type MaritalStatus string
type Relation string

const (
	MALE   Sex = "male"
	FEMALE Sex = "female"
)

const (
	EMPLOYED   EmploymentStatus = "employed"
	UNEMPLOYED EmploymentStatus = "unemployed"
)

const (
	SINGLE   MaritalStatus = "single"
	MARRIED  MaritalStatus = "married"
	WIDOWED  MaritalStatus = "widowed"
	DIVORCED MaritalStatus = "divorced"
)

const (
	PARENT  Relation = "parent"
	CHILD   Relation = "child"
	SIBLING Relation = "sibling"
	SPOUSE  Relation = "spouse"
	OTHER   Relation = "other"
)
