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

package types

import (
	"bytes"
	"testing"

	"github.com/kardiachain/go-kardia/lib/common"
	"github.com/kardiachain/go-kardia/lib/rlp"
)

func TestDualEventsEncoding(t *testing.T) {
	firstDualEvent := CreateNewDualEvent(100)
	secondDualEvent := CreateNewDualEvent(55)

	dualEvents := DualEvents{firstDualEvent, secondDualEvent}

	if returnbytes, err := rlp.EncodeToBytes(firstDualEvent); !bytes.Equal(dualEvents.GetRlp(0), returnbytes) || err != nil {
		t.Error("Dual Events GetRlp error")
	}
	if returnbytes, err := rlp.EncodeToBytes(secondDualEvent); !bytes.Equal(dualEvents.GetRlp(1), returnbytes) || err != nil {
		t.Error("Dual Events GetRlp error")
	}
}

func CreateNewDualEvent(nonce uint64) *DualEvent {
	dualAction := DualAction{Name:"dual_test"}
	return NewDualEvent(nonce, false, "KAI", new(common.Hash), new(EventSummary), &dualAction)
}
