package main

import (
	"fmt"
	"hash/fnv"
	"strings"
)

func main() {
	res := make([]SimpleImage, 3)

	for i := range res {
		res[i].Ser()
	}
	fmt.Println("hello wrod")

}

func GenerateUUID(strs ...string) uint32 {
	s := strings.Join(strs, "/")
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	return h.Sum32()
}

type SimpleImage struct {
	Library      string
	FullRepoName string
	Tag          string
}

func (si *SimpleImage) Ser() {
	si.Library = "hello"
	si.FullRepoName = "word"
	si.Tag = "what"

}
