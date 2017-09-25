package services

import "fmt"

type SomeService struct{}

func (s SomeService) Run(st string, in int) (string, int) {
	if in == 4 {
		st = fmt.Sprintf("%s%v", st, in)
	}
	return st, in
}
