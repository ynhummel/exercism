package airportrobot

import "fmt"

type Greeter interface {
    LanguageName() string
    Greet(a string) string
}

func SayHello(visitor string, greet Greeter) string {
    return fmt.Sprintf("I can speak %s: %s", greet.LanguageName(), greet.Greet(visitor))
}

type Italian struct {}

func(it Italian) LanguageName() string {
    return "Italian"
}

func(it Italian) Greet(name string) string {
    return "Ciao " + name + "!"
}

type Portuguese struct {}

func(pt Portuguese) LanguageName() string {
    return "Portuguese"
}

func(pt Portuguese) Greet(name string) string {
    return "Olá " + name + "!"
}

