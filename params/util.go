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

package params

import (
	"math"
	"math/big"

	"github.com/nekonium/go-nekonium/common"
)

var (
	TestNetGenesisHash = common.HexToHash("0x86a89bcce7783fec92fce4f7b85094db7702c3afc28881aa4a64fd633ecf3526") // Testnet genesis hash to enforce below configs on
	MainNetGenesisHash = common.HexToHash("0x1a505395bfe4b2a8eef2f80033d68228db70e82bb695dd4ffb20e6d0cf71cb73") // Mainnet genesis hash to enforce below configs on

	TestNetHomesteadBlock = big.NewInt(0) // Testnet homestead block
	MainNetHomesteadBlock = big.NewInt(0) // Mainnet homestead block

	TestNetHomesteadGasRepriceBlock = big.NewInt(0) // Testnet gas reprice block
	MainNetHomesteadGasRepriceBlock = big.NewInt(0) // Mainnet gas reprice block

	TestNetHomesteadGasRepriceHash = common.HexToHash("0x86a89bcce7783fec92fce4f7b85094db7702c3afc28881aa4a64fd633ecf3526") // Testnet gas reprice block hash (used by fast sync)
	MainNetHomesteadGasRepriceHash = common.HexToHash("0x1a505395bfe4b2a8eef2f80033d68228db70e82bb695dd4ffb20e6d0cf71cb73") // Mainnet gas reprice block hash (used by fast sync)

	TestNetSpuriousDragon = big.NewInt(100)
	MainNetSpuriousDragon = big.NewInt(100)

	TestNetMetropolisBlock = big.NewInt(math.MaxInt64)
	MainNetMetropolisBlock = big.NewInt(math.MaxInt64)

	TestNetChainID = big.NewInt(3) // Test net default chain ID
	MainNetChainID = big.NewInt(1) // main net default chain ID
)

var (
	//Homested難易度計算のHFです。Homestedフォークよりも大きい値をセットしてください。
	MainNetNekoniumFork01 = big.NewInt(7777)
	TestNetNekoniumFork01 = big.NewInt(1)
)
