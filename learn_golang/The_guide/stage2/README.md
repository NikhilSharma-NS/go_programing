### Deeper into GO

#### Project Overview

```
        Cards
newDeck -> create a list of playing cards, Essentially an array of strings
print   -> Log out the contents of a deck of cards
shuffle -> Shuffles all the cards in a deck
deal    -> create a hand of cards
saveToFile -> save a list of cards to file in system
newDeckFromFile -> load a list of cards from file
```

#### Variable Declarations

```
	var             card                  string =                           "Ace of Spades"


we are            the name              only a string will                  assign the value
about to          of the variable       ever be assigned to this            ace of spades to this varibel
declare a         will be geeeting      variable
variable 
```
#### Basic Go Types


```
bool
string
int
float64
```

#### Functions and Return Types

```
package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "my card"
}

```


```
func newCard() string {
}

newCard -> define a function 
called 'newCard'

string -> when executed this
function will return a value 
of type string
```


#### Slices and For loops

```
Array -> Fixed length 
list of things

Slice -> An array that can 
grow or shrink


```

```
package main

import "fmt"

func main() {

	card_slice := []string{newCard()}
	fmt.Println("slice", card_slice)
	card_slice = append(card_slice, newCard())
	card_slice = append(card_slice, newCard())
	fmt.Println("slice_after_add", card_slice)

	card_array := [2]string{newCard()}
	fmt.Println("array", card_array)
	card_array[1] = newCard()
	//	card_array[2] = newCard()
	// above line give error because size is only 2

	//card_array = append(card_array, newCard())
	// above line give error because this is not slice
	fmt.Println("array_after_add", card_array)

}

func newCard() string {
	return "mycard"
}

```

```
for index,card:= range cards{
    fmt.Println(card)
}

index-> index of this element in 
the array

card -> current card 
we are iterating over

range cards -> take the slice
of cards and loop over it

fmt.Println(card) -> run this
line one time for each card in 
the slice


```

#### OO Approach vs Go Approach

```


Deck class


Deck instance
    cards
    string

print   shuffle
func     func

saveToFile
func


```


```
Base type
string,integer,float
array,map

   
we can extend a base
type and add some extra
functionality to it   
    
     |

type deck []string

go want to create and array 
of strings and add a bunch 
of functions specifically
made to work with it


    |

Function with deck 
as a receiver

A function is like a method
a function that belongs to 
an instance


```


```

cards folder 

main.go

code to create 
and manipulate 
a deck

deck.go

code that describe
what a deck is and 
how it works

deck_test.go

code to 
automatically test the
deck



```

#### Custom Type declarations



```
package main


// create a new type of deck
// which is slice of string
type deck []string


```

```
package main

import "fmt"

// create a new type of deck
// which is slice of string
type deck []string

func (d deck) print() {
	for index, v := range d {
		fmt.Println("deck", index, v)
	}
}

```

#### Receiver functions

func (d deck) print(){

}

any variable of type deck
now gets access to the print 
method


```
func (d deck) print(){
}

d-> the actual copy of the 
deck we are working with 
is avaliable in the func
as a variable called 'd'


deck-> every variable of 
type `deck` can call 
this function on itself

```

#### Creating a New Deck

```
func newDeck() deck {
	newDeck := deck{"Initial"}
	return newDeck
}

```

#### Slice Range Syntax

```
fruits 

0     1       2      3
apple banana  grape  orange


fruits[startincludingindex:upToNotIncluding]


fruits[0:2]
apple banana


```

```
func main() {


	cardinput := deck{"apple" ,"banana" ," grape" ," orange"}

	fmt.Println(cardinput[:3])
	fmt.Println(cardinput[3:])
```
output:
```
[apple banana  grape]
[ orange]
```

#### Multiple Return Values


```
func (d deck) handSize(handSizei int) (start deck, end deck) {
	return d[:handSizei], d[handSizei:]
}
```

```
package main

import "fmt"

func main() {

	cardinput := deck{"apple", "banana", "grape", "orange"}

	start, end := cardinput.handSize(1)
	fmt.Println(start, end)

}

```

```
[apple] [banana grape orange]
```

#### Bytes Slices


```

hi 

[72 105]


```



#### Deck to string


```
[]byte("hi)

type  value
we    we 
want  have

```

```

bytes := []byte("hi")
fmt.Println(bytes)
```

```
(we have) 
deck
->[]string
->string->
[]byte(we want)

```
#### Joining a slice of Strings



```
func (d deck) toString() string {

	return strings.Join([]string(d), ",")
}

```

#### Saving Data to Hard Drive


```
func (d deck) saveToSystem(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
```

```
package main

import "fmt"

func main() {

	cardinput := deck{"apple", "banana", "grape", "orange"}
	stringVa := cardinput.toString()
	fmt.Println(stringVa)
	cardinput.saveToSystem("output.txt")

}

func newCard() string {
	return "new Card"
}

```

#### Reading Deck from Hard Drive

```

func readFromfile(fileName string) deck {
	bytes, err := ioutil.ReadFile(fileName)
	if err == nil {
		return deck(strings.Split(string(bytes), ","))
	}
	return nil
}
```

```
package main

import "fmt"

func main() {	

	cardsfromFile := readFromfile("output.txt")
	fmt.Println(cardsfromFile)

}



```

#### Shuffling a deck

```

func (d deck) shuffle() {
	for index, v := range d {
		newPos := rand.Intn(len(d) - 1)
		d[index], d[newPos] = d[newPos], d[index]
	}
}
```

```
package main

import "fmt"

func main() {

	cardsfromFile := readFromfile("output.txt")
	fmt.Println(cardsfromFile)

	cardsfromFile.shuffle()
	fmt.Println(cardsfromFile)

}


```

#### Testing with Go 

Go Testing is not RSpec,mocha,jasmine,selenium etc

```
go test ./... -coverprofile cover.out
go tool cover -html .\cover.out
```

```
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected length of deck is 16, but got %v", len(d))
	}
}
```