package versions

type Support struct {
	Delve DelveVersion
	MinGo GoVersion
	MaxGo GoVersion
}

func (spt Support) Compat(gov GoVersion) bool {
	return spt.MinGo.Cmp(gov) <= 0 && 0 <= spt.MaxGo.Cmp(gov)
}

// Table show that which delve supports which go.
var Supports = []Support{
	// sort new to old.

	{ // latest version supports go1.19+
		Delve: Delve(1, 22, 1, ""),
		MinGo: Go(1, 19, 0, ""),
		MaxGo: Go(1, 22, 0, ""),
	},
	{
		Delve: Delve(1, 22, 0, ""),
		MinGo: Go(1, 19, 0, ""),
		MaxGo: Go(1, 22, 0, ""),
	},
	{
		Delve: Delve(1, 21, 2, ""),
		MinGo: Go(1, 19, 0, ""),
		MaxGo: Go(1, 21, 0, ""),
	},
	{
		Delve: Delve(1, 21, 1, ""),
		MinGo: Go(1, 19, 0, ""),
		MaxGo: Go(1, 21, 0, ""),
	},
	{ // latest version supports go1.18
		Delve: Delve(1, 21, 0, ""),
		MinGo: Go(1, 18, 0, ""),
		MaxGo: Go(1, 21, 0, ""),
	},
	{
		Delve: Delve(1, 20, 2, ""),
		MinGo: Go(1, 18, 0, ""),
		MaxGo: Go(1, 20, 0, ""),
	},
	{
		Delve: Delve(1, 20, 1, ""),
		MinGo: Go(1, 18, 0, ""),
		MaxGo: Go(1, 20, 0, ""),
	},
	{
		Delve: Delve(1, 20, 0, ""),
		MinGo: Go(1, 18, 0, ""),
		MaxGo: Go(1, 20, 0, ""),
	},
	{ // latest version supports go1.17
		Delve: Delve(1, 9, 1, ""),
		MinGo: Go(1, 17, 0, ""),
		MaxGo: Go(1, 19, 0, ""),
	},
	{ // latest version supports go1.16
		Delve: Delve(1, 9, 0, ""),
		MinGo: Go(1, 16, 0, ""),
		MaxGo: Go(1, 19, 0, ""),
	},
	{
		Delve: Delve(1, 8, 3, ""),
		MinGo: Go(1, 16, 0, ""),
		MaxGo: Go(1, 18, 0, ""),
	},
	{
		Delve: Delve(1, 8, 2, ""),
		MinGo: Go(1, 16, 0, ""),
		MaxGo: Go(1, 18, 0, ""),
	},
	{
		Delve: Delve(1, 8, 1, ""),
		MinGo: Go(1, 16, 0, ""),
		MaxGo: Go(1, 18, 0, ""),
	},
	{
		Delve: Delve(1, 8, 0, ""),
		MinGo: Go(1, 16, 0, ""),
		MaxGo: Go(1, 18, 0, ""),
	},
	{ // latest version supports go1.15
		Delve: Delve(1, 7, 3, ""),
		MinGo: Go(1, 15, 0, ""),
		MaxGo: Go(1, 17, 0, ""),
	},
	{
		Delve: Delve(1, 7, 2, ""),
		MinGo: Go(1, 15, 0, ""),
		MaxGo: Go(1, 17, 0, ""),
	},
	{
		Delve: Delve(1, 7, 1, ""),
		MinGo: Go(1, 15, 0, ""),
		MaxGo: Go(1, 17, 0, ""),
	},
	{
		Delve: Delve(1, 7, 0, ""),
		MinGo: Go(1, 15, 0, ""),
		MaxGo: Go(1, 17, 0, ""),
	},
	{ // latest version supports go1.14
		Delve: Delve(1, 6, 1, ""),
		MinGo: Go(1, 14, 0, ""),
		MaxGo: Go(1, 16, 0, ""),
	},
	{ // latest version supports go1.13
		Delve: Delve(1, 6, 0, ""),
		MinGo: Go(1, 13, 0, ""),
		MaxGo: Go(1, 16, 0, ""),
	},
	{
		Delve: Delve(1, 5, 1, ""),
		MinGo: Go(1, 13, 0, ""),
		MaxGo: Go(1, 15, 0, ""),
	},
	{ // latest version supports go1.12
		Delve: Delve(1, 5, 0, ""),
		MinGo: Go(1, 12, 0, ""),
		MaxGo: Go(1, 15, 0, ""),
	},
	{
		Delve: Delve(1, 4, 1, ""),
		MinGo: Go(1, 12, 0, ""),
		MaxGo: Go(1, 14, 0, ""),
	},
	{ // latest version supports go1.12
		Delve: Delve(1, 4, 0, ""),
		MinGo: Go(1, 11, 0, ""),
		MaxGo: Go(1, 14, 0, ""),
	},
	// before that, go module is not supported.

	// NOTE: this list comes from
	//     https://github.com/go-delve/delve/blob/master/pkg/goversion/compat.go
	// for each release version tag.

}
