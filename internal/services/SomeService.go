package services

type SomeService struct{}

func (s SomeService) Run(st string, in int) (string, int) {
	return st, in
}
