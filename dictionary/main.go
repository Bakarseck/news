package main

import (
	"dictionary/dictionary"
	"flag"
	"fmt"
	"os"
)

func main() {

	action := flag.String("action", "list", "Action to perform on the dictionnary")

	d, err := dictionary.New("./badger")
	HandleErr(err)
	defer d.Close()

	flag.Parse()

	switch *action {
	case "list":
		actionList(d)
	case "add":
		actionAdd(d, flag.Args())
	case "remove":
		actionRemove(d, flag.Args())
	case "define":
		actionDefine(d, flag.Args())	
	default:
		fmt.Println("unknow action:%v", *action)
	}
}

func actionList(d *dictionary.Dictionary) {
	words, entries, err := d.List()
	HandleErr(err)
	fmt.Println("Dictionary content")
	for _, word := range words {
		fmt.Println(entries[word])
	}
}

func actionAdd(d *dictionary.Dictionary, args []string) {
	word := args[0]
	definition := args[1]
	err := d.Add(word, definition)
	HandleErr(err)
	fmt.Printf("%v added to the dictionnary\n", word)
}

func actionDefine(d *dictionary.Dictionary, args []string) {
	word := args[0]
	entry, err := d.Get(word)
	HandleErr(err)
	fmt.Println(entry)
}

func actionRemove(d *dictionary.Dictionary, args []string) {
	word := args[0]
	err := d.Remove(word)
	HandleErr(err)
	fmt.Printf("%v deleted to the dictionnary\n", word)
}

func HandleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionary error:%v\n", err.Error())
		os.Exit(1)
	}
}
