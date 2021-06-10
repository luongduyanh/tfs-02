package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"strconv"
	"strings"
)

func ReadNumber(r io.Reader) ([]float64, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []float64
	for scanner.Scan() {
		x, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

//bubble sort
func bubbleSort(sli []float64) {
	len := len(sli)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if sli[j] > sli[j+1] {
				sli[j], sli[j+1] = sli[j+1], sli[j]
			}
		}
	}
	fmt.Println("After Bubble Sorting")
	for _, val := range sli {
		fmt.Println(val)
	}
}

//kiem tra so nguyen to
func checkPrime(sli []float64) {
	len := len(sli)
	for i := 0; i < len; i++ {
		if sli[i]-float64(int64(sli[i])) != 0 { //kiem tra 1 so khong phai la so tu nhien
			fmt.Println(sli[i], "khong la so nguyen to")
		} else {
			if big.NewInt(int64(sli[i])).ProbablyPrime(0) {
				fmt.Println(sli[i], "la so nguyen to")
			} else {
				fmt.Println(sli[i], "khong la so nguyen to")
			}
		}

	}
}

//kiem tra ton tai
func checkInit(sli []float64, n float64) (int, bool) {
	var a int //a la so lan n xuat hien
	for i := 0; i < len(sli); i++ {
		if sli[i] == n {
			a++
		}
	}
	if a == 0 {
		return a, false
	} else {
		return a, true
	}
}

func main() {
	content, err := ioutil.ReadFile("test.txt") //doc file test.txt
	if err != nil {
		log.Fatal(err)
	}
	content_string := string(content) //convert []byte sang string
	numbers, err := ReadNumber(strings.NewReader(content_string))
	fmt.Println(numbers, err)

	//tim so lon nhat
	max := numbers[0]
	for i := 0; i < len(numbers); i++ {
		if max < numbers[i] {
			max = numbers[i]
		}
	}
	fmt.Println("MAX =", max)

	//tim so nho nhat
	min := numbers[0]
	for i := 0; i < len(numbers); i++ {
		if min > numbers[i] {
			min = numbers[i]
		}
	}
	fmt.Println("MIN =", min)

	//tinh trung binh
	avg := numbers[0]
	sum := 0.0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
		avg = sum / (float64(i) + 1)
	}
	fmt.Println("AVG =", avg)

	//bubble sort
	bubbleSort(numbers)

	//kiem tra so nguyen to
	checkPrime(numbers)

	//kiem tra ton tai
	k, found := checkInit(numbers, 5)
	if !found {
		fmt.Println("Value not found in slice")
	}
	fmt.Println("5 xuat hien", k, "lan")

	//in ra file
	

}
