package main

import (
	"fmt"


)

/*
-Arrays are fixed where as slices are dynamic || flexible in runtime && Slice is just array wraper

-In Go, a slice is a descriptor of a segment of an array. It doesn't hold the actual data itself.
  it contains a reference to the underlying array along with information about the length of the slice
  and its capacity.

-These might seem wierd now but when we learn about database we try to filter our data using this
 new array power(slice) we can filter the data in database.

 -
*/

func main()  {

	fmt.Println("Welcom to slices in Go")

	 fruitlist := []string{"banana","apple","peach","mango"}

	 fruitlist = append(fruitlist, "straberry","goldnuts")

	//  fruitlist = append(fruitlist[1:5]) Old way of slicing

	sliceFruit := fruitlist[0:3] // new way of slicing

	 fmt.Println("New fruitlist is:",sliceFruit)

	 //++++++++++++++++++++++++++++++++++++++++++++++++++++

	 /*
	  Slices in Go are references to sections of arrays.
	  Modifying a slice directly affects the underlying array and any other slices that reference
	  the same array segment.
	 */

	 anime := []string{"JJK","Naruto","Onepiece","ClassroomOfElite","MyHeroAcademia"}
	 fmt.Println("All anime",anime)


	 myFav := anime[0:3]
	 myNonFav := anime[2:]

	 fmt.Println("My Fav and non Fav",myFav,myNonFav)

	 myFav[2]="Boruto"

	 fmt.Println("New Fav,Non Fav",myFav,myNonFav)
	 fmt.Println("New All anime",anime)

	 anime = append(anime, "SoloLeveling")

	//  myNonFav[0] = "Shinchan"

	 fmt.Println("Append Check",anime,myNonFav)



	 // ++++++++++++++++++++++++++++++
	 fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	 a := make([]int,2,3)

	// var a []int

	 fmt.Println("Capacity of  a",cap(a))
	fmt.Println("Length of  a",len(a))


}
