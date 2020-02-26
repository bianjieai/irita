package keeper_test

import (
	"testing"

	"github.com/bianjieai/irita/modules/record"
	"github.com/bianjieai/irita/modules/record/internal/keeper"
	"github.com/bianjieai/irita/modules/record/internal/types"
	"github.com/bianjieai/irita/simapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

var (
	testCreator = sdk.AccAddress(tmhash.Sum([]byte("test-creator")))
)

type KeeperTestSuite struct {
	suite.Suite

	cdc    *codec.Codec
	ctx    sdk.Context
	keeper keeper.Keeper
}

func (suite *KeeperTestSuite) SetupTest() {
	app := simapp.Setup(false)

	suite.cdc = app.Codec()
	suite.ctx = app.BaseApp.NewContext(false, abci.Header{})
	suite.keeper = app.RecordKeeper
	suite.keeper.SetIntraTxCounter(suite.ctx, 0)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestAddRecord() {
	content := types.Content{
		Digest:     "test",
		DigestAlgo: "SHA256",
		URI:        "loalhost:1317",
		Meta:       "test",
	}
	testRecord := record.NewRecord([]byte("test"), []types.Content{content}, testCreator)

	recordID := suite.keeper.AddRecord(suite.ctx, testRecord)
	addedRecord, found := suite.keeper.GetRecord(suite.ctx, recordID)
	suite.True(found)
	suite.Equal(testRecord, addedRecord)

	recordsIterator := suite.keeper.RecordsIterator(suite.ctx)
	defer recordsIterator.Close()
	var records []types.Record
	for ; recordsIterator.Valid(); recordsIterator.Next() {
		var record types.Record
		suite.cdc.MustUnmarshalBinaryLengthPrefixed(recordsIterator.Value(), &record)
		records = append(records, record)
	}
	suite.Equal(1, len(records))
	suite.Contains(records, testRecord)
}
