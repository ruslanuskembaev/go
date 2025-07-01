package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(string) string
}

func SayHello(name string, g Greeter) string {
	return fmt.Sprintf("I can speak %s: %s %s!", g.LanguageName(), g.Greet(g.LanguageName()), name)
}

type Italian struct {
}

type Portuguese struct {
}

func (i Italian) LanguageName() string {
	return "Italian"
}
func (i Italian) Greet(string) string {
	return "Ciao"
}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}
func (p Portuguese) Greet(string) string {
	return "Olá"
}
