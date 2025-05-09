package tests

import (
	"github.com/dromara/carbon/v2"
	"gorm.io/gorm"
)

type RFC3339Layout string

func (t RFC3339Layout) Layout() string {
	return carbon.RFC3339Layout
}

type ISO8601Format string

func (t ISO8601Format) Format() string {
	return carbon.ISO8601Format
}

type MySQLModel1 struct {
	ID uint64 `json:"-" gorm:"column:id;primaryKey"`

	Carbon1 carbon.Carbon `gorm:"column:carbon1;type:varchar(50);" json:"carbon1"`
	Carbon2 carbon.Carbon `gorm:"column:carbon2;type:datetime;" json:"carbon2"`
	Carbon3 carbon.Carbon `gorm:"column:carbon3;type:timestamp;" json:"carbon3"`

	DateTime1 carbon.DateTime `gorm:"column:date_time1;type:varchar(50);" json:"date_time1"`
	DateTime2 carbon.DateTime `gorm:"column:date_time2;type:datetime;" json:"date_time2"`
	DateTime3 carbon.DateTime `gorm:"column:date_time3;type:timestamp;" json:"date_time3"`

	RFC3339Layout1 carbon.LayoutType[RFC3339Layout] `gorm:"column:rfc3339_layout1;type:varchar(50);" json:"rfc3339_layout1"`
	RFC3339Layout2 carbon.LayoutType[RFC3339Layout] `gorm:"column:rfc3339_layout2;type:datetime;" json:"rfc3339_layout2"`
	RFC3339Layout3 carbon.LayoutType[RFC3339Layout] `gorm:"column:rfc3339_layout3;type:timestamp;" json:"rfc3339_layout3"`

	ISO8601Format1 carbon.FormatType[ISO8601Format] `gorm:"column:iso8601_format1;type:varchar(50);" json:"iso8601_format1"`
	ISO8601Format2 carbon.FormatType[ISO8601Format] `gorm:"column:iso8601_format2;type:datetime;" json:"iso8601_format2"`
	ISO8601Format3 carbon.FormatType[ISO8601Format] `gorm:"column:iso8601_format3;type:timestamp;" json:"iso8601_format3"`

	Timestamp1 carbon.Timestamp `gorm:"column:timestamp1;type:bigint;" json:"timestamp1"`

	CreatedAt carbon.DateTime `gorm:"autoCreateTime;column:created_at;type:timestamp;" json:"-"`
	UpdatedAt carbon.DateTime `gorm:"autoUpdateTime;column:updated_at;type:timestamp;" json:"-"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at;type:datetime;" json:"-"`
}

func (MySQLModel1) TableName() string {
	return "carbon_mysql1"
}

type MySQLModel2 struct {
	ID uint64 `json:"-" gorm:"column:id;primaryKey"`

	Carbon1 *carbon.Carbon `gorm:"column:carbon1;type:varchar(50);" json:"carbon1"`
	Carbon2 *carbon.Carbon `gorm:"column:carbon2;type:datetime;" json:"carbon2"`
	Carbon3 *carbon.Carbon `gorm:"column:carbon3;type:timestamp;" json:"carbon3"`

	DateTime1 *carbon.DateTime `gorm:"column:date_time1;type:varchar(50);" json:"date_time1"`
	DateTime2 *carbon.DateTime `gorm:"column:date_time2;type:datetime;" json:"date_time2"`
	DateTime3 *carbon.DateTime `gorm:"column:date_time3;type:timestamp;" json:"date_time3"`

	RFC3339Layout1 *carbon.LayoutType[RFC3339Layout] `gorm:"column:rfc3339_layout1;type:varchar(50);" json:"rfc3339_layout1"`
	RFC3339Layout2 *carbon.LayoutType[RFC3339Layout] `gorm:"column:rfc3339_layout2;type:datetime;" json:"rfc3339_layout2"`
	RFC3339Layout3 *carbon.LayoutType[RFC3339Layout] `gorm:"column:rfc3339_layout3;type:timestamp;" json:"rfc3339_layout3"`

	ISO8601Format1 *carbon.FormatType[ISO8601Format] `gorm:"column:iso8601_format1;type:varchar(50);" json:"iso8601_format1"`
	ISO8601Format2 *carbon.FormatType[ISO8601Format] `gorm:"column:iso8601_format2;type:datetime;" json:"iso8601_format2"`
	ISO8601Format3 *carbon.FormatType[ISO8601Format] `gorm:"column:iso8601_format3;type:timestamp;" json:"iso8601_format3"`

	Timestamp1 *carbon.Timestamp `gorm:"column:timestamp1;type:bigint;" json:"timestamp1"`

	CreatedAt *carbon.DateTime `gorm:"autoCreateTime;column:created_at;type:timestamp;" json:"-"`
	UpdatedAt *carbon.DateTime `gorm:"autoUpdateTime;column:updated_at;type:timestamp;" json:"-"`
	DeletedAt *gorm.DeletedAt  `gorm:"column:deleted_at;type:datetime;" json:"-"`
}

func (MySQLModel2) TableName() string {
	return "carbon_mysql2"
}

type PgSQLModel1 struct {
	ID uint64 `json:"-" gorm:"column:id;primaryKey"`

	Carbon1 carbon.Carbon `gorm:"type:timestamp;" json:"carbon1"`
	Carbon2 carbon.Carbon `gorm:"type:timestamptz;" json:"carbon2"`

	Date1 carbon.Date `gorm:"type:date;" json:"date1"`

	Time1 carbon.Time `gorm:"type:time;" json:"time1"`
	Time2 carbon.Time `gorm:"type:timetz;" json:"time2"`

	DateTime1 carbon.DateTime `gorm:"type:timestamp;" json:"date_time1"`
	DateTime2 carbon.DateTime `gorm:"type:timestamptz;" json:"date_time2"`

	RFC3339Layout1 carbon.LayoutType[RFC3339Layout] `gorm:"column:rfc3339_layout1;type:timestamp;" json:"rfc3339_layout1"`
	RFC3339Layout2 carbon.LayoutType[RFC3339Layout] `gorm:"column:rfc3339_layout2;type:timestamptz;" json:"rfc3339_layout2"`

	ISO8601Format1 carbon.FormatType[ISO8601Format] `gorm:"column:iso8601_format1;type:timestamp;" json:"iso8601_format1"`
	ISO8601Format2 carbon.FormatType[ISO8601Format] `gorm:"column:iso8601_format2;type:timestamptz;" json:"iso8601_format2"`

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

	Date1 *carbon.Date `gorm:"type:date;" json:"date1"`

	Time1 *carbon.Time `gorm:"type:time;" json:"time1"`
	Time2 *carbon.Time `gorm:"type:timetz;" json:"time2"`

	DateTime1 *carbon.DateTime `gorm:"type:timestamp;" json:"date_time1"`
	DateTime2 *carbon.DateTime `gorm:"type:timestamptz;" json:"date_time2"`

	RFC3339Layout1 *carbon.LayoutType[RFC3339Layout] `gorm:"column:rfc3339_layout1;type:timestamp;" json:"rfc3339_layout1"`
	RFC3339Layout2 *carbon.LayoutType[RFC3339Layout] `gorm:"column:rfc3339_layout2;type:timestamptz;" json:"rfc3339_layout2"`

	ISO8601Format1 *carbon.FormatType[ISO8601Format] `gorm:"column:iso8601_format1;type:timestamp;" json:"iso8601_format1"`
	ISO8601Format2 *carbon.FormatType[ISO8601Format] `gorm:"column:iso8601_format2;type:timestamptz;" json:"iso8601_format2"`

	Timestamp1 *carbon.Timestamp `gorm:"column:timestamp1;type:int4;" json:"timestamp1"`

	CreatedAt *carbon.DateTime `gorm:"autoCreateTime;column:created_at;type:timestamptz;" json:"-"`
	UpdatedAt *carbon.DateTime `gorm:"autoUpdateTime;column:updated_at;type:timestamptz;" json:"-"`
	DeletedAt *gorm.DeletedAt  `gorm:"column:deleted_at;type:timestamptz;" json:"-"`
}

func (PgSQLModel2) TableName() string {
	return "carbon_pgsql2"
}
