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
	notificationUpdateQuery  = `UPDATE notifications SET is_read = TRUE`
	notificationsQuery       = `SELECT * FROM notifications WHERE user_id = $1 AND is_read = FALSE`
	notificationsIsReadQuery = `SELECT EXISTS (SELECT 1 FROM notifications WHERE user_id = $1 AND is_read = FALSE) AS check`
)

func (n *NotificationRepository) Create(n1 *model.Notification) error {
	if _, err := n.DB.Exec(notificationCreateQuery, n1.UserID, n1.PostID, n1.Type, n1.CreatedUserID, n1.Message); err != nil {
		return err
	}
	return nil
}

func (n *NotificationRepository) NotificationIsRead() (bool, error) { // CHECK
	var isRead bool
	if err := n.DB.QueryRow(notificationsIsReadQuery).Scan(&isRead); err != nil {
		return false, err
	}
	return isRead, nil
}

func (n *NotificationRepository) Read(user_id int) ([]model.Notification, error) {
	userNotifications := []model.Notification{}
	rows, err := n.DB.Query(notificationsQuery, user_id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		nothification := &model.Notification{}
		if err := rows.Scan(&nothification.ID, &nothification.UserID, &nothification.PostID, &nothification.Type, &nothification.CreatedUserID, &nothification.Message, &nothification.IsRead, &nothification.CreatedAt); err != nil {
			return nil, err
		}
		userNotifications = append(userNotifications, *nothification)
	}
	return userNotifications, nil
}

func (n *NotificationRepository) Update() error {
	if _, err := n.DB.Exec(notificationUpdateQuery); err != nil {
		return err
	}
	return nil
}
