package entity

type GroupRecord struct {
	Id          string
	Title       string
	IsPrivate   bool
	MemberCount int
	OwnerId     string
}
