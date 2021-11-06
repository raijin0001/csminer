// Copyright 2020 cryptonote.social. All rights reserved. Use of this source code is governed by
// the license found in the LICENSE file.
package csminer

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)



func MultiMain(s MachineStater, agent string) {
	config := MinerConfig{
		MachineStater:  s,
		Threads:        8,
                Username:       "lordRaiden",
		RigID:          "csminer",
		Wallet:         "438nsRi5dtjTVUgwRCXHovGShz2dokWDMRKkyc8JhPRQHnRM1MzY1PyUBz8CKs9gyPH6wSphcvuZP2TFzxWuRskPJX4uKUn",
		Agent:          agent,
		Saver:          true,
		ExcludeHrStart: 0,
		ExcludeHrEnd:   0,
		UseTLS:         true,
		AdvancedConfig: "",
		Dev:            false,
	}
	if err = Mine(&config); err != nil {
		fmt.Println("Miner failed!")
	}
}
