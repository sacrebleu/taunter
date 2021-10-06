package taunts

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type Taunt struct {
	Name string `json:"name"`
	Taunt string `json:"taunt"`
}

type Paths struct {
	Prefixes string
	Taunts string
	Verbs string
}

type Data struct {
	Prefixes []string
	Taunts []string
	Verbs []string
}


// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// LoadData reads files from a struct of Paths into a Data
func LoadData(paths Paths) Data {
	prefixes, err := readLines(paths.Prefixes)
	if err != nil {
		log.Fatalf("readLines: %s", err)
		os.Exit(2)
	}
	log.Printf("Loaded %d prefixes\n", len(prefixes))
	verbs, err := readLines(paths.Verbs)
	if err != nil {
		log.Fatalf("readLines: %s", err)
		os.Exit(2)
	}
	log.Printf("Loaded %d verbs\n", len(verbs))

	taunts, err := readLines(paths.Taunts)
	if err != nil {
		log.Fatalf("readLines: %s", err)
		os.Exit(2)
	}
	log.Printf("Loaded %d taunts\n", len(taunts))

	return Data{
		Prefixes: prefixes,
		Verbs: verbs,
		Taunts: taunts,
	}
}

// randS return a random string from source []string
func randS(source []string) string {
	return source[rand.Intn(len(source) - 1)]
}

// Generate a taunt targeting the specified target, using the reference data of type Data
func Generate(target string, data Data) Taunt {
	var taunt Taunt

	rand.Seed(time.Now().UnixNano())

	taunt.Name = target
	taunt.Taunt = fmt.Sprintf("%s %s %s %s", randS(data.Prefixes), target, randS(data.Verbs), randS(data.Taunts))

	return taunt
}
