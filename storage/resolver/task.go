// Copyright 2022 zGraph Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resolver

import (
	"github.com/vescale/zgraph/storage/kv"
	"github.com/vescale/zgraph/storage/mvcc"
)

// Task represents a resolve task.
type Task struct {
	Key       kv.Key
	StartVer  mvcc.Version
	CommitVer mvcc.Version
	Notifier  Notifier
}
