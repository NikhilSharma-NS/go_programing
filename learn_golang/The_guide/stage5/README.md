### Interface

#### Purpose of Interface

if same set of things are required in multiple places 
then we will go with interface

#### Problems without Interface

```
package main

import "fmt"

type englishBot struct {
}

type spanisBot struct {
}

func (eb englishBot) getGretting() string {
	return "Hi english"
}

func (sb spanisBot) getGretting() string {
	return "Hello spanish"
}

func printGretting(eb englishBot) {
	fmt.Println(eb.getGretting())
}

// func printGretting(sb spanisBot) {

// }

// we can't have same name function

func main() {

}

```

#### Interface in Practice


```
package main

import "fmt"

type englishBot struct {
}

type spanisBot struct {
}

type bot interface {
	getGretting() string
}

func (eb englishBot) getGretting() string {
	return "Hi english"
}

func (sb spanisBot) getGretting() string {
	return "Hello spanish"
}

func printGretting(b bot) {
	fmt.Println(b.getGretting())
}

// func printGretting(eb englishBot) {
// 	fmt.Println(eb.getGretting())
// }

// func printGretting(sb spanisBot) {

// }

func main() {
	eb := englishBot{}
	sb := spanisBot{}
	printGretting(sb)
	printGretting(eb)

}

```

```
type bot interface
Our Program has a new type
called bot

getGretting() string
if you are a type in this program
with a function called 'getGretting'
and you return a string then you are now 
an honorary member of type bot

func printGretting(b bot)
Now that you are also an honorary member of
type 'bot', you can now call this function
called 'printGreeting'
```

#### Rules of Interface



```
type bot interface {
	getGretting() string
}

type englishBot struct {
}

func (eb englishBot) getGretting() string {
	return "Hi english"
}

type spanisBot struct {
}


func (sb spanisBot) getGretting() string {
	return "Hello spanish"
}
```

```

type bot interface{
	getGretting() string

}

bot -> interface name
getGretting -> func name
() -> list of argument
string -> return type
```

```
Concrete Type
map
struct
int
string


interface Type
bot


```


Inferface are not generic types

Interface are implicit

Interfaces are a contract to help us
manage types

Interfaces are tough
(how to read them )