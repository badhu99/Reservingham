// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import mssql "github.com/microsoft/go-mssqldb"

const TableNameRole = "Role"

// Role mapped from table <Role>
type Role struct {
	ID          mssql.UniqueIdentifier `gorm:"column:Id;primaryKey;default:newid()" json:"Id"`
	Name        string                 `gorm:"column:Name;not null" json:"Name"`
	Level       int64                  `gorm:"column:Level;not null" json:"Level"`
	Permissions []Permission           `gorm:"foreignKey:RoleID" json:"permissions"`
}

// TableName Role's table name
func (*Role) TableName() string {
	return TableNameRole
}
