package escpos

var (
	Generic            = PrinterConfig{}
	Fujun              = PrinterConfig{QRCodeMaxSize: 8, DisableRotate: true}
	ConfigEpsonTMT20II = PrinterConfig{}
	ConfigEpsonTMT88II = PrinterConfig{DisableUpsideDown: true}
	ConfigSOL802       = PrinterConfig{DisableUpsideDown: true}
)
