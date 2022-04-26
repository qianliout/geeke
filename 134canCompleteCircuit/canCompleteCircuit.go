package main

func main() {

}

/*
在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
给定两个整数数组 gas 和 cost ，如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。
*/
func canCompleteCircuit2(gas []int, cost []int) int {
	if len(gas) == 0 || len(gas) != len(cost) {
		return -1
	}
	n := len(gas)
	for i := 0; i < len(gas); i++ {
		j, remain := i, gas[i]
		for remain-cost[j] >= 0 {
			remain = remain + gas[(j+1)%n] - cost[j]
			j = (j + 1) % n
			if j == i {
				return i
			}
		}
		if j < i {
			return -1
		}
		i = j
	}
	return -1
}

func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) == 0 || len(gas) != len(cost) {
		return -1
	}
	cur, total, start := 0, 0, 0

	for i := 0; i < len(gas); i++ {
		cur = cur + gas[i] - cost[i]
		total = total + gas[i] - cost[i]
		if cur < 0 {
			cur = 0
			start = i + 1
		}
	}

	if total < 0 {
		return -1
	}
	return start
}
