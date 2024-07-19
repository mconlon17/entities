package entities

import (
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	"os"
	"strconv"
)
	
var _ = fmt.Println // remove after test	

type Image struct {
	Key *Key
	FileName string
	Caption string
	AltText string
	height int
	width int
}

func NewImage(n string) (*Image,error) {
	p := new(Image)
	p.Key = makeKey("Image")

	if reader, err := os.Open(n); err == nil {
		defer reader.Close()
		im, _, err := image.DecodeConfig(reader)
		if err != nil {
			return nil, err
		}
		p.width = im.Width
		p.height = im.Height
		p.FileName = n
		return p, nil
	} else {
		err := errors.New("Can't open " + n)
		return nil, err
	}
}

func (a *Image) getHeight() int {
	return a.height
}

func (a *Image) getWidth() int {
	return a.width
}

func (a *Image) Triples () [][3]string {
	var t [][3]string
	t = append(t, makeTriple(Triple{(*a).Key,"hasType","Image",nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasFileName",(*a).FileName,nil}))
	if (*a).Caption != "" {t = append(t, makeTriple(Triple{(*a).Key,"hasCaption",(*a).Caption,nil}))}
	if (*a).AltText != "" {t = append(t, makeTriple(Triple{(*a).Key,"hasAltText",(*a).AltText,nil}))}
	if (*a).height != 0 {t = append(t, makeTriple(Triple{(*a).Key,"hasHeight",strconv.Itoa((*a).height),nil}))}
	if (*a).width != 0 {t = append(t, makeTriple(Triple{(*a).Key,"hasWidth",strconv.Itoa((*a).width),nil}))}
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
		Images[i].height = n
	case "Width":
		n,_ := strconv.Atoi(a[2])
		Images[i].width = n
	}
}

var Images []*Image