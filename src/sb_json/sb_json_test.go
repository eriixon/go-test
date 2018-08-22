package sb_json

import "testing"

func TestParseJSON(t *testing.T){
	var tests = []struct {
		c string; t float32; s string
	}{
		{"Boise", 300.83, string(`{"coord":{"lon":-116.2,"lat":43.62},"weather":[{"id":721,"main":"Haze","description":"haze","icon":"50n"}],"base":"stations","main":{"temp":300.83,"pressure":1013,"humidity":26,"temp_min":300.15,"temp_max":302.15},"visibility":12874,"wind":{"speed":3.6,"deg":320},"clouds":{"all":1},"dt":1534815300,"sys":{"type":1,"id":931,"message":0.0035,"country":"US","sunrise":1534856184,"sunset":1534905510},"id":5586437,"name":"Boise","cod":200}`)},
		{"Nampa", 299.83, string(`{"coord":{"lon":-116.2,"lat":43.62},"weather":[{"id":721,"main":"Haze","description":"haze","icon":"50n"}],"base":"stations","main":{"temp":299.83,"pressure":1013,"humidity":26,"temp_min":300.15,"temp_max":302.15},"visibility":12874,"wind":{"speed":3.6,"deg":320},"clouds":{"all":1},"dt":1534815300,"sys":{"type":1,"id":931,"message":0.0035,"country":"US","sunrise":1534856184,"sunset":1534905510},"id":5586437,"name":"Nampa","cod":200}`)},
	}

	for _, c := range tests{
		city, temp := ParseJSON(c.s)
		if (city != c.c && temp != int(c.t)) {
			t.Errorf("ParseJSON for %v and %v, want %v and %v", city, temp, c.c, c.t)
		}

	}
}

func BenchmarkParseJSON(b *testing.B){
	var tests = []struct {
		c string; t float32; s string
	}{
		{"Boise", 300.83, string(`{"coord":{"lon":-116.2,"lat":43.62},"weather":[{"id":721,"main":"Haze","description":"haze","icon":"50n"}],"base":"stations","main":{"temp":300.83,"pressure":1013,"humidity":26,"temp_min":300.15,"temp_max":302.15},"visibility":12874,"wind":{"speed":3.6,"deg":320},"clouds":{"all":1},"dt":1534815300,"sys":{"type":1,"id":931,"message":0.0035,"country":"US","sunrise":1534856184,"sunset":1534905510},"id":5586437,"name":"Boise","cod":200}`)},
		{"Nampa", 299.83, string(`{"coord":{"lon":-116.2,"lat":43.62},"weather":[{"id":721,"main":"Haze","description":"haze","icon":"50n"}],"base":"stations","main":{"temp":299.83,"pressure":1013,"humidity":26,"temp_min":300.15,"temp_max":302.15},"visibility":12874,"wind":{"speed":3.6,"deg":320},"clouds":{"all":1},"dt":1534815300,"sys":{"type":1,"id":931,"message":0.0035,"country":"US","sunrise":1534856184,"sunset":1534905510},"id":5586437,"name":"Nampa","cod":200}`)},
	}

	for _, c := range tests{
		ParseJSON(c.s)
	}
}

