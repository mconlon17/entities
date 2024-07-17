package entities

import (
	"errors"
	"fmt"
	"strconv"
)
	
var _ = fmt.Println // remove after test
	

type GeoLocation struct {
	Key *Key
	Latitude float64
	Longitude float64
}

func NewGeoLocation(lat float64, long float64) (*GeoLocation, error) {
	err := errors.New("")
	p := new(GeoLocation)
	p.Key = makeKey("GeoLocation")
	if lat >= -90 && lat <= 90 {
		p.Latitude = lat
	} else {
		err = errors.New("Invalid Latitidue " + strconv.FormatFloat(lat, 'f', -1, 64))
	}
	if long >= -180 && long <= 180 {
		p.Longitude = long
	} else {
		err = errors.New("Invalid Longitude " + strconv.FormatFloat(long, 'f', -1, 64))
	}
	return p, err
}

func (a *GeoLocation) Triples () [][3]string {
	var t [][3]string
	t = append(t, makeTriple(Triple{(*a).Key,"hasType","GeoLocation",nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasLatitude",strconv.FormatFloat((*a).Latitude, 'f', -1, 64),nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasLongitude",strconv.FormatFloat((*a).Longitude, 'f', -1, 64),nil}))
	return t
}

func FindGeoLocationKey (kf *Key) int {
	for i,a := range GeoLocations {
		if kf.s == a.Key.s {
			return i
		}
	}
	return -1
}

func AddGeoLocationFact(a []string) {
	key := new(Key)
	key.s = a[0]
	i := FindGeoLocationKey(key)
	switch a[1] {
	case "Latitude":
		f,_ := strconv.ParseFloat(a[2], 64)
		GeoLocations[i].Latitude = f
	case "Longitude":
		f,_ := strconv.ParseFloat(a[2], 64)
		GeoLocations[i].Longitude = f
	}
}


var GeoLocations []*GeoLocation