package entities

import (
	"fmt"
	"time"
)
	
var _ = fmt.Println // remove after test
	
func makeDateTypeSet() map[string]struct{} {
	a := map[string]struct{}{}
	a["Year"] = struct{}{}                   // Constructor Year
	a["YearMonth"] = struct{}{}              // Constructor YearMonth
	a["YearMonthDay"] = struct{}{}           // Constructor Date
	a["YearMonthDayHourMinute"] = struct{}{} // Constructor DateTime
	return a
}

var DateTypeSet = makeDateTypeSet()


type Date struct {
	Key *Key
	DateType string
	UTC time.Time
}

func NewYear(d string) (*Date, error) {
	p := new(Date)
	p.Key = makeKey("Date")
	p.DateType = "Year"
	date, err := time.Parse("2006",d)
	if err != nil {
		return nil, err
	} else {
		p.UTC = date
		return p, nil
	}
}

func NewYearMonth(d string) (*Date, error) {
	p := new(Date)
	p.Key = makeKey("Date")
	p.DateType = "YearMonth"
	date, err := time.Parse("2006-01",d)
	if err != nil {
		return nil, err
	} else {
		p.UTC = date
		return p, nil
	}
}

func NewDate(d string)(*Date, error) {
	p := new(Date)
	p.Key = makeKey("Date")
	p.DateType = "YearMonthDay"
	date, err := time.Parse(time.DateOnly,d)
	if err != nil {
		return nil, err
	} else {
		p.UTC = date
		return p, nil
	}
}

func NewDateTime(d string)(*Date, error) {
	p := new(Date)
	p.Key = makeKey("Date")
	p.DateType = "YearMonthDayHourMinute"
	date, err := time.Parse(time.DateTime,d) // yyyy-mm-dd hh:mm:ss
	if err != nil {
		return nil, err
	} else {
		p.UTC = date
		return p, nil
	}
}

func (a *Date) Triples () [][3]string {
	var t [][3]string
	t = append(t, makeTriple(Triple{(*a).Key,"hasType","Date",nil}))
	if (*a).DateType != "" {t = append(t, makeTriple(Triple{(*a).Key,"hasDateType",(*a).DateType,nil}))}
	t = append(t, makeTriple(Triple{(*a).Key,"hasUTC",((*a).UTC).Format(time.RFC3339),nil}))
	return t
}

func (a *Date) Row() []string {
	var t []string
	t = append(t, a.Key.s)
	t = append(t, a.DateType)
	t = append(t, (a.UTC).Format(time.RFC3339))
	return t
}

func FindDateKey (kf *Key) int {
	for i,a := range Dates {
		if kf.s == a.Key.s {
			return i
		}
	}
	return -1
}

func AddDateFact(a []string) {
	key := new(Key)
	key.s = a[0]
	i := FindDateKey(key)
	switch a[1] {
	case "DateType":
		Dates[i].DateType = a[2]
	case "UTC":
		date,_ := time.Parse(time.RFC3339,a[2])
		Dates[i].UTC = date
	}
}

var Dates []*Date