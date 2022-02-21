package main

import (
	"fmt"
	"strings"

	"github.com/google/go-containerregistry/pkg/name"
)

func main() {
	image := "https:ts-acr-registry.cn-hangzhou.cr.aliyuncs.cn/liuqianli/001100/spa_autobuilds:hello"
	simpleImage, err := parseImage(image)
	if err != nil {
		fmt.Println("err is ", err.Error())
		return
	}
	fmt.Println(simpleImage.Library, simpleImage.FullRepoName, simpleImage.Tag)

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
