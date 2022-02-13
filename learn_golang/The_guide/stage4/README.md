### Maps


Map GO
Hash Ruby
Object javascript
Dict Python


#### create the map 

```
import "fmt"

func main() {

	color := map[string]string{
		"key": "value",
		"ke1": "value1",
	}

	fmt.Println(color)

}

```

```
	var colors map[string]string
	fmt.Println(colors)
```

```
	colors := make(map[string]string)
	fmt.Println(colors)
```

#### Manipulate the map and iterate

```
package main

import "fmt"

func main() {
	colors := make(map[string]string)
	colors["key"] = "va"
	delete(colors, "key")
	colors["key1"] = "vaa"
	colors["key"] = "vaaa"
	fmt.Println(colors)

	delete(colors, "key")
	fmt.Println(colors)

}

```

```
package main

import "fmt"

func main() {


	colors := make(map[int]string)
	colors[1] = "va"

	delete(colors, 3)
	fmt.Println(colors)

	for key, v := range colors {
		fmt.Println(key, "-", v)
	}

}
```

#### Difference between maps and structs 


```
Map
1) All keys must be the same type
2) All values must be the same type
3) Keys are indexded we can iterate over them
4) Use to represent a collection of related propeties
5) Don't need to know all the key at compile time
6) reference Type

Struct
1) Values can be of different type 
2) key's don't support indexing
3) we need to know all the different fields 
at the compile time
4) Use to represent a thing with a lot of
different properties
5) Value Type



```

