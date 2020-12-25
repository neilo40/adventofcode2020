package main

import (
	"fmt"
)

func main() {
	part1()
	//part2()
}

func part1() {
	cardPublicKey := 6270530
	doorPublicKey := 14540258

	//cardLoopSize := getLoopSize(cardPublicKey, 7)
	doorLoopSize := getLoopSize(doorPublicKey, 7)

	encryptionKey := transform(cardPublicKey, doorLoopSize)

	fmt.Printf("Encryption Key: %d\n", encryptionKey)
}

func getLoopSize(publicKey int, subject int) (loopSize int) {
	loopSize = 1
	value := 1
	for {
		value *= subject
		value = value % 20201227
		if value == publicKey {
			return
		}
		loopSize++
	}
}

func transform(subject int, loopSize int) (key int) {
	key = 1
	for i := 0; i < loopSize; i++ {
		key *= subject
		key = key % 20201227
	}
	return
}

func part2() {
	fmt.Printf("Result: \n")
}
