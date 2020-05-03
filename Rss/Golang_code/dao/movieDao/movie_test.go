package movieDao

import "testing"

func TestGetAllMovieCount(t *testing.T) {
	num, err := GetAllMovieCount()
	if err != nil {
		t.Error("err")
	}

	t.Log(num)
}


func TestGetMovieById(t *testing.T) {
	num, err := GetMovieById(1)
	if err != nil {
		t.Error("err")
	}

	t.Log(num)
}
