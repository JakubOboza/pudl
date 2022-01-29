package pudl

import (
	"testing"
)

func TestPudlData(t *testing.T) {

	pd := PudlData{
		Config: struct{ TopCount int }{
			TopCount: 3,
		},
		Data: []Pudl{
			{
				Title: "Wow 1",
				Url:   "https://pudelek.pl/lolz/1",
			},
			{
				Title: "Wow 2",
				Url:   "https://pudelek.pl/lolz/2",
			},
		},
	}

	ser, err := pd.ToYaml()

	if err != nil {
		t.Fatal(err)
	}

	loaded, err := LoadDataFromYaml(ser)

	if loaded.Config.TopCount != 3 {
		t.Errorf("Expected TopCount to be 3 but got '%d'", loaded.Config.TopCount)
	}

	if loaded.Data[1].Title != "Wow 2" {
		t.Errorf("Expected '%s' but got '%s'", "Wow 2", loaded.Data[1].Title)
	}
}
