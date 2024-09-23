package schemecriteria

type Key int16

// Enum values for scheme criteria keys
const (
	EmploymentStatus Key = iota
	MaritalStatus
	HasChildren
	ChildrenSchoolLevel
)

// Can be generated using go generate
var KeyToString = map[Key]string{
	EmploymentStatus:    "employment_status",
	MaritalStatus:       "marital_status",
	HasChildren:         "has_children",
	ChildrenSchoolLevel: "children_school_level",
}

// Can be generated using go generate
var StringToKey = map[string]Key{
	"employment_status":     EmploymentStatus,
	"marital_status":        MaritalStatus,
	"has_children":          HasChildren,
	"children_school_level": ChildrenSchoolLevel,
}
