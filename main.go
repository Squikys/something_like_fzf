package main

import (
	"fmt"
	//	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/mattn/go-tty"
)

var name string
var dir string
var total int = 0
var display_limit int = 10
var founditems int = 0
var pos = display_limit
var files []string
var current_file string = ""

func backspace() {
	if len(name) > 0 {
		temp := strings.Split(name, "")
		temp2 := ""
		for i := 0; i < len(name)-2; i++ {
			temp2 += temp[i]
		}
		name = temp2
	}
}

func render2(items chan int) {
	var iteration int = 0
	itemcount := 0
	count := 0
	err := filepath.WalkDir(dir,
		func(path string, info os.DirEntry, err error) error {
			if !info.IsDir() && !os.IsPermission(err) && strings.Contains(info.Name(), name) {
				itemcount++
			}
			if !info.IsDir() && !os.IsPermission(err) && strings.Contains(info.Name(), name) && name != "" && pos >= count && count >= pos-display_limit {
				if iteration == 0 {
					fmt.Printf("\033[%d;0H[ -> ]%s %d\n", display_limit-iteration, path, founditems)
					current_file = path
				} else {

					fmt.Printf("\033[%d;0H[    ]%s %d\n", display_limit-iteration, path, count)
				}
				iteration++
				total += 1
			}
			count++
			return nil
		})

	if err != nil {
		log.Println(err, "this is error")
	}
	items <- itemcount

}

func render() {
	var iteration int = 0
	founditems = 0
	count := 0
	err := filepath.Walk("/home/kaustav/codes/fuzzy_search/test_dir",
		func(path string, info os.FileInfo, err error) error {
			if !os.IsPermission(err) && strings.Contains(info.Name(), name) {
				founditems++
			}
			if !os.IsPermission(err) && strings.Contains(info.Name(), name) && name != "" && pos >= count && count >= pos-display_limit {
				if iteration == 0 {
					fmt.Printf("\033[%d;0H[ -> ]%s %d\n", display_limit-iteration, path, founditems)
					current_file = path
				} else {

					fmt.Printf("\033[%d;0H[    ]%s %d\n", display_limit-iteration, path, count)
				}
				iteration++
				total += 1
			}
			count++
			return nil
		})

	if err != nil {
		log.Println(err, "this is error")
	}

}
func inputs() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()
	r, _ := tty.ReadRune()
	//127 delete
	//this if statement deletes characters

	// Arrow keys are sequences of multiple runes
	switch r {
	case 13:
		fmt.Print("\033[2J")
		fmt.Println(dir)
		os.Exit(0)
	case 27: // ESC
		r2, _ := tty.ReadRune()
		if r2 == 91 { // '[' character
			r3, _ := tty.ReadRune()
			switch r3 {
			case 66:
				fmt.Print("\033[2J")
				if pos > display_limit {
					pos--
				}
			case 65:
				fmt.Print("\033[2J")
				if pos <= founditems {
					pos++
				}
			}
		}
	case 127:
		name = name + string(r)
		total = 0
		fmt.Print("\033[2J")
		backspace()
		fmt.Printf("\033[%d;0H%d\n", display_limit+3, founditems)
		fmt.Printf("\033[%d;0H%s\n", display_limit+4, string(name))
	default:
		name = name + string(r)
		total = 0
		fmt.Print("\033[2J")
		fmt.Printf("\033[%d;0H%d/%d\n", display_limit+3, founditems, total)
		fmt.Printf("\033[%d;0H%s\n", display_limit+4, string(name))
	}
	items := make(chan int)
	go render2(items)
	temp := <-items
	founditems = int(temp)

}

func main() {
	arg := os.Args
	switch len(arg) {
	case 1:
		dir, _ = os.Getwd()
	case 2:
		dir = arg[1]
	case 3:

		dir = arg[1]
		re := regexp.MustCompile("[0-9]+")

		i, err := strconv.Atoi(re.FindAllString(arg[2], -1)[0])
		if err != nil {
			panic(err)
		}
		display_limit = i
	}
	if len(arg) > 1 {

	} else {
	}
	for {
		inputs()
	}

}
