// Copyright 2015,2016,2017,2018,2019,2020 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	h := sha256.New()
	for _, filepattern := range os.Args {
		matchs, err := filepath.Glob(filepattern)
		if err != nil {
			log.Fatal(err)
		}
		for _, filename := range matchs {
			if err := appendSum(filename, h); err != nil {
				log.Fatal(err)
			}
		}
	}
	fmt.Printf("%x", h.Sum(nil))
}

func appendSum(filename string, h hash.Hash) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := io.Copy(h, f); err != nil {
		return err
	}
	return nil
}
