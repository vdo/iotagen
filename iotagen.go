package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"os"
)

const seedSize int = 81
const seedChars string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ9"


func main() {
	args := os.Args[1:]
	var iterations int = 10

	if len(args) != 0 {
			i, err := strconv.Atoi(args[0])
			if err != nil {
					log.Println(err)
					os.Exit(1)
			}
			iterations = i
	}

	for i := 0; i < iterations; i++ {
			fmt.Println(randomIotaSeed(seedSize))
	}
}

func randomIotaSeed(size int) string {
    buf := make([]byte, size)
    for i := 0; i < size; i++ {
        nbig, err := rand.Int(rand.Reader, big.NewInt(int64(len(seedChars))))
        if err != nil {
            log.Println(err)
        }
        buf[i] = seedChars[nbig.Int64()]
    }
    return string(buf)
}
