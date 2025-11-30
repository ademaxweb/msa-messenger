package grpc

import (
	"social/internal/app/models"
	"social/internal/app/usecases/dto"
	pb "social/pkg/api/social/v1"
)

func FriendRequestStatusToPB(s models.FriendRequestStatus) pb.FriendRequestStatus {
	switch s {
	case models.RequestStatusPending:
		return pb.FriendRequestStatus_STATUS_PENDING
	case models.RequestStatusAccepted:
		return pb.FriendRequestStatus_STATUS_ACCEPTED
	case models.RequestStatusDeclined:
		return pb.FriendRequestStatus_STATUS_DECLINED
	case models.RequestStatusRemoved:
		return pb.FriendRequestStatus_STATUS_REMOVED
	default:
		return pb.FriendRequestStatus_STATUS_UNKNOWN
	}
}

func FriendRequestStatusFromPB(s pb.FriendRequestStatus) models.FriendRequestStatus {
	switch s {
	case pb.FriendRequestStatus_STATUS_PENDING:
		return models.RequestStatusPending
	case pb.FriendRequestStatus_STATUS_ACCEPTED:
		return models.RequestStatusAccepted
	case pb.FriendRequestStatus_STATUS_DECLINED:
		return models.RequestStatusDeclined
	case pb.FriendRequestStatus_STATUS_REMOVED:
		return models.RequestStatusRemoved
	default:
		return models.RequestStatusUnknown
	}
}

func dtoUpdateRequestStatusFromAcceptFriendRequestRequest(request *pb.AcceptFriendRequestRequest) dto.UpdateRequestStatus {
	return dto.UpdateRequestStatus{
		RequestId: request.GetRequestId(),
		Status:    models.RequestStatusAccepted,
	}
}

func dtoUpdateRequestStatusFromDeclineFriendRequestRequest(request *pb.DeclineFriendRequestRequest) dto.UpdateRequestStatus {
	return dto.UpdateRequestStatus{
		RequestId: request.GetRequestId(),
		Status:    models.RequestStatusDeclined,
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

func pbFriendRequestFromFriendRequestModel(m *models.FriendRequest) *pb.FriendRequest {
	return &pb.FriendRequest{
		Id:          m.Id,
		SenderId:    m.SenderId,
		RecipientId: m.RecipientID,
		Status:      FriendRequestStatusToPB(m.Status),
	}
}
