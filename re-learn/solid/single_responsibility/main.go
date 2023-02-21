package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

// we want to save the jounal's entries to the persistent storate,
// we don't want to add this functionaility to the Journal itself
// but use a decicated party for doing so to separate the concerns

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I cried today")
	j.AddEntry("I ate a bug")

	fmt.Println(j.String())

	p := Persistence{
		lineSeparator: "\n",
	}
	p.SaveToFile(&j, "my_diary.txt")
}
