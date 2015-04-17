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

  ddbSave := DwarfDBLoad("db.dwarf", true)
  ddbSave.Set("x", "v")
  ddbSave.Set("-", -278900000000000.0)
  ddbSave.Set("1", 1.0)
  ddbSave.Set("2", 2.2)
  ddbSave.Set("f", false)
  ddbSave.Set("n", nil)
  ddbSave.Set("t", true)
  ddbSave.Set("u", "日本語の手紙をテスト")
  ddbSave.Rem("x")

  ddbLoad := DwarfDBLoad("db.dwarf", false)

  values := ddbLoad.GetAll()
  expected := []string{"-", "1", "2", "f", "n", "t", "u"}

  if StrsEquals(values, expected) == false {
    t.Errorf("GetAll: %s != %s", values, expected)
  }

  _, err := ddbLoad.Get("x")
  if err == nil {
    t.Errorf("The key 'x' could not exist")
  }

  for _, test := range tests {
    if value, _ := ddbLoad.Get(test.key); value != test.value {
      t.Errorf("Key %s: %s != %s", test.key, value, test.value)
    }
  }

  ddbSave.DelDB()
  ddbDel := DwarfDBLoad("db.dwarf", false)

  if ddbDel.Len() > 0 {
    t.Errorf("The database should be empty")
  }
}
