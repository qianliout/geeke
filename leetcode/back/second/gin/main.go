package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {

	client := http.DefaultClient
	for i := 0; i < 200; i++ {
		req, err := http.NewRequest(http.MethodDelete, "http://172.21.0.3:30682/api/v1/images/bases/1", nil)
		if err != nil {
			log.Printf("删除基础镜像出错:%s", err.Error())
			continue
		}
		resp1, err := client.Do(req)
		if err != nil {
			log.Printf("删除基础镜像出错:%s", err.Error())
			continue
		}
		resp1.Body.Close()

		resp2, err := http.Get("http://172.21.0.3:30682/api/v1/scan/reportsByImageDetails?id=1")
		if err != nil {
			log.Printf("获取镜像信息出错:%s", err.Error())
			continue
		}
		s, err := ioutil.ReadAll(resp2.Body)
		if !strings.Contains(strings.Replace(string(s), " ", "", -1), "\"image_type\":0") {
			log.Printf("出现的延迟")
		}

		resp2.Body.Close()

		bys, _ := json.Marshal(map[string][]int{"image_ids": []int{1}})
		resp3, err := http.Post("http://172.21.0.3:30682/api/v1/images/bases", "application/x-www-form-urlencoded", bytes.NewReader(bys))
		if err != nil {
			log.Printf("增加基础镜像出错:%s", err.Error())
			continue
		}
		resp3.Body.Close()
		resp4, err := http.Get("http://172.21.0.3:30682/api/v1/scan/reportsByImageDetails?id=1")
		if err != nil {
			log.Printf("获取镜像信息出错:%s", err.Error())
			continue
		}
		s4, err := ioutil.ReadAll(resp4.Body)
		if !strings.Contains(strings.Replace(string(s4), " ", "", -1), "\"image_type\":1") {
			log.Printf("出现的延迟")
		}
	}
}
