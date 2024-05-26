package resultmodels

import "final/entities/resultset/resultsetmodels"

type Result struct {
	Status bool                      `json:"status"`
	Data   resultsetmodels.ResultSet `json:"data"`
	Error  string                    `json:"error"`
}
