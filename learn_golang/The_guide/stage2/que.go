/*
 * Complete the 'commonSubstring' function below.
 *
 * The function accepts following parameters:
 *  1. STRING_ARRAY a
 *  2. STRING_ARRAY b
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func commonSubstring(a []string, b []string) {
	// Write your code here

	for counter := 0; counter < len(b); counter++ {
		str := []rune(b[counter])
		count := 0
		for counter1 := 0; counter1 < len(str); counter1++ {
			if strings.Contains(a[counter], string(str[counter1])) {
				count = count + 1
			}
		}
		if count == 0 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")

		}
	}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

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

// /*
//  * Complete the 'commonSubstring' function below.
//  *
//  * The function accepts following parameters:
//  *  1. STRING_ARRAY a
//  *  2. STRING_ARRAY b
//  */

// func commonSubstring(a []string, b []string) {
// 	// Write your code here

// 	for counter := 0; counter < len(b); counter++ {
// 		str := []rune(b[counter])
// 		count := 0
// 		for counter1 := 0; counter1 < len(str); counter1++ {
// 			if strings.Contains(a[counter], string(str[counter1])) {
// 				count = count + 1
// 			}
// 		}
// 		if count == 0 {
// 			fmt.Println("NO")
// 		} else {
// 			fmt.Println("YES")

// 		}
// 	}

// }

// func main() {
// 	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

// 	aCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 	checkError(err)

// 	var a []string

// 	for i := 0; i < int(aCount); i++ {
// 		aItem := readLine(reader)
// 		a = append(a, aItem)
// 	}

// 	bCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 	checkError(err)

// 	var b []string

// 	for i := 0; i < int(bCount); i++ {
// 		bItem := readLine(reader)
// 		b = append(b, bItem)
// 	}

// 	commonSubstring(a, b)
// }

// func readLine(reader *bufio.Reader) string {
// 	str, _, err := reader.ReadLine()
// 	if err == io.EOF {
// 		return ""
// 	}

// 	return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
