package model

import "errors"

type person struct {
	name string
}

func NewPerson(n string) (p *person) {
	p = &person{
		name: n,
	}
	return
}

func (p *person) GetName() (n string, err error) {
	if p != nil {
		return p.name, nil
	} else {
		err = errors.New("aaaa")
		return "", err
	}
}

func (p *person) SetName(n string) (err error) {
	if p == nil {
		err = errors.New("sss")
		return
	}
	p.name = n
	return nil
}
