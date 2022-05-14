package main

import (
	"fmt"
	"hash/fnv"
	"strings"
)

func main() {
	ss := []string{"hello", "word", "fuck"}

	pp := ProjectSchedule{
		Total:     13,
		Completed: 34,
	}
	for _, v := range ss {
		pp.Project = v
		print(pp)
	}

}

func print(pp ProjectSchedule) {
	fmt.Println(pp.Project)
	fmt.Println(pp.Total)
	fmt.Println(pp.Completed)
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

type ImageSyncSchedule struct {
	ID           int64             `json:"id"`
	RegistryId   int64             `json:"registryId"`
	RegistryName string            `json:"registryName"`
	RegistryURL  string            `json:"registryURL"`
	Projects     []ProjectSchedule `json:"projects"`
	Total        int64             `json:"total"`
	Completed    int64             `json:"completed"`
	Done         bool              `json:"done"`
	CreatedAt    int64             `json:"createdAt"`
	UpdatedAt    int64             `json:"updatedAt"`
}

type ProjectSchedule struct {
	Project   string `json:"project"`
	Total     int64  `json:"total"`
	Completed int64  `json:"completed"`
}
