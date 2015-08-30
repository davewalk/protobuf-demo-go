package main

import (
	"encoding/json"
	"fmt"
	"github.com/davewalk/protobuf-demo-go/demo"
	"github.com/golang/protobuf/proto"
	"log"
)

func getdog() (data []byte, err error) {
	dog := &demo.Dog{
		Barks:    *proto.Bool(true),
		Yearsold: *proto.Int32(6),
	}

	data, err = proto.Marshal(dog)

	return
}

func getcat() (data []byte, err error) {
	dave := &demo.Cat_Parent{
		Name:  *proto.String("Dave"),
		Email: *proto.String("doesntmatter@gmail.com"),
	}
	cat := &demo.Cat{
		Name:    *proto.String("Rocky"),
		Age:     *proto.Int32(11),
		Parents: []*demo.Cat_Parent{dave},
	}

	data, err = proto.Marshal(cat)

	return
}

func main() {
	data, _ := getdog()
	cat := &demo.Cat{}
	err := proto.Unmarshal(data, cat)
	if err != nil {
		log.Fatal("problem unmarshaling: ", err)
	}

	fmt.Println(cat.String())

	var jsondata []byte
	jsondata, err = json.Marshal(cat)
	fmt.Println(string(jsondata))

}
