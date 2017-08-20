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
	//@J4Np
	"enode://24a1022814746eb3447c163b079c76b0e5e47190e865f4bc672bea558927a64aa80d8b32295147aa7e78e4f70e98a78c31a515abfa44219ef6c7c9f811c4b3ed@203.137.111.167:28568",
	//@ohac
	"enode://02bd416be704a5c84e323f350067f4850c227ac6b66fe18a8d571e8770525b9f1d07a262c39512304ee6e88e6dc92577c4aa24aaec817874b38a933a66e8dd60@157.7.152.130:28568",
	//@hxcoin#1
	"enode://2763414119225d208941358fe86fc6645cc183b9d1780eaf4aefb633afea912995f3c52890b396c34537b6e5be51db3d58d67601efc221a88dcc029926b4c826@76.28.205.190:28568",
	//@hxcoin#2
	"enode://399dde242064e038c9939fb6b69f2e8886b17d1518064ee3b130f99561d4b7833f82527291c3aa70815ba07113421f862f82d85566862472c31418941d5f9b6c@52.168.133.100:28568",
	//@fuyuton
	"enode://d72e22095cf0ee5b62ec9a8c8e103c178718a2b90b9c9194da8f54c71fe26037ea5fdb468d9e0056f1258e31dd3681f61f1f4814493e1afe8f90422b9af137c5@52.14.110.93:28568",
	//@w0life
	"enode://1117cdedbaaab7ea75567fd2f5b572c5c3cdd0d0812ff6a25e6bf7cdcc9b43e947e3dec447c3c906f436be48ba82812f325d13b03e0bb119a98c089bf7040327@212.24.104.24:28568",
	//@w0life#2
	"enode://8e1114411f2c8e1019cc10de66d74b043135cf4bf9355129e5d291b7f359a50b14130216a329955a405c638a5da384ec86dd4fbb2bc822cfcb7029089447ec45@35.197.49.146:28568",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
// "enode://6ce05930c72abc632c58e2e4324f7c7ea478cec0ed4fa2528982cf34483094e9cbc9216e7aa349691242576d552a2a56aaeae426c5303ded677ce455ba1acd9d@13.84.180.240:30303", // US-TX
// "enode://20c9ad97c081d63397d7b685a412227a40e23c8bdc6688c6f37e97cfbc22d2b4d1db1510d8f61e6a8866ad7f0e17c02b14182d37ea7c3c8b9c2683aeb6b733a1@52.169.14.227:30303", // IE
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
// "enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
}

// RinkebyV5Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network for the experimental RLPx v5 topic-discovery network.
var RinkebyV5Bootnodes = []string{
// "enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303?discport=30304", // IE
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
// "enode://0cc5f5ffb5d9098c8b8c62325f3797f56509bff942704687b6530992ac706e2cb946b90a34f1f19548cd3c7baccbcaea354531e5983c7d1bc0dee16ce4b6440b@40.118.3.223:30305",
// "enode://1c7a64d76c0334b0418c004af2f67c50e36a3be60b5e4790bdac0439d21603469a85fad36f2473c9a80eb043ae60936df905fa28f1ff614c3e5dc34f15dcd2dc@40.118.3.223:30308",
// "enode://85c85d7143ae8bb96924f2b54f1b3e70d8c4d367af305325d30a61385a432f247d2c75c45c6b4a60335060d072d7f5b35dd1d4c45f76941f62a4f83b6e75daaf@40.118.3.223:30309",
}
