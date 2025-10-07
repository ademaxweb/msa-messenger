package requests

import "social/internal/app/models"

func (r *Repository) UpdateRequestByUsers(request *models.FriendRequest) (*models.FriendRequest, error) {
	// WHERE (sender_id = request.SenderID AND recipient_id = request.RecipientID) OR (sender_id = request.RecipientID AND recipient_id = request.SenderID)
	return nil, nil
}
