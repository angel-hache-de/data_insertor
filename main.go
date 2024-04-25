package main

import (
	"data_inserter/config"
	"data_inserter/store"
	"fmt"
)

const FILE_NAME = "data.xlsx"

var actualStore store.IStore

func main() {
	envService := config.LoadENVService{}
	envService.NewLoadENV()
	actualStore = store.NewSqlServer()
	saveOneByOne()
	// saveOneMillion()
	// saveOneMillionBy100k()
	// saveOneMillionBy10k()
}

func saveOneByOne() {
	if err := InsertRegistriesOneByOne(); err != nil {
		fmt.Printf("Error: %+v \n", err)
	}
}

func saveOneMillion() {
	data := GetRegistries(30)
	if err := actualStore.SaveData(data); err != nil {
		fmt.Printf("Error: %+v \n", err)
		return
	}
}

func saveOneMillionBy100k() {
	data := GetRegistries(3)
	for i := 0; i < 10; i++ {
		if err := actualStore.SaveData(data); err != nil {
			fmt.Printf("Error: %+v \n", err)
			return
		}
	}
}

func saveOneMillionBy10k() {
	data := GetRegistries(0)[:10000]
	for i := 0; i < 100; i++ {
		if err := actualStore.SaveData(data); err != nil {
			fmt.Printf("Error: %+v \n", err)
			return
		}
	}
}
