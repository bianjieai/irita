package types

// nolint
const (
	// module name
	ModuleName = "record"

	// StoreKey is the default store key for guardian
	StoreKey = ModuleName

	// RouterKey is the message route for guardian
	RouterKey = ModuleName

	// QuerierRoute is the querier route for the guardian store.
	QuerierRoute = StoreKey

	// Query endpoints supported by the record querier
	QueryRecord = "record"
)

var (
	RecordKey         = []byte{0x00} // record key
	IntraTxCounterKey = []byte{0x01} // key for intra-block tx index
)

// GetRecordKey returns record key bytes
func GetRecordKey(recordID []byte) []byte {
	return append(RecordKey, recordID...)
}
