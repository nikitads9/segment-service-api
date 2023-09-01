package convert

import (
	"github.com/nikitads9/segment-service-api/internal/model"
	desc "github.com/nikitads9/segment-service-api/pkg/user_api"
)

func ToModifySegmentInfo(req *desc.ModifySegmentsRequest) *model.ModifySegmentInfo {
	if req == nil {
		return nil
	}

	return &model.ModifySegmentInfo{
		UserId:        req.GetId(),
		SlugsToAdd:    req.GetSlugToAdd(),
		SlugsToRemove: req.GetSlugToRemove(),
	}
}

func ToSetExpireTimeInfo(req *desc.SetExpireTimeRequest) *model.SetExpireTimeInfo {
	if req == nil {
		return nil
	}

	return &model.SetExpireTimeInfo{
		UserId:     req.GetId(),
		Slug:       req.GetSlug(),
		ExpireTime: req.GetExpirationTime().AsTime(),
	}
}
