package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"

	"github.com/vmihailenco/msgpack"
)

type s struct {
	data map[string]interface{}
}

type Person struct {
	Name   string
	Age    int
	Gender string
}

// Go语言中的json包在序列化空接口存放的数字类型（整型、浮点型等）都序列化成float64类型
func jsonDemo() {
	var s1 = s{
		data: make(map[string]interface{}, 8),
	}
	s1.data["count"] = 1
	s1.data["count2"] = 2
	ret, err := json.Marshal(s1.data)
	if err != nil {
		fmt.Println("marshal failed", err)
	}
	fmt.Printf("%#v\n", string(ret))
	var s2 = s{
		data: make(map[string]interface{}, 8),
	}
	err = json.Unmarshal(ret, &s2.data)
	if err != nil {
		fmt.Println("unmarshal failed", err)
	}
	fmt.Println(s2)
	for _, v := range s2.data {
		fmt.Printf("value:%v, type:%T\n", v, v)
	}
}

// 标准库gob是golang提供的“私有”的编解码方式，它的效率会比json，xml等更高，特别适合在Go语言程序间传递数据

func gobDemo() {
	var s1 = s{
		data: make(map[string]interface{}, 8),
	}
	s1.data["count"] = 1
	// encode
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(s1.data)
	if err != nil {
		fmt.Println("gob encode failed, err:", err)
		return
	}
	b := buf.Bytes()
	fmt.Println(b)
	var s2 = s{
		data: make(map[string]interface{}, 8),
	}
	// decode
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	err = dec.Decode(&s2.data)
	if err != nil {
		fmt.Println("gob decode failed, err", err)
		return
	}
	fmt.Println(s2.data)
	for _, v := range s2.data {
		fmt.Printf("value:%v, type:%T\n", v, v)
	}
}

func main() {
	// jsonDemo()
	// gobDemo()

	// MessagePack是一种高效的二进制序列化格式。它允许你在多种语言(如JSON)之间交换数据。但它更快更小。
	p1 := Person{
		Name:   "沙河娜扎",
		Age:    18,
		Gender: "男",
	}
	// marshal
	b, err := msgpack.Marshal(p1)
	if err != nil {
		fmt.Printf("msgpack marshal failed,err:%v", err)
		return
	}

	// unmarshal
	var p2 Person
	err = msgpack.Unmarshal(b, &p2)
	if err != nil {
		fmt.Printf("msgpack unmarshal failed,err:%v", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2) // p2:main.Person{Name:"沙河娜扎", Age:18, Gender:"男"}
}
