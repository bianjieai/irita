package app

import (
	"fmt"
	"os"
	"testing"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/simapp"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// Profile with:
// /usr/local/go/bin/go test -benchmem -run=^$ github.com/cosmos/cosmos-sdk/simapp -bench ^BenchmarkFullAppSimulation$ -Commit=true -cpuprofile cpu.out
func BenchmarkFullAppSimulation(b *testing.B) {
	config, db, dir, logger, _, err := simapp.SetupSimulation("goleveldb-app-sim", "Simulation")
	if err != nil {
		b.Fatalf("simulation setup failed: %s", err.Error())
	}

	defer func() {
		db.Close()
		err = os.RemoveAll(dir)
		if err != nil {
			b.Fatal(err)
		}
	}()

	app := NewIritaApp(logger, db, nil, true, map[int64]bool{}, DefaultNodeHome, simapp.FlagPeriodValue, MakeEncodingConfig(), simapp.EmptyAppOptions{}, interBlockCacheOpt())

	// run randomized simulation
	_, simParams, simErr := simulation.SimulateFromSeed(
		b,
		os.Stdout,
		app.BaseApp,
		simapp.AppStateFn(app.AppCodec(), app.SimulationManager()),
		simtypes.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
		simapp.SimulationOperations(app, app.AppCodec(), config),
		app.ModuleAccountAddrs(),
		config,
		app.AppCodec(),
	)

	// export state and simParams before the simulation error is checked
	if err = simapp.CheckExportSimulation(app, config, simParams); err != nil {
		b.Fatal(err)
	}

	if simErr != nil {
		b.Fatal(simErr)
	}

	if config.Commit {
		simapp.PrintStats(db)
	}
}

func BenchmarkInvariants(b *testing.B) {
	config, db, dir, logger, _, err := simapp.SetupSimulation("leveldb-app-invariant-bench", "Simulation")
	if err != nil {
		b.Fatalf("simulation setup failed: %s", err.Error())
	}

	config.AllInvariants = false

	defer func() {
		db.Close()
		if err = os.RemoveAll(dir); err != nil {
			b.Fatal(err)
		}
	}()

	app := NewIritaApp(logger, db, nil, true, map[int64]bool{}, DefaultNodeHome, simapp.FlagPeriodValue, MakeEncodingConfig(), simapp.EmptyAppOptions{}, interBlockCacheOpt())

	// run randomized simulation
	_, simParams, simErr := simulation.SimulateFromSeed(
		b,
		os.Stdout,
		app.BaseApp,
		simapp.AppStateFn(app.AppCodec(), app.SimulationManager()),
		simtypes.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
		simapp.SimulationOperations(app, app.AppCodec(), config),
		app.ModuleAccountAddrs(),
		config,
		app.AppCodec(),
	)

	// export state and simParams before the simulation error is checked
	if err = simapp.CheckExportSimulation(app, config, simParams); err != nil {
		b.Fatal(err)
	}

	if simErr != nil {
		b.Fatal(simErr)
	}

	if config.Commit {
		simapp.PrintStats(db)
	}

	ctx := app.NewContext(true, tmproto.Header{Height: app.LastBlockHeight() + 1})

	// 3. Benchmark each invariant separately
	//
	// NOTE: We use the crisis keeper as it has all the invariants registered with
	// their respective metadata which makes it useful for testing/benchmarking.
	for _, cr := range app.crisisKeeper.Routes() {
		cr := cr
		b.Run(fmt.Sprintf("%s/%s", cr.ModuleName, cr.Route), func(b *testing.B) {
			if res, stop := cr.Invar(ctx); stop {
				b.Fatalf(
					"broken invariant at block %d of %d\n%s",
					ctx.BlockHeight()-1, config.NumBlocks, res,
				)
			}
		})
	}
}
