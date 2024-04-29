package nothification

import (
	"database/sql"

	"gitea.com/lzhuk/forum/internal/model"
)

type NotificationRepository struct {
	DB *sql.DB
}

func InitNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{
		DB: db,
	}
}

const (
	notificationCreateQuery  = `INSERT INTO notifications (user_id, post_id, type, created_user_id, message) VALUES($1, $2, $3, $4, $5)`
	notificationUpdateQuery  = `UPDATE notifications SET is_read = TRUE WHERE user_id = $1`
	notificationsQuery       = `SELECT * FROM notifications WHERE user_id = $1 AND is_read = FALSE`
	notificationsIsReadQuery = `SELECT EXISTS (SELECT 1 FROM notifications WHERE user_id = $1 AND is_read = FALSE) AS has_unread_notifications`
)

func (n *NotificationRepository) Create(n1 *model.Notification) error {
	if _, err := n.DB.Exec(notificationCreateQuery, n1.UserID, n1.PostID, n1.Type, n1.CreatedUserID, n1.Message); err != nil {
		return err
	}
	return nil
}

func (n *NotificationRepository) NotificationIsRead(user_id int) (bool, error) { // CHECK
	var isRead bool
	if err := n.DB.QueryRow(notificationsIsReadQuery, user_id).Scan(&isRead); err != nil {
		return false, err
	}
	return isRead, nil
}

func (n *NotificationRepository) Update(user_id int) error {
	if _, err := n.DB.Exec(notificationUpdateQuery, user_id); err != nil {
		return err
	}
	return nil
}

func (n *NotificationRepository) Read(user_id int) ([]model.Notification, error) {
	userNotifications := []model.Notification{}
	rows, err := n.DB.Query(notificationsQuery, user_id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		notification := &model.Notification{}
		if err := rows.Scan(&notification.ID, &notification.UserID, &notification.PostID, &notification.Type, &notification.CreatedUserID, &notification.Message, &notification.IsRead, &notification.CreatedAt); err != nil {
			return nil, err
		}
		userNotifications = append(userNotifications, *notification)
	}
	return userNotifications, nil
}
