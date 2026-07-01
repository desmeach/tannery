package model

type Group struct {
	Id           string
	Title        string
	IsPrivate    bool
	MembersCount GroupMembersCount
	OwnerId      string
}

func NewGroup(id string, title string, isPrivate bool, membersCount GroupMembersCount, ownerId string) (Group, error) {
	return Group{
		Id:           id,
		Title:        title,
		IsPrivate:    isPrivate,
		MembersCount: membersCount,
		OwnerId:      ownerId,
	}, nil
}
