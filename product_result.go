// Copyright © 2016 OpenFoodFacts. All rights reserved.
// Use of this source code is governed by the MIT license which can be found in the LICENSE.txt file.

package openfoodfacts

// ProductResult is the response from a query against the OpenFoodFacts database.
type ProductResult struct {
	StatusVerbose string   `json:"status_verbose"`
	Status        int      `json:"status"`
	Code          string   `json:"code"`
	Product       *Product `json:"product"`
}
