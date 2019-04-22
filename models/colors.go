package models

import "fmt"

// HexCode represents an HTML color string
type HexCode string

// ToRGB converts a HexCode string color into its RGB equivalent
func (cs HexCode) ToRGB() (RGB, error) {
	var r, g, b uint8

	n, err := fmt.Sscanf(string(cs), "#%2x%2x%2x", &r, &g, &b)
	if err != nil || n != 3 {
		return RGB{}, err
	}

	return RGB{r, g, b}, nil
}

// RGB represents a traditional 24-bit color, having 8 bits
// for each of red, green, and blue.
type RGB struct {
	R, G, B uint8
}

// ToHexCode returns a Hex color string equivalent of the the
// RGB color
func (c RGB) ToHexCode() HexCode {
	return HexCode(fmt.Sprintf("%02x%02x%02x", c.R, c.G, c.B))
}

// ColorRequestBody is what the server is expecting from the
// POST /color call
type ColorRequestBody struct {
	Color string
}

// ColorResponse is the response on POST /color with a specific color
// in the body of the request.
type ColorResponse struct {
	RgbColor RGB `json:"rgbColor"`
}

// NewColorResponse creates a new color response out of the provided RGB color
func NewColorResponse(color RGB) ColorResponse {
	return ColorResponse{color}
}
