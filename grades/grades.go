package grades

import (
	"fmt"
	"sync"
)

type Student struct {
	ID        int    `faker:"boundary_start=100, boundary_end=1000"`
	FirstName string `faker:"first_name"`
	LastName  string `faker:"last_name"`
	Grades    []Grade
}

func (s Student) Average() float32 {
	var result float32
	for _, grade := range s.Grades {
		result += grade.Score
	}
	return result / float32(len(s.Grades))
}

type Students []Student

var (
	students     Students = GetFaker(3)
	studentMutex sync.Mutex
)

func (ss Students) GetByID(id int) (*Student, error) {
	for _, v := range ss {
		if v.ID == id {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("student with id '%d' not found", id)
}

type GradeType string

const (
	GradeQuiz = GradeType("Quiz")
	GradeTest = GradeType("Test")
	GradeExam = GradeType("Exam")
)

type Grade struct {
	Title string    `faker:"title"`
	Type  GradeType `faker:"gradetype"`
	Score float32   `faker:"boundary_start=0.00, boundary_end=100.00"`
}
