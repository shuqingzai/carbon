package gorm

import (
	"encoding/json"
	"testing"

	"github.com/dromara/carbon/v2"
	"github.com/stretchr/testify/suite"
)

type MySQLSuite struct {
	suite.Suite
}

func TestMySQLSuite(t *testing.T) {
	suite.Run(t, new(MySQLSuite))
}

func (s *MySQLSuite) SetupSuite() {
	carbon.SetTimezone(carbon.PRC)
	carbon.SetTestNow(carbon.Parse("2020-08-05 13:14:15.111111111"))
	db = connect(driverMySQL)
	if err = db.AutoMigrate(&MySQLModel1{}); err != nil {
		panic(err)
	}
	if err = db.AutoMigrate(&MySQLModel2{}); err != nil {
		panic(err)
	}
}

func (s *MySQLSuite) TearDownSuite() {
	carbon.ClearTestNow()
	//db.Unscoped().Where("1 = 1").Delete(&MySQLModel1{})
	//db.Unscoped().Where("1 = 1").Delete(&MySQLModel2{})
}

func (s *MySQLSuite) TestCurd1() {
	s.Run("unset carbon", func() {
		var model1 MySQLModel1

		// create
		if err = db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 MySQLModel1
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"carbon3":null,"date1":null,"date2":null,"date3":null,"time1":null,"time2":null,"time3":null,"date_time1":null,"date_time2":null,"date_time3":null,"rfc3339_layout1":null,"rfc3339_layout2":null,"rfc3339_layout3":null,"iso8601_format1":null,"iso8601_format2":null,"iso8601_format3":null,"timestamp1":null,"timestamp2":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("nil carbon", func() {
		var model1 MySQLModel1

		var c *carbon.Carbon
		c = nil

		model1.Date1 = *carbon.NewDate(c)
		model1.Date2 = *carbon.NewDate(c)
		model1.Date3 = *carbon.NewDate(c)

		model1.Time1 = *carbon.NewTime(c)
		model1.Time2 = *carbon.NewTime(c)
		model1.Time3 = *carbon.NewTime(c)

		model1.DateTime1 = *carbon.NewDateTime(c)
		model1.DateTime2 = *carbon.NewDateTime(c)
		model1.DateTime3 = *carbon.NewDateTime(c)

		model1.RFC3339Layout1 = *carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout2 = *carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout3 = *carbon.NewLayoutType[RFC3339Layout](c)

		model1.ISO8601Format1 = *carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format2 = *carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format3 = *carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp1 = *carbon.NewTimestamp(c)
		model1.Timestamp2 = *carbon.NewTimestamp(c)

		// create
		if err = db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 MySQLModel1
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"carbon3":null,"date1":null,"date2":null,"date3":null,"time1":null,"time2":null,"time3":null,"date_time1":null,"date_time2":null,"date_time3":null,"rfc3339_layout1":null,"rfc3339_layout2":null,"rfc3339_layout3":null,"iso8601_format1":null,"iso8601_format2":null,"iso8601_format3":null,"timestamp1":null,"timestamp2":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("zero carbon", func() {
		var model1 MySQLModel1

		c := carbon.NewCarbon()

		model1.Carbon1 = *c
		model1.Carbon2 = *c
		model1.Carbon3 = *c

		model1.Date1 = *carbon.NewDate(c)
		model1.Date2 = *carbon.NewDate(c)
		model1.Date3 = *carbon.NewDate(c)

		model1.Time1 = *carbon.NewTime(c)
		model1.Time2 = *carbon.NewTime(c)
		model1.Time3 = *carbon.NewTime(c)

		model1.DateTime1 = *carbon.NewDateTime(c)
		model1.DateTime2 = *carbon.NewDateTime(c)
		model1.DateTime3 = *carbon.NewDateTime(c)

		model1.RFC3339Layout1 = *carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout2 = *carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout3 = *carbon.NewLayoutType[RFC3339Layout](c)

		model1.ISO8601Format1 = *carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format2 = *carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format3 = *carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp1 = *carbon.NewTimestamp(c)
		model1.Timestamp2 = *carbon.NewTimestamp(c)

		// create
		if err = db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 MySQLModel1
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"carbon3":null,"date1":null,"date2":null,"date3":null,"time1":null,"time2":null,"time3":null,"date_time1":null,"date_time2":null,"date_time3":null,"rfc3339_layout1":null,"rfc3339_layout2":null,"rfc3339_layout3":null,"iso8601_format1":null,"iso8601_format2":null,"iso8601_format3":null,"timestamp1":null,"timestamp2":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("empty carbon", func() {
		var model1 MySQLModel1

		c := carbon.Parse("")

		model1.Carbon1 = *c
		model1.Carbon2 = *c
		model1.Carbon3 = *c

		model1.Date1 = *carbon.NewDate(c)
		model1.Date2 = *carbon.NewDate(c)
		model1.Date3 = *carbon.NewDate(c)

		model1.Time1 = *carbon.NewTime(c)
		model1.Time2 = *carbon.NewTime(c)
		model1.Time3 = *carbon.NewTime(c)

		model1.DateTime1 = *carbon.NewDateTime(c)
		model1.DateTime2 = *carbon.NewDateTime(c)
		model1.DateTime3 = *carbon.NewDateTime(c)

		model1.RFC3339Layout1 = *carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout2 = *carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout3 = *carbon.NewLayoutType[RFC3339Layout](c)

		model1.ISO8601Format1 = *carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format2 = *carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format3 = *carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp1 = *carbon.NewTimestamp(c)
		model1.Timestamp2 = *carbon.NewTimestamp(c)

		// create
		if err = db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 MySQLModel1
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"carbon3":null,"date1":null,"date2":null,"date3":null,"time1":null,"time2":null,"time3":null,"date_time1":null,"date_time2":null,"date_time3":null,"rfc3339_layout1":null,"rfc3339_layout2":null,"rfc3339_layout3":null,"iso8601_format1":null,"iso8601_format2":null,"iso8601_format3":null,"timestamp1":null,"timestamp2":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("valid carbon", func() {
		var model1 MySQLModel1

		c := carbon.Now()

		model1.Carbon1 = *c
		model1.Carbon2 = *c
		model1.Carbon3 = *c

		model1.Date1 = *carbon.NewDate(c)
		model1.Date2 = *carbon.NewDate(c)
		model1.Date3 = *carbon.NewDate(c)

		model1.Time1 = *carbon.NewTime(c)
		model1.Time2 = *carbon.NewTime(c)
		model1.Time3 = *carbon.NewTime(c)

		model1.DateTime1 = *carbon.NewDateTime(c)
		model1.DateTime2 = *carbon.NewDateTime(c)
		model1.DateTime3 = *carbon.NewDateTime(c)

		model1.RFC3339Layout1 = *carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout2 = *carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout3 = *carbon.NewLayoutType[RFC3339Layout](c)

		model1.ISO8601Format1 = *carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format2 = *carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format3 = *carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp1 = *carbon.NewTimestamp(c)
		model1.Timestamp2 = *carbon.NewTimestamp(c)

		// create
		if err = db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 MySQLModel1
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":"2020-08-05 13:14:15","carbon2":"2020-08-05 13:14:15","carbon3":"2020-08-05 13:14:15","date1":"2020-08-05","date2":"2020-08-05","date3":"2020-08-05","time1":"13:14:15","time2":"13:14:15","time3":"13:14:15","date_time1":"2020-08-05 13:14:15","date_time2":"2020-08-05 13:14:15","date_time3":"2020-08-05 13:14:15","rfc3339_layout1":"2020-08-05T13:14:15+08:00","rfc3339_layout2":"2020-08-05T13:14:15+08:00","rfc3339_layout3":"2020-08-05T13:14:15+08:00","iso8601_format1":"2020-08-05T13:14:15+08:00","iso8601_format2":"2020-08-05T13:14:15+08:00","iso8601_format3":"2020-08-05T13:14:15+08:00","timestamp1":1596604455,"timestamp2":1596604455}`, string(data1))

		c = c.Copy().AddDay()

		model2.Carbon1 = *c
		model2.Carbon2 = *c
		model2.Carbon3 = *c

		model2.Date1 = *carbon.NewDate(c)
		model2.Date2 = *carbon.NewDate(c)
		model2.Date3 = *carbon.NewDate(c)

		model2.Time1 = *carbon.NewTime(c)
		model2.Time2 = *carbon.NewTime(c)
		model2.Time3 = *carbon.NewTime(c)

		model2.DateTime1 = *carbon.NewDateTime(c)
		model2.DateTime2 = *carbon.NewDateTime(c)
		model2.DateTime3 = *carbon.NewDateTime(c)

		model2.RFC3339Layout1 = *carbon.NewLayoutType[RFC3339Layout](c)
		model2.RFC3339Layout2 = *carbon.NewLayoutType[RFC3339Layout](c)
		model2.RFC3339Layout3 = *carbon.NewLayoutType[RFC3339Layout](c)

		model2.ISO8601Format1 = *carbon.NewFormatType[ISO8601Format](c)
		model2.ISO8601Format2 = *carbon.NewFormatType[ISO8601Format](c)
		model2.ISO8601Format3 = *carbon.NewFormatType[ISO8601Format](c)

		model2.Timestamp1 = *carbon.NewTimestamp(c)
		model2.Timestamp2 = *carbon.NewTimestamp(c)

		// update
		db.Save(&model2)

		data2, err2 := json.Marshal(&model2)
		s.Nil(err2)
		s.Equal(`{"carbon1":"2020-08-06 13:14:15","carbon2":"2020-08-06 13:14:15","carbon3":"2020-08-06 13:14:15","date1":"2020-08-06","date2":"2020-08-06","date3":"2020-08-06","time1":"13:14:15","time2":"13:14:15","time3":"13:14:15","date_time1":"2020-08-06 13:14:15","date_time2":"2020-08-06 13:14:15","date_time3":"2020-08-06 13:14:15","rfc3339_layout1":"2020-08-06T13:14:15+08:00","rfc3339_layout2":"2020-08-06T13:14:15+08:00","rfc3339_layout3":"2020-08-06T13:14:15+08:00","iso8601_format1":"2020-08-06T13:14:15+08:00","iso8601_format2":"2020-08-06T13:14:15+08:00","iso8601_format3":"2020-08-06T13:14:15+08:00","timestamp1":1596690855,"timestamp2":1596690855}`, string(data2))

		// delete
		db.Delete(&model2)
	})
}

func (s *MySQLSuite) TestCurd2() {
	s.Run("unset carbon", func() {
		var model1 MySQLModel2

		// create
		if err = db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 MySQLModel2
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"carbon3":null,"date1":null,"date2":null,"date3":null,"time1":null,"time2":null,"time3":null,"date_time1":null,"date_time2":null,"date_time3":null,"rfc3339_layout1":null,"rfc3339_layout2":null,"rfc3339_layout3":null,"iso8601_format1":null,"iso8601_format2":null,"iso8601_format3":null,"timestamp1":null,"timestamp2":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("nil carbon", func() {
		var model1 MySQLModel2

		var c *carbon.Carbon
		c = nil

		model1.Carbon1 = c
		model1.Carbon2 = c
		model1.Carbon3 = c

		model1.Date1 = carbon.NewDate(c)
		model1.Date2 = carbon.NewDate(c)
		model1.Date3 = carbon.NewDate(c)

		model1.Time1 = carbon.NewTime(c)
		model1.Time2 = carbon.NewTime(c)
		model1.Time3 = carbon.NewTime(c)

		model1.DateTime1 = carbon.NewDateTime(c)
		model1.DateTime2 = carbon.NewDateTime(c)
		model1.DateTime3 = carbon.NewDateTime(c)

		model1.RFC3339Layout1 = carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout2 = carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout3 = carbon.NewLayoutType[RFC3339Layout](c)

		model1.ISO8601Format1 = carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format2 = carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format3 = carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp1 = carbon.NewTimestamp(c)
		model1.Timestamp2 = carbon.NewTimestamp(c)

		// create
		if err = db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 MySQLModel1
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"carbon3":null,"date1":null,"date2":null,"date3":null,"time1":null,"time2":null,"time3":null,"date_time1":null,"date_time2":null,"date_time3":null,"rfc3339_layout1":null,"rfc3339_layout2":null,"rfc3339_layout3":null,"iso8601_format1":null,"iso8601_format2":null,"iso8601_format3":null,"timestamp1":null,"timestamp2":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("zero carbon", func() {
		var model1 MySQLModel2

		c := carbon.NewCarbon()

		model1.Carbon1 = c
		model1.Carbon2 = c
		model1.Carbon3 = c

		model1.Date1 = carbon.NewDate(c)
		model1.Date2 = carbon.NewDate(c)
		model1.Date3 = carbon.NewDate(c)

		model1.Time1 = carbon.NewTime(c)
		model1.Time2 = carbon.NewTime(c)
		model1.Time3 = carbon.NewTime(c)

		model1.DateTime1 = carbon.NewDateTime(c)
		model1.DateTime2 = carbon.NewDateTime(c)
		model1.DateTime3 = carbon.NewDateTime(c)

		model1.RFC3339Layout1 = carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout2 = carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout3 = carbon.NewLayoutType[RFC3339Layout](c)

		model1.ISO8601Format1 = carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format2 = carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format3 = carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp1 = carbon.NewTimestamp(c)
		model1.Timestamp2 = carbon.NewTimestamp(c)

		// create
		if err = db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 MySQLModel2
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"carbon3":null,"date1":null,"date2":null,"date3":null,"time1":null,"time2":null,"time3":null,"date_time1":null,"date_time2":null,"date_time3":null,"rfc3339_layout1":null,"rfc3339_layout2":null,"rfc3339_layout3":null,"iso8601_format1":null,"iso8601_format2":null,"iso8601_format3":null,"timestamp1":null,"timestamp2":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("empty carbon", func() {
		var model1 MySQLModel2

		c := carbon.Parse("")

		model1.Carbon1 = c
		model1.Carbon2 = c
		model1.Carbon3 = c

		model1.Date1 = carbon.NewDate(c)
		model1.Date2 = carbon.NewDate(c)
		model1.Date3 = carbon.NewDate(c)

		model1.Time1 = carbon.NewTime(c)
		model1.Time2 = carbon.NewTime(c)
		model1.Time3 = carbon.NewTime(c)

		model1.DateTime1 = carbon.NewDateTime(c)
		model1.DateTime2 = carbon.NewDateTime(c)
		model1.DateTime3 = carbon.NewDateTime(c)

		model1.RFC3339Layout1 = carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout2 = carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout3 = carbon.NewLayoutType[RFC3339Layout](c)

		model1.ISO8601Format1 = carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format2 = carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format3 = carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp1 = carbon.NewTimestamp(c)
		model1.Timestamp2 = carbon.NewTimestamp(c)

		// create
		if err = db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 MySQLModel2
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"carbon3":null,"date1":null,"date2":null,"date3":null,"time1":null,"time2":null,"time3":null,"date_time1":null,"date_time2":null,"date_time3":null,"rfc3339_layout1":null,"rfc3339_layout2":null,"rfc3339_layout3":null,"iso8601_format1":null,"iso8601_format2":null,"iso8601_format3":null,"timestamp1":null,"timestamp2":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("valid carbon", func() {
		var model1 MySQLModel2

		c := carbon.Now()

		model1.Carbon1 = c
		model1.Carbon2 = c
		model1.Carbon3 = c

		model1.Date1 = carbon.NewDate(c)
		model1.Date2 = carbon.NewDate(c)
		model1.Date3 = carbon.NewDate(c)

		model1.Time1 = carbon.NewTime(c)
		model1.Time2 = carbon.NewTime(c)
		model1.Time3 = carbon.NewTime(c)

		model1.DateTime1 = carbon.NewDateTime(c)
		model1.DateTime2 = carbon.NewDateTime(c)
		model1.DateTime3 = carbon.NewDateTime(c)

		model1.RFC3339Layout1 = carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout2 = carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout3 = carbon.NewLayoutType[RFC3339Layout](c)

		model1.ISO8601Format1 = carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format2 = carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format3 = carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp1 = carbon.NewTimestamp(c)
		model1.Timestamp2 = carbon.NewTimestamp(c)

		// create
		if err = db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 MySQLModel2
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":"2020-08-05 13:14:15","carbon2":"2020-08-05 13:14:15","carbon3":"2020-08-05 13:14:15","date1":"2020-08-05","date2":"2020-08-05","date3":"2020-08-05","time1":"13:14:15","time2":"13:14:15","time3":"13:14:15","date_time1":"2020-08-05 13:14:15","date_time2":"2020-08-05 13:14:15","date_time3":"2020-08-05 13:14:15","rfc3339_layout1":"2020-08-05T13:14:15+08:00","rfc3339_layout2":"2020-08-05T13:14:15+08:00","rfc3339_layout3":"2020-08-05T13:14:15+08:00","iso8601_format1":"2020-08-05T13:14:15+08:00","iso8601_format2":"2020-08-05T13:14:15+08:00","iso8601_format3":"2020-08-05T13:14:15+08:00","timestamp1":1596604455,"timestamp2":1596604455}`, string(data1))

		c = c.Copy().AddDay()

		model2.Carbon1 = c
		model2.Carbon2 = c
		model2.Carbon3 = c

		model2.Date1 = carbon.NewDate(c)
		model2.Date2 = carbon.NewDate(c)
		model2.Date3 = carbon.NewDate(c)

		model2.Time1 = carbon.NewTime(c)
		model2.Time2 = carbon.NewTime(c)
		model2.Time3 = carbon.NewTime(c)

		model2.DateTime1 = carbon.NewDateTime(c)
		model2.DateTime2 = carbon.NewDateTime(c)
		model2.DateTime3 = carbon.NewDateTime(c)

		model2.RFC3339Layout1 = carbon.NewLayoutType[RFC3339Layout](c)
		model2.RFC3339Layout2 = carbon.NewLayoutType[RFC3339Layout](c)
		model2.RFC3339Layout3 = carbon.NewLayoutType[RFC3339Layout](c)

		model2.ISO8601Format1 = carbon.NewFormatType[ISO8601Format](c)
		model2.ISO8601Format2 = carbon.NewFormatType[ISO8601Format](c)
		model2.ISO8601Format3 = carbon.NewFormatType[ISO8601Format](c)

		model2.Timestamp1 = carbon.NewTimestamp(c)
		model2.Timestamp2 = carbon.NewTimestamp(c)

		// update
		db.Save(&model2)

		data2, err2 := json.Marshal(&model2)
		s.Nil(err2)
		s.Equal(`{"carbon1":"2020-08-06 13:14:15","carbon2":"2020-08-06 13:14:15","carbon3":"2020-08-06 13:14:15","date1":"2020-08-06","date2":"2020-08-06","date3":"2020-08-06","time1":"13:14:15","time2":"13:14:15","time3":"13:14:15","date_time1":"2020-08-06 13:14:15","date_time2":"2020-08-06 13:14:15","date_time3":"2020-08-06 13:14:15","rfc3339_layout1":"2020-08-06T13:14:15+08:00","rfc3339_layout2":"2020-08-06T13:14:15+08:00","rfc3339_layout3":"2020-08-06T13:14:15+08:00","iso8601_format1":"2020-08-06T13:14:15+08:00","iso8601_format2":"2020-08-06T13:14:15+08:00","iso8601_format3":"2020-08-06T13:14:15+08:00","timestamp1":1596690855,"timestamp2":1596690855}`, string(data2))

		// delete
		db.Delete(&model2)
	})
}
