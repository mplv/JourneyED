package eventhandler

import (
	"bufio"
	"encoding/json"
	"event"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestHandleEvent(t *testing.T) {
	var e event.Event
	handleLine := HandleLine{Handlers: make(map[string]func(*event.Event, bool) string)}
	HandleEventInit(&handleLine)
	f, err := os.Open("test/Journal.160726141417.01.log")
	if err != nil {
		t.Fatal("File open", err)
	}
	r := bufio.NewReader(f)

	for err != io.EOF {
		b, err := r.ReadBytes('\n')
		if err != nil && err != io.EOF {
			t.Fatal("Read Bytes", err)
		}
		if err != nil && err == io.EOF {
			break
		}
		if !testing.Short() {
			fmt.Println(string(b))
		}
		err = json.Unmarshal(b, &e)
		if err != nil {
			t.Fatal("Unmarshal", err)
		}
		handleLine.HandleEvent(&e)
	}
}
