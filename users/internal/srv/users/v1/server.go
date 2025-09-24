package v1

import (
	"context"
	pb "users/pkg/api/users/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedUsersServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s Server) CreateProfile(ctx context.Context, request *pb.CreateProfileRequest) (*pb.CreateProfileResponse, error) {
	//TODO implement me
	return nil, status.New(codes.Unimplemented, codes.Unimplemented.String()).Err()
}

func (s Server) UpdateProfile(ctx context.Context, request *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	//TODO implement me
	return nil, status.New(codes.Unimplemented, codes.Unimplemented.String()).Err()
}

func (s Server) GetProfileByID(ctx context.Context, request *pb.GetProfileByIDRequest) (*pb.GetProfileByIDResponse, error) {
	//TODO implement me
	return nil, status.New(codes.Unimplemented, codes.Unimplemented.String()).Err()
}

func (s Server) GetProfileByNickname(ctx context.Context, request *pb.GetProfileByNicknameRequest) (*pb.GetProfileByNicknameResponse, error) {
	//TODO implement me
	return nil, status.New(codes.Unimplemented, codes.Unimplemented.String()).Err()
}

func (s Server) SearchProfileByNickname(ctx context.Context, request *pb.SearchProfileByNicknameRequest) (*pb.SearchProfileByNicknameResponse, error) {
	//TODO implement me
	return nil, status.New(codes.Unimplemented, codes.Unimplemented.String()).Err()
}
