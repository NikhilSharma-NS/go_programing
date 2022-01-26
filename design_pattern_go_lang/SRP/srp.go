package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

var entryCount = 0

type Journel struct {
	entries []string
}

func (j *Journel) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journel) RemoveEntry(index int) {

}

func (j *Journel) String() string {
	return strings.Join(j.entries, "\n")
}

// sepration of concerns

func (j *Journel) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journel) Load(filename string) {

}
func (j *Journel) LoadFromWeb(url *url.URL) {
}

var LineSeprator = "\n"

func (p *Persistence) SaveToFile(j *Journel, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSep)), 0644)

}

type Persistence struct {
	lineSep string
}

func main() {

	j := Journel{}
	j.AddEntry("first")
	j.AddEntry("second")
	fmt.Println(j.String())

	p := Persistence{"\r\n"}
	p.SaveToFile(&j, "jon.txt")

}
