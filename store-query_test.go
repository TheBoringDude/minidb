package minidb

import (
	"reflect"
	"sort"
	"testing"
)

func TestGetQuery(t *testing.T) {
	defer cleanFileAfter("getstore.json", t)

	db := NewStore("getstore.json")
	db.Set("hello", "world")

	if db.GetString("hello") != "world" {
		t.Fatal("`hello` key is not equal to world")
	}
}

func TestRemoveQuery(t *testing.T) {
	filename := "removestore.json"

	defer cleanFileAfter(filename, t)

	db := NewStore(filename)
	db.Set("value", false)
	db.Set("string", "123")

	err := db.Remove("value")
	if err != nil {
		t.Fatal("key is not removed")
	}
}

func TestUpdateQuery(t *testing.T) {
	filename := "updatestore.json"

	defer cleanFileAfter(filename, t)

	db := NewStore(filename)
	db.Set("value", false)
	db.Set("string", "123")

	db.Update("value", true)
	if db.GetBool("value") != true {
		t.Fatal("update is not working ")
	}
}

func TestFindKeyQuery(t *testing.T) {
	filename := "findstore.json"

	defer cleanFileAfter(filename, t)

	db := NewStore(filename)
	db.Set("hello", "world")
	db.Set("hellox", "x hello")
	db.Set("hell", false)
	db.Set("sample", false)
	db.Set("number", 100)

	expected := []string{"hell", "hello", "hellox"}
	r := db.FindKey("hell")
	sort.Strings(r)

	if !reflect.DeepEqual(expected, r) {
		t.Fatal("values returned from findkey are not equal")
	}

}
