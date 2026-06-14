package model

import "tanneryApp/internal/errors"

const (
	minMemberRestriction = 1
	maxMemberRestriction = 10
)

type GroupMembersCount int

func NewGroupMembersCount(count int) (GroupMembersCount, error) {
	if count > maxMemberRestriction || count < minMemberRestriction {
		return 0, errors.ErrGroupMemberCountInvalid
	}

	return GroupMembersCount(count), nil
}

func (gmc GroupMembersCount) Value() int {
	return int(gmc)
}
