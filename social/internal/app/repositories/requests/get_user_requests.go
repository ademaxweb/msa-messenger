package requests

import "social/internal/app/models"

func (r *Repository) GetUserRequests(filters *models.FriendRequest) ([]models.FriendRequest, error) {
	// WHERE (recipient_id = filters.RecipientID OR sender_id = filters.SenderID) AND status_id = filters.Status
	return nil, nil
}
