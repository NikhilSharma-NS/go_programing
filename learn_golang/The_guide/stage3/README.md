### Oragnizing Data with Structs


Struct 

Data Structure 
Collection of properties 
that are related together


#### Define the strct 

```
type Person struct {
	firstName string
	lastName  string
}
```

#### Declare the struct 

```
	p := person{firstName: "nikhi", lastName: "sharma"}
	p1 := person{"nikhi", "sharma"}

	fmt.Println(p)
	fmt.Println(p1)
```

#### Embedding structs


```
package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {

	p := person{
		firstName: "Nikhi",
		lastName:  "Sharma",
		contact: contactInfo{
			email:   "nk@gmail.com",
			zipCode: 411057,
		},
	}
	fmt.Println(p)

}

```

#### struct with receiver functions


```
package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {


	p := person{
		firstName: "Nikhi",
		lastName:  "Sharma",
		contactInfo: contactInfo{
			email:   "nk@gmail.com",
			zipCode: 411057,
		},
	}
	fmt.Println(p)

}

```

```
package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	fmt.Printf("%v", p)
}

func main() {

	p := person{
		firstName: "Nikhi",
		lastName:  "Sharma",
		contactInfo: contactInfo{
			email:   "nk@gmail.com",
			zipCode: 411057,
		},
	}
	fmt.Println(p)
	p.print()

}

```

#### Pass by Value 
```
func (p *person) updateName(newfirstname string) {
	p.firstName = newfirstname
}
```


 ```
 package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	fmt.Printf("%v", p)
}

func (p *person) updateName(newfirstname string) {
	p.firstName = newfirstname
}

func main() {


	p := person{
		firstName: "Nikhi",
		lastName:  "Sharma",
		contactInfo: contactInfo{
			email:   "nk@gmail.com",
			zipCode: 411057,
		},
	}
	fmt.Println(p)
	p.updateName("nicky")
	p.print()

}

 ```

 ```
 package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	fmt.Printf("%v", p)
}

func (p1 *person) updateName(newfirstname string) {
	p1.firstName = newfirstname
}

func main() {
	p := person{
		firstName: "Nikhi",
		lastName:  "Sharma",
		contactInfo: contactInfo{
			email:   "nk@gmail.com",
			zipCode: 411057,
		},
	}
	fmt.Println(p)
	p2 := &p
	p.updateName("nicky")
	p.print()
	fmt.Println()
	p2.print()

}

 ```

 #### Value Types vs Reference Types

Values Types

 int
 float
 string
 bool
 structs


Reference Types

 slices
 maps
 channels
 pointers
 functions

