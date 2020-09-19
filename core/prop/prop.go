package prop

// Prop provides methods to validate SGF properties scanned in by the parser.

// Field is used to store label parsed from SGF
type Field string

// ValidProperties is a comprehensive list of valid SGF properties
var ValidProperties = map[Field]bool{"AB": true, "AE": true, "AN": true, "AP": true, "AR": true, "AW": true, "B": true, "BL": true, "BM": true, "BR": true, "BT": true, "C": true, "CA": true, "CP": true, "CR": true, "DD": true, "DM": true, "DO": true, "DT": true, "EV": true, "FF": true, "FG": true, "GB": true, "GC": true, "GM": true, "GN": true, "GW": true, "HA": true, "HO": true, "IT": true, "KM": true, "KO": true, "LB": true, "LN": true, "MA": true, "MN": true, "N": true, "OB": true, "ON": true, "OT": true, "OW": true, "PB": true, "PC": true, "PL": true, "PM": true, "PW": true, "RE": true, "RO": true, "RU": true, "SL": true, "SO": true, "SQ": true, "ST": true, "SZ": true, "TB": true, "TE": true, "TM": true, "TR": true, "TW": true, "UC": true, "US": true, "V": true, "VW": true, "W": true, "WL": true, "WR": true, "WT": true}

// Validate returns true if Prop is an accepted SGF property.
func Validate(field Field) bool {
	_, ok := ValidProperties[field]
	return ok
}
