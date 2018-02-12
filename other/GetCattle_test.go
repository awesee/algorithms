package other

import "testing"

//println(GetCattle(3)) // 1
//
//println(GetCattle(4)) // 2
//
//println(GetCattle(10)) // 15
//
//println(GetCattle(20)) // 346
//
//println(GetCattle(30)) // 8317

func TestGetCattle(t *testing.T) {

	m := map[uint]uint{
		3:  1,
		4:  2,
		10: 15,
		20: 346,
		30: 8317,
	}

	for k, v := range m {
		if GetCattle(k) != v {
			t.Errorf("GetCattle(%v) failed. Expected %v", k, v)
		}
	}

}
