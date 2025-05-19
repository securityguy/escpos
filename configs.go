package escpos

var (
	Generic            = PrinterConfig{}
	ConfigEpsonTMT20II = PrinterConfig{}
	ConfigEpsonTMT88II = PrinterConfig{DisableUpsideDown: true}
	ConfigSOL802       = PrinterConfig{DisableUpsideDown: true}
)
