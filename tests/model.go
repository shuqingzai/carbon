package tests

import (
	"github.com/dromara/carbon/v2"
	"gorm.io/gorm"
)

type ISO8601Type string

func (t ISO8601Type) Format() string {
	return carbon.ISO8601Format
}

type MySQLModel1 struct {
	ID uint64 `json:"-" gorm:"column:id;primaryKey"`

	DateTimeCarbon1 carbon.Carbon `gorm:"type:varchar(50);" json:"date_time_carbon1"`
	DateTimeCarbon2 carbon.Carbon `gorm:"type:datetime;" json:"date_time_carbon2"`
	DateTimeCarbon3 carbon.Carbon `gorm:"type:timestamp;" json:"date_time_carbon3"`

	DateTimeLayout1 carbon.DateTime `gorm:"type:varchar(50);" json:"date_time_layout1"`
	DateTimeLayout2 carbon.DateTime `gorm:"type:datetime;" json:"date_time_layout2"`
	DateTimeLayout3 carbon.DateTime `gorm:"type:timestamp;" json:"date_time_layout3"`

	DateTimeFormat1 carbon.FormatType[ISO8601Type] `gorm:"type:varchar(50);" json:"date_time_format1"`
	DateTimeFormat2 carbon.FormatType[ISO8601Type] `gorm:"type:datetime;" json:"date_time_format2"`
	DateTimeFormat3 carbon.FormatType[ISO8601Type] `gorm:"type:timestamp;" json:"date_time_format3"`

	Timestamp1 carbon.Timestamp `gorm:"column:timestamp1;type:bigint;" json:"timestamp1"`

	CreatedAt carbon.DateTime `gorm:"column:created_at;type:timestamp;" json:"-"`
	UpdatedAt carbon.DateTime `gorm:"column:updated_at;type:timestamp;" json:"-"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at;type:datetime;" json:"-"`
}

func (MySQLModel1) TableName() string {
	return "carbon_mysql1"
}

type MySQLModel2 struct {
	ID uint64 `json:"-" gorm:"column:id;primaryKey"`

	DateTimeCarbon1 *carbon.Carbon `gorm:"type:varchar(50);" json:"date_time_carbon1"`
	DateTimeCarbon2 *carbon.Carbon `gorm:"type:datetime;" json:"date_time_carbon2"`
	DateTimeCarbon3 *carbon.Carbon `gorm:"type:timestamp;" json:"date_time_carbon3"`

	DateTimeLayout1 *carbon.DateTime `gorm:"type:varchar(50);" json:"date_time_layout1"`
	DateTimeLayout2 *carbon.DateTime `gorm:"type:datetime;" json:"date_time_layout2"`
	DateTimeLayout3 *carbon.DateTime `gorm:"type:timestamp;" json:"date_time_layout3"`

	DateTimeFormat1 *carbon.FormatType[ISO8601Type] `gorm:"type:varchar(50);" json:"date_time_format1"`
	DateTimeFormat2 *carbon.FormatType[ISO8601Type] `gorm:"type:datetime;" json:"date_time_format2"`
	DateTimeFormat3 *carbon.FormatType[ISO8601Type] `gorm:"type:timestamp;" json:"date_time_format3"`

	Timestamp1 *carbon.Timestamp `gorm:"column:timestamp1;type:bigint;" json:"timestamp1"`

	CreatedAt *carbon.DateTime `gorm:"column:created_at;type:timestamp;" json:"-"`
	UpdatedAt *carbon.DateTime `gorm:"column:updated_at;type:timestamp;" json:"-"`
	DeletedAt *gorm.DeletedAt  `gorm:"column:deleted_at;type:datetime;" json:"-"`
}

func (MySQLModel2) TableName() string {
	return "carbon_mysql2"
}

type PgSQLModel1 struct {
	ID uint64 `json:"-" gorm:"column:id;primaryKey"`

	Carbon1 carbon.Carbon `gorm:"type:timestamp;" json:"carbon1"`
	Carbon2 carbon.Carbon `gorm:"type:timestamptz;" json:"carbon2"`

	DateLayout1 carbon.Date `gorm:"type:date;" json:"date_carbon2"`

	TimeLayout1 carbon.Time `gorm:"type:time;" json:"time_layout1"`
	TimeLayout2 carbon.Time `gorm:"type:timetz;" json:"time_layout2"`

	DateTimeLayout1 carbon.DateTime `gorm:"type:timestamp;" json:"date_time_layout1"`
	DateTimeLayout2 carbon.DateTime `gorm:"type:timestamptz;" json:"date_time_layout2"`

	CustomerFormat1 carbon.FormatType[ISO8601Type] `gorm:"type:timestamp;" json:"customer_format1"`
	CustomerFormat2 carbon.FormatType[ISO8601Type] `gorm:"type:timestamptz;" json:"customer_format2"`

	Timestamp1 carbon.Timestamp `gorm:"column:timestamp1;type:int4;" json:"timestamp1"`

	CreatedAt carbon.DateTime `gorm:"autoCreateTime;column:created_at;type:timestamptz;" json:"-"`
	UpdatedAt carbon.DateTime `gorm:"autoUpdateTime;column:updated_at;type:timestamptz;" json:"-"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at;type:timestamptz;" json:"-"`
}

func (PgSQLModel1) TableName() string {
	return "carbon_pgsql1"
}

type PgSQLModel2 struct {
	ID uint64 `json:"-" gorm:"column:id;primaryKey"`

	Carbon1 *carbon.Carbon `gorm:"type:timestamp;" json:"carbon1"`
	Carbon2 *carbon.Carbon `gorm:"type:timestamptz;" json:"carbon2"`

	DateLayout1 *carbon.Date `gorm:"type:date;" json:"date_carbon2"`

	TimeLayout1 *carbon.Time `gorm:"type:time;" json:"time_layout1"`
	TimeLayout2 *carbon.Time `gorm:"type:timetz;" json:"time_layout2"`

	DateTimeLayout1 *carbon.DateTime `gorm:"type:timestamp;" json:"date_time_layout1"`
	DateTimeLayout2 *carbon.DateTime `gorm:"type:timestamptz;" json:"date_time_layout2"`

	CustomerFormat1 *carbon.FormatType[ISO8601Type] `gorm:"type:timestamp;" json:"customer_format1"`
	CustomerFormat2 *carbon.FormatType[ISO8601Type] `gorm:"type:timestamptz;" json:"customer_format2"`

	Timestamp1 *carbon.Timestamp `gorm:"column:timestamp1;type:int4;" json:"timestamp1"`

	CreatedAt *carbon.DateTime `gorm:"autoCreateTime;column:created_at;type:timestamptz;" json:"-"`
	UpdatedAt *carbon.DateTime `gorm:"autoUpdateTime;column:updated_at;type:timestamptz;" json:"-"`
	DeletedAt *gorm.DeletedAt  `gorm:"column:deleted_at;type:timestamptz;" json:"-"`
}

func (PgSQLModel2) TableName() string {
	return "carbon_pgsql2"
}
