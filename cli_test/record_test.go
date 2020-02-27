package clitest

import (
	"fmt"
	"testing"

	"github.com/bianjieai/irita/app"
	"github.com/bianjieai/irita/modules/record"
	clientkeys "github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/tests"
	"github.com/stretchr/testify/require"
)

func TestIritaCLICreateRecord(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	fooAddr := f.KeyAddress(keyFoo)

	// start irita server
	proc := f.GDStart()
	defer proc.Stop(false)

	digest := "test"
	digestAlgo := "SHA256"
	uri := "localhost:1317"
	meta := "test"

	success, _, stderr := f.TxCreateRecord(fooAddr.String(), digest, digestAlgo, uri, meta, "-y")
	require.True(f.T, success)
	require.Empty(f.T, stderr)

	tests.WaitForNextNBlocksTM(1, f.Port)
	// Ensure transaction tags can be queried
	searchResult := f.QueryTxs(1, 50, "message.action=create_record", fmt.Sprintf("message.sender=%s", fooAddr))
	require.Len(t, searchResult.Txs, 1)

	var recordID string
	for _, log := range searchResult.Txs[0].Logs {
		for _, event := range log.Events {
			if event.Type == record.EventTypeCreateRecord {
				for _, attribute := range event.Attributes {
					if attribute.Key == record.AttributeKeyRecordID {
						recordID = attribute.Value
					}
				}
			}
		}
	}

	expRecord := record.RecordOutput{
		TxHash: searchResult.Txs[0].TxHash,
		Contents: []record.Content{{
			Digest:     digest,
			DigestAlgo: digestAlgo,
			URI:        uri,
			Meta:       meta,
		}},
		Creator: fooAddr,
	}

	res := f.QueryRecord(recordID)
	require.NotEmpty(f.T, res)
	require.Equal(f.T, expRecord, res)

	// Cleanup testing directories
	f.Cleanup()
}

// TxCreateRecord is iritacli tx record create
func (f *Fixtures) TxCreateRecord(from, digest, digestAlgo, uri, meta string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx record create %s %s --uri=%s --meta=%s %v --keyring-backend=test --from=%s", f.IritaCLIBinary, digest, digestAlgo, uri, meta, f.Flags(), from)
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), clientkeys.DefaultKeyPass)
}

// QueryRecord is iritacli query record record
func (f *Fixtures) QueryRecord(recordID string) (result record.RecordOutput) {
	cmd := fmt.Sprintf("%s query record record %s --output=%s %v", f.IritaCLIBinary, recordID, "json", f.Flags())
	out, _ := tests.ExecuteT(f.T, cmd, "")
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &result)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return
}
