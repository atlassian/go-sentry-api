package datatype

import (
	"reflect"
	"testing"
)

func TestGetMapping(t *testing.T) {
	t.Run("Test a known good type: Message", func(t *testing.T) {
		name, inter := GetMapping("message")

		if reflect.TypeOf(inter) != reflect.TypeOf(Message{}) {
			t.Error("Returned interface is not a message")
		}

		if name != "message" {
			t.Error("Message as not sent back as message. I dunno how...")
		}
	})

	t.Run("Test nil for unknown datatype", func(t *testing.T) {
		name, datatype := GetMapping("sillyclowntown")
		if datatype != nil {
			t.Error("Should have been nil for sillyclowntown")
		}
		if name != "sillyclowntown" {
			t.Error("Name should have been sillyclowntown")
		}
	})
}
