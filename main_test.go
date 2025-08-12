package main

import (
    "os"
    "testing"
    "sauce-service/src/utils"
)

func TestMain(m *testing.M) {
    utils.SetupTestDB()
    code := m.Run()
    os.Exit(code)
}
