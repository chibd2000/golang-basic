package main

import "fmt"

type Values map[string][]string

func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

// Add adds the value to key.
// It appends to any existing values associated with key.
func (v Values) Add(key, value string) {
	defer func() {
		if t := recover(); t != nil {
			fmt.Println(t)
		}
	}()

	v[key] = append(v[key], value)
}

func main() {
	m := Values{"lang": {"en"}} // direct construction
	m.Add("item", "1")
	m.Add("item", "2")
	m.Add("asd", "1")

	p := m

	m = nil

	fmt.Println(p)
	fmt.Println(m)

}
