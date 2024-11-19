package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("urls.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
	repos := make(Repos, len(lines))
	for i, line := range lines {
		spt := strings.Split(line, " ")
		url := spt[0]
		star, _ := strconv.Atoi(spt[1])
		repos[i] = Repo{Url: url, Star: star}
	}

	sort.Sort(repos)

	for _, repo := range repos {
		fmt.Println("| ", repo.Url, " | ", repo.Star, " |")
	}
}

type Repo struct {
	Url  string
	Star int
}

type Repos []Repo

func (r Repos) Len() int {
	return len(r)
}

func (r Repos) Less(i, j int) bool {
	return r[i].Star > r[j].Star
}

func (r Repos) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
