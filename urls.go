package entities

// TODO URL regex checking
// TODO URL existence checking (whoa!)


// ^(ht|f)tp(s?)\:\/\/[0-9a-zA-Z]([-.\w]*[0-9a-zA-Z])*(:(0-9)*)*(\/?)([a-zA-Z0-9\-\.\?\,\'\/\\\+&amp;%\$#_]*)?$
import (
	"fmt"
)
	
var _ = fmt.Println // remove after test

// *** Entities and their Methods ***

type URL struct {
	Key *Key
	Name string
	URLText string
}

func NewURL(n string, u string) *URL {
	p := new(URL)
	p.Key = makeKey("URL")
	p.Name = n
	p.URLText = u
	return p
}

func (a *URL) Triples () [][3]string {
	var t [][3]string
	t = append(t, makeTriple(Triple{(*a).Key,"hasType","URL",nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasName",(*a).Name,nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasURLText",(*a).URLText,nil}))
	return t
}

func FindURLKey (kf *Key) int {
	for i,a := range URLs {
		if kf.s == a.Key.s {
			return i
		}
	}
	return -1
}

func AddURLFact(a []string) {
	key := new(Key)
	key.s = a[0]
	i := FindURLKey(key)
	switch a[1] {
	case "Name":
		URLs[i].Name = a[2]
	case "URLText":
		URLs[i].URLText = a[2]
	}
}

var URLs []*URL