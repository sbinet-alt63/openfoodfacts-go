// Copyright © 2016 OpenFoodFacts. All rights reserved.
// Use of this source code is governed by the MIT license which can be found in the LICENSE.txt file.

package openfoodfacts

// Ingredient lists informations about a Product's ingredient.
type Ingredient struct {
	ID   string `json:"id"`
	Rank int    `json:"rank,omitempty"`
	Text string `json:"text"`
}
