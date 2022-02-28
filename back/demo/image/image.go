package main

import (
	"fmt"
	"hash/fnv"
	"strings"

	"github.com/google/go-containerregistry/pkg/name"
)

func main() {
	imageName := "harbor"
	imageName = strings.TrimLeft(imageName, "https")
	fmt.Println(imageName)
}

func GenerateUUID(strs ...string) uint32 {
	s := strings.Join(strs, "/")
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	return h.Sum32()
}

func parseImage(imageName string) (*SimpleImage, error) {
	imageName = strings.TrimLeft(imageName, "https:")
	imageName = strings.TrimLeft(imageName, "http:")
	imageName = strings.TrimLeft(imageName, ":")
	imageName = strings.TrimLeft(imageName, "//")
	imageName = strings.TrimLeft(imageName, "/")

	var nameOpts []name.Option
	nameOpts = append(nameOpts, name.Insecure)

	ref, err := name.ParseReference(imageName, nameOpts...)
	if err != nil {
		return nil, err
	}
	repo := ref.Context()

	lib := repo.RegistryStr()

	tag := ref.Identifier()
	repoName := ref.Context().RepositoryStr()

	return &SimpleImage{Library: lib, FullRepoName: repoName, Tag: tag}, nil
}

type SimpleImage struct {
	Library      string
	FullRepoName string
	Tag          string
}
