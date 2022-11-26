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

package storage

import (
	"github.com/cockroachdb/pebble"
)

type pebbleStorage struct {
	db *pebble.DB
}

func New() Storage {
	return &pebbleStorage{}

}

func (s *pebbleStorage) Open(dirname string, options ...Option) error {
	opt := &pebble.Options{}
	for _, op := range options {
		op(opt)
	}
	db, err := pebble.Open(dirname, opt)
	if err != nil {
		return err
	}
	s.db = db

	return nil
}

func (s *pebbleStorage) Begin() (Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (s *pebbleStorage) Snapshot(ver Version) (Snapshot, error) {
	//TODO implement me
	panic("implement me")
}

func (s *pebbleStorage) CurrentVersion() (Version, error) {
	//TODO implement me
	panic("implement me")
}

func (s *pebbleStorage) Close() error {
	if s.db == nil {
		return nil
	}
	return s.db.Close()
}
