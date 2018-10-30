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

package blockchain

import (
	"github.com/kardiachain/go-kardia/common/state"
	"github.com/kardiachain/go-kardia/configs"
	"github.com/kardiachain/go-kardia/lib/common"
	"github.com/kardiachain/go-kardia/types"
)

type BaseBlockChain interface {
	Genesis() *types.Block
	CurrentHeader() *types.Header
	CurrentBlock() *types.Block
	WriteBlockWithoutState(block *types.Block) error
	WriteCommit(height uint64, commit *types.Commit)
	GetBlockByHeight(height uint64) *types.Block
	State() (*state.StateDB, error)
	CommitTrie(root common.Hash) error
	WriteReceipts(receipts types.Receipts, block *types.Block)
	ReadCommit(height uint64) *types.Commit
	Config() *configs.ChainConfig
}
