// Copyright © 2021 Ettore Di Giacinto <mudler@mocaccino.org>
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 2 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program; if not, see <http://www.gnu.org/licenses/>.

package types

type Context interface {
	Logger
	GarbageCollector
	GetConfig() LuetConfig
	Copy() Context
	// SetAnnotation sets generic annotations to hold in a context
	SetAnnotation(s string, i interface{})

	// GetAnnotation gets generic annotations to hold in a context
	GetAnnotation(s string) interface{}

	WithLoggingContext(s string) Context
}
