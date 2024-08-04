package main

import "fmt"

func main()  {
	fmt.Println("Welcome Maps in Go")

	// Maps sre like similar to objects in javascript we can store our multiple data in maps

	myLang := make(map[string]string)

	// Adding iteams to map
	myLang["CPP"] = "C++"
	myLang["JS"] = "Javascript"
	myLang["TS"] = "Typescript"
	myLang["GO"] = "GoLang"

	//Update
	myLang["GO"] = "Golng🦫"

	//Delete
	delete(myLang,"CPP")

	//Retrive or Read 
	myFavLang := myLang["JS"]

	fmt.Println("My Languages",myLang)
	fmt.Println("My Fav Languages",myFavLang)

	// Retriving data which does not exist
	v1 := myLang["PY"]

	fmt.Println("v1:",v1) // for strings it returns "" & for numbers it returns 0

	// Looping through maps
	fmt.Println("================================================")

	for key,value := range myLang {
		fmt.Printf("The key is %v, the value is %v \n",key,value)
	}


	// Declaring and initializing in maps diffrently
	m := map[string]int{"foo":1,"bar":2}

	for _,value := range m{
		fmt.Printf("the value is %v \n",value)
	}
	
}
