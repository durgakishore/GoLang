package main

import (
	simplepb "PACKT/simple"
	complexpb "Packt/complex"
	enumpb "Packt/enum-example"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	sn := doSimple()
	writeToFile("simple.bin", sn)

	s := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", s)
	smJSON := toJSON(s)
	fmt.Println("JSON String ", smJSON)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smJSON, sm2)
	fmt.Println("PB String ", sm2)

	doEnum()

	doComplex()
}

func doComplex() {

	cm := complexpb.ComplexMessage{
			OneDummy: &complexpb.DummyMessage{
				Id:   1,
				Name: "Kishore",
			},
			MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Dummy 1",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "Dummy 2",
			},
			&complexpb.DummyMessage{
				Id:   4,
				Name: "Dummy 3",
			},
		},
	}

	fmt.Println("Complex Message: ", cm)
}
func doEnum() {

	em := enumpb.EnumMessage{
		Id:           1,
		DayOfTheWeek: enumpb.DayOfTheWeek_THURSDAY,
	}

	fmt.Println("enum message ", em)
}
func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}

	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Couldn't convert JSON data to pb structure", err)
	}
}
func writeToFile(name string, pb proto.Message) error {

	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("can't serialize to byte", err)
		return err
	}
	if err := ioutil.WriteFile(name, out, 8644); err != nil {
		log.Fatalln("can't write to file ", err)
		return err
	}
	fmt.Println("Data has been writeen")
	return nil
}

func readFromFile(name string, pb proto.Message) error {
	in, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatalln("Can't read the file", err)
		return err
	}

	err1 := proto.Unmarshal(in, pb)
	if err1 != nil {
		log.Fatalln("Couldn't put bytes into protocol buffers struct ", err1)
		return err1
	}
	fmt.Println("readFromFile : ", pb)
	return nil
}
func doSimple() *simplepb.SimpleMessage {

	sn := &simplepb.SimpleMessage{
		Id:         10,
		IsSimple:   true,
		Name:       "Kishore",
		SimpleList: []int32{1, 2, 3, 4},
	}

	fmt.Println(sn)

	sn.Name = "Rishika"
	fmt.Println(sn)

	return sn
}
