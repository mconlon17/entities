package entities

import (
	"errors"
	"fmt"
)
	
var _ = fmt.Println // remove after test
	
func makeCourseTypeSet() map[string]struct{} {
	a := map[string]struct{}{}
	a["ShortCourse"] = struct{}{}
	a["AcademicCourse"] = struct{}{}
	a["Seminar"] = struct{}{}
	a["ThesisAdvising"] = struct{}{}
	return a
}

var CourseTypeSet = makeCourseTypeSet()

type Course struct {
	Key *Key
	CourseType string
	CourseNumber string
	CourseTitle string
	CourseDescription string
	CourseOrganization *Organization
}

func NewCourse(nt string, nu string, nm string) (*Course,error) {
	p := new(Course)
	p.Key = makeKey("Course")
	_, ok := CourseTypeSet[nt]
	if ok {
		p.CourseType = nt
		p.CourseNumber = nu
		p.CourseTitle = nm
		return p, nil
	} else {
		err := errors.New("Unknown course type: " + nt)
		return nil, err
	}
}

func NewAcademicCourse(nu string, nm string)  (*Course,error) {
	p,e := NewCourse("AcademicCourse",nu,nm)
	return p,e
}

func NewShortCourse(nu string, nm string)  (*Course,error) {
	p,e := NewCourse("ShortCourse",nu,nm)
	return p,e
}

func NewSeminar(nu string, nm string)  (*Course,error) {
	p,e := NewCourse("Seminar",nu,nm)
	return p,e
}

func NewThesisAdvising(nu string, nm string)  (*Course,error) {
	p,e := NewCourse("ThesisAdvising",nu,nm)
	return p,e
}

func (a *Course) Triples () [][3]string {
	var t [][3]string
	t = append(t, makeTriple(Triple{(*a).Key,"hasType","Course",nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasCourseType",(*a).CourseType,nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasCourseNumber",(*a).CourseNumber,nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasCourseTitle",(*a).CourseTitle,nil}))
	if (*a).CourseDescription != "" {t = append(t, makeTriple(Triple{(*a).Key,"hasCourseDescription",(*a).CourseDescription,nil}))}
	if (*a).CourseOrganization != nil { t=append(t, makeTriple(Triple{(*a).Key,"hasCourseOrganization","",(*a.CourseOrganization).Key}))}
	return t
}

func FindCourseKey (kf *Key) int {
	for i,a := range Courses {
		if kf.s == a.Key.s {
			return i
		}
	}
	return -1
}

func AddCourseFact(a []string) {
	key := new(Key)
	key.s = a[0]
	i := FindCourseKey(key)
	switch a[1] {
	case "CourseType":
		Courses[i].CourseType = a[2]
	case "CourseNumber":
		Courses[i].CourseNumber = a[2]
	case "CourseTitle":
		Courses[i].CourseTitle = a[2]
	case "CourseDescription":
		Courses[i].CourseDescription = a[2]
	case "CourseOrganization":
		key.s = a[2]
		j := FindOrganizationKey(key)
		Courses[i].CourseOrganization = Organizations[j]
	}
}


var Courses []*Course