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

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



/*
 * Complete the 'commonSubstring' function below.
 *
 * The function accepts following parameters:
 *  1. STRING_ARRAY a
 *  2. STRING_ARRAY b
 */

func commonSubstring(a []string, b []string) {
    // Write your code here

    for counter:=0;counter<len(b);counter++{
    str:=[]rune(b[counter])    
    count:=0
    for counter1:=0;counter1<len(str);counter1++{
        if strings.Contains(a[counter],string(str[counter1])){
            count=count+1
        }
    }
    if count==0{
      fmt.Println("NO")  
    }else{
           fmt.Println("YES")  
   
    }
    }
    

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    aCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var a []string

    for i := 0; i < int(aCount); i++ {
        aItem := readLine(reader)
        a = append(a, aItem)
    }

    bCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var b []string

    for i := 0; i < int(bCount); i++ {
        bItem := readLine(reader)
        b = append(b, bItem)
    }

    commonSubstring(a, b)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
