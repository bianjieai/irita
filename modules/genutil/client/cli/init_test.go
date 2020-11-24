package cli_test

//var testMbm = module.NewBasicManager(genutil.AppModuleBasic{})
//
//func TestInitCmd(t *testing.T) {
//	tests := []struct {
//		name      string
//		flags     func(dir string) []string
//		shouldErr bool
//		err       error
//	}{{
//		name: "happy path",
//		flags: func(dir string) []string {
//			return []string{"appnode-test"}
//		},
//		shouldErr: false,
//		err:       nil,
//	}}
//
//	for _, tt := range tests {
//		tt := tt
//		t.Run(tt.name, func(t *testing.T) {
//			home, cleanup := testutil.NewTestCaseDir(t)
//			defer cleanup()
//
//			logger := log.NewNopLogger()
//			cfg, err := genutiltest.CreateDefaultTendermintConfig(home)
//			require.NoError(t, err)
//
//			serverCtx := server.NewContext(viper.New(), cfg, logger)
//			interfaceRegistry := types.NewInterfaceRegistry()
//			marshaler := codec.NewProtoCodec(interfaceRegistry)
//			clientCtx := client.Context{}.
//				WithJSONMarshaler(marshaler).
//				WithLegacyAmino(makeCodec()).
//				WithHomeDir(home)
//
//			ctx := context.Background()
//			ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)
//			ctx = context.WithValue(ctx, server.ServerContextKey, serverCtx)
//
//			cmd := genutilcli.InitCmd(testMbm, home)
//			cmd.SetArgs(
//				tt.flags(home),
//			)
//
//			if tt.shouldErr {
//				err := cmd.ExecuteContext(ctx)
//				require.EqualError(t, err, tt.err.Error())
//			} else {
//				require.NoError(t, cmd.ExecuteContext(ctx))
//			}
//		})
//	}
//
//}
//
//func setupClientHome(t *testing.T) func() {
//	_, cleanup := testutil.NewTestCaseDir(t)
//	return cleanup
//}
//
//func TestEmptyState(t *testing.T) {
//	t.Cleanup(setupClientHome(t))
//
//	home, cleanup := testutil.NewTestCaseDir(t)
//	t.Cleanup(cleanup)
//
//	logger := log.NewNopLogger()
//	cfg, err := genutiltest.CreateDefaultTendermintConfig(home)
//	require.NoError(t, err)
//
//	serverCtx := server.NewContext(viper.New(), cfg, logger)
//	interfaceRegistry := types.NewInterfaceRegistry()
//	marshaler := codec.NewProtoCodec(interfaceRegistry)
//	clientCtx := client.Context{}.
//		WithJSONMarshaler(marshaler).
//		WithLegacyAmino(makeCodec()).
//		WithHomeDir(home)
//
//	ctx := context.Background()
//	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)
//	ctx = context.WithValue(ctx, server.ServerContextKey, serverCtx)
//
//	cmd := genutilcli.InitCmd(testMbm, home)
//	cmd.SetArgs([]string{"appnode-test", fmt.Sprintf("--%s=%s", cli.HomeFlag, home)})
//
//	require.NoError(t, cmd.ExecuteContext(ctx))
//
//	old := os.Stdout
//	r, w, _ := os.Pipe()
//	os.Stdout = w
//
//	cmd = server.ExportCmd(nil, home)
//	cmd.SetArgs([]string{fmt.Sprintf("--%s=%s", cli.HomeFlag, home)})
//	require.NoError(t, cmd.ExecuteContext(ctx))
//
//	outC := make(chan string)
//	go func() {
//		var buf bytes.Buffer
//		_, _ = io.Copy(&buf, r)
//		outC <- buf.String()
//	}()
//
//	w.Close()
//	os.Stdout = old
//	out := <-outC
//
//	require.Contains(t, out, "genesis_time")
//	require.Contains(t, out, "chain_id")
//	require.Contains(t, out, "consensus_params")
//	require.Contains(t, out, "app_hash")
//	require.Contains(t, out, "app_state")
//}
//
//func TestStartStandAlone(t *testing.T) {
//	home, cleanup := testutil.NewTestCaseDir(t)
//	t.Cleanup(cleanup)
//	t.Cleanup(setupClientHome(t))
//
//	logger := log.NewNopLogger()
//	interfaceRegistry := types.NewInterfaceRegistry()
//	marshaler := codec.NewProtoCodec(interfaceRegistry)
//	err := genutiltest.ExecInitCmd(testMbm, home, marshaler)
//	require.NoError(t, err)
//
//	app, err := mock.NewApp(home, logger)
//	require.NoError(t, err)
//
//	svrAddr, _, err := server.FreeTCPAddr()
//	require.NoError(t, err)
//
//	svr, err := abci_server.NewServer(svrAddr, "socket", app)
//	require.NoError(t, err, "error creating listener")
//
//	svr.SetLogger(logger.With("module", "abci-server"))
//	err = svr.Start()
//	require.NoError(t, err)
//
//	timer := time.NewTimer(time.Duration(2) * time.Second)
//	for range timer.C {
//		err = svr.Stop()
//		require.NoError(t, err)
//		break
//	}
//}
//
//func TestInitNodeValidatorFiles(t *testing.T) {
//	home, cleanup := testutil.NewTestCaseDir(t)
//	cfg, _ := genutiltest.CreateDefaultTendermintConfig(home)
//	t.Cleanup(cleanup)
//	nodeID, valPubKey, err := genutil.InitializeNodeValidatorFiles(cfg)
//	require.Nil(t, err)
//	require.NotEqual(t, "", nodeID)
//	require.NotEqual(t, 0, len(valPubKey.Bytes()))
//}
//
//// custom tx codec
//func makeCodec() *codec.LegacyAmino {
//	var cdc = codec.NewLegacyAmino()
//	sdk.RegisterLegacyAminoCodec(cdc)
//	cryptocodec.RegisterCrypto(cdc)
//	return cdc
//}
