package main

import (
	"fmt"
	"strconv"
)

func main() {
	menu_of_cafe()
	BillCalculator()
}

func menu_of_cafe() {
	fmt.Println("  WELCOME TO GOLANG CAFE  ")
	fmt.Println("``````````````````````````")
	fmt.Println("CODE      ITEM          PRICE")
	fmt.Println("`````````````````````````````")
	fmt.Print("C         COFFEE          40/-\n" +
		"D          DOSA           80/-\n" +
		"T       TOMATO SOUP       20/-\n" +
		"I          IDLI           20/-\n" +
		"V          VADA           25/-\n" +
		"B      CHOLE BHATURE      30/-\n" +
		"P      PANEER PAKODA      30/-\n" +
		"M       MANCHURIAN        90/-\n" +
		"H      HAKKA NOODLE       70/-\n" +
		"F      FRENCH FRIES       30/-\n" +
		"J        JALEBI           30/-\n" +
		"L       LEMONADE          15/-\n" +
		"S      SPRING ROLL        20/-\n")
	fmt.Println("NOTE:- Enter END when order is complete.")
	fmt.Println("Please Give Your Order:- ")
}

func BillCalculator() {
	price := map[string]int{
		"C": 40,
		"D": 80,
		"T": 20,
		"I": 20,
		"V": 25,
		"B": 30,
		"P": 30,
		"M": 90,
		"H": 70,
		"F": 30,
		"J": 30,
		"L": 15,
		"S": 20,
	}
	total := 0
	Bill := 0
	for true {
		var quantity, item string
		fmt.Scan(&quantity)
		if quantity == "END" {
			fmt.Println("TOTAL EARNING of GOLANG CAFE is:- ", total)
			break
		}
		fmt.Scan(&item)
		q, err := strconv.Atoi(item)
		if quantity == "END" {
			fmt.Print(err)
		}
		Bill = q * price[quantity]
		fmt.Println("Amount of Items ordered:- ", Bill)
		total = total + Bill

	}
}
