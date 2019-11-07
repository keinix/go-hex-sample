package login

import "testing"

func TestPasswordMatch(t *testing.T) {
	tables := []struct {
		name      string
		password1 string
		password2 string
		expected  bool
	}{
		{name: "same passwords match", password1: "zelda123", password2: "zelda123", expected: true},
		{name: "different passwords don't match", password1: "zelda123", password2: "link123", expected: false},
	}
	for _, table := range tables {
		hash, err := newPasswordHash(table.password1)
		if err != nil {
			t.Errorf("error in test case: %v", err)
			t.FailNow()
		}
		match, err := checkPlainTextMatchesHash(table.password2, hash)
		if err != nil {
			t.Errorf("error in test case: %v", err)
			t.FailNow()
		}
		if match != table.expected {
			t.Errorf("%v, password1 = %v password2 = %v expectedMatch = %v got = %v",
				table.name, table.password1, table.password2, table.expected, match)
		}
	}
}

func BenchmarkNewPasswordHash(b *testing.B) {
	_, _ = newPasswordHash("helloWorld1%^&#)NIOEIFH))(@*!#&*)(#*JDJOSOIJFIABUASODIIBASDI)@(*)*$#(*")
}
