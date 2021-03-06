// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package tests

import (
	"bytes"
	"fmt"
	"io"
	"math/big"
	"strconv"
	"strings"
	"testing"

	"github.com/nekonium/go-nekonium/common"
	"github.com/nekonium/go-nekonium/common/math"
	"github.com/nekonium/go-nekonium/core"
	"github.com/nekonium/go-nekonium/core/state"
	"github.com/nekonium/go-nekonium/core/types"
	"github.com/nekonium/go-nekonium/ethdb"
	"github.com/nekonium/go-nekonium/log"
	"github.com/nekonium/go-nekonium/params"
)

func RunStateTestWithReader(chainConfig *params.ChainConfig, r io.Reader, skipTests []string) error {
	tests := make(map[string]VmTest)
	if err := readJson(r, &tests); err != nil {
		return err
	}

	if err := runStateTests(chainConfig, tests, skipTests); err != nil {
		return err
	}

	return nil
}

func RunStateTest(chainConfig *params.ChainConfig, p string, skipTests []string) error {
	tests := make(map[string]VmTest)
	if err := readJsonFile(p, &tests); err != nil {
		return err
	}

	if err := runStateTests(chainConfig, tests, skipTests); err != nil {
		return err
	}

	return nil

}

func BenchStateTest(chainConfig *params.ChainConfig, p string, conf bconf, b *testing.B) error {
	tests := make(map[string]VmTest)
	if err := readJsonFile(p, &tests); err != nil {
		return err
	}
	test, ok := tests[conf.name]
	if !ok {
		return fmt.Errorf("test not found: %s", conf.name)
	}

	// XXX Yeah, yeah...
	env := make(map[string]string)
	env["currentCoinbase"] = test.Env.CurrentCoinbase
	env["currentDifficulty"] = test.Env.CurrentDifficulty
	env["currentGasLimit"] = test.Env.CurrentGasLimit
	env["currentNumber"] = test.Env.CurrentNumber
	env["previousHash"] = test.Env.PreviousHash
	if n, ok := test.Env.CurrentTimestamp.(float64); ok {
		env["currentTimestamp"] = strconv.Itoa(int(n))
	} else {
		env["currentTimestamp"] = test.Env.CurrentTimestamp.(string)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchStateTest(chainConfig, test, env, b)
	}

	return nil
}

func benchStateTest(chainConfig *params.ChainConfig, test VmTest, env map[string]string, b *testing.B) {
	b.StopTimer()
	db, _ := ethdb.NewMemDatabase()
	statedb := makePreState(db, test.Pre)
	b.StartTimer()

	RunState(chainConfig, statedb, env, test.Exec)
}

func runStateTests(chainConfig *params.ChainConfig, tests map[string]VmTest, skipTests []string) error {
	skipTest := make(map[string]bool, len(skipTests))
	for _, name := range skipTests {
		skipTest[name] = true
	}

	for name, test := range tests {
		if skipTest[name] /*|| name != "JUMPDEST_Attack"*/ {
			log.Info(fmt.Sprint("Skipping state test", name))
			continue
		}

		//fmt.Println("StateTest:", name)
		if err := runStateTest(chainConfig, test); err != nil {
			return fmt.Errorf("%s: %s\n", name, err.Error())
		}

		//log.Info(fmt.Sprint("State test passed: ", name))
		//fmt.Println(string(statedb.Dump()))
	}
	return nil

}

func runStateTest(chainConfig *params.ChainConfig, test VmTest) error {
	db, _ := ethdb.NewMemDatabase()
	statedb := makePreState(db, test.Pre)

	// XXX Yeah, yeah...
	env := make(map[string]string)
	env["currentCoinbase"] = test.Env.CurrentCoinbase
	env["currentDifficulty"] = test.Env.CurrentDifficulty
	env["currentGasLimit"] = test.Env.CurrentGasLimit
	env["currentNumber"] = test.Env.CurrentNumber
	env["previousHash"] = test.Env.PreviousHash
	if n, ok := test.Env.CurrentTimestamp.(float64); ok {
		env["currentTimestamp"] = strconv.Itoa(int(n))
	} else {
		env["currentTimestamp"] = test.Env.CurrentTimestamp.(string)
	}

	var (
		ret []byte
		// gas  *big.Int
		// err  error
		logs []*types.Log
	)

	ret, logs, _, _ = RunState(chainConfig, statedb, env, test.Transaction)

	// Compare expected and actual return
	var rexp []byte
	if strings.HasPrefix(test.Out, "#") {
		n, _ := strconv.Atoi(test.Out[1:])
		rexp = make([]byte, n)
	} else {
		rexp = common.FromHex(test.Out)
	}
	if !bytes.Equal(rexp, ret) {
		return fmt.Errorf("return failed. Expected %x, got %x\n", rexp, ret)
	}

	// check post state
	for addr, account := range test.Post {
		address := common.HexToAddress(addr)
		if !statedb.Exist(address) {
			return fmt.Errorf("did not find expected post-state account: %s", addr)
		}

		if balance := statedb.GetBalance(address); balance.Cmp(math.MustParseBig256(account.Balance)) != 0 {
			return fmt.Errorf("(%x) balance failed. Expected: %v have: %v\n", address[:4], math.MustParseBig256(account.Balance), balance)
		}

		if nonce := statedb.GetNonce(address); nonce != math.MustParseUint64(account.Nonce) {
			return fmt.Errorf("(%x) nonce failed. Expected: %v have: %v\n", address[:4], account.Nonce, nonce)
		}

		for addr, value := range account.Storage {
			v := statedb.GetState(address, common.HexToHash(addr))
			vexp := common.HexToHash(value)

			if v != vexp {
				return fmt.Errorf("storage failed:\n%x: %s:\nexpected: %x\nhave:     %x\n(%v %v)\n", address[:4], addr, vexp, v, vexp.Big(), v.Big())
			}
		}
	}

	root, _ := statedb.Commit(false)
	if common.HexToHash(test.PostStateRoot) != root {
		return fmt.Errorf("Post state root error. Expected: %s have: %x", test.PostStateRoot, root)
	}

	// check logs
	if len(test.Logs) > 0 {
		if err := checkLogs(test.Logs, logs); err != nil {
			return err
		}
	}

	return nil
}

func RunState(chainConfig *params.ChainConfig, statedb *state.StateDB, env, tx map[string]string) ([]byte, []*types.Log, *big.Int, error) {
	environment, msg := NewEVMEnvironment(false, chainConfig, statedb, env, tx)
	gaspool := new(core.GasPool).AddGas(math.MustParseBig256(env["currentGasLimit"]))

	root, _ := statedb.Commit(false)
	statedb.Reset(root)

	snapshot := statedb.Snapshot()

	ret, gasUsed, err := core.ApplyMessage(environment, msg, gaspool)
	if err != nil {
		statedb.RevertToSnapshot(snapshot)
	}
	statedb.Commit(chainConfig.IsEIP158(environment.Context.BlockNumber))

	return ret, statedb.Logs(), gasUsed, err
}
