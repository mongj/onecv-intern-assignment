package enums

type Sex string
type EmploymentStatus string
type MaritalStatus string
type Relation string
type SchoolLevel string
type ApplicationStatus string

const (
	SexMale   Sex = "male"
	SexFemale Sex = "female"
)

const (
	EmploymentStatusEmployed   EmploymentStatus = "employed"
	EmploymentStatusUnemployed EmploymentStatus = "unemployed"
)

const (
	MaritalStatusSingle   MaritalStatus = "single"
	MaritalStatusMarried  MaritalStatus = "married"
	MaritalStatusWidowed  MaritalStatus = "widowed"
	MaritalStatusDivorced MaritalStatus = "divorced"
)

const (
	RelationParent  Relation = "parent"
	RelationChild   Relation = "child"
	RelationSibling Relation = "sibling"
	RelationSpouse  Relation = "spouse"
	RelationOther   Relation = "other"
)

const (
	SchoolLevelPreschool     SchoolLevel = "preschool"
	SchoolLevelPrimary       SchoolLevel = "primary"
	SchoolLevelSecondary     SchoolLevel = "secondary"
	SchoolLevelPostsecondary SchoolLevel = "post-secondary"
)

const (
	ApplicationStatusPending  ApplicationStatus = "pending"
	ApplicationStatusApproved ApplicationStatus = "approved"
	ApplicationStatusRejected ApplicationStatus = "rejected"
)
