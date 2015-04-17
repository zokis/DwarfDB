package dwarfdb

import (
  "testing"
)

type Test struct {
  key   string
  value interface{}
}

func StrsEquals(a, b []string) bool {
  if len(a) != len(b) {
    return false
  }
  for i, v := range a {
    if v != b[i] {
      return false
    }
  }
  return true
}

func TestSomething(t *testing.T) {

  var tests = []Test{
    {"1", 1.0},
    {"2", 2.2},
    {"t", true},
    {"f", false},
    {"n", nil},
    {"u", "日本語の手紙をテスト"},
    {"-", -2.789e+14},
  }

  ddb_save := DwarfDBLoad("db.dwarf", true)
  ddb_save.Set("x", "v")
  ddb_save.Set("-", -278900000000000.0)
  ddb_save.Set("1", 1.0)
  ddb_save.Set("2", 2.2)
  ddb_save.Set("f", false)
  ddb_save.Set("n", nil)
  ddb_save.Set("t", true)
  ddb_save.Set("u", "日本語の手紙をテスト")
  ddb_save.Rem("x")

  ddb_load := DwarfDBLoad("db.dwarf", false)

  values := ddb_load.GetAll()
  expected := []string{"-", "1", "2", "f", "n", "t", "u"}

  if StrsEquals(values, expected) == false {
    t.Errorf("GetAll: %s != %s", values, expected)
  }

  _, err := ddb_load.Get("x")
  if err == nil {
    t.Errorf("The key 'x' could not exist")
  }

  for _, test := range tests {
    if value, _ := ddb_load.Get(test.key); value != test.value {
      t.Errorf("Key %s: %s != %s", test.key, value, test.value)
    }
  }

  ddb_save.DelDB()
  ddb_del := DwarfDBLoad("db.dwarf", false)

  if ddb_del.Len() > 0 {
    t.Errorf("The database should be empty")
  }
}
