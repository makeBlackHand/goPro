package testExercise

import "testing"

func TestStore(t *testing.T) {
	monster := &Monster{
		Name:  "牛魔王",
		Age:   90,
		Skill: "哈哈哈哈",
	}
	res := monster.store()
	if res != true {
		t.Fatalf("store failed")
	}
	t.Logf("store success")
}

func TestRestore(t *testing.T) {
	var monster = &Monster{}
	res := monster.restore()
	if res != true {
		t.Fatalf("restore failed")
	}
	if monster.Name != "牛魔王" {
		t.Fatalf("restore failed")
	}
	t.Logf("restore success")
}
