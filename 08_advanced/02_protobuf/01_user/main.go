package main

import (
	"log"

	"google.golang.org/protobuf/proto"

	"github.com/JimLee0921/GolangTraining/08_advanced/02_protobuf/01_user/pb"
)

func main() {
	// 使用 protobuf 序列化和反序列化
	u := &pb.User{
		Id:      1,
		Name:    "Tom",
		Role:    pb.Role_ADMIN,
		Tags:    []string{"vip", "tester"},
		Profile: &pb.Profile{Email: "tom@example.com"},
	}

	// 序列化
	data, err := proto.Marshal(u)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newU := &pb.User{}
	err = proto.Unmarshal(data, newU)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)

	}
	// 被序列化的和反序列化后的实例，包含相同的数据
	if u.GetName() != newU.GetName() {
		log.Fatalf("data mismatch %q != %q", u.GetName(), newU.GetName())
	}

}
