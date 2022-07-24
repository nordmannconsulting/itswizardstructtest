package basic_structs

import "errors"

type Teststtruct struct {
	Name  string
	Value int
	init  bool
}

// Was macht die Funktion?
func NewTestStructure(name string, value int) *Teststtruct {
	p := new(Teststtruct)
	p.Name = name
	p.Value = value
	p.init = true
	return p
}

func (p *Teststtruct) GetName() (string, error) {
	if !p.init {
		return "", errors.New("The Teststruct is not initial, call NewTestStruct")
	} else {
		return p.Name, nil
	}
}

func (p *Teststtruct) GetValue() (int, error) {
	if !p.init {
		return 0, errors.New("The Teststruct is not initial, call NewTestStruct")
	} else {
		return p.Value, nil
	}
}

func (p *Teststtruct) Increment() error {
	if !p.init {
		return errors.New("The Teststruct is not initial, call NewTestStruct")
	} else {
		p.Value++
		return nil
	}
}
