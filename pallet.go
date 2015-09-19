// +build !windows

package main

func init() {
	for i := 16; i <= 255; i++ {
		palette = append(palette, attrColor(i))
	}
}
