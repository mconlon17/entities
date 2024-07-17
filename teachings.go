package entities

import (
	"fmt"
)
	
var _ = fmt.Println // remove after test

type Teaching struct {
	Key *Key
	Description string
	Course *Course
	Teacher *Person
	StartDate *Date
	EndDate *Date
	Students []*Person
}

func NewTeaching(c *Course, pe *Person) *Teaching {
	p := new(Teaching)
	p.Key = makeKey("Teaching")
	p.Course = c
	p.Teacher = pe
	return p
}

func (a *Teaching) Triples () [][3]string {
	var t [][3]string
	t = append(t, makeTriple(Triple{(*a).Key,"hasType","Teaching",nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasDescription",(*a).Description,nil}))
	if (*a).Course != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasCourse","",(*a.Course).Key}))}
	if (*a).Teacher != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasTeacher","",(*a.Teacher).Key}))}
    if (*a).StartDate != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasStartDate","",(*a.StartDate).Key}))}
	if (*a).EndDate != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasEndDate","",(*a.EndDate).Key}))}
	for _,cptr := range (*a).Students {
		t = append(t, makeTriple(Triple{(*a).Key,"hasStudent","",cptr.Key}))
	}
	return t
}

func FindTeachingKey (kf *Key) int {
	for i,a := range Teachings {
		if kf.s == a.Key.s {
			return i
		}
	}
	return -1
}

func AddTeachingFact(a []string) {
	key := new(Key)
	key.s = a[0]
	i := FindTeachingKey(key)
	switch a[1] {
	case "Description":
		Teachings[i].Description = a[2]
	case "Course":
		key.s = a[2]
		j := FindCourseKey(key)
		Teachings[i].Course = Courses[j]
	case "Teacher":
		key.s = a[2]
		j := FindPersonKey(key)
		Teachings[i].Teacher = People[j]	
	case "StartDate":
		key.s = a[2]
		j := FindDateKey(key)
		Teachings[i].StartDate = Dates[j]
	case "EndDate":
		key.s = a[2]
		j := FindDateKey(key)
		Teachings[i].EndDate = Dates[j]
	case "Student":
		key.s = a[2]
		j := FindPersonKey(key)
		Teachings[i].Students = append(Teachings[i].Students,People[j])
	}
}

var Teachings []*Teaching