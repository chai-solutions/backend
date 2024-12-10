package repos

import (
	"context"

	"chai/database/sqlc"
)

type NotificationsRepo interface {
	CreateNotifications(userIDs []int32, title string, message string) error
}

type notificationsRepoImpl struct {
	db *sqlc.Queries
}

func NewNotificationsRepo(db *sqlc.Queries) NotificationsRepo {
	return &notificationsRepoImpl{db: db}
}

func (r *notificationsRepoImpl) CreateNotifications(userIDs []int32, title string, message string) error {
	rows := make([]sqlc.CreateNotificationsParams, len(userIDs))

	for i, v := range userIDs {
		rows[i] = sqlc.CreateNotificationsParams{
			User:    v,
			Title:   title,
			Message: message,
		}
	}

	_, err := r.db.CreateNotifications(context.Background(), rows)
	return err
}
