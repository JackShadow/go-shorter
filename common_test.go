package hash

import "testing"

func TestGetId(t *testing.T) {

	m := New("lraGYMz8DZ0Uqw7jAe1WVE9LQdCcN4FOgPXfkvHpI2SxhJRKB6yuTm5t3inobs")
	data := []int{1, 2, 3, 61, 62, 63, 64, 120, 140, 200, 1024, 51000, 51001, 99999999999999,}
	for _, val := range data {

		str := m.GetString(val)

		id := m.GetId(str)

		if id != val {
			t.Errorf("Incorect val %d in string %s", id, str)
		}
	}

	for i := 200; i < 7000; i++ {
		str := m.GetString(i)

		id := m.GetId(str)

		if id != i {
			t.Errorf("Incorect val %d in string %s", id, str)
		}
	}
}
