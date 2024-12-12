package clog_test

import (
	"errors"
	"testing"

	"github.com/scalarorg/scalar-core/utils/clog"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

type TestStruct struct {
	Value int    `json:"value"`
	Text  string `json:"text"`
}

func TestColorLog(t *testing.T) {
	// Test basic logging
	clog.Green("This is an info message", map[string]interface{}{
		"key1": "value1",
		"key2": 123,
	})

	clog.Red("This is an error message", map[string]interface{}{
		"error_code": 500,
		"details":    "Something went wrong",
	})

	clog.Blue("This is a debug message", map[string]interface{}{
		"debug_info": "Testing",
		"value":      42.5,
	})

	clog.Yellow("This is a warning message", map[string]interface{}{
		"warning_type": "performance",
		"threshold":    95,
	})

	clog.Cyan("This is a trace message", map[string]interface{}{
		"trace_id": "123e4567-e89b-12d3-a456-426614174000",
	})

	clog.Green("This is a string message", "hello")
	clog.Red("This is an error message", errors.New("an error"))

	clog.Blue("This is a float message", 1.23)
	clog.Yellow("This is an int message", 42)
	clog.Cyan("This is a bool message", true)
	clog.Blue("This is a struct message", TestStruct{Value: 42, Text: "hello"})

	var chain nexus.ChainName = "btc"

	clog.Red("chain", chain)

	clog.Red("chain", &TestStruct{Value: 42, Text: "hello"})

	clog.Yellow("block_height_not_found", "block_hash:", 123)
	clog.Yellowf("User %s logged in from %s", "John", "192.168.1.1")
	clog.Yellowf("Test struct %+v", TestStruct{Value: 42, Text: "hello"})
}
