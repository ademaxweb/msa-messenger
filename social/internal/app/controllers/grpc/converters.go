package grpc

import (
	"social/internal/app/models"
	"social/internal/app/usecases/dto"
	pb "social/pkg/api/social/v1"
)

func dtoUpdateRequestStatusFromAcceptFriendRequestRequest(request *pb.AcceptFriendRequestRequest) dto.UpdateRequestStatus {
	return dto.UpdateRequestStatus{
		RequestId: request.GetRequestId(),
		Status:    pb.FriendRequestStatus_STATUS_ACCEPTED,
	}
}

func dtoUpdateRequestStatusFromDeclineFriendRequestRequest(request *pb.DeclineFriendRequestRequest) dto.UpdateRequestStatus {
	return dto.UpdateRequestStatus{
		RequestId: request.GetRequestId(),
		Status:    pb.FriendRequestStatus_STATUS_DECLINED,
	}
}

func dtoListFriendsFromListFriendsRequest(request *pb.ListFriendsRequest) dto.ListFriends {
	return dto.ListFriends{
		UserID: request.GetUserId(),
	}
}

func dtoListRequestsFromListRequestsRequest(request *pb.ListRequestsRequest) dto.ListRequests {
	return dto.ListRequests{
		UserID: request.GetUserId(),
	}
}

func pbFriendRequestFromFriendRequestModel(request *models.FriendRequest) *pb.FriendRequest {
	return &pb.FriendRequest{
		Id:          request.Id,
		SenderId:    request.SenderId,
		RecipientId: request.RecipientID,
		Status:      request.Status,
	}
}
