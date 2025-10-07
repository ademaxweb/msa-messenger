package dto

import pb "social/pkg/api/social/v1"

type UpdateRequestStatus struct {
	RequestId uint32
	Status    pb.FriendRequestStatus
}
