package jsonw

import "testing"

func TestUnmarshal(t *testing.T) {
	var s struct {
		Name      string   `json:"name"`
		Age       int      `json:"age"`
		IsHuman   bool     `json:"is_human"`
		Languages []string `json:"languages"`
	}

	t.Run("success", func(t *testing.T) {
		data := `
{
	"name": "gorilla",
	"age": 26,
	"is_human": false,
	"languages": [
		"Go",
		"PHP",
		"Java",
		"Java Script",
		"HTML"
	]
}`

		b := []byte(data)

		if err := Unmarshal(b, &s); err != nil {
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	t.Run("faild when invalid byte", func(t *testing.T) {
		data := `
{
	"name": "gorilla,
	"age": 26,
	"is_human": false,
	"languages": [
		"Go",
		"PHP",
		"Java",
		"Java Script",
		"HTML"
	]
}`

		b := []byte(data)

		if err := Unmarshal(b, &s); err != nil {
			t.Log(err)
		} else {
			t.Fatal("no error")
		}
	})
}
