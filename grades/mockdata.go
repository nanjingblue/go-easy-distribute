package grades

import (
	"github.com/bxcodec/faker/v3"
	"math/rand"
	"reflect"
	"time"
)

func GetFaker(num int) []Student {
	CustomGenerator()
	_ = faker.SetRandomMapAndSliceMinSize(3)
	_ = faker.SetRandomMapAndSliceMaxSize(3)
	var students []Student
	for i := 0; i < num; i++ {
		var s Student
		_ = faker.FakeData(&s)
		students = append(students, s)
	}
	return students
}

func CustomGenerator() {
	rand.Seed(time.Now().UnixNano())
	_ = faker.AddProvider("title", func(v reflect.Value) (interface{}, error) {
		courses := []string{"Chinese", "Math", "English"}
		return courses[rand.Intn(3)], nil
	})
	_ = faker.AddProvider("gradetype", func(v reflect.Value) (interface{}, error) {
		types := []GradeType{GradeQuiz, GradeTest, GradeExam}
		return string(types[rand.Intn(3)]), nil
	})
}
