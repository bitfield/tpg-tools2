package battery

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
)

type Battery struct {
	Name             string
	ID               int64
	ChargePercent    int
	TimeToFullCharge string
	Present          bool
}

func (b Battery) ToJSON() string {
	output, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	pretty := jsontext.Value(output)
	pretty.Indent()
	return string(pretty)
}
