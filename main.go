/*
 * Author Zoe Wrubleski
 * Version 1.0.0
 * Date 2025-11-19
 * Fileoverview This program displays the happy birthday song with the users name
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// set variables
	var userName string

	reader := bufio.NewReader(os.Stdin)

	// prompt user for their name
	fmt.Print("What is your name? ")
	userName, _ = reader.ReadString('\n')
	userName = strings.TrimSpace(userName)

	// display happy birthday
	fmt.Println()
	fmt.Println("Happy birthday to you!")
	fmt.Println("Happy birthday to you!")
	fmt.Println("Happy birthday dear,", userName, "!")
	fmt.Println("Happy birthday to you!")

	fmt.Println("\nDone.")
}
