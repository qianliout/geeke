package main

import (
	"fmt"
)

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	res := canCompleteCircuit(gas, cost)
	fmt.Println("res is ", res)
}

func canCompleteCircuit(gas []int, cost []int) int {
	// 先验证是否可以走完
	if len(gas) == 0 || len(cost) == 0 || len(gas) != len(cost) {
		return -1
	}
	currGas, totalGas, start := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		currGas = currGas + gas[i] - cost[i]
		totalGas = totalGas + gas[i] - cost[i]
		if currGas < 0 {
			start = i + 1
			currGas = 0
		}
	}
	if totalGas < 0 {
		return -1
	}
	return start

}
