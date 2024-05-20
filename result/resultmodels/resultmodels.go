package resultmodels

import "final/resultset/resultsetmodels"

type Result struct {
	Status bool                      `json:"status"`
	Data   resultsetmodels.ResultSet `json:"data"`
	Error  string                    `json:"error"`
}
