package 单元测试

import "testing"

func HelloTom() string {
	return "Jerry"
}

func TestHelloTom(t *testing.T) {
	output := HelloTom()
	expectOutput := "Tom"
	if output != expectOutput {
		t.Errorf("Expected %s do not match actual %s", expectOutput, output)
	}
}

func TestMain(m *testing.M) {

}
