package model

/*

  Copyright 2024, JAFAX, Inc.

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

type InvalidStatusValue struct {
	Err error
}

func (i *InvalidStatusValue) Error() string {
	return "Invalid value! Must be either 'enabled' or 'locked'"
}

type PasswordHashMismatch struct {
	Err error
}

func (p *PasswordHashMismatch) Error() string {
	return "Password hashes do not match!"
}

type SchedulingConflict struct {
	Err error
}

func (s *SchedulingConflict) Error() string {
	return "Scheduling conflict: Start or end of event conflicts with existing scheduled event"
}
