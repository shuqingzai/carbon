package tests

import (
	"encoding/json"
	"testing"

	"gitee.com/golang-package/carbon/v2"
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
	carbon.SetTestNow(carbon.Parse("2020-08-05 13:14:15"))
	db = connect("mysql")
	if err := db.AutoMigrate(&MySQLModel1{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&MySQLModel2{}); err != nil {
		panic(err)
	}
}

func (s *MySQLSuite) TearDownSuite() {
	carbon.CleanTestNow()
	db.Unscoped().Where("1 = 1").Delete(&MySQLModel1{})
	db.Unscoped().Where("1 = 1").Delete(&MySQLModel2{})
}

func (s *MySQLSuite) TestCurd1() {
	s.Run("unset carbon", func() {
		var model1 MySQLModel1

		// create
		if err := db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 MySQLModel1
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"date_time_carbon1":null,"date_time_carbon2":null,"date_time_carbon3":null,"date_time_layout1":null,"date_time_layout2":null,"date_time_layout3":null,"date_time_format1":null,"date_time_format2":null,"date_time_format3":null,"timestamp1":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("nil carbon", func() {
		var model1 MySQLModel1

		var c *carbon.Carbon
		c = nil

		model1.DateTimeLayout1 = *carbon.NewDateTime(c)
		model1.DateTimeLayout2 = *carbon.NewDateTime(c)
		model1.DateTimeLayout3 = *carbon.NewDateTime(c)

		model1.DateTimeFormat1 = *carbon.NewFormatType[ISO8601Type](c)
		model1.DateTimeFormat2 = *carbon.NewFormatType[ISO8601Type](c)
		model1.DateTimeFormat3 = *carbon.NewFormatType[ISO8601Type](c)

		model1.Timestamp1 = *carbon.NewTimestamp(c)

		// create
		if err := db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 MySQLModel1
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"date_time_carbon1":null,"date_time_carbon2":null,"date_time_carbon3":null,"date_time_layout1":null,"date_time_layout2":null,"date_time_layout3":null,"date_time_format1":null,"date_time_format2":null,"date_time_format3":null,"timestamp1":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("zero carbon", func() {
		var model1 MySQLModel1

		c := carbon.NewCarbon()

		model1.DateTimeCarbon1 = *c
		model1.DateTimeCarbon2 = *c
		model1.DateTimeCarbon3 = *c

		model1.DateTimeLayout1 = *carbon.NewDateTime(c)
		model1.DateTimeLayout2 = *carbon.NewDateTime(c)
		model1.DateTimeLayout3 = *carbon.NewDateTime(c)

		model1.DateTimeFormat1 = *carbon.NewFormatType[ISO8601Type](c)
		model1.DateTimeFormat2 = *carbon.NewFormatType[ISO8601Type](c)
		model1.DateTimeFormat3 = *carbon.NewFormatType[ISO8601Type](c)

		model1.Timestamp1 = *carbon.NewTimestamp(c)

		// create
		if err := db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 MySQLModel1
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"date_time_carbon1":null,"date_time_carbon2":null,"date_time_carbon3":null,"date_time_layout1":null,"date_time_layout2":null,"date_time_layout3":null,"date_time_format1":null,"date_time_format2":null,"date_time_format3":null,"timestamp1":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("empty carbon", func() {
		var model1 MySQLModel1

		c := carbon.Parse("")

		model1.DateTimeCarbon1 = *c
		model1.DateTimeCarbon2 = *c
		model1.DateTimeCarbon3 = *c

		model1.DateTimeLayout1 = *carbon.NewDateTime(c)
		model1.DateTimeLayout2 = *carbon.NewDateTime(c)
		model1.DateTimeLayout3 = *carbon.NewDateTime(c)

		model1.DateTimeFormat1 = *carbon.NewFormatType[ISO8601Type](c)
		model1.DateTimeFormat2 = *carbon.NewFormatType[ISO8601Type](c)
		model1.DateTimeFormat3 = *carbon.NewFormatType[ISO8601Type](c)

		model1.Timestamp1 = *carbon.NewTimestamp(c)

		// create
		if err := db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 MySQLModel1
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"date_time_carbon1":null,"date_time_carbon2":null,"date_time_carbon3":null,"date_time_layout1":null,"date_time_layout2":null,"date_time_layout3":null,"date_time_format1":null,"date_time_format2":null,"date_time_format3":null,"timestamp1":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("valid carbon", func() {
		var model1 MySQLModel1

		c := carbon.Now()

		model1.DateTimeCarbon1 = *c
		model1.DateTimeCarbon2 = *c
		model1.DateTimeCarbon3 = *c

		model1.DateTimeLayout1 = *carbon.NewDateTime(c)
		model1.DateTimeLayout2 = *carbon.NewDateTime(c)
		model1.DateTimeLayout3 = *carbon.NewDateTime(c)

		model1.DateTimeFormat1 = *carbon.NewFormatType[ISO8601Type](c)
		model1.DateTimeFormat2 = *carbon.NewFormatType[ISO8601Type](c)
		model1.DateTimeFormat3 = *carbon.NewFormatType[ISO8601Type](c)

		model1.Timestamp1 = *carbon.NewTimestamp(c)

		// create
		if err := db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 MySQLModel1
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"date_time_carbon1":"2020-08-05 13:14:15","date_time_carbon2":"2020-08-05 13:14:15","date_time_carbon3":"2020-08-05 13:14:15","date_time_layout1":"2020-08-05 13:14:15","date_time_layout2":"2020-08-05 13:14:15","date_time_layout3":"2020-08-05 13:14:15","date_time_format1":"2020-08-05T13:14:15+08:00","date_time_format2":"2020-08-05T13:14:15+08:00","date_time_format3":"2020-08-05T13:14:15+08:00","timestamp1":1596604455}`, string(data1))

		model2.DateTimeCarbon1 = *c.Copy().AddDay()
		model2.DateTimeCarbon2 = *c.Copy().AddDay()
		model2.DateTimeCarbon3 = *c.Copy().AddDay()

		model2.DateTimeLayout1 = *carbon.NewDateTime(c.Copy().AddDay())
		model2.DateTimeLayout2 = *carbon.NewDateTime(c.Copy().AddDay())
		model2.DateTimeLayout3 = *carbon.NewDateTime(c.Copy().AddDay())

		model2.DateTimeFormat1 = *carbon.NewFormatType[ISO8601Type](c.Copy().AddDay())
		model2.DateTimeFormat2 = *carbon.NewFormatType[ISO8601Type](c.Copy().AddDay())
		model2.DateTimeFormat3 = *carbon.NewFormatType[ISO8601Type](c.Copy().AddDay())

		model2.Timestamp1 = *carbon.NewTimestamp(c.Copy().AddDay())

		// update
		db.Save(&model2)

		data2, err2 := json.Marshal(&model2)
		s.Nil(err2)
		s.Equal(`{"date_time_carbon1":"2020-08-06 13:14:15","date_time_carbon2":"2020-08-06 13:14:15","date_time_carbon3":"2020-08-06 13:14:15","date_time_layout1":"2020-08-06 13:14:15","date_time_layout2":"2020-08-06 13:14:15","date_time_layout3":"2020-08-06 13:14:15","date_time_format1":"2020-08-06T13:14:15+08:00","date_time_format2":"2020-08-06T13:14:15+08:00","date_time_format3":"2020-08-06T13:14:15+08:00","timestamp1":1596690855}`, string(data2))

		// delete
		db.Delete(&model2)
	})
}

func (s *MySQLSuite) TestCurd2() {
	s.Run("unset carbon", func() {
		var model1 MySQLModel2

		// create
		if err := db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 MySQLModel1
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"date_time_carbon1":null,"date_time_carbon2":null,"date_time_carbon3":null,"date_time_layout1":null,"date_time_layout2":null,"date_time_layout3":null,"date_time_format1":null,"date_time_format2":null,"date_time_format3":null,"timestamp1":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("nil carbon", func() {
		var model1 MySQLModel2

		var c *carbon.Carbon
		c = nil

		model1.DateTimeCarbon1 = c
		model1.DateTimeCarbon2 = c
		model1.DateTimeCarbon3 = c

		model1.DateTimeLayout1 = carbon.NewDateTime(c)
		model1.DateTimeLayout2 = carbon.NewDateTime(c)
		model1.DateTimeLayout3 = carbon.NewDateTime(c)

		model1.DateTimeFormat1 = carbon.NewFormatType[ISO8601Type](c)
		model1.DateTimeFormat2 = carbon.NewFormatType[ISO8601Type](c)
		model1.DateTimeFormat3 = carbon.NewFormatType[ISO8601Type](c)

		model1.Timestamp1 = carbon.NewTimestamp(c)

		// create
		if err := db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 MySQLModel1
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"date_time_carbon1":null,"date_time_carbon2":null,"date_time_carbon3":null,"date_time_layout1":null,"date_time_layout2":null,"date_time_layout3":null,"date_time_format1":null,"date_time_format2":null,"date_time_format3":null,"timestamp1":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("zero carbon", func() {
		var model1 MySQLModel2

		c := carbon.NewCarbon()

		model1.DateTimeCarbon1 = c
		model1.DateTimeCarbon2 = c
		model1.DateTimeCarbon3 = c

		model1.DateTimeLayout1 = carbon.NewDateTime(c)
		model1.DateTimeLayout2 = carbon.NewDateTime(c)
		model1.DateTimeLayout3 = carbon.NewDateTime(c)

		model1.DateTimeFormat1 = carbon.NewFormatType[ISO8601Type](c)
		model1.DateTimeFormat2 = carbon.NewFormatType[ISO8601Type](c)
		model1.DateTimeFormat3 = carbon.NewFormatType[ISO8601Type](c)

		model1.Timestamp1 = carbon.NewTimestamp(c)

		// create
		if err := db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 MySQLModel1
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"date_time_carbon1":null,"date_time_carbon2":null,"date_time_carbon3":null,"date_time_layout1":null,"date_time_layout2":null,"date_time_layout3":null,"date_time_format1":null,"date_time_format2":null,"date_time_format3":null,"timestamp1":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("empty carbon", func() {
		var model1 MySQLModel2

		c := carbon.Parse("")

		model1.DateTimeCarbon1 = c
		model1.DateTimeCarbon2 = c
		model1.DateTimeCarbon3 = c

		model1.DateTimeLayout1 = carbon.NewDateTime(c)
		model1.DateTimeLayout2 = carbon.NewDateTime(c)
		model1.DateTimeLayout3 = carbon.NewDateTime(c)

		model1.DateTimeFormat1 = carbon.NewFormatType[ISO8601Type](c)
		model1.DateTimeFormat2 = carbon.NewFormatType[ISO8601Type](c)
		model1.DateTimeFormat3 = carbon.NewFormatType[ISO8601Type](c)

		model1.Timestamp1 = carbon.NewTimestamp(c)

		// create
		if err := db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 MySQLModel1
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"date_time_carbon1":null,"date_time_carbon2":null,"date_time_carbon3":null,"date_time_layout1":null,"date_time_layout2":null,"date_time_layout3":null,"date_time_format1":null,"date_time_format2":null,"date_time_format3":null,"timestamp1":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("valid carbon", func() {
		var model1 MySQLModel2

		c := carbon.Now()

		model1.DateTimeCarbon1 = c
		model1.DateTimeCarbon2 = c
		model1.DateTimeCarbon3 = c

		model1.DateTimeLayout1 = carbon.NewDateTime(c)
		model1.DateTimeLayout2 = carbon.NewDateTime(c)
		model1.DateTimeLayout3 = carbon.NewDateTime(c)

		model1.DateTimeFormat1 = carbon.NewFormatType[ISO8601Type](c)
		model1.DateTimeFormat2 = carbon.NewFormatType[ISO8601Type](c)
		model1.DateTimeFormat3 = carbon.NewFormatType[ISO8601Type](c)

		model1.Timestamp1 = carbon.NewTimestamp(c)

		// create
		if err := db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 MySQLModel2
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"date_time_carbon1":"2020-08-05 13:14:15","date_time_carbon2":"2020-08-05 13:14:15","date_time_carbon3":"2020-08-05 13:14:15","date_time_layout1":"2020-08-05 13:14:15","date_time_layout2":"2020-08-05 13:14:15","date_time_layout3":"2020-08-05 13:14:15","date_time_format1":"2020-08-05T13:14:15+08:00","date_time_format2":"2020-08-05T13:14:15+08:00","date_time_format3":"2020-08-05T13:14:15+08:00","timestamp1":1596604455}`, string(data1))

		model2.DateTimeCarbon1 = c.Copy().AddDay()
		model2.DateTimeCarbon2 = c.Copy().AddDay()
		model2.DateTimeCarbon3 = c.Copy().AddDay()

		model2.DateTimeLayout1 = carbon.NewDateTime(c.Copy().AddDay())
		model2.DateTimeLayout2 = carbon.NewDateTime(c.Copy().AddDay())
		model2.DateTimeLayout3 = carbon.NewDateTime(c.Copy().AddDay())

		model2.DateTimeFormat1 = carbon.NewFormatType[ISO8601Type](c.Copy().AddDay())
		model2.DateTimeFormat2 = carbon.NewFormatType[ISO8601Type](c.Copy().AddDay())
		model2.DateTimeFormat3 = carbon.NewFormatType[ISO8601Type](c.Copy().AddDay())

		model2.Timestamp1 = carbon.NewTimestamp(c.Copy().AddDay())

		// update
		db.Save(&model2)

		data2, err2 := json.Marshal(&model2)
		s.Nil(err2)
		s.Equal(`{"date_time_carbon1":"2020-08-06 13:14:15","date_time_carbon2":"2020-08-06 13:14:15","date_time_carbon3":"2020-08-06 13:14:15","date_time_layout1":"2020-08-06 13:14:15","date_time_layout2":"2020-08-06 13:14:15","date_time_layout3":"2020-08-06 13:14:15","date_time_format1":"2020-08-06T13:14:15+08:00","date_time_format2":"2020-08-06T13:14:15+08:00","date_time_format3":"2020-08-06T13:14:15+08:00","timestamp1":1596690855}`, string(data2))

		// delete
		db.Delete(&model2)
	})
}
