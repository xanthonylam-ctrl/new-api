package controller

import "testing"

func TestPQAPIHMAC(t *testing.T) {
	got := pqapiHMAC("secret", "1700000000\nnonce\n{\"amount\":3990}")
	want := "9ea50b93cbbdcc31e1bc859e072100278e34115f745c1c68d983e318c71a4957"
	if got != want {
		t.Fatalf("unexpected signature: got %s want %s", got, want)
	}
}
