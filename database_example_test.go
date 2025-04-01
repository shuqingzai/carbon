package carbon_test

import (
	"encoding/json"
	"fmt"

	"github.com/dromara/carbon/v2"
)

type W3cLayout string

func (W3cLayout) SetLayout() string {
	return carbon.W3cLayout
}

func ExampleLayoutType_MarshalJSON() {
	//type W3cLayout string
	//
	//func (W3cLayout) SetLayout() string {
	//	return carbon.W3cLayout
	//}

	type User struct {
		Date     carbon.LayoutType[carbon.Date]     `json:"date"`
		Time     carbon.LayoutType[carbon.Time]     `json:"time"`
		DateTime carbon.LayoutType[carbon.DateTime] `json:"date_time"`
		Customer carbon.LayoutType[W3cLayout]       `json:"customer"`
	}

	var user User

	c := carbon.Parse("2020-08-05 13:14:15")

	user.Date = carbon.NewLayoutType[carbon.Date](c)
	user.Time = carbon.NewLayoutType[carbon.Time](c)
	user.DateTime = carbon.NewLayoutType[carbon.DateTime](c)
	user.Customer = carbon.NewLayoutType[W3cLayout](c)

	data, _ := json.Marshal(&user)
	fmt.Printf("%s", data)

	// Output:
	// {"date":"2020-08-05","time":"13:14:15","date_time":"2020-08-05 13:14:15","customer":"2020-08-05T13:14:15Z"}
}

func ExampleLayoutType_UnmarshalJSON() {
	//type W3cLayout string
	//
	//func (W3cLayout) SetLayout() string {
	//	return carbon.W3cLayout
	//}
	type User struct {
		Date     carbon.LayoutType[carbon.Date]     `json:"date"`
		Time     carbon.LayoutType[carbon.Time]     `json:"time"`
		DateTime carbon.LayoutType[carbon.DateTime] `json:"date_time"`
		Customer carbon.LayoutType[W3cLayout]       `json:"customer"`
	}

	var user User

	value := `{"date":"2020-08-05","time":"13:14:15","date_time":"2020-08-05 13:14:15","customer":"2020-08-05T13:14:15Z"}`
	_ = json.Unmarshal([]byte(value), &user)

	fmt.Println(user.Date.String())
	fmt.Println(user.Time.String())
	fmt.Println(user.DateTime.String())
	fmt.Println(user.Customer.String())

	// Output:
	// 2020-08-05
	// 13:14:15
	// 2020-08-05 13:14:15
	// 2020-08-05T13:14:15Z
}

type RFC3339Format string

func (RFC3339Format) SetFormat() string {
	return carbon.RFC3339Format
}

func ExampleFormatType_MarshalJSON() {
	//type RFC3339Format string
	//
	//func (RFC3339Format) SetFormat() string {
	//	return carbon.RFC3339Format
	//}
	type User struct {
		Date     carbon.FormatType[carbon.Date]     `json:"date"`
		Time     carbon.FormatType[carbon.Time]     `json:"time"`
		DateTime carbon.FormatType[carbon.DateTime] `json:"date_time"`
		Customer carbon.FormatType[RFC3339Format]   `json:"customer"`
	}

	var user User

	c := carbon.Parse("2020-08-05 13:14:15")

	user.Date = carbon.NewFormatType[carbon.Date](c)
	user.Time = carbon.NewFormatType[carbon.Time](c)
	user.DateTime = carbon.NewFormatType[carbon.DateTime](c)
	user.Customer = carbon.NewFormatType[RFC3339Format](c)

	data, _ := json.Marshal(&user)
	fmt.Printf("%s", data)

	// Output:
	// {"date":"2020-08-05","time":"13:14:15","date_time":"2020-08-05 13:14:15","customer":"2020-08-05T13:14:15Z"}
}

func ExampleFormatType_UnmarshalJSON() {
	//type RFC3339Format string
	//
	//func (RFC3339Format) SetFormat() string {
	//	return carbon.RFC3339Format
	//}

	type User struct {
		Date     carbon.FormatType[carbon.Date]     `json:"date"`
		Time     carbon.FormatType[carbon.Time]     `json:"time"`
		DateTime carbon.FormatType[carbon.DateTime] `json:"date_time"`
		Customer carbon.FormatType[RFC3339Format]   `json:"customer"`
	}

	var user User

	value := `{"date":"2020-08-05","time":"13:14:15","date_time":"2020-08-05 13:14:15","customer":"2020-08-05T13:14:15Z"}`
	_ = json.Unmarshal([]byte(value), &user)

	fmt.Println(user.Date.String())
	fmt.Println(user.Time.String())
	fmt.Println(user.DateTime.String())
	fmt.Println(user.Customer.String())

	// Output:
	// 2020-08-05
	// 13:14:15
	// 2020-08-05 13:14:15
	// 2020-08-05T13:14:15Z
}

func ExampleTimestampType_MarshalJSON() {
	type User struct {
		Timestamp      carbon.TimestampType[carbon.Timestamp]      `json:"timestamp"`
		TimestampMilli carbon.TimestampType[carbon.TimestampMilli] `json:"timestamp_milli"`
		TimestampMicro carbon.TimestampType[carbon.TimestampMicro] `json:"timestamp_micro"`
		TimestampNano  carbon.TimestampType[carbon.TimestampNano]  `json:"timestamp_nano"`
	}

	var user User

	c := carbon.Parse("2020-08-05 13:14:15")

	user.Timestamp = carbon.NewTimestampType[carbon.Timestamp](c)
	user.TimestampMilli = carbon.NewTimestampType[carbon.TimestampMilli](c)
	user.TimestampMicro = carbon.NewTimestampType[carbon.TimestampMicro](c)
	user.TimestampNano = carbon.NewTimestampType[carbon.TimestampNano](c)

	data, _ := json.Marshal(&user)
	fmt.Printf("%s", data)

	// Output:
	// {"timestamp":1596633255,"timestamp_milli":1596633255000,"timestamp_micro":1596633255000000,"timestamp_nano":1596633255000000000}
}

func ExampleTimestampType_UnmarshalJSON() {
	type User struct {
		Timestamp      carbon.TimestampType[carbon.Timestamp]      `json:"timestamp"`
		TimestampMilli carbon.TimestampType[carbon.TimestampMilli] `json:"timestamp_milli"`
		TimestampMicro carbon.TimestampType[carbon.TimestampMicro] `json:"timestamp_micro"`
		TimestampNano  carbon.TimestampType[carbon.TimestampNano]  `json:"timestamp_nano"`
	}

	var user User

	value := `{"timestamp":1596633255,"timestamp_milli":1596633255000,"timestamp_micro":1596633255000000,"timestamp_nano":1596633255000000000}`
	_ = json.Unmarshal([]byte(value), &user)

	fmt.Printf("string(%s)\n", user.Timestamp.String())
	fmt.Printf("string(%s)\n", user.TimestampMilli.String())
	fmt.Printf("string(%s)\n", user.TimestampMicro.String())
	fmt.Printf("string(%s)\n", user.TimestampNano.String())

	fmt.Printf("int64(%d)\n", user.Timestamp.Int64())
	fmt.Printf("int64(%d)\n", user.TimestampMilli.Int64())
	fmt.Printf("int64(%d)\n", user.TimestampMicro.Int64())
	fmt.Printf("int64(%d)\n", user.TimestampNano.Int64())

	// Output:
	// string(1596633255)
	// string(1596633255000)
	// string(1596633255000000)
	// string(1596633255000000000)
	// int64(1596633255)
	// int64(1596633255000)
	// int64(1596633255000000)
	// int64(1596633255000000000)
}
