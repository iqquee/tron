package tron

type Tron struct{}

const (
	MainNet       = "https://api.trongrid.io"
	ShastaTestNet = "https://api.shasta.trongrid.io"
	NileTestNet   = "https://nile.trongrid.io"
)

func New() *Tron {
	return &Tron{}
}
