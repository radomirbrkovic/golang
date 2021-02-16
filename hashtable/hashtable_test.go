package hastable

import (
    "fmt"
    "testing"
)

func populateHashTable(count int, start int) *ValueHashTable {
    dict := ValueHashTable{}
    for i := start; i < (start + count); i++ {
        dict.Put(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
    }
    return &dict
}

func TestPut(t *testing.T) {
    dict := populateHashTable(3, 0)
    if size := dict.Size(); size != 3 {
        t.Errorf("wrong count, expected 3 and got %d", size)
    }
    dict.Put("key1", "value1") //should not add a new one, just change the existing one
    if size := dict.Size(); size != 3 {
        t.Errorf("wrong count, expected 3 and got %d", size)
    }
    dict.Put("key4", "value4") //should add it
    if size := dict.Size(); size != 4 {
        t.Errorf("wrong count, expected 4 and got %d", size)
    }
}

func TestRemove(t *testing.T) {
    dict := populateHashTable(3, 0)
    dict.Remove("key2")
    if size := dict.Size(); size != 2 {
        t.Errorf("wrong count, expected 2 and got %d", size)
    }
}