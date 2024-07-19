package entities

import (
	"errors"
	"fmt"
	"strconv"
)
	
var _ = fmt.Println // remove after test
	

type GeoLocation struct {
	Key *Key
	latitude float64
	longitude float64
}

func NewGeoLocation(lat float64, long float64) (*GeoLocation, error) {
	p := new(GeoLocation)
	p.Key = makeKey("GeoLocation")
	if lat >= -90 && lat <= 90 && long >= -180 && long <= 180 {
		p.latitude = lat
		p.longitude = long
		return p,nil
	} else {
		err := errors.New("Invalid geolocation " + strconv.FormatFloat(lat, 'f', -1, 64) + " " + strconv.FormatFloat(long, 'f', -1, 64))
		return nil,err
	}
}

func (a *GeoLocation) setLatitude(n float64) {
	if n >= -90 && n <= 90 {
		a.latitude = n
	} else {
		err := errors.New("Invalid latitude: " + strconv.FormatFloat(n, 'f', -1, 64))
		fmt.Println(err)
	}
}

func (a *GeoLocation) setLongitude(n float64) {
	if n >= -180 && n <= 180 {
		a.longitude = n
	} else {
		err := errors.New("Invalid longitude: " + strconv.FormatFloat(n, 'f', -1, 64))
		fmt.Println(err)
	}
}

func (a *GeoLocation) getLatitude() float64 {
	return a.latitude
}

func (a *GeoLocation) getLongitude() float64 {
	return a.longitude
}


func (a *GeoLocation) Triples () [][3]string {
	var t [][3]string
	t = append(t, makeTriple(Triple{(*a).Key,"hasType","GeoLocation",nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasLatitude",strconv.FormatFloat((*a).latitude, 'f', -1, 64),nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasLongitude",strconv.FormatFloat((*a).longitude, 'f', -1, 64),nil}))
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
		GeoLocations[i].setLatitude(f)
	case "Longitude":
		f,_ := strconv.ParseFloat(a[2], 64)
		GeoLocations[i].setLongitude(f)
	}
}

var GeoLocations []*GeoLocation