package main

import (
	"MyProgram/trie"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

var t *trie.Trie

type school struct {
	Name     string
	City     string
	Province string
}

func init() {
	t = trie.NewTrie()
}

func loadSchool(filename string) (err error) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			return err
		}

		var s school
		fields := strings.Fields(string(line))
		if len(fields) != 3 {
			continue
		}

		s.Name = fields[0]
		s.City = fields[1]
		s.Province = fields[2]

		t.Add(s.Name, s)
	}
	return
}

func loadBadWord(filename string) (err error) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		l, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			return err
		}
		line := string(l)
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		t.Add(line, line)
	}
	return
}

// 前缀搜索 输入前缀 输出结果
func prefixSearch(key string) string {
	res := t.PrefixSearch(key)
	result, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		return string(result)
	}
	return string(result)
}

func badWord(key string) (r string, find bool) {
	r, find = t.BadWord(key)
	return
}

func main() {

	err := loadSchool("data.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	res := prefixSearch("北")
	fmt.Println(res)

	// err := loadBadWord("CensorWords.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// r, b := badWord("金鳞岂是池中物, 一遇风云便化龙")
	// fmt.Println(r, b)
}
