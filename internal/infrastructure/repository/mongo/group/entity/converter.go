package entity

import (
	model2 "tanneryApp/internal/domain/model"
)

func ToGroupRecord(group model2.Group) GroupRecord {
	return GroupRecord{
		Id:          group.Id,
		Title:       group.Title,
		IsPrivate:   group.IsPrivate,
		MemberCount: group.MembersCount.Value(),
		OwnerId:     group.OwnerId,
	}
}

func toGroup(record GroupRecord) (model2.Group, error) {
	memberCount, err := model2.NewGroupMembersCount(record.MemberCount)
	if err != nil {
		return model2.Group{}, err
	}

	return model2.Group{
		Id:           record.Id,
		Title:        record.Title,
		IsPrivate:    record.IsPrivate,
		MembersCount: memberCount,
		OwnerId:      record.OwnerId,
	}, nil
}
