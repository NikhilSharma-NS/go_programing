package exampletabletest

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

// These constant gives us checkboxes for visualization.
const (
	succeed = "\u2713"
	failed  = "\u2717"
)

func TestTable(t *testing.T) {
	// This table is a slice of anonymous struct type. It is the URL we are gonna call and
	// statusCode are what we expect.
	tests := []struct {
		url        string
		statusCode int
	}{
		{"https://www.google.com/", http.StatusOK},
		{"http://rss.cnn.com/rss/cnn_topstorie.rss", http.StatusNotFound},
	}

	t.Log("Given the need to test downloading different content.")
	{
		for i, tt := range tests {
			t.Logf("\tTest: %d\tWhen checking %q for status code %d", i, tt.url, tt.statusCode)
			{
				resp, err := http.Get(tt.url)
				if err != nil {
					t.Fatalf("\t%s\tShould be able to make the Get call : %v", failed, err)
				}
				t.Logf("\t%s\tShould be able to make the Get call.", succeed)

				defer resp.Body.Close()

				if resp.StatusCode == tt.statusCode {
					t.Logf("\t%s\tShould receive a %d status code.", succeed, tt.statusCode)
				} else {
					t.Errorf("\t%s\tShould receive a %d status code : %v", failed, tt.statusCode, resp.StatusCode)
				}
			}
		}
	}
}

func TestTable2(t *testing.T) {
	x := []struct {
		name string
	}{
		{"jayes"}, {"rames"},
	}

	{
		for i, v := range x {
			fmt.Println(i, v)
			if strings.Compare(v.name, "jayes") == 0 || strings.Compare(v.name, "rames") == 0 {
				t.Logf("passed case:  %d %v", i, succeed)
			} else {
				t.Errorf("failed case: %d %v", i, failed)
			}
		}
	}

}
