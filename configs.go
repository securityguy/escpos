package escpos

var (
	Generic            = PrinterConfig{QRCodeMaxSize: 8}
	ConfigEpsonTMT20II = PrinterConfig{}
	ConfigEpsonTMT88II = PrinterConfig{DisableUpsideDown: true}
	ConfigSOL802       = PrinterConfig{DisableUpsideDown: true}
)
