package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type TAIBroadcastNRItem struct {
	TAI                   TAI `aper:"valueExt"`
	CompletedCellsInTAINR CompletedCellsInTAINR
	IEExtensions          *ProtocolExtensionContainerTAIBroadcastNRItemExtIEs `aper:"optional"`
}
