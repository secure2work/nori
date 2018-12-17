// Copyright © 2018 Secure2Work info@secure2work.com
//
// This program is free software: you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation, either version 3
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package interfaces

import (
	"errors"
	"html/template"
	"io"
)

type Templates interface {
	Set(name string, body string) error
	Get(name string) (*template.Template, error)
	Delete(name string) error
	Execute(name string, data interface{}) (string, error)
	ExecuteWriter(name string, data interface{}, w io.Writer) error
}

var ErrTemplateNotFound = errors.New("template not found")