// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017 The go-ovh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ContactAddress Representation of an Address
// swagger:model contact.Address
type ContactAddress struct {

	// City
	City string `json:"city,omitempty"`

	// Country
	Country string `json:"country,omitempty"`

	// First line of the address
	Line1 string `json:"line1,omitempty"`

	// Second line of the address
	Line2 string `json:"line2,omitempty"`

	// Third line of the address
	Line3 string `json:"line3,omitempty"`

	// Others details
	OtherDetails string `json:"otherDetails,omitempty"`

	// Province name
	Province string `json:"province,omitempty"`

	// Zipcode
	Zip string `json:"zip,omitempty"`
}

// Validate validates this contact address
func (m *ContactAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountry(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var contactAddressTypeCountryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AC","AD","AE","AF","AG","AI","AL","AM","AO","AQ","AR","AS","AT","AU","AW","AX","AZ","BA","BB","BD","BE","BF","BG","BH","BI","BJ","BL","BM","BN","BO","BQ","BR","BS","BT","BW","BY","BZ","CA","CC","CD","CF","CG","CH","CI","CK","CL","CM","CN","CO","CR","CU","CV","CW","CX","CY","CZ","DE","DG","DJ","DK","DM","DO","DZ","EA","EC","EE","EG","EH","ER","ES","ET","FI","FJ","FK","FM","FO","FR","GA","GB","GD","GE","GF","GG","GH","GI","GL","GM","GN","GP","GQ","GR","GS","GT","GU","GW","GY","HK","HN","HR","HT","HU","IC","ID","IE","IL","IM","IN","IO","IQ","IR","IS","IT","JE","JM","JO","JP","KE","KG","KH","KI","KM","KN","KP","KR","KW","KY","KZ","LA","LB","LC","LI","LK","LR","LS","LT","LU","LV","LY","MA","MC","MD","ME","MF","MG","MH","MK","ML","MM","MN","MO","MP","MQ","MR","MS","MT","MU","MV","MW","MX","MY","MZ","NA","NC","NE","NF","NG","NI","NL","NO","NP","NR","NU","NZ","OM","PA","PE","PF","PG","PH","PK","PL","PM","PN","PR","PS","PT","PW","PY","QA","RE","RO","RS","RU","RW","SA","SB","SC","SD","SE","SG","SH","SI","SJ","SK","SL","SM","SN","SO","SR","SS","ST","SV","SX","SY","SZ","TA","TC","TD","TF","TG","TH","TJ","TK","TL","TM","TN","TO","TR","TT","TV","TW","TZ","UA","UG","UM","UNKNOWN","US","UY","UZ","VA","VC","VE","VG","VI","VN","VU","WF","WS","XK","YE","YT","ZA","ZM","ZW"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		contactAddressTypeCountryPropEnum = append(contactAddressTypeCountryPropEnum, v)
	}
}

const (
	// ContactAddressCountryAC captures enum value "AC"
	ContactAddressCountryAC string = "AC"
	// ContactAddressCountryAD captures enum value "AD"
	ContactAddressCountryAD string = "AD"
	// ContactAddressCountryAE captures enum value "AE"
	ContactAddressCountryAE string = "AE"
	// ContactAddressCountryAF captures enum value "AF"
	ContactAddressCountryAF string = "AF"
	// ContactAddressCountryAG captures enum value "AG"
	ContactAddressCountryAG string = "AG"
	// ContactAddressCountryAI captures enum value "AI"
	ContactAddressCountryAI string = "AI"
	// ContactAddressCountryAL captures enum value "AL"
	ContactAddressCountryAL string = "AL"
	// ContactAddressCountryAM captures enum value "AM"
	ContactAddressCountryAM string = "AM"
	// ContactAddressCountryAO captures enum value "AO"
	ContactAddressCountryAO string = "AO"
	// ContactAddressCountryAQ captures enum value "AQ"
	ContactAddressCountryAQ string = "AQ"
	// ContactAddressCountryAR captures enum value "AR"
	ContactAddressCountryAR string = "AR"
	// ContactAddressCountryAS captures enum value "AS"
	ContactAddressCountryAS string = "AS"
	// ContactAddressCountryAT captures enum value "AT"
	ContactAddressCountryAT string = "AT"
	// ContactAddressCountryAU captures enum value "AU"
	ContactAddressCountryAU string = "AU"
	// ContactAddressCountryAW captures enum value "AW"
	ContactAddressCountryAW string = "AW"
	// ContactAddressCountryAX captures enum value "AX"
	ContactAddressCountryAX string = "AX"
	// ContactAddressCountryAZ captures enum value "AZ"
	ContactAddressCountryAZ string = "AZ"
	// ContactAddressCountryBA captures enum value "BA"
	ContactAddressCountryBA string = "BA"
	// ContactAddressCountryBB captures enum value "BB"
	ContactAddressCountryBB string = "BB"
	// ContactAddressCountryBD captures enum value "BD"
	ContactAddressCountryBD string = "BD"
	// ContactAddressCountryBE captures enum value "BE"
	ContactAddressCountryBE string = "BE"
	// ContactAddressCountryBF captures enum value "BF"
	ContactAddressCountryBF string = "BF"
	// ContactAddressCountryBG captures enum value "BG"
	ContactAddressCountryBG string = "BG"
	// ContactAddressCountryBH captures enum value "BH"
	ContactAddressCountryBH string = "BH"
	// ContactAddressCountryBI captures enum value "BI"
	ContactAddressCountryBI string = "BI"
	// ContactAddressCountryBJ captures enum value "BJ"
	ContactAddressCountryBJ string = "BJ"
	// ContactAddressCountryBL captures enum value "BL"
	ContactAddressCountryBL string = "BL"
	// ContactAddressCountryBM captures enum value "BM"
	ContactAddressCountryBM string = "BM"
	// ContactAddressCountryBN captures enum value "BN"
	ContactAddressCountryBN string = "BN"
	// ContactAddressCountryBO captures enum value "BO"
	ContactAddressCountryBO string = "BO"
	// ContactAddressCountryBQ captures enum value "BQ"
	ContactAddressCountryBQ string = "BQ"
	// ContactAddressCountryBR captures enum value "BR"
	ContactAddressCountryBR string = "BR"
	// ContactAddressCountryBS captures enum value "BS"
	ContactAddressCountryBS string = "BS"
	// ContactAddressCountryBT captures enum value "BT"
	ContactAddressCountryBT string = "BT"
	// ContactAddressCountryBW captures enum value "BW"
	ContactAddressCountryBW string = "BW"
	// ContactAddressCountryBY captures enum value "BY"
	ContactAddressCountryBY string = "BY"
	// ContactAddressCountryBZ captures enum value "BZ"
	ContactAddressCountryBZ string = "BZ"
	// ContactAddressCountryCA captures enum value "CA"
	ContactAddressCountryCA string = "CA"
	// ContactAddressCountryCC captures enum value "CC"
	ContactAddressCountryCC string = "CC"
	// ContactAddressCountryCD captures enum value "CD"
	ContactAddressCountryCD string = "CD"
	// ContactAddressCountryCF captures enum value "CF"
	ContactAddressCountryCF string = "CF"
	// ContactAddressCountryCG captures enum value "CG"
	ContactAddressCountryCG string = "CG"
	// ContactAddressCountryCH captures enum value "CH"
	ContactAddressCountryCH string = "CH"
	// ContactAddressCountryCI captures enum value "CI"
	ContactAddressCountryCI string = "CI"
	// ContactAddressCountryCK captures enum value "CK"
	ContactAddressCountryCK string = "CK"
	// ContactAddressCountryCL captures enum value "CL"
	ContactAddressCountryCL string = "CL"
	// ContactAddressCountryCM captures enum value "CM"
	ContactAddressCountryCM string = "CM"
	// ContactAddressCountryCN captures enum value "CN"
	ContactAddressCountryCN string = "CN"
	// ContactAddressCountryCO captures enum value "CO"
	ContactAddressCountryCO string = "CO"
	// ContactAddressCountryCR captures enum value "CR"
	ContactAddressCountryCR string = "CR"
	// ContactAddressCountryCU captures enum value "CU"
	ContactAddressCountryCU string = "CU"
	// ContactAddressCountryCV captures enum value "CV"
	ContactAddressCountryCV string = "CV"
	// ContactAddressCountryCW captures enum value "CW"
	ContactAddressCountryCW string = "CW"
	// ContactAddressCountryCX captures enum value "CX"
	ContactAddressCountryCX string = "CX"
	// ContactAddressCountryCY captures enum value "CY"
	ContactAddressCountryCY string = "CY"
	// ContactAddressCountryCZ captures enum value "CZ"
	ContactAddressCountryCZ string = "CZ"
	// ContactAddressCountryDE captures enum value "DE"
	ContactAddressCountryDE string = "DE"
	// ContactAddressCountryDG captures enum value "DG"
	ContactAddressCountryDG string = "DG"
	// ContactAddressCountryDJ captures enum value "DJ"
	ContactAddressCountryDJ string = "DJ"
	// ContactAddressCountryDK captures enum value "DK"
	ContactAddressCountryDK string = "DK"
	// ContactAddressCountryDM captures enum value "DM"
	ContactAddressCountryDM string = "DM"
	// ContactAddressCountryDO captures enum value "DO"
	ContactAddressCountryDO string = "DO"
	// ContactAddressCountryDZ captures enum value "DZ"
	ContactAddressCountryDZ string = "DZ"
	// ContactAddressCountryEA captures enum value "EA"
	ContactAddressCountryEA string = "EA"
	// ContactAddressCountryEC captures enum value "EC"
	ContactAddressCountryEC string = "EC"
	// ContactAddressCountryEE captures enum value "EE"
	ContactAddressCountryEE string = "EE"
	// ContactAddressCountryEG captures enum value "EG"
	ContactAddressCountryEG string = "EG"
	// ContactAddressCountryEH captures enum value "EH"
	ContactAddressCountryEH string = "EH"
	// ContactAddressCountryER captures enum value "ER"
	ContactAddressCountryER string = "ER"
	// ContactAddressCountryES captures enum value "ES"
	ContactAddressCountryES string = "ES"
	// ContactAddressCountryET captures enum value "ET"
	ContactAddressCountryET string = "ET"
	// ContactAddressCountryFI captures enum value "FI"
	ContactAddressCountryFI string = "FI"
	// ContactAddressCountryFJ captures enum value "FJ"
	ContactAddressCountryFJ string = "FJ"
	// ContactAddressCountryFK captures enum value "FK"
	ContactAddressCountryFK string = "FK"
	// ContactAddressCountryFM captures enum value "FM"
	ContactAddressCountryFM string = "FM"
	// ContactAddressCountryFO captures enum value "FO"
	ContactAddressCountryFO string = "FO"
	// ContactAddressCountryFR captures enum value "FR"
	ContactAddressCountryFR string = "FR"
	// ContactAddressCountryGA captures enum value "GA"
	ContactAddressCountryGA string = "GA"
	// ContactAddressCountryGB captures enum value "GB"
	ContactAddressCountryGB string = "GB"
	// ContactAddressCountryGD captures enum value "GD"
	ContactAddressCountryGD string = "GD"
	// ContactAddressCountryGE captures enum value "GE"
	ContactAddressCountryGE string = "GE"
	// ContactAddressCountryGF captures enum value "GF"
	ContactAddressCountryGF string = "GF"
	// ContactAddressCountryGG captures enum value "GG"
	ContactAddressCountryGG string = "GG"
	// ContactAddressCountryGH captures enum value "GH"
	ContactAddressCountryGH string = "GH"
	// ContactAddressCountryGI captures enum value "GI"
	ContactAddressCountryGI string = "GI"
	// ContactAddressCountryGL captures enum value "GL"
	ContactAddressCountryGL string = "GL"
	// ContactAddressCountryGM captures enum value "GM"
	ContactAddressCountryGM string = "GM"
	// ContactAddressCountryGN captures enum value "GN"
	ContactAddressCountryGN string = "GN"
	// ContactAddressCountryGP captures enum value "GP"
	ContactAddressCountryGP string = "GP"
	// ContactAddressCountryGQ captures enum value "GQ"
	ContactAddressCountryGQ string = "GQ"
	// ContactAddressCountryGR captures enum value "GR"
	ContactAddressCountryGR string = "GR"
	// ContactAddressCountryGS captures enum value "GS"
	ContactAddressCountryGS string = "GS"
	// ContactAddressCountryGT captures enum value "GT"
	ContactAddressCountryGT string = "GT"
	// ContactAddressCountryGU captures enum value "GU"
	ContactAddressCountryGU string = "GU"
	// ContactAddressCountryGW captures enum value "GW"
	ContactAddressCountryGW string = "GW"
	// ContactAddressCountryGY captures enum value "GY"
	ContactAddressCountryGY string = "GY"
	// ContactAddressCountryHK captures enum value "HK"
	ContactAddressCountryHK string = "HK"
	// ContactAddressCountryHN captures enum value "HN"
	ContactAddressCountryHN string = "HN"
	// ContactAddressCountryHR captures enum value "HR"
	ContactAddressCountryHR string = "HR"
	// ContactAddressCountryHT captures enum value "HT"
	ContactAddressCountryHT string = "HT"
	// ContactAddressCountryHU captures enum value "HU"
	ContactAddressCountryHU string = "HU"
	// ContactAddressCountryIC captures enum value "IC"
	ContactAddressCountryIC string = "IC"
	// ContactAddressCountryID captures enum value "ID"
	ContactAddressCountryID string = "ID"
	// ContactAddressCountryIE captures enum value "IE"
	ContactAddressCountryIE string = "IE"
	// ContactAddressCountryIL captures enum value "IL"
	ContactAddressCountryIL string = "IL"
	// ContactAddressCountryIM captures enum value "IM"
	ContactAddressCountryIM string = "IM"
	// ContactAddressCountryIN captures enum value "IN"
	ContactAddressCountryIN string = "IN"
	// ContactAddressCountryIO captures enum value "IO"
	ContactAddressCountryIO string = "IO"
	// ContactAddressCountryIQ captures enum value "IQ"
	ContactAddressCountryIQ string = "IQ"
	// ContactAddressCountryIR captures enum value "IR"
	ContactAddressCountryIR string = "IR"
	// ContactAddressCountryIS captures enum value "IS"
	ContactAddressCountryIS string = "IS"
	// ContactAddressCountryIT captures enum value "IT"
	ContactAddressCountryIT string = "IT"
	// ContactAddressCountryJE captures enum value "JE"
	ContactAddressCountryJE string = "JE"
	// ContactAddressCountryJM captures enum value "JM"
	ContactAddressCountryJM string = "JM"
	// ContactAddressCountryJO captures enum value "JO"
	ContactAddressCountryJO string = "JO"
	// ContactAddressCountryJP captures enum value "JP"
	ContactAddressCountryJP string = "JP"
	// ContactAddressCountryKE captures enum value "KE"
	ContactAddressCountryKE string = "KE"
	// ContactAddressCountryKG captures enum value "KG"
	ContactAddressCountryKG string = "KG"
	// ContactAddressCountryKH captures enum value "KH"
	ContactAddressCountryKH string = "KH"
	// ContactAddressCountryKI captures enum value "KI"
	ContactAddressCountryKI string = "KI"
	// ContactAddressCountryKM captures enum value "KM"
	ContactAddressCountryKM string = "KM"
	// ContactAddressCountryKN captures enum value "KN"
	ContactAddressCountryKN string = "KN"
	// ContactAddressCountryKP captures enum value "KP"
	ContactAddressCountryKP string = "KP"
	// ContactAddressCountryKR captures enum value "KR"
	ContactAddressCountryKR string = "KR"
	// ContactAddressCountryKW captures enum value "KW"
	ContactAddressCountryKW string = "KW"
	// ContactAddressCountryKY captures enum value "KY"
	ContactAddressCountryKY string = "KY"
	// ContactAddressCountryKZ captures enum value "KZ"
	ContactAddressCountryKZ string = "KZ"
	// ContactAddressCountryLA captures enum value "LA"
	ContactAddressCountryLA string = "LA"
	// ContactAddressCountryLB captures enum value "LB"
	ContactAddressCountryLB string = "LB"
	// ContactAddressCountryLC captures enum value "LC"
	ContactAddressCountryLC string = "LC"
	// ContactAddressCountryLI captures enum value "LI"
	ContactAddressCountryLI string = "LI"
	// ContactAddressCountryLK captures enum value "LK"
	ContactAddressCountryLK string = "LK"
	// ContactAddressCountryLR captures enum value "LR"
	ContactAddressCountryLR string = "LR"
	// ContactAddressCountryLS captures enum value "LS"
	ContactAddressCountryLS string = "LS"
	// ContactAddressCountryLT captures enum value "LT"
	ContactAddressCountryLT string = "LT"
	// ContactAddressCountryLU captures enum value "LU"
	ContactAddressCountryLU string = "LU"
	// ContactAddressCountryLV captures enum value "LV"
	ContactAddressCountryLV string = "LV"
	// ContactAddressCountryLY captures enum value "LY"
	ContactAddressCountryLY string = "LY"
	// ContactAddressCountryMA captures enum value "MA"
	ContactAddressCountryMA string = "MA"
	// ContactAddressCountryMC captures enum value "MC"
	ContactAddressCountryMC string = "MC"
	// ContactAddressCountryMD captures enum value "MD"
	ContactAddressCountryMD string = "MD"
	// ContactAddressCountryME captures enum value "ME"
	ContactAddressCountryME string = "ME"
	// ContactAddressCountryMF captures enum value "MF"
	ContactAddressCountryMF string = "MF"
	// ContactAddressCountryMG captures enum value "MG"
	ContactAddressCountryMG string = "MG"
	// ContactAddressCountryMH captures enum value "MH"
	ContactAddressCountryMH string = "MH"
	// ContactAddressCountryMK captures enum value "MK"
	ContactAddressCountryMK string = "MK"
	// ContactAddressCountryML captures enum value "ML"
	ContactAddressCountryML string = "ML"
	// ContactAddressCountryMM captures enum value "MM"
	ContactAddressCountryMM string = "MM"
	// ContactAddressCountryMN captures enum value "MN"
	ContactAddressCountryMN string = "MN"
	// ContactAddressCountryMO captures enum value "MO"
	ContactAddressCountryMO string = "MO"
	// ContactAddressCountryMP captures enum value "MP"
	ContactAddressCountryMP string = "MP"
	// ContactAddressCountryMQ captures enum value "MQ"
	ContactAddressCountryMQ string = "MQ"
	// ContactAddressCountryMR captures enum value "MR"
	ContactAddressCountryMR string = "MR"
	// ContactAddressCountryMS captures enum value "MS"
	ContactAddressCountryMS string = "MS"
	// ContactAddressCountryMT captures enum value "MT"
	ContactAddressCountryMT string = "MT"
	// ContactAddressCountryMU captures enum value "MU"
	ContactAddressCountryMU string = "MU"
	// ContactAddressCountryMV captures enum value "MV"
	ContactAddressCountryMV string = "MV"
	// ContactAddressCountryMW captures enum value "MW"
	ContactAddressCountryMW string = "MW"
	// ContactAddressCountryMX captures enum value "MX"
	ContactAddressCountryMX string = "MX"
	// ContactAddressCountryMY captures enum value "MY"
	ContactAddressCountryMY string = "MY"
	// ContactAddressCountryMZ captures enum value "MZ"
	ContactAddressCountryMZ string = "MZ"
	// ContactAddressCountryNA captures enum value "NA"
	ContactAddressCountryNA string = "NA"
	// ContactAddressCountryNC captures enum value "NC"
	ContactAddressCountryNC string = "NC"
	// ContactAddressCountryNE captures enum value "NE"
	ContactAddressCountryNE string = "NE"
	// ContactAddressCountryNF captures enum value "NF"
	ContactAddressCountryNF string = "NF"
	// ContactAddressCountryNG captures enum value "NG"
	ContactAddressCountryNG string = "NG"
	// ContactAddressCountryNI captures enum value "NI"
	ContactAddressCountryNI string = "NI"
	// ContactAddressCountryNL captures enum value "NL"
	ContactAddressCountryNL string = "NL"
	// ContactAddressCountryNO captures enum value "NO"
	ContactAddressCountryNO string = "NO"
	// ContactAddressCountryNP captures enum value "NP"
	ContactAddressCountryNP string = "NP"
	// ContactAddressCountryNR captures enum value "NR"
	ContactAddressCountryNR string = "NR"
	// ContactAddressCountryNU captures enum value "NU"
	ContactAddressCountryNU string = "NU"
	// ContactAddressCountryNZ captures enum value "NZ"
	ContactAddressCountryNZ string = "NZ"
	// ContactAddressCountryOM captures enum value "OM"
	ContactAddressCountryOM string = "OM"
	// ContactAddressCountryPA captures enum value "PA"
	ContactAddressCountryPA string = "PA"
	// ContactAddressCountryPE captures enum value "PE"
	ContactAddressCountryPE string = "PE"
	// ContactAddressCountryPF captures enum value "PF"
	ContactAddressCountryPF string = "PF"
	// ContactAddressCountryPG captures enum value "PG"
	ContactAddressCountryPG string = "PG"
	// ContactAddressCountryPH captures enum value "PH"
	ContactAddressCountryPH string = "PH"
	// ContactAddressCountryPK captures enum value "PK"
	ContactAddressCountryPK string = "PK"
	// ContactAddressCountryPL captures enum value "PL"
	ContactAddressCountryPL string = "PL"
	// ContactAddressCountryPM captures enum value "PM"
	ContactAddressCountryPM string = "PM"
	// ContactAddressCountryPN captures enum value "PN"
	ContactAddressCountryPN string = "PN"
	// ContactAddressCountryPR captures enum value "PR"
	ContactAddressCountryPR string = "PR"
	// ContactAddressCountryPS captures enum value "PS"
	ContactAddressCountryPS string = "PS"
	// ContactAddressCountryPT captures enum value "PT"
	ContactAddressCountryPT string = "PT"
	// ContactAddressCountryPW captures enum value "PW"
	ContactAddressCountryPW string = "PW"
	// ContactAddressCountryPY captures enum value "PY"
	ContactAddressCountryPY string = "PY"
	// ContactAddressCountryQA captures enum value "QA"
	ContactAddressCountryQA string = "QA"
	// ContactAddressCountryRE captures enum value "RE"
	ContactAddressCountryRE string = "RE"
	// ContactAddressCountryRO captures enum value "RO"
	ContactAddressCountryRO string = "RO"
	// ContactAddressCountryRS captures enum value "RS"
	ContactAddressCountryRS string = "RS"
	// ContactAddressCountryRU captures enum value "RU"
	ContactAddressCountryRU string = "RU"
	// ContactAddressCountryRW captures enum value "RW"
	ContactAddressCountryRW string = "RW"
	// ContactAddressCountrySA captures enum value "SA"
	ContactAddressCountrySA string = "SA"
	// ContactAddressCountrySB captures enum value "SB"
	ContactAddressCountrySB string = "SB"
	// ContactAddressCountrySC captures enum value "SC"
	ContactAddressCountrySC string = "SC"
	// ContactAddressCountrySD captures enum value "SD"
	ContactAddressCountrySD string = "SD"
	// ContactAddressCountrySE captures enum value "SE"
	ContactAddressCountrySE string = "SE"
	// ContactAddressCountrySG captures enum value "SG"
	ContactAddressCountrySG string = "SG"
	// ContactAddressCountrySH captures enum value "SH"
	ContactAddressCountrySH string = "SH"
	// ContactAddressCountrySI captures enum value "SI"
	ContactAddressCountrySI string = "SI"
	// ContactAddressCountrySJ captures enum value "SJ"
	ContactAddressCountrySJ string = "SJ"
	// ContactAddressCountrySK captures enum value "SK"
	ContactAddressCountrySK string = "SK"
	// ContactAddressCountrySL captures enum value "SL"
	ContactAddressCountrySL string = "SL"
	// ContactAddressCountrySM captures enum value "SM"
	ContactAddressCountrySM string = "SM"
	// ContactAddressCountrySN captures enum value "SN"
	ContactAddressCountrySN string = "SN"
	// ContactAddressCountrySO captures enum value "SO"
	ContactAddressCountrySO string = "SO"
	// ContactAddressCountrySR captures enum value "SR"
	ContactAddressCountrySR string = "SR"
	// ContactAddressCountrySS captures enum value "SS"
	ContactAddressCountrySS string = "SS"
	// ContactAddressCountryST captures enum value "ST"
	ContactAddressCountryST string = "ST"
	// ContactAddressCountrySV captures enum value "SV"
	ContactAddressCountrySV string = "SV"
	// ContactAddressCountrySX captures enum value "SX"
	ContactAddressCountrySX string = "SX"
	// ContactAddressCountrySY captures enum value "SY"
	ContactAddressCountrySY string = "SY"
	// ContactAddressCountrySZ captures enum value "SZ"
	ContactAddressCountrySZ string = "SZ"
	// ContactAddressCountryTA captures enum value "TA"
	ContactAddressCountryTA string = "TA"
	// ContactAddressCountryTC captures enum value "TC"
	ContactAddressCountryTC string = "TC"
	// ContactAddressCountryTD captures enum value "TD"
	ContactAddressCountryTD string = "TD"
	// ContactAddressCountryTF captures enum value "TF"
	ContactAddressCountryTF string = "TF"
	// ContactAddressCountryTG captures enum value "TG"
	ContactAddressCountryTG string = "TG"
	// ContactAddressCountryTH captures enum value "TH"
	ContactAddressCountryTH string = "TH"
	// ContactAddressCountryTJ captures enum value "TJ"
	ContactAddressCountryTJ string = "TJ"
	// ContactAddressCountryTK captures enum value "TK"
	ContactAddressCountryTK string = "TK"
	// ContactAddressCountryTL captures enum value "TL"
	ContactAddressCountryTL string = "TL"
	// ContactAddressCountryTM captures enum value "TM"
	ContactAddressCountryTM string = "TM"
	// ContactAddressCountryTN captures enum value "TN"
	ContactAddressCountryTN string = "TN"
	// ContactAddressCountryTO captures enum value "TO"
	ContactAddressCountryTO string = "TO"
	// ContactAddressCountryTR captures enum value "TR"
	ContactAddressCountryTR string = "TR"
	// ContactAddressCountryTT captures enum value "TT"
	ContactAddressCountryTT string = "TT"
	// ContactAddressCountryTV captures enum value "TV"
	ContactAddressCountryTV string = "TV"
	// ContactAddressCountryTW captures enum value "TW"
	ContactAddressCountryTW string = "TW"
	// ContactAddressCountryTZ captures enum value "TZ"
	ContactAddressCountryTZ string = "TZ"
	// ContactAddressCountryUA captures enum value "UA"
	ContactAddressCountryUA string = "UA"
	// ContactAddressCountryUG captures enum value "UG"
	ContactAddressCountryUG string = "UG"
	// ContactAddressCountryUM captures enum value "UM"
	ContactAddressCountryUM string = "UM"
	// ContactAddressCountryUNKNOWN captures enum value "UNKNOWN"
	ContactAddressCountryUNKNOWN string = "UNKNOWN"
	// ContactAddressCountryUS captures enum value "US"
	ContactAddressCountryUS string = "US"
	// ContactAddressCountryUY captures enum value "UY"
	ContactAddressCountryUY string = "UY"
	// ContactAddressCountryUZ captures enum value "UZ"
	ContactAddressCountryUZ string = "UZ"
	// ContactAddressCountryVA captures enum value "VA"
	ContactAddressCountryVA string = "VA"
	// ContactAddressCountryVC captures enum value "VC"
	ContactAddressCountryVC string = "VC"
	// ContactAddressCountryVE captures enum value "VE"
	ContactAddressCountryVE string = "VE"
	// ContactAddressCountryVG captures enum value "VG"
	ContactAddressCountryVG string = "VG"
	// ContactAddressCountryVI captures enum value "VI"
	ContactAddressCountryVI string = "VI"
	// ContactAddressCountryVN captures enum value "VN"
	ContactAddressCountryVN string = "VN"
	// ContactAddressCountryVU captures enum value "VU"
	ContactAddressCountryVU string = "VU"
	// ContactAddressCountryWF captures enum value "WF"
	ContactAddressCountryWF string = "WF"
	// ContactAddressCountryWS captures enum value "WS"
	ContactAddressCountryWS string = "WS"
	// ContactAddressCountryXK captures enum value "XK"
	ContactAddressCountryXK string = "XK"
	// ContactAddressCountryYE captures enum value "YE"
	ContactAddressCountryYE string = "YE"
	// ContactAddressCountryYT captures enum value "YT"
	ContactAddressCountryYT string = "YT"
	// ContactAddressCountryZA captures enum value "ZA"
	ContactAddressCountryZA string = "ZA"
	// ContactAddressCountryZM captures enum value "ZM"
	ContactAddressCountryZM string = "ZM"
	// ContactAddressCountryZW captures enum value "ZW"
	ContactAddressCountryZW string = "ZW"
)

// prop value enum
func (m *ContactAddress) validateCountryEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, contactAddressTypeCountryPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ContactAddress) validateCountry(formats strfmt.Registry) error {

	if swag.IsZero(m.Country) { // not required
		return nil
	}

	// value enum
	if err := m.validateCountryEnum("country", "body", m.Country); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContactAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactAddress) UnmarshalBinary(b []byte) error {
	var res ContactAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
