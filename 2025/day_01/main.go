package main

import (
	_ "bufio" // For reading file line by line
	"fmt"
	_ "os" // For opening files
	_ "strconv"
	_ "strings"
)

type Rotation struct {
	Direction string // "L" or "R"
	Distance  int    // How many clicks
}

// Slices: []Rotation for storing rotations
// For loops:
// for scanner.Scan() { ... } for reading file
// for _, rotation := range rotations { ... } for iterating
// Modulo arithmetic: For circular dial (remember negative numbers!)

func parseRotation(line string) Rotation {
	// Hint: Use strings.TrimSpace(), string indexing [0], and strconv.Atoi()
	// Takes a line like "L68"
	// This might give you negative results:
	// newPos := (currentPos - distance) % 100

	// To handle negatives properly in a 0-99 range:
	// newPos := ((currentPos - distance) % 100 + 100) % 100
	// Returns a Rotation struct
	return Rotation{}
}

func applyRotation(curPos int, rotation Rotation) int {
	// Takes current dial position and a rotation
	// Returns new position
	// Hint: Use modulo operator % to handle wrapping around (0-99 range)
	// Remember: Left decreases, Right increases
	return 0
}

func readRotations(filename string) ([]Rotation, error) {
	// Opens the file
	// Error handling: Check if err != nil after file operations

	// Reads line by line using bufio.Scanner
	// Returns slice of Rotation structs
	// Hint: Use bufio.NewScanner(file), scanner.Scan(), scanner.Text()
	return nil, nil
}

func main() {
	// 1. Read rotations from "input.txt"
	rotations, err := readRotations("input.txt")
	if err != nil {
		panic(err)
	}

	// 2. Initialize starting position (50)
	position := 50

	// 3. Initialize counter for zeros
	zeroCount := 0

	// 4. Loop through each rotation:
	//    - Apply rotation
	//    - Check if position is 0
	//    - Increment counter if yes

	// 5. Print the password (total count)

	for i := 1; i <= 5; i++ {
		// TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
		// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
		fmt.Println("i =", 100/i)
	}
}
