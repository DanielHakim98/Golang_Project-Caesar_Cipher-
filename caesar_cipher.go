package main

import (
	"fmt"
	"strconv"
)

func checkWrapCapital(shiftedNum int64) int64 {
	if shiftedNum > 90 {
		for shiftedNum > 90 {
			shiftedNum = shiftedNum - 90 + 65 - 1
		}
	} else {
		for shiftedNum < 65 {
			shiftedNum = 90 - (65 - shiftedNum) + 1
		}
	}

	return shiftedNum
}

func checkWrapLower(shiftedNum int64) int64 {
	if shiftedNum > 122 {
		for shiftedNum > 122 {
			shiftedNum = shiftedNum - 122 + 97 - 1
		}
	} else {
		for shiftedNum < 97 {
			shiftedNum = 122 - (97 - shiftedNum) + 1
		}
	}

	return shiftedNum
}

func convertByteStringInt64(b byte) int64 {

	// Convert from byte to string and convert to int64
	byteString := fmt.Sprintf("%v", b)
	num, err := strconv.ParseInt(byteString, 10, 0)
	if err != nil {
		fmt.Println(err)
	}
	return num
}

func caesar(word string, shift int64) (string, error) {

	wordByte := []byte(word)
	shiftedByte := make([]byte, 0)

	for _, charByte := range wordByte {
		var shiftedNum int64
		byteLiteral := convertByteStringInt64(charByte)

		if byteLiteral >= 65 && byteLiteral <= 90 {
			// Make shift by adding 'shift' to the num
			shiftedNum = byteLiteral + shift

			// To check if the integer is out of range from ASCII range for alphabetic characters
			if shiftedNum < 65 || shiftedNum > 90 {
				shiftedNum = checkWrapCapital(shiftedNum)
			}

		} else if byteLiteral >= 97 && byteLiteral <= 122 {
			// Make shift by adding 'shift' to the num
			shiftedNum = convertByteStringInt64(charByte) + shift

			// To check if the integer is out of range from ASCII range for alphabetic characters
			if shiftedNum < 97 || shiftedNum > 122 {
				shiftedNum = checkWrapLower(shiftedNum)
			}

		} else {
			shiftedNum = byteLiteral
		}

		shiftedByte = append(shiftedByte, byte(shiftedNum))
	}
	return string(shiftedByte), nil
}

func main() {
	ans, err := caesar("Hello, World!", 5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println([]byte("Hello, World!"))
	fmt.Println([]byte(ans))
	fmt.Println(ans)
}
