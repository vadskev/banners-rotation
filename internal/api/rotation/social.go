package rotation

import (
	"context"
	"fmt"

	"github.com/vadskev/banners-rotation/internal/converter"
	"github.com/vadskev/banners-rotation/internal/logger"
	desc "github.com/vadskev/banners-rotation/pkg/rotation_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i Implementation) AddSocialGroup(ctx context.Context, request *desc.AddSocialGroupRequest) (*desc.AddSocialGroupResponse, error) {
	if len(request.Description) == 0 {
		return nil, status.Error(codes.Aborted, "description is required")
	}

	socialGroup, err := i.storageService.AddSocialGroup(ctx, converter.ToSocialGroupFromAddSocialGroupRequest(request))
	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	}

	logger.Info(fmt.Sprintf("inserted banner: %#v", socialGroup))

	return &desc.AddSocialGroupResponse{
		Id: socialGroup.ID,
	}, nil
}

func (i Implementation) DeleteSocialGroup(ctx context.Context, request *desc.DeleteSocialGroupRequest) (*desc.DeleteSocialGroupResponse, error) {
	if request.Id <= 0 {
		return nil, status.Error(codes.Aborted, "id is required")

	}

	st, err := i.storageService.DeleteSocialGroup(ctx, converter.ToSocialGroupFromSocialGroupRequest(request))
	if err != nil {
		return nil, status.Error(codes.NotFound, "social group not found")
	}

	logger.Info(fmt.Sprintf("deleted social group: %#v", st))

	return &desc.DeleteSocialGroupResponse{
		Status: "OK",
	}, nil
}
