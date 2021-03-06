// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package viewertest

import (
	"context"

	"github.com/facebookincubator/symphony/graph/ent"
	"github.com/facebookincubator/symphony/graph/viewer"
)

// DefaultViewer defines the default viewer set by this package.
var DefaultViewer = viewer.Viewer{
	Tenant: "test",
	User:   "tester@example.com",
	Role:   "superuser",
}

// Option enables viewer customization.
type Option func(*viewer.Viewer)

// WithViewer overrides default viewer.
func WithViewer(override *viewer.Viewer) Option {
	return func(v *viewer.Viewer) {
		*v = *override
	}
}

// NewContext returns viewer context for tests.
func NewContext(c *ent.Client, opts ...Option) context.Context {
	v := DefaultViewer
	for _, opt := range opts {
		opt(&v)
	}
	ctx := viewer.NewContext(context.Background(), &v)
	return ent.NewContext(ctx, c)
}
