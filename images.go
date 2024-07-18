package entities

// TODO: Calculate dimensions from image
// TODO: Hide height and width. Create getters

import (
	"fmt"
	"strconv"
)
	
var _ = fmt.Println // remove after test	

type Image struct {
	Key *Key
	FileName string
	Caption string
	AltText string
	Height int
	Width int
}

func NewImage(n string) *Image {
	p := new(Image)
	p.Key = makeKey("Image")
	p.FileName = n
	return p
}

func (a *Image) Triples () [][3]string {
	var t [][3]string
	t = append(t, makeTriple(Triple{(*a).Key,"hasType","Image",nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasFileName",(*a).FileName,nil}))
	if (*a).Caption != "" {t = append(t, makeTriple(Triple{(*a).Key,"hasCaption",(*a).Caption,nil}))}
	if (*a).AltText != "" {t = append(t, makeTriple(Triple{(*a).Key,"hasAltText",(*a).AltText,nil}))}
	if (*a).Height != 0 {t = append(t, makeTriple(Triple{(*a).Key,"hasHeight",strconv.Itoa((*a).Height),nil}))}
	if (*a).Width != 0 {t = append(t, makeTriple(Triple{(*a).Key,"hasWidth",strconv.Itoa((*a).Width),nil}))}
	return t
}

func FindImageKey (kf *Key) int {
	for i,a := range Images {
		if kf.s == a.Key.s {
			return i
		}
	}
	return -1
}

func AddImageFact(a []string) {
	key := new(Key)
	key.s = a[0]
	i := FindImageKey(key)
	switch a[1] {
	case "FileName":
		Images[i].FileName = a[2]
	case "Caption":
		Images[i].Caption = a[2]
	case "AltText":
		Images[i].AltText = a[2]
	case "Height":
		n,_ := strconv.Atoi(a[2])
		Images[i].Height = n
	case "Width":
		n,_ := strconv.Atoi(a[2])
		Images[i].Width = n
	}
}

var Images []*Image