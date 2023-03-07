package people

import "time"

type People struct {
	Id         uint
	Name       string
	Age        int32
	Number     int64
	Job        string
	Created_at time.Time
}

func NewPeople(name string, age int32, number int64, job string) (*People, error) {
	people := &People{
		Name:       name,
		Age:        age,
		Number:     number,
		Job:        job,
		Created_at: time.Now(),
	}

	return people, nil
}
