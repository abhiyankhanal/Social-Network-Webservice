package utilities

import "testing"

func TestGreet(t *testing.T) {
	expected := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJ1c2VybmFtZSI6InRlc3QtdG9rZW4ifQ.AbumL1-fVQj8ISRKhkwTf0VAWQPk_F2aMXh-QfsewAc"
	actual := CreateToken("test-token")
	if actual != expected {
		t.Errorf("expected %s, but was %s", expected, actual)
	}
}
