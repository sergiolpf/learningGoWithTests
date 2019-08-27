package structsMethodsInter

import "testing"

func TestPerimeter(t *testing.T){
	testRect := Rectangle{10.0, 10.0}
	got := testRect.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)

	}
}


func TestArea(t *testing.T){

	checkArea := func(t *testing.T, shape Shape, want float64){
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("calculating area for rectangle", func(t *testing.T){
		testRect := Rectangle{10.0, 10.0}
		checkArea(t, testRect, 100.0)
	})

	t.Run("calculating area of a circle", func(t *testing.T){
		testCircle := Circle{10.0}
		checkArea(t, testCircle, 314.1592653589793)

	})

}
