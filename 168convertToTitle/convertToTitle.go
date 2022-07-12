package main

func main() {

}

func convertToTitle(columnNumber int) string {
	ans := ""
	for columnNumber > 26 {
		num := columnNumber % 26
		if num == 0 {
			ans = "Z" + ans
			columnNumber = (columnNumber / 26) - 1
		} else {
			ans = string(byte(64+num)) + ans
			columnNumber = columnNumber / 26
		}
	}
	ans = string(byte(columnNumber+64)) + ans
	return ans
}
