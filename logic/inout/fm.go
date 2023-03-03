package inout

type FM struct {
	Path  string `json:"path"  query:"path"`
	Match string `json:"match" query:"match"`
	Limit int    `json:"limit" query:"limit"`
}
