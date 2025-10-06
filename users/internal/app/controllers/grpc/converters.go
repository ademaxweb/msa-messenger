package grpc

import (
	"users/internal/app/models"
	"users/internal/app/usecases/dto"
	pb "users/pkg/api/users/v1"
)

func dtoCreateProfileFromCreateProfileRequest(r *pb.CreateProfileRequest) dto.CreateProfile {
	return dto.CreateProfile{
		Name: r.GetNickname(),
		Bio:  r.GetBio(),
	}
}

func dtoGetProfileByIDFromGetProfileByIDRequest(r *pb.GetProfileByIDRequest) dto.GetProfileByID {
	return dto.GetProfileByID{
		Id: r.GetUserId(),
	}
}

func dtoGetProfileByNicknameFromGetProfileByNicknameRequest(r *pb.GetProfileByNicknameRequest) dto.GetProfileByNickname {
	return dto.GetProfileByNickname{
		Name: r.GetNickname(),
	}
}

func dtoSearchProfileByNicknameFromSearchProfileRequest(r *pb.SearchProfileByNicknameRequest) dto.SearchProfileByNickname {
	return dto.SearchProfileByNickname{
		Query: r.GetQuery(),
	}
}

func dtoUpdateProfileFromUpdateProfileRequest(r *pb.UpdateProfileRequest) dto.UpdateProfile {
	return dto.UpdateProfile{
		Id:        r.GetUserId(),
		Name:      r.GetNickname(),
		Bio:       r.GetBio(),
		AvatarUrl: r.GetAvatarUrl(),
	}
}

func modelUserToUserProfile(m *models.User) *pb.UserProfile {
	return &pb.UserProfile{
		UserId:    m.Id,
		Nickname:  m.Name,
		Bio:       m.Bio,
		AvatarUrl: m.AvatarUrl,
	}
}
