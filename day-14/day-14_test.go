package main




import "testing"




func testAdd(t *testing.T)  {

	result := AddNum(2,10)

	expected_answer :=  12

	if result != expected_answer{
		t.Errorf("Add(2,3)= %d, expected answer = %d",result,expected_answer)
	}
	
}
