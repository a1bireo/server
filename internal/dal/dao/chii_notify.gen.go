// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"github.com/bangumi/server/internal/model"
)

const TableNameNotification = "chii_notify"

// Notification mapped from table <chii_notify>
type Notification struct {
	ID          model.NotificationID `gorm:"column:nt_id;type:mediumint(8) unsigned;primaryKey;autoIncrement:true"`
	ReceiverID  model.UserID         `gorm:"column:nt_uid;type:mediumint(8) unsigned;not null"`
	SenderID    model.UserID         `gorm:"column:nt_from_uid;type:mediumint(8) unsigned;not null"`
	Status      uint8                `gorm:"column:nt_status;type:tinyint(1) unsigned;not null;default:1"`
	Type        uint8                `gorm:"column:nt_type;type:tinyint(3) unsigned;not null"`
	FieldID     uint32               `gorm:"column:nt_mid;type:mediumint(8) unsigned;not null"` // ID in notify_field
	RelatedID   uint32               `gorm:"column:nt_related_id;type:int(10) unsigned;not null"`
	CreatedTime uint32               `gorm:"column:nt_dateline;type:int(10) unsigned;not null"`
}

// TableName Notification's table name
func (*Notification) TableName() string {
	return TableNameNotification
}
