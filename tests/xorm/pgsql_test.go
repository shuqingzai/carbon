package xorm

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
	carbon.SetTestNow(carbon.Parse("2020-08-05 13:14:15.111111111"))
	db = connect(driverPgSQL)
	if err = db.Sync(&PgSQLModel1{}); err != nil {
		panic(err)
	}
	if err = db.Sync(&PgSQLModel2{}); err != nil {
		panic(err)
	}
}

func (s *PgSQLSuite) TearDownSuite() {
	carbon.ClearTestNow()
	//if _, err = db.Unscoped().Where("1 = 1").Delete(PgSQLModel1{}); err != nil {
	//	panic(err)
	//}
	//if _, err = db.Unscoped().Where("1 = 1").Delete(PgSQLModel2{}); err != nil {
	//	panic(err)
	//}
}

func (s *PgSQLSuite) TestCurd1() {
	s.Run("unset carbon", func() {
		var model1 PgSQLModel1

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 PgSQLModel1
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"date1":null,"time1":null,"time2":null,"date_time1":null,"date_time2":null,"rfc3339_layout1":null,"rfc3339_layout2":null,"iso8601_format1":null,"iso8601_format2":null,"timestamp1":null,"timestamp2":null}`, string(data1))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})

	s.Run("nil carbon", func() {
		var model1 PgSQLModel1

		var c *carbon.Carbon
		c = nil

		model1.Date1 = *carbon.NewDate(c)

		model1.Time1 = *carbon.NewTime(c)
		model1.Time2 = *carbon.NewTime(c)

		model1.DateTime1 = *carbon.NewDateTime(c)
		model1.DateTime2 = *carbon.NewDateTime(c)

		model1.RFC3339Layout1 = *carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout2 = *carbon.NewLayoutType[RFC3339Layout](c)

		model1.ISO8601Format1 = *carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format2 = *carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp1 = *carbon.NewTimestamp(c)
		model1.Timestamp2 = *carbon.NewTimestamp(c)

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 PgSQLModel1
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"date1":null,"time1":null,"time2":null,"date_time1":null,"date_time2":null,"rfc3339_layout1":null,"rfc3339_layout2":null,"iso8601_format1":null,"iso8601_format2":null,"timestamp1":null,"timestamp2":null}`, string(data1))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})

	s.Run("zero carbon", func() {
		var model1 PgSQLModel1

		c := carbon.NewCarbon()

		model1.Carbon1 = *c
		model1.Carbon2 = *c

		model1.Date1 = *carbon.NewDate(c)

		model1.Time1 = *carbon.NewTime(c)
		model1.Time2 = *carbon.NewTime(c)

		model1.DateTime1 = *carbon.NewDateTime(c)
		model1.DateTime2 = *carbon.NewDateTime(c)

		model1.RFC3339Layout1 = *carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout2 = *carbon.NewLayoutType[RFC3339Layout](c)

		model1.ISO8601Format1 = *carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format2 = *carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp1 = *carbon.NewTimestamp(c)
		model1.Timestamp2 = *carbon.NewTimestamp(c)

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 PgSQLModel1
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"date1":null,"time1":null,"time2":null,"date_time1":null,"date_time2":null,"rfc3339_layout1":null,"rfc3339_layout2":null,"iso8601_format1":null,"iso8601_format2":null,"timestamp1":null,"timestamp2":null}`, string(data1))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})

	s.Run("empty carbon", func() {
		var model1 PgSQLModel1

		c := carbon.Parse("")

		model1.Carbon1 = *c
		model1.Carbon2 = *c

		model1.Date1 = *carbon.NewDate(c)

		model1.Time1 = *carbon.NewTime(c)
		model1.Time2 = *carbon.NewTime(c)

		model1.DateTime1 = *carbon.NewDateTime(c)
		model1.DateTime2 = *carbon.NewDateTime(c)

		model1.RFC3339Layout1 = *carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout2 = *carbon.NewLayoutType[RFC3339Layout](c)

		model1.ISO8601Format1 = *carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format2 = *carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp1 = *carbon.NewTimestamp(c)
		model1.Timestamp2 = *carbon.NewTimestamp(c)

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 PgSQLModel1
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"date1":null,"time1":null,"time2":null,"date_time1":null,"date_time2":null,"rfc3339_layout1":null,"rfc3339_layout2":null,"iso8601_format1":null,"iso8601_format2":null,"timestamp1":null,"timestamp2":null}`, string(data1))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})

	s.Run("valid carbon", func() {
		var model1 PgSQLModel1

		c := carbon.Now()

		model1.Carbon1 = *c
		model1.Carbon2 = *c

		model1.Date1 = *carbon.NewDate(c)

		model1.Time1 = *carbon.NewTime(c)
		model1.Time2 = *carbon.NewTime(c)

		model1.DateTime1 = *carbon.NewDateTime(c)
		model1.DateTime2 = *carbon.NewDateTime(c)

		model1.RFC3339Layout1 = *carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout2 = *carbon.NewLayoutType[RFC3339Layout](c)

		model1.ISO8601Format1 = *carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format2 = *carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp1 = *carbon.NewTimestamp(c)
		model1.Timestamp2 = *carbon.NewTimestamp(c)

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 PgSQLModel1
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":"2020-08-05 13:14:15","carbon2":"2020-08-05 13:14:15","date1":"2020-08-05","time1":"13:14:15","time2":"13:14:15","date_time1":"2020-08-05 13:14:15","date_time2":"2020-08-05 13:14:15","rfc3339_layout1":"2020-08-05T13:14:15Z","rfc3339_layout2":"2020-08-05T13:14:15+08:00","iso8601_format1":"2020-08-05T13:14:15+00:00","iso8601_format2":"2020-08-05T13:14:15+08:00","timestamp1":1596633255,"timestamp2":1596604455}`, string(data1))

		c = c.Copy().AddDay()

		model2.Carbon1 = *c
		model2.Carbon2 = *c

		model2.Date1 = *carbon.NewDate(c)

		model2.Time1 = *carbon.NewTime(c)
		model2.Time2 = *carbon.NewTime(c)

		model2.DateTime1 = *carbon.NewDateTime(c)
		model2.DateTime2 = *carbon.NewDateTime(c)

		model2.RFC3339Layout1 = *carbon.NewLayoutType[RFC3339Layout](c)
		model2.RFC3339Layout2 = *carbon.NewLayoutType[RFC3339Layout](c)

		model2.ISO8601Format1 = *carbon.NewFormatType[ISO8601Format](c)
		model2.ISO8601Format2 = *carbon.NewFormatType[ISO8601Format](c)

		model2.Timestamp1 = *carbon.NewTimestamp(c)
		model2.Timestamp2 = *carbon.NewTimestamp(c)

		// update
		_, err = db.Update(model2)
		s.Nil(err)

		data2, err2 := json.Marshal(&model2)
		s.Nil(err2)
		s.Equal(`{"carbon1":"2020-08-06 13:14:15","carbon2":"2020-08-06 13:14:15","date1":"2020-08-06","time1":"13:14:15","time2":"13:14:15","date_time1":"2020-08-06 13:14:15","date_time2":"2020-08-06 13:14:15","rfc3339_layout1":"2020-08-06T13:14:15+08:00","rfc3339_layout2":"2020-08-06T13:14:15+08:00","iso8601_format1":"2020-08-06T13:14:15+08:00","iso8601_format2":"2020-08-06T13:14:15+08:00","timestamp1":1596690855,"timestamp2":1596690855}`, string(data2))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})
}

func (s *PgSQLSuite) TestCurd2() {
	s.Run("unset carbon", func() {
		var model1 PgSQLModel2

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 PgSQLModel2
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"date1":null,"time1":null,"time2":null,"date_time1":null,"date_time2":null,"rfc3339_layout1":null,"rfc3339_layout2":null,"iso8601_format1":null,"iso8601_format2":null,"timestamp1":null,"timestamp2":null}`, string(data1))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})

	s.Run("nil carbon", func() {
		var model1 PgSQLModel2

		var c *carbon.Carbon
		c = nil

		model1.Carbon1 = c
		model1.Carbon2 = c

		model1.Date1 = carbon.NewDate(c)

		model1.Time1 = carbon.NewTime(c)
		model1.Time2 = carbon.NewTime(c)

		model1.DateTime1 = carbon.NewDateTime(c)
		model1.DateTime2 = carbon.NewDateTime(c)

		model1.RFC3339Layout1 = carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout2 = carbon.NewLayoutType[RFC3339Layout](c)

		model1.ISO8601Format1 = carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format2 = carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp1 = carbon.NewTimestamp(c)
		model1.Timestamp2 = carbon.NewTimestamp(c)

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 PgSQLModel2
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"date1":null,"time1":null,"time2":null,"date_time1":null,"date_time2":null,"rfc3339_layout1":null,"rfc3339_layout2":null,"iso8601_format1":null,"iso8601_format2":null,"timestamp1":null,"timestamp2":null}`, string(data1))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})

	s.Run("zero carbon", func() {
		var model1 PgSQLModel2

		c := carbon.NewCarbon()

		model1.Carbon1 = c
		model1.Carbon2 = c

		model1.Date1 = carbon.NewDate(c)

		model1.Time1 = carbon.NewTime(c)
		model1.Time2 = carbon.NewTime(c)

		model1.DateTime1 = carbon.NewDateTime(c)
		model1.DateTime2 = carbon.NewDateTime(c)

		model1.RFC3339Layout1 = carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout2 = carbon.NewLayoutType[RFC3339Layout](c)

		model1.ISO8601Format1 = carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format2 = carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp1 = carbon.NewTimestamp(c)
		model1.Timestamp2 = carbon.NewTimestamp(c)

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 PgSQLModel2
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"date1":null,"time1":null,"time2":null,"date_time1":null,"date_time2":null,"rfc3339_layout1":null,"rfc3339_layout2":null,"iso8601_format1":null,"iso8601_format2":null,"timestamp1":null,"timestamp2":null}`, string(data1))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})

	s.Run("empty carbon", func() {
		var model1 PgSQLModel2

		c := carbon.Parse("")

		model1.Carbon1 = c
		model1.Carbon2 = c

		model1.Date1 = carbon.NewDate(c)

		model1.Time1 = carbon.NewTime(c)
		model1.Time2 = carbon.NewTime(c)

		model1.DateTime1 = carbon.NewDateTime(c)
		model1.DateTime2 = carbon.NewDateTime(c)

		model1.RFC3339Layout1 = carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout2 = carbon.NewLayoutType[RFC3339Layout](c)

		model1.ISO8601Format1 = carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format2 = carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp1 = carbon.NewTimestamp(c)
		model1.Timestamp2 = carbon.NewTimestamp(c)

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 PgSQLModel2
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"date1":null,"time1":null,"time2":null,"date_time1":null,"date_time2":null,"rfc3339_layout1":null,"rfc3339_layout2":null,"iso8601_format1":null,"iso8601_format2":null,"timestamp1":null,"timestamp2":null}`, string(data1))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})

	s.Run("valid carbon", func() {
		var model1 PgSQLModel2

		c := carbon.Now()

		model1.Carbon1 = c
		model1.Carbon2 = c

		model1.Date1 = carbon.NewDate(c)

		model1.Time1 = carbon.NewTime(c)
		model1.Time2 = carbon.NewTime(c)

		model1.DateTime1 = carbon.NewDateTime(c)
		model1.DateTime2 = carbon.NewDateTime(c)

		model1.RFC3339Layout1 = carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout2 = carbon.NewLayoutType[RFC3339Layout](c)

		model1.ISO8601Format1 = carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format2 = carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp1 = carbon.NewTimestamp(c)
		model1.Timestamp2 = carbon.NewTimestamp(c)

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 PgSQLModel2
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":"2020-08-05 13:14:15","carbon2":"2020-08-05 13:14:15","date1":"2020-08-05","time1":"13:14:15","time2":"13:14:15","date_time1":"2020-08-05 13:14:15","date_time2":"2020-08-05 13:14:15","rfc3339_layout1":"2020-08-05T13:14:15Z","rfc3339_layout2":"2020-08-05T13:14:15+08:00","iso8601_format1":"2020-08-05T13:14:15+00:00","iso8601_format2":"2020-08-05T13:14:15+08:00","timestamp1":1596633255,"timestamp2":1596604455}`, string(data1))

		c = c.Copy().AddDay()

		model2.Carbon1 = c
		model2.Carbon2 = c

		model2.Date1 = carbon.NewDate(c)

		model2.Time1 = carbon.NewTime(c)
		model2.Time2 = carbon.NewTime(c)

		model2.DateTime1 = carbon.NewDateTime(c)
		model2.DateTime2 = carbon.NewDateTime(c)

		model2.RFC3339Layout1 = carbon.NewLayoutType[RFC3339Layout](c)
		model2.RFC3339Layout2 = carbon.NewLayoutType[RFC3339Layout](c)

		model2.ISO8601Format1 = carbon.NewFormatType[ISO8601Format](c)
		model2.ISO8601Format2 = carbon.NewFormatType[ISO8601Format](c)

		model2.Timestamp1 = carbon.NewTimestamp(c)
		model2.Timestamp2 = carbon.NewTimestamp(c)

		// update
		_, err = db.Update(model2)
		s.Nil(err)

		data2, err2 := json.Marshal(&model2)
		s.Nil(err2)
		s.Equal(`{"carbon1":"2020-08-06 13:14:15","carbon2":"2020-08-06 13:14:15","date1":"2020-08-06","time1":"13:14:15","time2":"13:14:15","date_time1":"2020-08-06 13:14:15","date_time2":"2020-08-06 13:14:15","rfc3339_layout1":"2020-08-06T13:14:15+08:00","rfc3339_layout2":"2020-08-06T13:14:15+08:00","iso8601_format1":"2020-08-06T13:14:15+08:00","iso8601_format2":"2020-08-06T13:14:15+08:00","timestamp1":1596690855,"timestamp2":1596690855}`, string(data2))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})
}
