package qrgen

// ゠゠゠゠゠゠゠゠゠゠ここから追加する゠゠゠゠゠゠゠゠゠゠
import (
	"bytes"
	"image"
	"image/png"
)

// CreateImage generates image from the given byte format of the image
func CreateImage(b []byte) (image.Image, error) {
	return png.Decode(bytes.NewReader(b))
}
// ゠゠゠゠゠゠゠゠゠゠ここまで追加する゠゠゠゠゠゠゠゠゠゠