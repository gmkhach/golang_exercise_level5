package main

import "fmt"

func main() {
	/*
		Exercise 1
			1. Create your own type “person” which will have an underlying type of “struct” so that it can store the following data: first name, last name, favorite ice cream flavors.
			2. Create two VALUES of TYPE person. 
			3. Print out the values, ranging over the elements in the slice which stores the favorite flavors. 

    */
	type person struct {
		firstName				string
		lastName				string
		favIceCreamFlavors	[]string
	}

	var a = person{
		firstName: "Deanna", 
		lastName: "Troi",
		favIceCreamFlavors: []string{"Chocolate", "Vanilla"},
	}

	var b = person{
		firstName: "Jean-Luc",
		lastName: "Picard",
		favIceCreamFlavors: []string{"None"},
	}

	fmt.Printf("first name: %v\nlast name: %v\nfavorite flavors: ", a.firstName, a.lastName)
	for i, v := range a.favIceCreamFlavors {
		fmt.Printf("%v", v)
		if i == len(a.favIceCreamFlavors) - 1 {
			fmt.Printf("\n")
		} else {
			fmt.Printf(", ")
		}
	}
	
	fmt.Printf("first name: %v\nlast name: %v\nfavorite flavors: ", b.firstName, b.lastName)
	for i, v := range b.favIceCreamFlavors {
		fmt.Printf("%v", v)
		if i == len(b.favIceCreamFlavors) - 1 {
			fmt.Printf("\n")
		} else {
			fmt.Printf(", ")
		}
	}

	/*
		Exercise 2
		Take the code from the previous exercise, then ...
			1. Store the values of type person in a map with the key of last name. 
			2. Access each value in the map. Print out the values, ranging over the slice.
    */
	var m = map[string]person{
		a.lastName: a,
		b.lastName: b,
	}

	fmt.Printf("first name: %v\nlast name: %v\nfavorite flavors: ", m["Troi"].firstName, m["Troi"].lastName)
	for i, v := range m["Troi"].favIceCreamFlavors {
		fmt.Printf("%v", v)
		if i == len(a.favIceCreamFlavors) - 1 {
			fmt.Printf("\n")
		} else {
			fmt.Printf(", ")
		}
	}
	
	fmt.Printf("first name: %v\nlast name: %v\nfavorite flavors: ", m["Picard"].firstName, m["Picard"].lastName)
	for i, v := range m["Picard"].favIceCreamFlavors {
		fmt.Printf("%v", v)
		if i == len(b.favIceCreamFlavors) - 1 {
			fmt.Printf("\n")
		} else {
			fmt.Printf(", ")
		}
	}

	/*
		Exercise 3
			1. Create a new type: vehicle. The underlying type is a struct with the following fields: doors, color.
			2. Create two new types: truck & sedan. The underlying type of each of these new types is a struct. 
			3. Embed the “vehicle” type in both truck & sedan. 
			4. Give truck the field “fourWheel” which will be set to bool. 
			5. Give sedan the field “luxury” which will be set to bool.
			6. Using the vehicle, truck, and sedan structs: 
				a. using a composite literal, create a value of type truck and assign values to the fields; 
				b. using a composite literal, create a value of type sedan and assign values to the fields. 
				c. Print out each of these values. 
				d. Print out a single field from each of these values.

	*/
	type vehicle struct {
		doors		int
		color		string
	}

	type truck struct {
		vehicle
		fourWheel	bool
	}

	type sedan struct {
		vehicle
		luxury 		bool
	}

	f150 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "blue",
		},
		fourWheel: true,
	}

	s_class := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}

	fmt.Println(f150)
	fmt.Println(s_class)

	fmt.Printf("Some f150s are still offered in a %v-door configuration\n", f150.doors)
	fmt.Printf("The Mercedes S Class is a leader in a %v leader in the luxury segment\n", s_class.luxury)

	/*
		Exercise 4
		Create and use an anonymous struct.
	*/
	autoMakers := struct {
		makes 	[]string
		models	map[string][]string
	}{
		makes: []string{"Ford", "Mercedes"},
		models: map[string][]string{
			"Ford": []string{
				"Branco", "Branco Sport", "Ecosport", "Edge", "Escape", "Explorer", "Expedition", 
				"Flex", "Fusion", "Mustang", "Mustang Mach-E", "F-150", "F-150 Raptor", "Super Duty", 
				"Ranger", "Transit", "Transit Connect", "Ford GT", "Shelby GT500",
			},
			"Mercedes": []string{
				"A-Class", "B-Class", "C-Class", "CLA", "CLS", "E-Class", "EQC", "G-Class", "GLA", 
				"GLB", "GLC", "GLS", "S-Class", "SL", "SLC", "V-Class", "AMG GT", "Maybach",
			},
		},
	}
  
	fmt.Printf("makes: %v\n", autoMakers.makes)
	for i, models := range autoMakers.models {
		fmt.Printf("make: %v\n", i)
		for _, v := range models {
			fmt.Printf("\t%v\n", v)
		} 
	}
}