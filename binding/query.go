// Copyright 2017 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"fmt"
	"net/http"
)

type queryBinding struct{}

func (queryBinding) Name() string {
	return "query"
}

func (queryBinding) Bind(req *http.Request, obj any) error {
	values := req.URL.Query()
	fmt.Printf("tag %s query values %+v\n", getTagFromMimes(req.Header.Get("Content-Type")), values)
	if err := mapForm(obj, values, getTagFromMimes(req.Header.Get("Content-Type"))); err != nil {
		return err
	}
	return validate(obj)
}

func getTagFromMimes(contentType string) string {
	var tag string
	switch contentType {
	case MIMEJSON:
		tag = "json"
	default:
		tag = "form"
	}
	return tag
}
