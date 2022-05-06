package main

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"strings"
	"time"
)

func main() {

	iss := ImageSyncSchedule{
		ID:         12,
		RegistryId: 3,
		Projects: []ProjectSchedule{{
			Project:   "tensorsec",
			Total:     30,
			Completed: 20,
		}},
		Total:     100,
		Completed: 30,
		Done:      false,
		CreatedAt: time.Now().UTC().Unix(),
		UpdatedAt: time.Now().UTC().Unix(),
	}
	bys, _ := json.Marshal(iss)
	fmt.Println(string(bys))

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
