package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type AllocationAndRetentionPriority struct {
	PriorityLevelARP        PriorityLevelARP
	PreEmptionCapability    PreEmptionCapability
	PreEmptionVulnerability PreEmptionVulnerability
	IEExtensions            *ProtocolExtensionContainerAllocationAndRetentionPriorityExtIEs `aper:"optional"`
}
