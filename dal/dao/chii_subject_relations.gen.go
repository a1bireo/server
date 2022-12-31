// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"github.com/bangumi/server/internal/model"
)

const TableNameSubjectRelation = "chii_subject_relations"

// SubjectRelation mapped from table <chii_subject_relations>
type SubjectRelation struct {
	SubjectID            model.SubjectID `gorm:"column:rlt_subject_id;type:mediumint(8) unsigned;primaryKey"` // 关联主 ID
	SubjectTypeID        uint8           `gorm:"column:rlt_subject_type_id;type:tinyint(3) unsigned;not null"`
	RelationType         uint16          `gorm:"column:rlt_relation_type;type:smallint(5) unsigned;not null"`          // 关联类型
	RelatedSubjectID     model.SubjectID `gorm:"column:rlt_related_subject_id;type:mediumint(8) unsigned;primaryKey"`  // 关联目标 ID
	RelatedSubjectTypeID uint8           `gorm:"column:rlt_related_subject_type_id;type:tinyint(3) unsigned;not null"` // 关联目标类型
	ViceVersa            bool            `gorm:"column:rlt_vice_versa;type:tinyint(1) unsigned;primaryKey"`
	Order                uint8           `gorm:"column:rlt_order;type:tinyint(3) unsigned;not null"` // 关联排序
	Subject              Subject         `gorm:"foreignKey:rlt_related_subject_id;references:subject_id" json:"subject"`
}

// TableName SubjectRelation's table name
func (*SubjectRelation) TableName() string {
	return TableNameSubjectRelation
}