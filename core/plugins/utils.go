// Copyright © 2018 Nori info@nori.io
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

package plugins

import (
	"github.com/nori-io/nori-common/logger"
	"github.com/nori-io/nori-common/meta"
)

type FileTable map[string]meta.Meta

func (ft FileTable) Find(id meta.ID) string {
	for p, i := range ft {
		if i.Id() == id {
			return p
		}
	}
	return ""
}

func LogFieldsMeta(m meta.Meta) logger.Fields {
	return logger.Fields{
		"plugin_id":      string(m.Id().ID),
		"plugin_version": m.Id().Version,
		"interface":      string(m.GetInterface()),
	}
}

func LogFieldsId(id meta.ID) logger.Fields {
	return logger.Fields{
		"plugin_id":      string(id.ID),
		"plugin_version": id.Version,
	}
}
