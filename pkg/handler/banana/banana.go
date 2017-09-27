package banana

import (
	"encoding/json"
	"fmt"

	"github.com/makarski/mybanana/pkg/db/banana"
)

type (
	// Banana describes a banana API model
	Banana struct {
		ID      uint64 `json:"id"`
		Name    string `json:"name"`
		SelfURL string `json:"self_url"`
	}
)

func (b *Banana) MarshalJSON() ([]byte, error) {
	b.SelfURL = fmt.Sprintf("https://mybanana.io/bananas/%d", b.ID)
	type alias Banana

	return json.Marshal(alias(*b))
}

func (b *Banana) fromDB(dbB *banana.Banana) {
	b.ID = dbB.ID
	b.Name = dbB.Name
}
