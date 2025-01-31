/*
   GoToSocial
   Copyright (C) 2021 GoToSocial Authors admin@gotosocial.org

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package nodeinfo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// NodeInfoGETHandler returns a compliant nodeinfo response to node info queries.
// See: https://nodeinfo.diaspora.software/
func (m *Module) NodeInfoGETHandler(c *gin.Context) {
	l := logrus.WithFields(logrus.Fields{
		"func":       "NodeInfoGETHandler",
		"user-agent": c.Request.UserAgent(),
	})

	ni, err := m.processor.GetNodeInfo(c.Request.Context(), c.Request)
	if err != nil {
		l.Debugf("error with get node info request: %s", err)
		c.JSON(err.Code(), err.Safe())
		return
	}

	c.JSON(http.StatusOK, ni)
}
