package packages

import "testing"

func TestSearch(t *testing.T) {
    address := "http://128.0.255.10:9200"
    Search(address, "")
}