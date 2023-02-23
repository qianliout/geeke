package main

import (
	"strconv"
	"strings"
)

func ipToUint8Slice(ip string) []uint8 {
	ans := make([]uint8, 0)
	split := strings.Split(ip, "/")
	if len(split) == 0 {
		return ans
	}
	cnt, err := strconv.Atoi(split[1])
	if err != nil {
		return ans
	}
	ips := strings.Split(split[0], ".")
	ipNumber := make([]uint8, 0)
	for _, p := range ips {
		atoi, err := strconv.Atoi(p)
		if err != nil {
			return ans
		}
		ipNumber = append(ipNumber, uint8(atoi))
	}
	for _, ip := range ipNumber {
		uint8s := uint8ToUint8slice(ip)
		for _, u8 := range uint8s {
			ans = append(ans, u8)
			if len(ans) >= cnt {
				return ans
			}
		}
	}
	return ans
}

func uint8ToUint8slice(ip uint8) []uint8 {
	ans := make([]uint8, 0)
	for i := 0; i < 8; i++ {
		a := (ip & (1 << (8 - i - 1))) >> (8 - i - 1)
		ans = append(ans, a)
	}
	return ans
}
