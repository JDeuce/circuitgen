// Copyright 2019 Twitch Interactive, Inc.  All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the License is
// located at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file. This file is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package circuittestmulti

// Test generation in a separate package. gen_test.go contains comprehensive test on generated circuit wrappers.
//
// Disable goimports to catch any import bugs

//go:generate circuitgen circuit --goimports=true --pkg ../ --name Publisher --name Aggregator --out ./
