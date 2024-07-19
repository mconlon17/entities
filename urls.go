package entities

// TODO URL existence checking
 
import (
	"errors"
	"fmt"
	"regexp"
)
	
var _ = fmt.Println // remove after test
var r7 = regexp.MustCompile(`^(?P<URL>[A-Za-z0-9\:\/\/\.\-]*)$`)

// *** Entities and their Methods ***

type URL struct {
	Key *Key
	Name string
	URLText string
}

func NewURL(n string, url string) (*URL,error) {
	p := new(URL)
	p.Key = makeKey("URL")
	p.Name = n
	u := r7.FindStringSubmatch(url)
	fmt.Println(url,u)
	if len(u) == 2 {
		p.URLText = url
		return p,nil
	} else {
		err := errors.New("Invalid URL: " + url)
		return nil,err
	}
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