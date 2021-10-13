package main

import (
	"time"
)

func main() {
	
}

type ManifestV1 struct {
	// Name is the name of the image's repository
	Name string `json:"name"`
	
	// Tag is the tag of the image specified by this manifest
	Tag string `json:"tag"`
	
	// Architecture is the host architecture on which this image is intended to
	// run
	Architecture string `json:"architecture"`
	
	// FSLayers is a list of filesystem layer blobSums contained in this image
	FSLayers []FSLayerV1 `json:"fsLayers"`
	
	// History is a list of unstructured historical data for v1 compatibility
	History []map[string]string `json:"history"`
}

type FSLayerV1 struct {
	BlobSum string `json:"blobSum"`
}

type HistoryV1 struct {
	Throwaway       bool              `json:"throwaway"`
	Created         time.Time         `json:"created"` // "2020-12-22T18:04:24.760519058Z",
	LayerDegest     string            `json:"id"`
	ContainerConfig ContainerConfigV1 `json:"container_config"`
}

type ContainerConfigV1 struct {
	Cmd []string `json:"Cmd"`
}
