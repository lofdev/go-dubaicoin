// Copyright 2015 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package utils

import "github.com/dbix-project/go-dubaicoin/p2p/discover"

// FrontierBootNodes are the enode URLs of the P2P bootstrap nodes running on
// the Frontier network.
var FrontierBootNodes = []*discover.Node{
	// DBIX/DEV Go Bootnodes
	discover.MustParseNode("enode://ecd0e22008d07dfb56487a59a1a707b8806c18677d790a675a7ff5e19155ad6e379fcb27b77115e16f33021757527d097df922773d56336e5e05f9efc91f1238@104.207.152.207:57955"), //node1
	discover.MustParseNode("enode://e5bd453779a941730929f77564cb8a053b845401ca1902b7bc66aba89631ce6423fb657905c806102b780cff31b0ff8a0ad652d3d3188181414d4b397067b6ef@80.241.222.77:57955"), //dbix b server
	discover.MustParseNode("enode://14a622abe7f93f589b1edca25e52b0a52ac0a2baf1387a197bd44bf5521068bd0c3687d6128afd410eed072b78332ced4ce4a550099b4e76710c9a89ce242088@45.32.65.104:57955"), //dbix_node1
	discover.MustParseNode("enode://39bb1c137d4d53017c39b5db3dd1b1dd8e44f710a2b55e79fdc0a1ebd0f139e6973420b1358f7e40f23eb5630b7605d25529fbac05aabdecdc2cfd4d0d2452d8@45.32.220.235:57955"), //dbix_node2
	discover.MustParseNode("enode://8cda27e3d924155bc8a823b95e327041161cdc876627921b4a690cae2946580cf0358a8035417f9c0621b4735570e0867171869c1989613a3793c18c65126d18@45.76.35.129:57955"), //dbix_node3
	discover.MustParseNode("enode://9073cc28395f32d7a85cf8a0295a353453dd9e682afca7cd968a4e5b5d660385069ecae0f7026916f0f5dda471284114ae6cb8e16d94c69a8fb49733f0b8aa21@45.76.138.199:57955"), //dbix_node4
	discover.MustParseNode("enode://1d1affca2ea45105897b2fc2de696cb95dd5dd06516db9a088b9d3ba1deefb81b3a7982ef5b728e09251e005af75c9db33ef1969bee0ed4a6df64b4452039be3@45.32.158.190:57955"), //dbix_node5
}

// TestNetBootNodes are the enode URLs of the P2P bootstrap nodes running on the
// Morden test network.
var TestNetBootNodes = []*discover.Node{
	// ETH/DEV Go Bootnodes
	/*discover.MustParseNode("enode://n4533109cc9bd7604e4ff6c095f7a1d807e15b38e9bfeb05d3b7c423ba86af0a9e89abbf40bd9dde4250fef114cd09270fa4e224cbeef8b7bf05a51e8260d7uv@198.158.98.185.231:54920"),*/

	// ETH/DEV Cpp Bootnodes
}
