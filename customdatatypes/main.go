package main

import (
	"customdatatypes/organization"
	"fmt"
)

func main() {

	p := organization.NewPerson("Kishore", "Y", organization.NewSocialSecurityNumber("1234-567-890"))

	err := p.SetTwitterhandler("@durgakishore")

	if err != nil {
		fmt.Printf("An error occured setting twitter handler %s", err.Error())
	}

	fmt.Println(p.TwitterHandler())
	fmt.Println(p.ID())
	fmt.Println(p.FullName())
}
