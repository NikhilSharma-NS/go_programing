package example_basic

import (
	"net/http"
	"testing"
)

func TestDownload(t *testing.T) {
	url := "https://www.google.com/"
	statusCode := 200

	resp, err := http.Get(url)

	if err != nil {
		t.Fatalf("unable to issue Get on url%s %s", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != statusCode {
		t.Log("exp", statusCode)
		t.Log("Goot", resp.StatusCode)
		t.Fatal("status code don't match")
	} else {
		t.Logf("Should receive a %d status code.", statusCode)
	}

}
