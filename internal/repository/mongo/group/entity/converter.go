package entity

import "tanneryApp/internal/model"

func ToGroupRecord(group model.Group) GroupRecord {
	return GroupRecord{
		Id:          group.Id,
		Title:       group.Title,
		IsPrivate:   group.IsPrivate,
		MemberCount: group.MembersCount.Value(),
		OwnerId:     group.OwnerId,
	}
}

func toGroup(record GroupRecord) (model.Group, error) {
	memberCount, err := model.NewGroupMembersCount(record.MemberCount)
	if err != nil {
		return model.Group{}, err
	}

	return model.Group{
		Id:           record.Id,
		Title:        record.Title,
		IsPrivate:    record.IsPrivate,
		MembersCount: memberCount,
		OwnerId:      record.OwnerId,
	}, nil
}
