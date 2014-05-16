package painter

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

func CountryCodes() map[string]string {
	Codes := make(map[string]string)
	Codes["AND"] = "AD"
	Codes["ARE"] = "AE"
	Codes["AFG"] = "AF"
	Codes["ATG"] = "AG"
	Codes["AIA"] = "AI"
	Codes["ALB"] = "AL"
	Codes["ARM"] = "AM"
	Codes["ANT"] = "AN"
	Codes["AGO"] = "AO"
	Codes["ATA"] = "AQ"
	Codes["ARG"] = "AR"
	Codes["ASM"] = "AS"
	Codes["AUT"] = "AT"
	Codes["AUS"] = "AU"
	Codes["ABW"] = "AW"
	Codes["AZE"] = "AZ"
	Codes["BIH"] = "BA"
	Codes["BRB"] = "BB"
	Codes["BGD"] = "BD"
	Codes["BEL"] = "BE"
	Codes["BFA"] = "BF"
	Codes["BGR"] = "BG"
	Codes["BHR"] = "BH"
	Codes["BDI"] = "BI"
	Codes["BEN"] = "BJ"
	Codes["BLM"] = "BL"
	Codes["BMU"] = "BM"
	Codes["BRN"] = "BN"
	Codes["BOL"] = "BO"
	Codes["BRA"] = "BR"
	Codes["BHS"] = "BS"
	Codes["BTN"] = "BT"
	Codes["BWA"] = "BW"
	Codes["BLR"] = "BY"
	Codes["BLZ"] = "BZ"
	Codes["CAN"] = "CA"
	Codes["CCK"] = "CC"
	Codes["COD"] = "CD"
	Codes["CAF"] = "CF"
	Codes["COG"] = "CG"
	Codes["CHE"] = "CH"
	Codes["CIV"] = "CI"
	Codes["COK"] = "CK"
	Codes["CHL"] = "CL"
	Codes["CMR"] = "CM"
	Codes["CHN"] = "CN"
	Codes["COL"] = "CO"
	Codes["CRC"] = "CR"
	Codes["CUB"] = "CU"
	Codes["CPV"] = "CV"
	Codes["CXR"] = "CX"
	Codes["CYP"] = "CY"
	Codes["CZE"] = "CZ"
	Codes["DEU"] = "DE"
	Codes["DJI"] = "DJ"
	Codes["DNK"] = "DK"
	Codes["DMA"] = "DM"
	Codes["DOM"] = "DO"
	Codes["DZA"] = "DZ"
	Codes["ECU"] = "EC"
	Codes["EST"] = "EE"
	Codes["EGY"] = "EG"
	Codes["ESH"] = "EH"
	Codes["ERI"] = "ER"
	Codes["ESP"] = "ES"
	Codes["ETH"] = "ET"
	Codes["FIN"] = "FI"
	Codes["FJI"] = "FJ"
	Codes["FLK"] = "FK"
	Codes["FSM"] = "FM"
	Codes["FRO"] = "FO"
	Codes["FRA"] = "FR"
	Codes["GAB"] = "GA"
	Codes["GBR"] = "GB"
	Codes["GRD"] = "GD"
	Codes["GEO"] = "GE"
	Codes["GHA"] = "GH"
	Codes["GIB"] = "GI"
	Codes["GRL"] = "GL"
	Codes["GMB"] = "GM"
	Codes["GIN"] = "GN"
	Codes["GNQ"] = "GQ"
	Codes["GRC"] = "GR"
	Codes["GTM"] = "GT"
	Codes["GUM"] = "GU"
	Codes["GNB"] = "GW"
	Codes["GUY"] = "GY"
	Codes["HKG"] = "HK"
	Codes["HND"] = "HN"
	Codes["HRV"] = "HR"
	Codes["HTI"] = "HT"
	Codes["HUN"] = "HU"
	Codes["IDN"] = "ID"
	Codes["IRL"] = "IE"
	Codes["ISR"] = "IL"
	Codes["IMN"] = "IM"
	Codes["IND"] = "IN"
	Codes["IOT"] = "IO"
	Codes["IRQ"] = "IQ"
	Codes["IRN"] = "IR"
	Codes["IS"] = "IS"
	Codes["ITA"] = "IT"
	Codes["JEY"] = "JE"
	Codes["JAM"] = "JM"
	Codes["JOR"] = "JO"
	Codes["JPN"] = "JP"
	Codes["KEN"] = "KE"
	Codes["KGZ"] = "KG"
	Codes["KHM"] = "KH"
	Codes["KIR"] = "KI"
	Codes["COM"] = "KM"
	Codes["KNA"] = "KN"
	Codes["PRK"] = "KP"
	Codes["KOR"] = "KR"
	Codes["KWT"] = "KW"
	Codes["CYM"] = "KY"
	Codes["KAZ"] = "KZ"
	Codes["LAO"] = "LA"
	Codes["LBN"] = "LB"
	Codes["LCA"] = "LC"
	Codes["LIE"] = "LI"
	Codes["LKA"] = "LK"
	Codes["LBR"] = "LR"
	Codes["LSO"] = "LS"
	Codes["LTU"] = "LT"
	Codes["LUX"] = "LU"
	Codes["LVA"] = "LV"
	Codes["LBY"] = "LY"
	Codes["MAR"] = "MA"
	Codes["MCO"] = "MC"
	Codes["MDA"] = "MD"
	Codes["MNE"] = "ME"
	Codes["MAF"] = "MF"
	Codes["MDG"] = "MG"
	Codes["MHL"] = "MH"
	Codes["MKD"] = "MK"
	Codes["MLI"] = "ML"
	Codes["MMR"] = "MM"
	Codes["MNG"] = "MN"
	Codes["MAC"] = "MO"
	Codes["MNP"] = "MP"
	Codes["MRT"] = "MR"
	Codes["MSR"] = "MS"
	Codes["MLT"] = "MT"
	Codes["MUS"] = "MU"
	Codes["MDV"] = "MV"
	Codes["MWI"] = "MW"
	Codes["MEX"] = "MX"
	Codes["MYS"] = "MY"
	Codes["MOZ"] = "MZ"
	Codes["NAM"] = "NA"
	Codes["NCL"] = "NC"
	Codes["NER"] = "NE"
	Codes["NGA"] = "NG"
	Codes["NIC"] = "NI"
	Codes["NLD"] = "NL"
	Codes["NOR"] = "NO"
	Codes["NPL"] = "NP"
	Codes["NRU"] = "NR"
	Codes["NIU"] = "NU"
	Codes["NZL"] = "NZ"
	Codes["OMN"] = "OM"
	Codes["PAN"] = "PA"
	Codes["PER"] = "PE"
	Codes["PYF"] = "PF"
	Codes["PNG"] = "PG"
	Codes["PHL"] = "PH"
	Codes["PAK"] = "PK"
	Codes["POL"] = "PL"
	Codes["SPM"] = "PM"
	Codes["PCN"] = "PN"
	Codes["PRI"] = "PR"
	Codes["PRT"] = "PT"
	Codes["PLW"] = "PW"
	Codes["PRY"] = "PY"
	Codes["QAT"] = "QA"
	Codes["ROU"] = "RO"
	Codes["SRB"] = "RS"
	Codes["RUS"] = "RU"
	Codes["RWA"] = "RW"
	Codes["SAU"] = "SA"
	Codes["SLB"] = "SB"
	Codes["SYC"] = "SC"
	Codes["SDN"] = "SD"
	Codes["SWE"] = "SE"
	Codes["SGP"] = "SG"
	Codes["SHN"] = "SH"
	Codes["SVN"] = "SI"
	Codes["SJM"] = "SJ"
	Codes["SVK"] = "SK"
	Codes["SLE"] = "SL"
	Codes["SMR"] = "SM"
	Codes["SEN"] = "SN"
	Codes["SOM"] = "SO"
	Codes["SUR"] = "SR"
	Codes["STP"] = "ST"
	Codes["SLV"] = "SV"
	Codes["SYR"] = "SY"
	Codes["SWZ"] = "SZ"
	Codes["TCA"] = "TC"
	Codes["TCD"] = "TD"
	Codes["TGO"] = "TG"
	Codes["THA"] = "TH"
	Codes["TJK"] = "TJ"
	Codes["TKL"] = "TK"
	Codes["TLS"] = "TL"
	Codes["TKM"] = "TM"
	Codes["TUN"] = "TN"
	Codes["TON"] = "TO"
	Codes["TUR"] = "TR"
	Codes["TTO"] = "TT"
	Codes["TUV"] = "TV"
	Codes["TWN"] = "TW"
	Codes["TZA"] = "TZ"
	Codes["UKR"] = "UA"
	Codes["UGA"] = "UG"
	Codes["USA"] = "US"
	Codes["URY"] = "UY"
	Codes["UZB"] = "UZ"
	Codes["VAT"] = "VA"
	Codes["VCT"] = "VC"
	Codes["VEN"] = "VE"
	Codes["VGB"] = "VG"
	Codes["VIR"] = "VI"
	Codes["VNM"] = "VN"
	Codes["VUT"] = "VU"
	Codes["WLF"] = "WF"
	Codes["WSM"] = "WS"
	Codes["YEM"] = "YE"
	Codes["MYT"] = "YT"
	Codes["ZAF"] = "ZA"
	Codes["ZMB"] = "ZM"
	Codes["ZWE"] = "ZW"
	return Codes
}

type SVG struct {
	Colors map[string]string
	Base   string
}

func NewSVG() *SVG {

	svg := new(SVG)

	svg.Colors = make(map[string]string)

	return svg
}

func (svg *SVG) Generate(destination string) {
	file, _ := os.Open("base.svg")
	defer file.Close()
	decoder := xml.NewDecoder(file)

	newFile, _ := os.Create(destination)
	defer newFile.Close()

	encoder := xml.NewEncoder(newFile)
	encoder.Indent("", "	")

	head := "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"no\"?>\n<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:dc=\"http://purl.org/dc/elements/1.1/\" xmlns:cc=\"http://creativecommons.org/ns#\" xmlns:rdf=\"http://www.w3.org/1999/02/22-rdf-syntax-ns#\" xmlns:svg=\"http://www.w3.org/2000/svg\" xmlns:sodipodi=\"http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd\" xmlns:inkscape=\"http://www.inkscape.org/namespaces/inkscape\" viewBox=\"-100 -450 3000 2000\" version=\"1.0\" height=\"800\" width=\"1400\">"

	land := ""

	for token, err := decoder.Token(); err == nil; token, err = decoder.Token() {

		switch elem := token.(type) {
		case xml.StartElement:
			if elem.Name.Local == "svg" {
				fmt.Fprintln(newFile, head)
			} else {
				for _, attr := range elem.Attr {
					if attr.Name.Local == "class" {
						if strings.Contains(attr.Value, "land") {
							land = strings.ToUpper(attr.Value[5:])
						}
						break
					}
				}
				if land != "" {
					color, ok := svg.Colors[land]
					if ok {
						//fmt.Println(land, color)
					}
					for i := range elem.Attr {
						if elem.Attr[i].Name.Local == "style" {
							if ok {
								elem.Attr[i].Value = "fill:" + color
							} else {
								elem.Attr[i].Value = ""
							}
							break
						}
					}
				}
				if land != "AQ" {
					encoder.EncodeToken(elem)
				}
			}

		case xml.EndElement:
			old := land
			if elem.Name.Local == "g" {
				land = ""
			}
			if old != "AQ" {
				encoder.EncodeToken(elem)
			}
		case xml.CharData:
			if land != "AQ" {
				elem = []byte(strings.Replace(string(elem), "\n", "", -1))
				encoder.EncodeToken(elem)
			}
		}

	}

	encoder.Flush()
	fmt.Fprintln(newFile, "</svg>")

}
