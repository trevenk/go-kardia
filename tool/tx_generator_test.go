/*
 *  Copyright 2018 KardiaChain
 *  This file is part of the go-kardia library.
 *
 *  The go-kardia library is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU Lesser General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.
 *
 *  The go-kardia library is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 *  GNU Lesser General Public License for more details.
 *
 *  You should have received a copy of the GNU Lesser General Public License
 *  along with the go-kardia library. If not, see <http://www.gnu.org/licenses/>.
 */

package tool

import (
	"testing"

	"github.com/kardiachain/go-kardia/types"
	"github.com/kardiachain/go-kardia/configs"
)

func TestGenerateTx(t *testing.T) {
	genTool := NewGeneratorTool(GetAccounts(configs.GenesisAddrKeys))

	result := genTool.GenerateTx(1000)
	for _, tx := range result {
		from, _ := types.Sender(tx)
		if !containsInGenesis(from.String()) {
			t.Error("Sender addr should be in genesis block")
		}
		if !containsInGenesis(tx.To().String()) {
			t.Error("Receiver addr should be in genesis")
		}
		if from.String() == configs.KardiaAccountToCallSmc {
			t.Error("Sender should not be the account to call smc")
		}
		if from == *tx.To() {
			t.Error("Sender & receiver addrs should not be the same")
		}
	}
}

/*func TestGenerateRandomTxWithState(t *testing.T) {
	genTool := NewGeneratorTool(GetAccounts(configs.GenesisAddrKeys))
	statedb, _ := state.New(log.New(), common.Hash{}, state.NewDatabase(kaidb.NewMemStore()))
	result := genTool.GenerateRandomTxWithState(10, statedb)
	for _, txInterface := range result {
		tx := txInterface.(*types.Transaction)
		from, _ := types.Sender(tx)
		if !containsInGenesis(from.String()) {
			t.Error("Sender addr should be in genesis block")
		}
		if !containsInGenesis(tx.To().String()) {
			t.Error("Receiver addr should be in genesis")
		}
		if from.String() == configs.KardiaAccountToCallSmc {
			t.Error("Sender should not be the account to call smc")
		}
		if from == *tx.To() {
			t.Error("Sender & receiver addrs should not be the same")
		}
	}
}*/

func containsInGenesis(address string) bool {
	for k := range configs.GenesisAccounts {
		if k == address {
			return true
		}
	}
	return false
}
