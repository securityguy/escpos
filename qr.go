package escpos

import (
	"fmt"
	"math"
)

// Prints a QR Code.
// code specifies the data to be printed
// model specifies the qr code model. false for model 1, true for model 2
// size specifies the size in dots. It needs to be between 1 and 16

func (e *Escpos) QRCode(code string, model bool, size uint8, correctionLevel uint8) (int, error) {
	if len(code) > 7089 {
		return 0, fmt.Errorf("the code is too long, it's length should be smaller than 7090")
	}

	var maxSize uint8 = 16
	if e.config.QRCodeMaxSize > 0 {
		maxSize = e.config.QRCodeMaxSize
	}

	if size < 1 {
		size = 1
	}

	if size > maxSize {
		size = maxSize
	}

	var m byte = 49
	var err error
	// set the qr code model
	if model {
		m = 50
	}
	_, err = e.WriteRaw([]byte{gs, '(', 'k', 4, 0, 49, 65, m, 0})
	if err != nil {
		return 0, err
	}

	// set the qr code size
	_, err = e.WriteRaw([]byte{gs, '(', 'k', 3, 0, 49, 67, size})
	if err != nil {
		return 0, err
	}

	// set the qr code error correction level
	if correctionLevel < 48 {
		correctionLevel = 48
	}
	if correctionLevel > 51 {
		correctionLevel = 51
	}
	_, err = e.WriteRaw([]byte{gs, '(', 'k', 3, 0, 49, 69, size})
	if err != nil {
		return 0, err
	}

	// store the data in the buffer
	// we now write stuff to the printer, so lets save it for returning

	// pL and pH define the size of the data. Data ranges from 1 to (pL + pH*256)-3
	// 3 < pL + pH*256 < 7093
	var codeLength = len(code) + 3
	var pL, pH byte
	pH = byte(int(math.Floor(float64(codeLength) / 256)))
	pL = byte(codeLength - 256*int(pH))

	written, err := e.WriteRaw(append([]byte{gs, '(', 'k', pL, pH, 49, 80, 48}, []byte(code)...))
	if err != nil {
		return written, err
	}

	// finally print the buffer
	_, err = e.WriteRaw([]byte{gs, '(', 'k', 3, 0, 49, 81, 48})
	if err != nil {
		return written, err
	}

	return written, nil
}
