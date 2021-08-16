package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sum-metric-service-problem/src/internal/memstorage"
)

func main() {
	store := memstorage.NewStorage()
	defer store.Purge()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			continue
		}

		strs := strings.Split(s, " ")
		if strs[0] == "print" {
			fmt.Printf("Current store is: %v\n", store)
			continue
		}

		if len(strs) > 1 && strs[1] == "" {
			continue
		}

		backet := strs[0]
		if value, err := strconv.Atoi(strs[1]); err == nil {
			store.AddToBacket(backet, value)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("exit")
	}
}
