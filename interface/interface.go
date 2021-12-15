package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//	store input argumemts:
//	capacity of backpack
//	weight and value of items as slice of maps
type InputArguments struct {
	backpack_capacity int
	item_map          map[int]Item
}
type Item struct {
	weight int
	value  int
}

func getArguments() InputArguments {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Welcome to Backpacker!\n")
	fmt.Print("Let's help you decide what to bring\n")
	capacity := getCapacity(reader)
	num_items := getNumItems(reader)
	item_info := getItem(reader, num_items)
	info := InputArguments{capacity, item_info}
	return info
}

func getCapacity(reader *bufio.Reader) int {
	fmt.Print("First. How much weight can your bag carry?: ")
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return getCapacity(reader)
	}
	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	int_input, int_err := strconv.Atoi(input)
	if (int_err != nil) || (int_input < 1) || (int_input > 100) {
		fmt.Printf("%q :( Please enter an integer between 1-100.\n", input)
		return getCapacity(reader)
	}
	fmt.Printf("Cool, we've got %s to fill.\n", strconv.Itoa(int_input))
	return int_input
}

func getNumItems(reader *bufio.Reader) int {
	fmt.Print("How many items are you considering bringing? ")
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return getNumItems(reader)
	}
	input = strings.TrimSuffix(input, "\n")
	int_input, int_err := strconv.Atoi(input)
	if (int_err != nil) || (int_input < 1) || (int_input > 100) {
		fmt.Printf("%q Please enter an integer between 1-100.\n", input)
		return getNumItems(reader)
	}
	fmt.Printf("Cool, we've got %s items to consider.\n", input)
	return int_input
}

func getItem(reader *bufio.Reader, num_items int) map[int]Item {
	data := make(map[int]Item)
	i := 1
	for i <= num_items {
		fmt.Printf("Please enter the weight for item %s: ", strconv.Itoa(i))
		// ReadString will block until the delimiter is entered
		w_input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
		}
		w_input = strings.TrimSuffix(w_input, "\n")
		w_int_input, w_int_err := strconv.Atoi(w_input)
		if (w_int_err != nil) || (w_int_input < 1) || (w_int_input > 100) {
			fmt.Printf("%q Please enter an integer between 1-100.\n", w_input)
		}
		fmt.Printf("Please enter the value for item %s: ", strconv.Itoa(i))
		// ReadString will block until the delimiter is entered
		v_input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
		}
		v_input = strings.TrimSuffix(v_input, "\n")
		v_int_input, v_int_err := strconv.Atoi(v_input)
		if (v_int_err != nil) || (v_int_input < 1) || (v_int_input > 100) {
			fmt.Printf("%q Please enter an integer between 1-100.\n", v_input)
		}
		data[i] = Item{w_int_input, v_int_input}
		i = i + 1
	}
	return data
}

func main() {
	info := getArguments()
	result := Solve(info)
	fmt.Println(result)
}
