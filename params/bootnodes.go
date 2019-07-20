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

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
		// nekonium Bootnodes
		"enode://de730651d37a3ad090c1a19f9ba098dcd1d50dcc28ce05b0c13481346fe77a6598399c24ddd44b474a37e5f14580a35da5cbdd71fb5c9269a8c1e086d71b5566@150.95.148.243:28568",
		//@fuyuton
		"enode://d72e22095cf0ee5b62ec9a8c8e103c178718a2b90b9c9194da8f54c71fe26037ea5fdb468d9e0056f1258e31dd3681f61f1f4814493e1afe8f90422b9af137c5@52.14.110.93:28568",
		//@hxcoin#3
		"enode://9b65f365aff76a784f79d3e01d1a35f3cd6f2e3b36876367f78f2a3b0529363581121601dace248440ce22b63a7dff8f947c3606aaa38b537ab2bb0846f78d9b@104.196.183.246:28568",
		//@ohac#2
		"enode://6ccee0c32237e9eb6441c412d577902fc70cb6d94811e2f80a0412969b2ddb866eca45b01be53e1d243a7a734da8ea89c347b9d4fcd041a18c5ca5de60243c8b@34.216.230.29:28568",
		//@namyang
		"enode://761472e1485201b8e83398cbc5d514250aee94219b6e45399608a129e04d39113f1f1440ebb574bbc6526b6149e42e66efc0cee3f650bacdf693996ec7b2c3d4@212.24.97.43:28568",	
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
