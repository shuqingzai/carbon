package tests

import (
	"encoding/json"
	"testing"

	"github.com/dromara/carbon/v2"
	"github.com/stretchr/testify/suite"
)

type PgSQLSuite struct {
	suite.Suite
}

func TestPgSQLSuite(t *testing.T) {
	suite.Run(t, new(PgSQLSuite))
}

func (s *PgSQLSuite) SetupSuite() {
	carbon.SetTimezone(carbon.PRC)
	carbon.SetTestNow(carbon.Parse("2020-08-05 13:14:15"))
	db = connect("pgsql")
	if err := db.AutoMigrate(&PgSQLModel1{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&PgSQLModel2{}); err != nil {
		panic(err)
	}
}

func (s *PgSQLSuite) TearDownSuite() {
	carbon.CleanTestNow()
	db.Unscoped().Where("1 = 1").Delete(&PgSQLModel1{})
	db.Unscoped().Where("1 = 1").Delete(&PgSQLModel2{})
}

func (s *PgSQLSuite) TestCurd1() {
	s.Run("unset carbon", func() {
		var model1 PgSQLModel1

		// create
		if err := db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 PgSQLModel1
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"date_carbon2":null,"time_layout1":null,"time_layout2":null,"date_time_layout1":null,"date_time_layout2":null,"customer_format1":null,"customer_format2":null,"timestamp1":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("nil carbon", func() {
		var model1 PgSQLModel1

		var c *carbon.Carbon
		c = nil

		model1.DateLayout1 = *carbon.NewDate(c)

		model1.TimeLayout1 = *carbon.NewTime(c)
		model1.TimeLayout2 = *carbon.NewTime(c)

		model1.DateTimeLayout1 = *carbon.NewDateTime(c)
		model1.DateTimeLayout2 = *carbon.NewDateTime(c)

		model1.CustomerFormat1 = *carbon.NewFormatType[ISO8601Type](c)
		model1.CustomerFormat2 = *carbon.NewFormatType[ISO8601Type](c)

		model1.Timestamp1 = *carbon.NewTimestamp(c)

		// create
		if err := db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 PgSQLModel1
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"date_carbon2":null,"time_layout1":null,"time_layout2":null,"date_time_layout1":null,"date_time_layout2":null,"customer_format1":null,"customer_format2":null,"timestamp1":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("zero carbon", func() {
		var model1 PgSQLModel1

		c := carbon.NewCarbon()

		model1.Carbon1 = *c
		model1.Carbon2 = *c

		model1.DateLayout1 = *carbon.NewDate(c)

		model1.TimeLayout1 = *carbon.NewTime(c)
		model1.TimeLayout2 = *carbon.NewTime(c)

		model1.DateTimeLayout1 = *carbon.NewDateTime(c)
		model1.DateTimeLayout2 = *carbon.NewDateTime(c)

		model1.CustomerFormat1 = *carbon.NewFormatType[ISO8601Type](c)
		model1.CustomerFormat2 = *carbon.NewFormatType[ISO8601Type](c)

		model1.Timestamp1 = *carbon.NewTimestamp(c)

		// create
		if err := db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 PgSQLModel1
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"date_carbon2":null,"time_layout1":null,"time_layout2":null,"date_time_layout1":null,"date_time_layout2":null,"customer_format1":null,"customer_format2":null,"timestamp1":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("empty carbon", func() {
		var model1 PgSQLModel1

		c := carbon.Parse("")

		model1.Carbon1 = *c
		model1.Carbon2 = *c

		model1.DateLayout1 = *carbon.NewDate(c)

		model1.TimeLayout1 = *carbon.NewTime(c)
		model1.TimeLayout2 = *carbon.NewTime(c)

		model1.DateTimeLayout1 = *carbon.NewDateTime(c)
		model1.DateTimeLayout2 = *carbon.NewDateTime(c)

		model1.CustomerFormat1 = *carbon.NewFormatType[ISO8601Type](c)
		model1.CustomerFormat2 = *carbon.NewFormatType[ISO8601Type](c)

		model1.Timestamp1 = *carbon.NewTimestamp(c)

		// create
		if err := db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 PgSQLModel1
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"date_carbon2":null,"time_layout1":null,"time_layout2":null,"date_time_layout1":null,"date_time_layout2":null,"customer_format1":null,"customer_format2":null,"timestamp1":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("valid carbon", func() {
		var model1 PgSQLModel1

		c := carbon.Now()

		model1.Carbon1 = *c
		model1.Carbon2 = *c

		model1.DateLayout1 = *carbon.NewDate(c)

		model1.TimeLayout1 = *carbon.NewTime(c)
		model1.TimeLayout2 = *carbon.NewTime(c)

		model1.DateTimeLayout1 = *carbon.NewDateTime(c)
		model1.DateTimeLayout2 = *carbon.NewDateTime(c)

		model1.CustomerFormat1 = *carbon.NewFormatType[ISO8601Type](c)
		model1.CustomerFormat2 = *carbon.NewFormatType[ISO8601Type](c)

		model1.Timestamp1 = *carbon.NewTimestamp(c)

		// create
		if err := db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 PgSQLModel1
		db.Last(&model2)
		//
		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":"2020-08-05 21:14:15","carbon2":"2020-08-05 13:14:15","date_carbon2":"2020-08-05","time_layout1":"13:14:15","time_layout2":"13:19:58","date_time_layout1":"2020-08-05 21:14:15","date_time_layout2":"2020-08-05 13:14:15","customer_format1":"2020-08-05T21:14:15+08:00","customer_format2":"2020-08-05T13:14:15+08:00","timestamp1":1596604455}`, string(data1))

		newCarbon := c.Copy().AddDay()

		model2.Carbon1 = *newCarbon
		model2.Carbon2 = *newCarbon

		model2.DateLayout1 = *carbon.NewDate(newCarbon)

		model2.TimeLayout1 = *carbon.NewTime(newCarbon)
		model2.TimeLayout2 = *carbon.NewTime(newCarbon)

		model2.DateTimeLayout1 = *carbon.NewDateTime(newCarbon)
		model2.DateTimeLayout2 = *carbon.NewDateTime(newCarbon)

		model2.CustomerFormat1 = *carbon.NewFormatType[ISO8601Type](newCarbon)
		model2.CustomerFormat2 = *carbon.NewFormatType[ISO8601Type](newCarbon)

		model2.Timestamp1 = *carbon.NewTimestamp(newCarbon)

		// update
		db.Save(&model2)

		data2, err2 := json.Marshal(&model2)
		s.Nil(err2)
		s.Equal(`{"carbon1":"2020-08-06 13:14:15","carbon2":"2020-08-06 13:14:15","date_carbon2":"2020-08-06","time_layout1":"13:14:15","time_layout2":"13:14:15","date_time_layout1":"2020-08-06 13:14:15","date_time_layout2":"2020-08-06 13:14:15","customer_format1":"2020-08-06T13:14:15+08:00","customer_format2":"2020-08-06T13:14:15+08:00","timestamp1":1596690855}`, string(data2))

		// delete
		db.Delete(&model2)
	})
}

func (s *PgSQLSuite) TestCurd2() {
	s.Run("unset carbon", func() {
		var model1 PgSQLModel2

		// create
		if err := db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 PgSQLModel2
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"date_carbon2":null,"time_layout1":null,"time_layout2":null,"date_time_layout1":null,"date_time_layout2":null,"customer_format1":null,"customer_format2":null,"timestamp1":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("nil carbon", func() {
		var model1 PgSQLModel2

		var c *carbon.Carbon
		c = nil

		model1.Carbon1 = c
		model1.Carbon2 = c

		model1.DateLayout1 = carbon.NewDate(c)

		model1.TimeLayout1 = carbon.NewTime(c)
		model1.TimeLayout2 = carbon.NewTime(c)

		model1.DateTimeLayout1 = carbon.NewDateTime(c)
		model1.DateTimeLayout2 = carbon.NewDateTime(c)

		model1.CustomerFormat1 = carbon.NewFormatType[ISO8601Type](c)
		model1.CustomerFormat2 = carbon.NewFormatType[ISO8601Type](c)

		model1.Timestamp1 = carbon.NewTimestamp(c)

		// create
		if err := db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 PgSQLModel2
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"date_carbon2":null,"time_layout1":null,"time_layout2":null,"date_time_layout1":null,"date_time_layout2":null,"customer_format1":null,"customer_format2":null,"timestamp1":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("zero carbon", func() {
		var model1 PgSQLModel2

		c := carbon.NewCarbon()

		model1.Carbon1 = c
		model1.Carbon2 = c

		model1.DateLayout1 = carbon.NewDate(c)

		model1.TimeLayout1 = carbon.NewTime(c)
		model1.TimeLayout2 = carbon.NewTime(c)

		model1.DateTimeLayout1 = carbon.NewDateTime(c)
		model1.DateTimeLayout2 = carbon.NewDateTime(c)

		model1.CustomerFormat1 = carbon.NewFormatType[ISO8601Type](c)
		model1.CustomerFormat2 = carbon.NewFormatType[ISO8601Type](c)

		model1.Timestamp1 = carbon.NewTimestamp(c)

		// create
		if err := db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 PgSQLModel2
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"date_carbon2":null,"time_layout1":null,"time_layout2":null,"date_time_layout1":null,"date_time_layout2":null,"customer_format1":null,"customer_format2":null,"timestamp1":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("empty carbon", func() {
		var model1 PgSQLModel2

		c := carbon.Parse("")

		model1.Carbon1 = c
		model1.Carbon2 = c

		model1.DateLayout1 = carbon.NewDate(c)

		model1.TimeLayout1 = carbon.NewTime(c)
		model1.TimeLayout2 = carbon.NewTime(c)

		model1.DateTimeLayout1 = carbon.NewDateTime(c)
		model1.DateTimeLayout2 = carbon.NewDateTime(c)

		model1.CustomerFormat1 = carbon.NewFormatType[ISO8601Type](c)
		model1.CustomerFormat2 = carbon.NewFormatType[ISO8601Type](c)

		model1.Timestamp1 = carbon.NewTimestamp(c)

		// create
		if err := db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 PgSQLModel2
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"date_carbon2":null,"time_layout1":null,"time_layout2":null,"date_time_layout1":null,"date_time_layout2":null,"customer_format1":null,"customer_format2":null,"timestamp1":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("valid carbon", func() {
		var model1 PgSQLModel2

		c := carbon.Now()

		model1.Carbon1 = c
		model1.Carbon2 = c

		model1.DateLayout1 = carbon.NewDate(c)

		model1.TimeLayout1 = carbon.NewTime(c)
		model1.TimeLayout2 = carbon.NewTime(c)

		model1.DateTimeLayout1 = carbon.NewDateTime(c)
		model1.DateTimeLayout2 = carbon.NewDateTime(c)

		model1.CustomerFormat1 = carbon.NewFormatType[ISO8601Type](c)
		model1.CustomerFormat2 = carbon.NewFormatType[ISO8601Type](c)

		model1.Timestamp1 = carbon.NewTimestamp(c)

		// create
		if err := db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 PgSQLModel2
		db.Last(&model2)
		//
		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":"2020-08-05 21:14:15","carbon2":"2020-08-05 13:14:15","date_carbon2":"2020-08-05","time_layout1":"13:14:15","time_layout2":"13:19:58","date_time_layout1":"2020-08-05 21:14:15","date_time_layout2":"2020-08-05 13:14:15","customer_format1":"2020-08-05T21:14:15+08:00","customer_format2":"2020-08-05T13:14:15+08:00","timestamp1":1596604455}`, string(data1))

		newCarbon := c.Copy().AddDay()

		model2.Carbon1 = newCarbon
		model2.Carbon2 = newCarbon

		model2.DateLayout1 = carbon.NewDate(newCarbon)

		model2.TimeLayout1 = carbon.NewTime(newCarbon)
		model2.TimeLayout2 = carbon.NewTime(newCarbon)

		model2.DateTimeLayout1 = carbon.NewDateTime(newCarbon)
		model2.DateTimeLayout2 = carbon.NewDateTime(newCarbon)

		model2.CustomerFormat1 = carbon.NewFormatType[ISO8601Type](newCarbon)
		model2.CustomerFormat2 = carbon.NewFormatType[ISO8601Type](newCarbon)

		model2.Timestamp1 = carbon.NewTimestamp(newCarbon)

		// update
		db.Save(&model2)

		data2, err2 := json.Marshal(&model2)
		s.Nil(err2)
		s.Equal(`{"carbon1":"2020-08-06 13:14:15","carbon2":"2020-08-06 13:14:15","date_carbon2":"2020-08-06","time_layout1":"13:14:15","time_layout2":"13:14:15","date_time_layout1":"2020-08-06 13:14:15","date_time_layout2":"2020-08-06 13:14:15","customer_format1":"2020-08-06T13:14:15+08:00","customer_format2":"2020-08-06T13:14:15+08:00","timestamp1":1596690855}`, string(data2))

		// delete
		db.Delete(&model2)
	})
}
