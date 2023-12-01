package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	InvalidDecode = -1
)

type Integer struct {
	v          string
	discovered bool
}

func NewInteger() *Integer {
	return &Integer{}
}

func isInt(v string) bool {
	// Checks if the current character can be represented as an integer
	_, err := strconv.Atoi(v)
	if err != nil {
		return false
	}
	return true
}

func convergeInts(x, y *Integer) (int, error) {

	// if x.v == "" {
	// 	y_int, err := strconv.Atoi(y.v)
	// 	if err != nil {
	// 		return InvalidDecode, err
	// 	}
	// 	return y_int, nil
	// }

	// if y.v == "" {
	// 	x_int, err := strconv.Atoi(x.v)
	// 	if err != nil {
	// 		return InvalidDecode, err
	// 	}
	// 	return x_int, nil
	// }

	x_int, err := strconv.Atoi(x.v)
	if err != nil {
		return InvalidDecode, err
	}
	// check if there is only one integer present within the array
	y_int, err := strconv.Atoi(y.v)
	if err != nil {
		return InvalidDecode, err
	}

	z, err := strconv.Atoi(fmt.Sprintf("%d%d", x_int, y_int))
	if err != nil {
		return InvalidDecode, err
	}

	return z, nil
}

func DecodeCalibrationValue(v string) (int, error) {
	stream := strings.Split(v, "")

	fmt.Println("String:", v, "String length:", len(v))

	// Divide the array into 2 halves
	// hops := int(math.Ceil(float64(len(stream)) / 2.0))
	// fmt.Println("Hops: ", hops)

	x := NewInteger()
	y := NewInteger()

	// Walk forward from the start and walk backwards from the back end of the string
	for i := 0; i <= len(stream); i++ {
		if !x.discovered {
			ok := isInt(stream[i])
			if ok {
				x.v = stream[i]
				x.discovered = true
			}
		}
		if !y.discovered {
			y_index := (len(stream) - 1) - i
			fmt.Println(stream[y_index])
			ok := isInt(stream[y_index])
			if ok {
				y.v = stream[y_index]
				y.discovered = true
			}
		}
	}

	z, err := convergeInts(x, y)
	if err != nil {
		return InvalidDecode, err
	}

	return z, nil
}

func run() error {
	buf, err := os.ReadFile("1.txt")
	if err != nil {
		return err
	}

	stream := strings.Split(string(buf), "\n")

	sum := 0
	current := 0

	for _, v := range stream {
		z, err := DecodeCalibrationValue(v)
		if err != nil {
			return err
		}
		fmt.Println("Value:", z)
		sum += z
		current++
	}

	fmt.Println("Final Sum:", sum, "Current Line:", current)

	// TODO DecodeCalibrationValue breaks if there is only 1 integer present within the string
	// z, err := DecodeCalibrationValue("85oneighttx")
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(z)

	return nil
}

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}
