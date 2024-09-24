package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type Store struct {
	data map[string]string
	//在 map[string]string 中：
	//
	//string：表示键的类型是字符串。
	//string：表示值的类型也是字符串。
	mu   sync.RWMutex
	sets map[string]map[string]bool
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
		sets: make(map[string]map[string]bool),
	}
}

func (s *Store) Load(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &s.data)
}
func (s *Store) Save(filename string) error {
	file, err := json.MarshalIndent(s.data, "", " ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, file, 0644)
}
func (s *Store) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}
func (s *Store) SetNx(key, value string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.data[key]; exists {
		return 0
	}
	s.data[key] = value
	return 1
}
func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, exists := s.data[key]
	return value, exists
}
func (s *Store) Del(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}
func (s *Store) SADD(setName string, value string) int {
	if _, exists := s.sets[setName]; !exists {
		s.sets[setName] = make(map[string]bool)
	}
	if s.sets[setName][value] {
		return 0
	}
	s.sets[setName][value] = true
	return 1
}
func showMenu() {
	fmt.Println("1.命令行")
	fmt.Println("2.程序说明")
	fmt.Println("3.退出程序")
}
func showhelp() {
	fmt.Println("1.就是让你个ldx输入点东西实现简单redis")
	fmt.Println("2.self")
	fmt.Println("3.看如此屎的代码看不下去了就溜走")
}
func main() {
	store := NewStore()
	const filename = "store.json"
	err := store.Load(filename)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Error loading data:", err)
		return
	}
	for {
		showMenu()
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var command string
			var key, value string
			fmt.Print(">")
			n, err := fmt.Scanf("%s %s %s", &command, &key, &value)
			if err := store.Save(filename); err != nil {
				fmt.Println("Error saving data:", err)
			}
			if err != nil && err.Error() != "expected newline" {
				fmt.Println("Error reading command:", err)
				continue
			}
			if n == 0 {
				fmt.Println("No input recived.")
				continue
			}
			switch command {
			case "SET":
				store.Set(key, value)
				store.Save(filename)
				fmt.Println("ok了")
			case "SETNX":
				result := store.SetNx(key, value)
				store.Save(filename)
				if result == 0 {
					fmt.Println("0")
				} else {
					fmt.Println("1")
				}
			case "GET":
				if val, exists := store.Get(key); exists {
					fmt.Println(val)
				} else {
					fmt.Println("not found")
				}
			case "DEL":
				store.Del(key)
				store.Save(filename)
				fmt.Println("ok")
			default:
				fmt.Println("Unknown command")
			case "SADD":
				Result1 := store.SADD(key, value)
				fmt.Println(Result1)
			}

		case 2:
			showhelp()
		case 3:
			fmt.Println("Exiting program.")
			return
		}
	}
}
