package test

import (
	"resetgoapi.com/go-admin-api/utils/cypher"
	"testing"
)

func TestGenPassword(test *testing.T) {
	password, _ := cypher.HashPassword("123456")
	test.Logf(password)
}
