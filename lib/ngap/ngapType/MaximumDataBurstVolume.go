package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type MaximumDataBurstVolume struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:4095"`
}
