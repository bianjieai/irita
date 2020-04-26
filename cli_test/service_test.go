package clitest

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	tmtypes "github.com/tendermint/tendermint/types"

	clientkeys "github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irismod/service"

	"github.com/bianjieai/irita/app"
)

func TestIritaCLIService(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	cdc := app.MakeCodec()

	serviceFeeTax := sdk.NewDecWithPrec(1, 2) // 1%
	maxRequestTimeout := int64(10)

	// Update service params for test
	genesisState := f.GenesisState()
	var serviceData service.GenesisState
	err := cdc.UnmarshalJSON(genesisState[service.ModuleName], &serviceData)
	require.NoError(t, err)
	serviceData.Params.ServiceFeeTax = serviceFeeTax
	serviceData.Params.MaxRequestTimeout = maxRequestTimeout
	serviceDataBz, err := cdc.MarshalJSON(serviceData)
	require.NoError(t, err)
	genesisState[service.ModuleName] = serviceDataBz

	genFile := filepath.Join(f.IritadHome, "config", "genesis.json")
	genDoc, err := tmtypes.GenesisDocFromFile(genFile)
	require.NoError(t, err)
	genDoc.AppState, err = cdc.MarshalJSON(genesisState)
	require.NoError(t, genDoc.SaveAs(genFile))

	proc := f.GDStart()
	defer proc.Stop(false)

	tests.WaitForTMStart(f.Port)
	tests.WaitForNextNBlocksTM(2, f.Port)

	fooAddr := f.KeyAddress(keyFoo)
	barAddr := f.KeyAddress(keyBar)

	sendTokens := sdk.TokensFromConsensusPower(10)
	f.TxSend(keyFoo, barAddr, sdk.NewCoin(denom, sendTokens), "-y")
	tests.WaitForNextNBlocksTM(1, f.Port)

	fooAcc := f.QueryAccount(fooAddr)
	fooCoinAmt := fooAcc.Coins.AmountOf(denom)
	require.Equal(t, "40000000", fooCoinAmt.String())

	barAcc := f.QueryAccount(barAddr)
	barCoinAmt := barAcc.Coins.AmountOf(denom)
	require.Equal(t, "10000000", barCoinAmt.String())

	// testing variables
	serviceName := "test"
	serviceDesc := "test"
	serviceTags := []string{"tag1", "tag2"}
	authorDesc := "author"
	serviceSchemas := `{"input":{"type":"object"},"output":{"type":"object"}}`
	deposit := fmt.Sprintf("10000%s", denom)
	priceAmt := 10
	price := fmt.Sprintf("%d%s", priceAmt, denom)
	pricing := fmt.Sprintf(`{"price":"%s"}`, price)
	minRespTime := uint64(5)
	addedDeposit := fmt.Sprintf("1%s", denom)
	serviceFeeCap := fmt.Sprintf("10%s", denom)
	input := `{"pair":"iris-usdt"}`
	timeout := int64(7)
	repeatedFreq := uint64(20)
	repeatedTotal := int64(10)
	result := `{"code":200,"message":""}`
	output := `{"last":"100"}`

	author := fooAddr.String()
	provider := fooAddr.String()
	consumer := barAddr.String()

	// define service
	success, _, _ := f.TxServiceDefine(serviceName, serviceDesc, serviceTags, authorDesc, serviceSchemas, author, "-y")
	require.True(t, success)

	tests.WaitForNextNBlocksTM(1, f.Port)

	svcDef := f.QueryServiceDefinition(serviceName)
	require.Equal(t, serviceName, svcDef.Name)
	require.Equal(t, serviceDesc, svcDef.Description)
	require.Equal(t, serviceTags, svcDef.Tags)
	require.Equal(t, author, svcDef.Author.String())
	require.Equal(t, authorDesc, svcDef.AuthorDescription)
	require.Equal(t, serviceSchemas, svcDef.Schemas)

	// bind service
	success, _, _ = f.TxServiceBind(serviceName, deposit, pricing, minRespTime, provider, "-y")
	require.True(t, success)

	tests.WaitForNextNBlocksTM(1, f.Port)

	binding := f.QueryServiceBinding(serviceName, provider)
	require.Equal(t, serviceName, binding.ServiceName)
	require.Equal(t, provider, binding.Provider)
	require.Equal(t, deposit, binding.Deposit.String())
	require.Equal(t, pricing, binding.Pricing)
	require.Equal(t, minRespTime, binding.MinRespTime)
	require.True(t, binding.Available)

	bindings := f.QueryServiceBindings(serviceName)
	require.Equal(t, 1, len(bindings))
	require.Equal(t, binding, bindings[0])

	// update binding
	success, _, _ = f.TxServiceUpdateBinding(serviceName, addedDeposit, "", 0, provider, "-y")
	require.True(t, success)

	tests.WaitForNextNBlocksTM(1, f.Port)

	binding = f.QueryServiceBinding(serviceName, provider)
	require.Equal(t, "11"+denom, binding.Deposit.String())

	// set withdrawal address
	success, _, _ = f.TxServiceSetWithdrawAddr(barAddr.String(), provider, "-y")
	require.True(t, success)

	tests.WaitForNextNBlocksTM(1, f.Port)

	withdrawAddr := f.QueryServiceWithdrawAddr(provider)
	require.Equal(t, barAddr.String(), withdrawAddr.String())

	// disable binding
	success, _, _ = f.TxServiceDisable(serviceName, provider, "-y")
	require.True(t, success)

	tests.WaitForNextNBlocksTM(1, f.Port)

	binding = f.QueryServiceBinding(serviceName, provider)
	require.Equal(t, false, binding.Available)

	// refund deposit
	success, _, _ = f.TxServiceRefundDeposit(serviceName, provider, "-y")
	require.True(t, success)

	tests.WaitForNextNBlocksTM(1, f.Port)

	binding = f.QueryServiceBinding(serviceName, provider)
	require.Equal(t, "", binding.Deposit.String())

	// enable binding
	success, _, _ = f.TxServiceEnable(serviceName, deposit, provider, "-y")
	require.True(t, success)

	tests.WaitForNextNBlocksTM(1, f.Port)

	binding = f.QueryServiceBinding(serviceName, provider)
	require.Equal(t, deposit, binding.Deposit.String())
	require.Equal(t, true, binding.Available)

	// service call
	success, _, _ = f.TxServiceCall(serviceName, provider, serviceFeeCap, input, timeout, repeatedFreq, repeatedTotal, consumer, "-y")
	require.True(t, success)

	tests.WaitForNextNBlocksTM(1, f.Port)

	// query the txs with the call_service event
	searchResult := f.QueryTxs(1, 50, "message.action=call_service", fmt.Sprintf("message.sender=%s", consumer))
	require.Len(t, searchResult.Txs, 1)

	var requestContextID string

	events := searchResult.Txs[0].Logs[0].Events
	for _, e := range events {
		for _, attribute := range e.Attributes {
			if attribute.Key == service.AttributeKeyRequestID {
				requestContextID = attribute.Value
				break
			}
		}
	}

	require.NotEmpty(t, requestContextID)

	requestContext := f.QueryServiceRequestContext(requestContextID)
	require.Equal(t, serviceName, requestContext.ServiceName)
	require.Equal(t, []sdk.AccAddress{fooAddr}, requestContext.Providers)
	require.Equal(t, consumer, requestContext.Consumer.String())
	require.Equal(t, input, requestContext.Input)
	require.Equal(t, serviceFeeCap, requestContext.ServiceFeeCap.String())
	require.Equal(t, timeout, requestContext.Timeout)
	require.Equal(t, false, requestContext.SuperMode)
	require.Equal(t, true, requestContext.Repeated)
	require.Equal(t, repeatedFreq, requestContext.RepeatedFrequency)
	require.Equal(t, repeatedTotal, requestContext.RepeatedTotal)
	require.Equal(t, uint64(1), requestContext.BatchCounter)
	require.Equal(t, uint16(1), requestContext.BatchRequestCount)
	require.Equal(t, uint16(0), requestContext.BatchResponseCount)
	require.Equal(t, service.BATCHRUNNING, requestContext.BatchState)
	require.Equal(t, service.RUNNING, requestContext.State)
	require.Equal(t, uint16(0), requestContext.ResponseThreshold)
	require.Equal(t, "", requestContext.ModuleName)

	// query requests by binding
	requests := f.QueryServiceRequests(serviceName, provider)
	require.Equal(t, 1, len(requests))
	require.Equal(t, consumer, requests[0].Consumer.String())
	require.Equal(t, provider, requests[0].Provider.String())
	require.Equal(t, input, requests[0].Input)
	require.Equal(t, price, requests[0].ServiceFee)
	require.Equal(t, requestContextID, requests[0].RequestContextID)
	require.Equal(t, uint64(1), requests[0].RequestContextBatchCounter)

	// query requests by request context
	requestsByCtx := f.QueryServiceRequests(requestContextID, "1")
	require.Equal(t, 1, len(requestsByCtx))
	require.Equal(t, requests[0], requestsByCtx[0])

	requestID := requests[0].ID.String()

	// respond to service request
	success, _, _ = f.TxServiceRespond(requestID, result, output, provider, "-y")
	require.True(t, success)

	tests.WaitForNextNBlocksTM(1, f.Port)

	// query response
	response := f.QueryServiceResponse(requestID)
	require.Equal(t, provider, response.Provider.String())
	require.Equal(t, consumer, response.Consumer.String())
	require.Equal(t, result, response.Result)
	require.Equal(t, output, response.Output)
	require.Equal(t, requestContextID, response.RequestContextID.String())
	require.Equal(t, uint64(1), response.RequestContextBatchCounter)

	// query request context
	requestContext = f.QueryServiceRequestContext(requestContextID)
	require.Equal(t, uint16(1), requestContext.BatchResponseCount)
	require.Equal(t, service.BATCHCOMPLETED, requestContext.BatchState)
	require.Equal(t, service.RUNNING, requestContext.State)

	// query responses by request context
	responses := f.QueryServiceResponses(requestContextID, 1)
	require.Equal(t, 1, len(responses))
	require.Equal(t, response, responses[0])

	// responses deleted on expiration height
	tests.WaitForHeightTM(requests[0].ExpirationHeight, f.Port)
	responses = f.QueryServiceResponses(requestContextID, 1)
	require.Equal(t, 0, len(responses))

	// pause the request context
	success, _, _ = f.TxServicePause(requestContextID, consumer, "-y")
	require.True(t, success)

	tests.WaitForNextNBlocksTM(1, f.Port)

	// query request context
	requestContext = f.QueryServiceRequestContext(requestContextID)
	require.Equal(t, service.PAUSED, requestContext.State)

	// start the request context
	success, _, _ = f.TxServiceStart(requestContextID, consumer, "-y")
	require.True(t, success)

	tests.WaitForNextNBlocksTM(1, f.Port)

	// query request context
	requestContext = f.QueryServiceRequestContext(requestContextID)
	require.Equal(t, service.RUNNING, requestContext.State)

	// kill the request context
	success, _, _ = f.TxServiceKill(requestContextID, consumer, "-y")
	require.True(t, success)

	tests.WaitForNextNBlocksTM(1, f.Port)

	// query request context
	requestContext = f.QueryServiceRequestContext(requestContextID)
	require.Equal(t, service.COMPLETED, requestContext.State)

	// query the earned fees
	earnedFees := f.QueryServiceFees(provider)

	earnedFeesAmtDec := sdk.NewDecFromInt(sdk.NewInt(int64(priceAmt))).Mul(sdk.OneDec().Sub(serviceFeeTax))
	earnedFeesAmt, _ := sdk.NewIntFromString(earnedFeesAmtDec.String())
	require.Equal(t, fmt.Sprintf("%s%s", earnedFeesAmt.String(), denom), earnedFees.Coins.String())

	// withdraw the earned fees (the provider's withdrawal address is bar)
	barAcc = f.QueryAccount(barAddr)
	oldBarCoinAmt := barAcc.Coins.AmountOf(denom)

	success, _, _ = f.TxServiceWithdrawFees(provider, "-y")
	require.True(t, success)

	tests.WaitForNextNBlocksTM(1, f.Port)

	barAcc = f.QueryAccount(barAddr)
	newBarCoinAmt := barAcc.Coins.AmountOf(denom)

	require.Equal(t, oldBarCoinAmt.Add(earnedFeesAmt), newBarCoinAmt)
}

// TxServiceDefine is iritacli tx service define
func (f *Fixtures) TxServiceDefine(serviceName, serviceDesc string, tags []string, serviceAuthorDesc, serviceSchemas, from string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx service define --keyring-backend=test --name %s --description %s --tags %s --author-description %s --schemas %s --from=%s %v", f.IritaCLIBinary, serviceName, serviceDesc, tags, serviceAuthorDesc, serviceSchemas, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), clientkeys.DefaultKeyPass)
}

// TxServiceBind is iritacli tx service bind
func (f *Fixtures) TxServiceBind(serviceName, deposit, pricing string, minRespTime uint64, from string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx service bind --keyring-backend=test --service-name %s --deposit %s --pricing %s --min-resp-time %d --from=%s %v", f.IritaCLIBinary, serviceName, deposit, pricing, minRespTime, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), clientkeys.DefaultKeyPass)
}

// TxServiceUpdateBinding is iritacli tx service update-binding
func (f *Fixtures) TxServiceUpdateBinding(serviceName, deposit, pricing string, minRespTime uint64, from string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx service update-binding --keyring-backend=test --service-name %s --deposit %s --pricing %s --min-resp-time %d --from=%s %v", f.IritaCLIBinary, serviceName, deposit, pricing, minRespTime, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), clientkeys.DefaultKeyPass)
}

// TxServiceSetWithdrawAddr is iritacli tx service set-withdraw-addr
func (f *Fixtures) TxServiceSetWithdrawAddr(withdrawalAddr string, from string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx service set-withdraw-addr --keyring-backend=test %s --from=%s %v", f.IritaCLIBinary, withdrawalAddr, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), clientkeys.DefaultKeyPass)
}

// TxServiceDisable is iritacli tx service disable
func (f *Fixtures) TxServiceDisable(serviceName, from string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx service disable --keyring-backend=test --service-name %s --from=%s %v", f.IritaCLIBinary, serviceName, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), clientkeys.DefaultKeyPass)
}

// TxServiceEnable is iritacli tx service enable
func (f *Fixtures) TxServiceEnable(serviceName, deposit, from string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx service enable --keyring-backend=test --service-name %s --deposit %s --from=%s %v", f.IritaCLIBinary, serviceName, deposit, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), clientkeys.DefaultKeyPass)
}

// TxServiceRefundDeposit is iritacli tx service refund-deposit
func (f *Fixtures) TxServiceRefundDeposit(serviceName, from string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx service refund-deposit --keyring-backend=test --service-name %s --from=%s %v", f.IritaCLIBinary, serviceName, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), clientkeys.DefaultKeyPass)
}

// TxServiceCall is iritacli tx service call
func (f *Fixtures) TxServiceCall(serviceName, providers, serviceFeeCap, input string, timeout int64, frequency uint64, total int64, from string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx service call --keyring-backend=test --service-name %s --providers %s --service-fee-cap %s --data %s --timeout %d --repeated --frequency %d --total %d --from=%s %v", f.IritaCLIBinary, serviceName, providers, serviceFeeCap, input, timeout, frequency, total, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), clientkeys.DefaultKeyPass)
}

// TxServiceRespond is iritacli tx service respond
func (f *Fixtures) TxServiceRespond(requestID, result, output, from string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx service respond --keyring-backend=test --request-id %s --result %s --data %s --from=%s %v", f.IritaCLIBinary, requestID, result, output, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), clientkeys.DefaultKeyPass)
}

// TxServicePause is iritacli tx service pause
func (f *Fixtures) TxServicePause(requestContextID, from string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx service pause %s --keyring-backend=test --from=%s %v", f.IritaCLIBinary, requestContextID, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), clientkeys.DefaultKeyPass)
}

// TxServiceStart is iritacli tx service start
func (f *Fixtures) TxServiceStart(requestContextID, from string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx service start %s --keyring-backend=test --from=%s %v", f.IritaCLIBinary, requestContextID, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), clientkeys.DefaultKeyPass)
}

// TxServiceKill is iritacli tx service kill
func (f *Fixtures) TxServiceKill(requestContextID, from string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx service kill %s --keyring-backend=test --from=%s %v", f.IritaCLIBinary, requestContextID, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), clientkeys.DefaultKeyPass)
}

// TxServiceWithdrawFees is iritacli tx service withdraw-fees
func (f *Fixtures) TxServiceWithdrawFees(from string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx service withdraw-fees --keyring-backend=test --from=%s %v", f.IritaCLIBinary, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), clientkeys.DefaultKeyPass)
}

// QueryServiceDefinition is iritacli query service definition
func (f *Fixtures) QueryServiceDefinition(serviceName string) service.ServiceDefinition {
	cmd := fmt.Sprintf("%s query service definition %s %v", f.IritaCLIBinary, serviceName, f.Flags())
	res, errStr := tests.ExecuteT(f.T, cmd, "")
	require.Empty(f.T, errStr)
	cdc := app.MakeCodec()
	var svcDef service.ServiceDefinition
	err := cdc.UnmarshalJSON([]byte(res), &svcDef)
	require.NoError(f.T, err)
	return svcDef
}

// QueryServiceBinding is iritacli query service binding
func (f *Fixtures) QueryServiceBinding(serviceName, provider string) service.ServiceBinding {
	cmd := fmt.Sprintf("%s query service binding %s %s %v", f.IritaCLIBinary, serviceName, provider, f.Flags())
	res, errStr := tests.ExecuteT(f.T, cmd, "")
	require.Empty(f.T, errStr)
	cdc := app.MakeCodec()
	var binding service.ServiceBinding
	err := cdc.UnmarshalJSON([]byte(res), &binding)
	require.NoError(f.T, err)
	return binding
}

// QueryServiceBindings is iritacli query service bindings
func (f *Fixtures) QueryServiceBindings(serviceName string) []service.ServiceBinding {
	cmd := fmt.Sprintf("%s query service bindings %s %v", f.IritaCLIBinary, serviceName, f.Flags())
	res, errStr := tests.ExecuteT(f.T, cmd, "")
	require.Empty(f.T, errStr)
	cdc := app.MakeCodec()
	var bindings []service.ServiceBinding
	err := cdc.UnmarshalJSON([]byte(res), &bindings)
	require.NoError(f.T, err)
	return bindings
}

// QueryServiceWithdrawAddr is iritacli query service withdraw-addr
func (f *Fixtures) QueryServiceWithdrawAddr(provider string) sdk.AccAddress {
	cmd := fmt.Sprintf("%s query service withdraw-addr %s %v", f.IritaCLIBinary, provider, f.Flags())
	res, errStr := tests.ExecuteT(f.T, cmd, "")
	require.Empty(f.T, errStr)
	cdc := app.MakeCodec()
	var withdrawAddr sdk.AccAddress
	err := cdc.UnmarshalJSON([]byte(res), &withdrawAddr)
	require.NoError(f.T, err)
	return withdrawAddr
}

// QueryServiceRequests is iritacli query service requests
func (f *Fixtures) QueryServiceRequests(arg1, arg2 string) []service.Request {
	cmd := fmt.Sprintf("%s query service requests %s %s %v", f.IritaCLIBinary, arg1, arg2, f.Flags())
	res, errStr := tests.ExecuteT(f.T, cmd, "")
	require.Empty(f.T, errStr)
	cdc := app.MakeCodec()
	var reqs []service.Request
	err := cdc.UnmarshalJSON([]byte(res), &reqs)
	require.NoError(f.T, err)
	return reqs
}

// QueryServiceResponse is iritacli query service response
func (f *Fixtures) QueryServiceResponse(requestID string) service.Response {
	cmd := fmt.Sprintf("%s query service response %s %v", f.IritaCLIBinary, requestID, f.Flags())
	res, errStr := tests.ExecuteT(f.T, cmd, "")
	require.Empty(f.T, errStr)
	cdc := app.MakeCodec()
	var response service.Response
	err := cdc.UnmarshalJSON([]byte(res), &response)
	require.NoError(f.T, err)
	return response
}

// QueryServiceResponses is iritacli query service responses
func (f *Fixtures) QueryServiceResponses(requestContextID string, batchCounter uint64) []service.Response {
	cmd := fmt.Sprintf("%s query service responses %s %d %v", f.IritaCLIBinary, requestContextID, batchCounter, f.Flags())
	res, errStr := tests.ExecuteT(f.T, cmd, "")
	require.Empty(f.T, errStr)
	cdc := app.MakeCodec()
	var responses []service.Response
	err := cdc.UnmarshalJSON([]byte(res), &responses)
	require.NoError(f.T, err)
	return responses
}

// QueryServiceRequestContext is iritacli query service request-context
func (f *Fixtures) QueryServiceRequestContext(requestContextID string) service.RequestContext {
	cmd := fmt.Sprintf("%s query service request-context %s %v", f.IritaCLIBinary, requestContextID, f.Flags())
	res, errStr := tests.ExecuteT(f.T, cmd, "")
	require.Empty(f.T, errStr)
	cdc := app.MakeCodec()
	var reqCtx service.RequestContext
	err := cdc.UnmarshalJSON([]byte(res), &reqCtx)
	require.NoError(f.T, err)
	return reqCtx
}

// QueryServiceFees is iritacli query service fees
func (f *Fixtures) QueryServiceFees(provider string) service.EarnedFees {
	cmd := fmt.Sprintf("%s query service fees %s %v", f.IritaCLIBinary, provider, f.Flags())
	res, errStr := tests.ExecuteT(f.T, cmd, "")
	require.Empty(f.T, errStr)
	cdc := app.MakeCodec()
	var fees service.EarnedFees
	err := cdc.UnmarshalJSON([]byte(res), &fees)
	require.NoError(f.T, err)
	return fees
}
