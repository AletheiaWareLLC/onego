/*
 * Copyright 2020 Aletheia Ware LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"github.com/AletheiaWareLLC/bcgo"
	"github.com/AletheiaWareLLC/cryptogo"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	t, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(time.Since(start), err)
	}
	threshold := uint64(t)
	block := &bcgo.Block{
		Timestamp: bcgo.Timestamp(),
	}
	var max uint64
	for nonce := uint64(1); nonce > 0; nonce++ {
		block.Nonce = nonce
		hash, err := cryptogo.HashProtobuf(block)
		if err != nil {
			log.Fatal(time.Since(start), err)
		}
		ones := bcgo.Ones(hash)
		if ones > max {
			log.Println(time.Since(start), nonce, ones)
			max = ones
		}
		if ones > threshold {
			log.Println(time.Since(start), "Success")
			return
		}
	}
	log.Fatal(time.Since(start), "Error: Nonce Wrap Around")
}
