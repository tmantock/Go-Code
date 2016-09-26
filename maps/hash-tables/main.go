package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	bs, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	str := string(bs)
	fmt.Println(str)

	words := make(map[string]string)

	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		words[sc.Text()] = ""
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input: ", err)
	}

	i := 0

	for k, _ := range words {
		fmt.Println(k)
		if i == 200 {
			break
		}
		i++
	}
}
