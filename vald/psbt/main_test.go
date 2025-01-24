package psbt_test

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	psbtMgr "github.com/scalarorg/scalar-core/vald/psbt"
)

var mockMgr = &psbtMgr.Mgr{}

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env.test")
	if err != nil {
		log.Fatal("Error loading .env.test file")
	}
	os.Exit(m.Run())
}
