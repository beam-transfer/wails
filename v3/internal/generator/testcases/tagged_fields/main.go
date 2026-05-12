package main

import (
	_ "embed"
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
)

// Sample mirrors a protoc-gen-go-generated message, carrying both
// protobuf and json tags. The bindings generator's FieldNameTag option
// chooses which tag wins.
type Sample struct {
	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FullName string `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	IsActive bool   `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Status   int32  `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`

	// PlainGoField has no protobuf tag; its json tag is used in both modes.
	PlainGoField string `json:"plain_go_field,omitempty"`
}

// GreetService is great
type GreetService struct{}

// Describe returns a Sample value.
func (*GreetService) Describe() Sample {
	return Sample{}
}

func main() {
	app := application.New(application.Options{
		Services: []application.Service{
			application.NewService(&GreetService{}),
		},
	})

	app.Window.New()

	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
