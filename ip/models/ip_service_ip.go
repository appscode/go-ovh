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

// IPServiceIP Your IP linked to service
// swagger:model ip.ServiceIp
type IPServiceIP struct {

	// can be terminated
	// Required: true
	// Read Only: true
	CanBeTerminated bool `json:"canBeTerminated"`

	// country
	// Read Only: true
	Country string `json:"country,omitempty"`

	// Custom description on your ip
	Description string `json:"description,omitempty"`

	// ip
	// Required: true
	// Read Only: true
	IP string `json:"ip"`

	// IP block organisation Id
	// Read Only: true
	OrganisationID string `json:"organisationId,omitempty"`

	// routed to
	RoutedTo *IPRoutedTo `json:"routedTo,omitempty"`

	// type
	// Required: true
	// Read Only: true
	Type string `json:"type"`
}

// Validate validates this ip service Ip
func (m *IPServiceIP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCanBeTerminated(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRoutedTo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPServiceIP) validateCanBeTerminated(formats strfmt.Registry) error {

	if err := validate.Required("canBeTerminated", "body", bool(m.CanBeTerminated)); err != nil {
		return err
	}

	return nil
}

var ipServiceIpTypeCountryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ac","ad","ae","af","ag","ai","al","am","an","ao","aq","ar","as","at","au","aw","ax","az","ba","bb","bd","be","bf","bg","bh","bi","bj","bl","bm","bn","bo","bq","br","bs","bt","bv","bw","by","bz","ca","cc","cd","cf","cg","ch","ci","ck","cl","cm","cn","co","cr","cs","cu","cv","cw","cx","cy","cz","de","dj","dk","dm","do","dz","ec","ee","eg","eh","er","es","et","fc","fd","fi","fj","fk","fm","fo","fr","fx","ga","gb","gd","ge","gf","gg","gh","gi","gl","gm","gn","gp","gq","gr","gs","gt","gu","gw","gy","hk","hm","hn","hr","ht","hu","id","ie","il","im","in","io","iq","ir","is","it","je","jm","jo","jp","ke","kg","kh","ki","km","kn","kp","kr","kw","ky","kz","la","lb","lc","li","lk","lr","ls","lt","lu","lv","ly","ma","mc","md","me","mf","mg","mh","mk","ml","mm","mn","mo","mp","mq","mr","ms","mt","mu","mv","mw","mx","my","mz","na","nc","ne","nf","ng","ni","nl","no","np","nr","nu","nz","om","pa","pe","pf","pg","ph","pk","pl","pm","pn","pr","ps","pt","pw","py","qa","qc","re","ro","rs","ru","rw","sa","sb","sc","sd","se","sg","sh","si","sj","sk","sl","sm","sn","so","sr","ss","st","sv","sx","sy","sz","tc","td","tf","tg","th","tj","tk","tl","tm","tn","to","tp","tr","tt","tv","tw","tz","ua","ug","uk","um","us","uy","uz","va","vc","ve","vg","vi","vn","vu","we","wf","ws","ye","yt","yu","za","zm","zw"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipServiceIpTypeCountryPropEnum = append(ipServiceIpTypeCountryPropEnum, v)
	}
}

const (
	// IPServiceIPCountryAc captures enum value "ac"
	IPServiceIPCountryAc string = "ac"
	// IPServiceIPCountryAd captures enum value "ad"
	IPServiceIPCountryAd string = "ad"
	// IPServiceIPCountryAe captures enum value "ae"
	IPServiceIPCountryAe string = "ae"
	// IPServiceIPCountryAf captures enum value "af"
	IPServiceIPCountryAf string = "af"
	// IPServiceIPCountryAg captures enum value "ag"
	IPServiceIPCountryAg string = "ag"
	// IPServiceIPCountryAi captures enum value "ai"
	IPServiceIPCountryAi string = "ai"
	// IPServiceIPCountryAl captures enum value "al"
	IPServiceIPCountryAl string = "al"
	// IPServiceIPCountryAm captures enum value "am"
	IPServiceIPCountryAm string = "am"
	// IPServiceIPCountryAn captures enum value "an"
	IPServiceIPCountryAn string = "an"
	// IPServiceIPCountryAo captures enum value "ao"
	IPServiceIPCountryAo string = "ao"
	// IPServiceIPCountryAq captures enum value "aq"
	IPServiceIPCountryAq string = "aq"
	// IPServiceIPCountryAr captures enum value "ar"
	IPServiceIPCountryAr string = "ar"
	// IPServiceIPCountryAs captures enum value "as"
	IPServiceIPCountryAs string = "as"
	// IPServiceIPCountryAt captures enum value "at"
	IPServiceIPCountryAt string = "at"
	// IPServiceIPCountryAu captures enum value "au"
	IPServiceIPCountryAu string = "au"
	// IPServiceIPCountryAw captures enum value "aw"
	IPServiceIPCountryAw string = "aw"
	// IPServiceIPCountryAx captures enum value "ax"
	IPServiceIPCountryAx string = "ax"
	// IPServiceIPCountryAz captures enum value "az"
	IPServiceIPCountryAz string = "az"
	// IPServiceIPCountryBa captures enum value "ba"
	IPServiceIPCountryBa string = "ba"
	// IPServiceIPCountryBb captures enum value "bb"
	IPServiceIPCountryBb string = "bb"
	// IPServiceIPCountryBd captures enum value "bd"
	IPServiceIPCountryBd string = "bd"
	// IPServiceIPCountryBe captures enum value "be"
	IPServiceIPCountryBe string = "be"
	// IPServiceIPCountryBf captures enum value "bf"
	IPServiceIPCountryBf string = "bf"
	// IPServiceIPCountryBg captures enum value "bg"
	IPServiceIPCountryBg string = "bg"
	// IPServiceIPCountryBh captures enum value "bh"
	IPServiceIPCountryBh string = "bh"
	// IPServiceIPCountryBi captures enum value "bi"
	IPServiceIPCountryBi string = "bi"
	// IPServiceIPCountryBj captures enum value "bj"
	IPServiceIPCountryBj string = "bj"
	// IPServiceIPCountryBl captures enum value "bl"
	IPServiceIPCountryBl string = "bl"
	// IPServiceIPCountryBm captures enum value "bm"
	IPServiceIPCountryBm string = "bm"
	// IPServiceIPCountryBn captures enum value "bn"
	IPServiceIPCountryBn string = "bn"
	// IPServiceIPCountryBo captures enum value "bo"
	IPServiceIPCountryBo string = "bo"
	// IPServiceIPCountryBq captures enum value "bq"
	IPServiceIPCountryBq string = "bq"
	// IPServiceIPCountryBr captures enum value "br"
	IPServiceIPCountryBr string = "br"
	// IPServiceIPCountryBs captures enum value "bs"
	IPServiceIPCountryBs string = "bs"
	// IPServiceIPCountryBt captures enum value "bt"
	IPServiceIPCountryBt string = "bt"
	// IPServiceIPCountryBv captures enum value "bv"
	IPServiceIPCountryBv string = "bv"
	// IPServiceIPCountryBw captures enum value "bw"
	IPServiceIPCountryBw string = "bw"
	// IPServiceIPCountryBy captures enum value "by"
	IPServiceIPCountryBy string = "by"
	// IPServiceIPCountryBz captures enum value "bz"
	IPServiceIPCountryBz string = "bz"
	// IPServiceIPCountryCa captures enum value "ca"
	IPServiceIPCountryCa string = "ca"
	// IPServiceIPCountryCc captures enum value "cc"
	IPServiceIPCountryCc string = "cc"
	// IPServiceIPCountryCd captures enum value "cd"
	IPServiceIPCountryCd string = "cd"
	// IPServiceIPCountryCf captures enum value "cf"
	IPServiceIPCountryCf string = "cf"
	// IPServiceIPCountryCg captures enum value "cg"
	IPServiceIPCountryCg string = "cg"
	// IPServiceIPCountryCh captures enum value "ch"
	IPServiceIPCountryCh string = "ch"
	// IPServiceIPCountryCi captures enum value "ci"
	IPServiceIPCountryCi string = "ci"
	// IPServiceIPCountryCk captures enum value "ck"
	IPServiceIPCountryCk string = "ck"
	// IPServiceIPCountryCl captures enum value "cl"
	IPServiceIPCountryCl string = "cl"
	// IPServiceIPCountryCm captures enum value "cm"
	IPServiceIPCountryCm string = "cm"
	// IPServiceIPCountryCn captures enum value "cn"
	IPServiceIPCountryCn string = "cn"
	// IPServiceIPCountryCo captures enum value "co"
	IPServiceIPCountryCo string = "co"
	// IPServiceIPCountryCr captures enum value "cr"
	IPServiceIPCountryCr string = "cr"
	// IPServiceIPCountryCs captures enum value "cs"
	IPServiceIPCountryCs string = "cs"
	// IPServiceIPCountryCu captures enum value "cu"
	IPServiceIPCountryCu string = "cu"
	// IPServiceIPCountryCv captures enum value "cv"
	IPServiceIPCountryCv string = "cv"
	// IPServiceIPCountryCw captures enum value "cw"
	IPServiceIPCountryCw string = "cw"
	// IPServiceIPCountryCx captures enum value "cx"
	IPServiceIPCountryCx string = "cx"
	// IPServiceIPCountryCy captures enum value "cy"
	IPServiceIPCountryCy string = "cy"
	// IPServiceIPCountryCz captures enum value "cz"
	IPServiceIPCountryCz string = "cz"
	// IPServiceIPCountryDe captures enum value "de"
	IPServiceIPCountryDe string = "de"
	// IPServiceIPCountryDj captures enum value "dj"
	IPServiceIPCountryDj string = "dj"
	// IPServiceIPCountryDk captures enum value "dk"
	IPServiceIPCountryDk string = "dk"
	// IPServiceIPCountryDm captures enum value "dm"
	IPServiceIPCountryDm string = "dm"
	// IPServiceIPCountryDo captures enum value "do"
	IPServiceIPCountryDo string = "do"
	// IPServiceIPCountryDz captures enum value "dz"
	IPServiceIPCountryDz string = "dz"
	// IPServiceIPCountryEc captures enum value "ec"
	IPServiceIPCountryEc string = "ec"
	// IPServiceIPCountryEe captures enum value "ee"
	IPServiceIPCountryEe string = "ee"
	// IPServiceIPCountryEg captures enum value "eg"
	IPServiceIPCountryEg string = "eg"
	// IPServiceIPCountryEh captures enum value "eh"
	IPServiceIPCountryEh string = "eh"
	// IPServiceIPCountryEr captures enum value "er"
	IPServiceIPCountryEr string = "er"
	// IPServiceIPCountryEs captures enum value "es"
	IPServiceIPCountryEs string = "es"
	// IPServiceIPCountryEt captures enum value "et"
	IPServiceIPCountryEt string = "et"
	// IPServiceIPCountryFc captures enum value "fc"
	IPServiceIPCountryFc string = "fc"
	// IPServiceIPCountryFd captures enum value "fd"
	IPServiceIPCountryFd string = "fd"
	// IPServiceIPCountryFi captures enum value "fi"
	IPServiceIPCountryFi string = "fi"
	// IPServiceIPCountryFj captures enum value "fj"
	IPServiceIPCountryFj string = "fj"
	// IPServiceIPCountryFk captures enum value "fk"
	IPServiceIPCountryFk string = "fk"
	// IPServiceIPCountryFm captures enum value "fm"
	IPServiceIPCountryFm string = "fm"
	// IPServiceIPCountryFo captures enum value "fo"
	IPServiceIPCountryFo string = "fo"
	// IPServiceIPCountryFr captures enum value "fr"
	IPServiceIPCountryFr string = "fr"
	// IPServiceIPCountryFx captures enum value "fx"
	IPServiceIPCountryFx string = "fx"
	// IPServiceIPCountryGa captures enum value "ga"
	IPServiceIPCountryGa string = "ga"
	// IPServiceIPCountryGb captures enum value "gb"
	IPServiceIPCountryGb string = "gb"
	// IPServiceIPCountryGd captures enum value "gd"
	IPServiceIPCountryGd string = "gd"
	// IPServiceIPCountryGe captures enum value "ge"
	IPServiceIPCountryGe string = "ge"
	// IPServiceIPCountryGf captures enum value "gf"
	IPServiceIPCountryGf string = "gf"
	// IPServiceIPCountryGg captures enum value "gg"
	IPServiceIPCountryGg string = "gg"
	// IPServiceIPCountryGh captures enum value "gh"
	IPServiceIPCountryGh string = "gh"
	// IPServiceIPCountryGi captures enum value "gi"
	IPServiceIPCountryGi string = "gi"
	// IPServiceIPCountryGl captures enum value "gl"
	IPServiceIPCountryGl string = "gl"
	// IPServiceIPCountryGm captures enum value "gm"
	IPServiceIPCountryGm string = "gm"
	// IPServiceIPCountryGn captures enum value "gn"
	IPServiceIPCountryGn string = "gn"
	// IPServiceIPCountryGp captures enum value "gp"
	IPServiceIPCountryGp string = "gp"
	// IPServiceIPCountryGq captures enum value "gq"
	IPServiceIPCountryGq string = "gq"
	// IPServiceIPCountryGr captures enum value "gr"
	IPServiceIPCountryGr string = "gr"
	// IPServiceIPCountryGs captures enum value "gs"
	IPServiceIPCountryGs string = "gs"
	// IPServiceIPCountryGt captures enum value "gt"
	IPServiceIPCountryGt string = "gt"
	// IPServiceIPCountryGu captures enum value "gu"
	IPServiceIPCountryGu string = "gu"
	// IPServiceIPCountryGw captures enum value "gw"
	IPServiceIPCountryGw string = "gw"
	// IPServiceIPCountryGy captures enum value "gy"
	IPServiceIPCountryGy string = "gy"
	// IPServiceIPCountryHk captures enum value "hk"
	IPServiceIPCountryHk string = "hk"
	// IPServiceIPCountryHm captures enum value "hm"
	IPServiceIPCountryHm string = "hm"
	// IPServiceIPCountryHn captures enum value "hn"
	IPServiceIPCountryHn string = "hn"
	// IPServiceIPCountryHr captures enum value "hr"
	IPServiceIPCountryHr string = "hr"
	// IPServiceIPCountryHt captures enum value "ht"
	IPServiceIPCountryHt string = "ht"
	// IPServiceIPCountryHu captures enum value "hu"
	IPServiceIPCountryHu string = "hu"
	// IPServiceIPCountryID captures enum value "id"
	IPServiceIPCountryID string = "id"
	// IPServiceIPCountryIe captures enum value "ie"
	IPServiceIPCountryIe string = "ie"
	// IPServiceIPCountryIl captures enum value "il"
	IPServiceIPCountryIl string = "il"
	// IPServiceIPCountryIm captures enum value "im"
	IPServiceIPCountryIm string = "im"
	// IPServiceIPCountryIn captures enum value "in"
	IPServiceIPCountryIn string = "in"
	// IPServiceIPCountryIo captures enum value "io"
	IPServiceIPCountryIo string = "io"
	// IPServiceIPCountryIq captures enum value "iq"
	IPServiceIPCountryIq string = "iq"
	// IPServiceIPCountryIr captures enum value "ir"
	IPServiceIPCountryIr string = "ir"
	// IPServiceIPCountryIs captures enum value "is"
	IPServiceIPCountryIs string = "is"
	// IPServiceIPCountryIt captures enum value "it"
	IPServiceIPCountryIt string = "it"
	// IPServiceIPCountryJe captures enum value "je"
	IPServiceIPCountryJe string = "je"
	// IPServiceIPCountryJm captures enum value "jm"
	IPServiceIPCountryJm string = "jm"
	// IPServiceIPCountryJo captures enum value "jo"
	IPServiceIPCountryJo string = "jo"
	// IPServiceIPCountryJp captures enum value "jp"
	IPServiceIPCountryJp string = "jp"
	// IPServiceIPCountryKe captures enum value "ke"
	IPServiceIPCountryKe string = "ke"
	// IPServiceIPCountryKg captures enum value "kg"
	IPServiceIPCountryKg string = "kg"
	// IPServiceIPCountryKh captures enum value "kh"
	IPServiceIPCountryKh string = "kh"
	// IPServiceIPCountryKi captures enum value "ki"
	IPServiceIPCountryKi string = "ki"
	// IPServiceIPCountryKm captures enum value "km"
	IPServiceIPCountryKm string = "km"
	// IPServiceIPCountryKn captures enum value "kn"
	IPServiceIPCountryKn string = "kn"
	// IPServiceIPCountryKp captures enum value "kp"
	IPServiceIPCountryKp string = "kp"
	// IPServiceIPCountryKr captures enum value "kr"
	IPServiceIPCountryKr string = "kr"
	// IPServiceIPCountryKw captures enum value "kw"
	IPServiceIPCountryKw string = "kw"
	// IPServiceIPCountryKy captures enum value "ky"
	IPServiceIPCountryKy string = "ky"
	// IPServiceIPCountryKz captures enum value "kz"
	IPServiceIPCountryKz string = "kz"
	// IPServiceIPCountryLa captures enum value "la"
	IPServiceIPCountryLa string = "la"
	// IPServiceIPCountryLb captures enum value "lb"
	IPServiceIPCountryLb string = "lb"
	// IPServiceIPCountryLc captures enum value "lc"
	IPServiceIPCountryLc string = "lc"
	// IPServiceIPCountryLi captures enum value "li"
	IPServiceIPCountryLi string = "li"
	// IPServiceIPCountryLk captures enum value "lk"
	IPServiceIPCountryLk string = "lk"
	// IPServiceIPCountryLr captures enum value "lr"
	IPServiceIPCountryLr string = "lr"
	// IPServiceIPCountryLs captures enum value "ls"
	IPServiceIPCountryLs string = "ls"
	// IPServiceIPCountryLt captures enum value "lt"
	IPServiceIPCountryLt string = "lt"
	// IPServiceIPCountryLu captures enum value "lu"
	IPServiceIPCountryLu string = "lu"
	// IPServiceIPCountryLv captures enum value "lv"
	IPServiceIPCountryLv string = "lv"
	// IPServiceIPCountryLy captures enum value "ly"
	IPServiceIPCountryLy string = "ly"
	// IPServiceIPCountryMa captures enum value "ma"
	IPServiceIPCountryMa string = "ma"
	// IPServiceIPCountryMc captures enum value "mc"
	IPServiceIPCountryMc string = "mc"
	// IPServiceIPCountryMd captures enum value "md"
	IPServiceIPCountryMd string = "md"
	// IPServiceIPCountryMe captures enum value "me"
	IPServiceIPCountryMe string = "me"
	// IPServiceIPCountryMf captures enum value "mf"
	IPServiceIPCountryMf string = "mf"
	// IPServiceIPCountryMg captures enum value "mg"
	IPServiceIPCountryMg string = "mg"
	// IPServiceIPCountryMh captures enum value "mh"
	IPServiceIPCountryMh string = "mh"
	// IPServiceIPCountryMk captures enum value "mk"
	IPServiceIPCountryMk string = "mk"
	// IPServiceIPCountryMl captures enum value "ml"
	IPServiceIPCountryMl string = "ml"
	// IPServiceIPCountryMm captures enum value "mm"
	IPServiceIPCountryMm string = "mm"
	// IPServiceIPCountryMn captures enum value "mn"
	IPServiceIPCountryMn string = "mn"
	// IPServiceIPCountryMo captures enum value "mo"
	IPServiceIPCountryMo string = "mo"
	// IPServiceIPCountryMp captures enum value "mp"
	IPServiceIPCountryMp string = "mp"
	// IPServiceIPCountryMq captures enum value "mq"
	IPServiceIPCountryMq string = "mq"
	// IPServiceIPCountryMr captures enum value "mr"
	IPServiceIPCountryMr string = "mr"
	// IPServiceIPCountryMs captures enum value "ms"
	IPServiceIPCountryMs string = "ms"
	// IPServiceIPCountryMt captures enum value "mt"
	IPServiceIPCountryMt string = "mt"
	// IPServiceIPCountryMu captures enum value "mu"
	IPServiceIPCountryMu string = "mu"
	// IPServiceIPCountryMv captures enum value "mv"
	IPServiceIPCountryMv string = "mv"
	// IPServiceIPCountryMw captures enum value "mw"
	IPServiceIPCountryMw string = "mw"
	// IPServiceIPCountryMx captures enum value "mx"
	IPServiceIPCountryMx string = "mx"
	// IPServiceIPCountryMy captures enum value "my"
	IPServiceIPCountryMy string = "my"
	// IPServiceIPCountryMz captures enum value "mz"
	IPServiceIPCountryMz string = "mz"
	// IPServiceIPCountryNa captures enum value "na"
	IPServiceIPCountryNa string = "na"
	// IPServiceIPCountryNc captures enum value "nc"
	IPServiceIPCountryNc string = "nc"
	// IPServiceIPCountryNe captures enum value "ne"
	IPServiceIPCountryNe string = "ne"
	// IPServiceIPCountryNf captures enum value "nf"
	IPServiceIPCountryNf string = "nf"
	// IPServiceIPCountryNg captures enum value "ng"
	IPServiceIPCountryNg string = "ng"
	// IPServiceIPCountryNi captures enum value "ni"
	IPServiceIPCountryNi string = "ni"
	// IPServiceIPCountryNl captures enum value "nl"
	IPServiceIPCountryNl string = "nl"
	// IPServiceIPCountryNo captures enum value "no"
	IPServiceIPCountryNo string = "no"
	// IPServiceIPCountryNp captures enum value "np"
	IPServiceIPCountryNp string = "np"
	// IPServiceIPCountryNr captures enum value "nr"
	IPServiceIPCountryNr string = "nr"
	// IPServiceIPCountryNu captures enum value "nu"
	IPServiceIPCountryNu string = "nu"
	// IPServiceIPCountryNz captures enum value "nz"
	IPServiceIPCountryNz string = "nz"
	// IPServiceIPCountryOm captures enum value "om"
	IPServiceIPCountryOm string = "om"
	// IPServiceIPCountryPa captures enum value "pa"
	IPServiceIPCountryPa string = "pa"
	// IPServiceIPCountryPe captures enum value "pe"
	IPServiceIPCountryPe string = "pe"
	// IPServiceIPCountryPf captures enum value "pf"
	IPServiceIPCountryPf string = "pf"
	// IPServiceIPCountryPg captures enum value "pg"
	IPServiceIPCountryPg string = "pg"
	// IPServiceIPCountryPh captures enum value "ph"
	IPServiceIPCountryPh string = "ph"
	// IPServiceIPCountryPk captures enum value "pk"
	IPServiceIPCountryPk string = "pk"
	// IPServiceIPCountryPl captures enum value "pl"
	IPServiceIPCountryPl string = "pl"
	// IPServiceIPCountryPm captures enum value "pm"
	IPServiceIPCountryPm string = "pm"
	// IPServiceIPCountryPn captures enum value "pn"
	IPServiceIPCountryPn string = "pn"
	// IPServiceIPCountryPr captures enum value "pr"
	IPServiceIPCountryPr string = "pr"
	// IPServiceIPCountryPs captures enum value "ps"
	IPServiceIPCountryPs string = "ps"
	// IPServiceIPCountryPt captures enum value "pt"
	IPServiceIPCountryPt string = "pt"
	// IPServiceIPCountryPw captures enum value "pw"
	IPServiceIPCountryPw string = "pw"
	// IPServiceIPCountryPy captures enum value "py"
	IPServiceIPCountryPy string = "py"
	// IPServiceIPCountryQa captures enum value "qa"
	IPServiceIPCountryQa string = "qa"
	// IPServiceIPCountryQc captures enum value "qc"
	IPServiceIPCountryQc string = "qc"
	// IPServiceIPCountryRe captures enum value "re"
	IPServiceIPCountryRe string = "re"
	// IPServiceIPCountryRo captures enum value "ro"
	IPServiceIPCountryRo string = "ro"
	// IPServiceIPCountryRs captures enum value "rs"
	IPServiceIPCountryRs string = "rs"
	// IPServiceIPCountryRu captures enum value "ru"
	IPServiceIPCountryRu string = "ru"
	// IPServiceIPCountryRw captures enum value "rw"
	IPServiceIPCountryRw string = "rw"
	// IPServiceIPCountrySa captures enum value "sa"
	IPServiceIPCountrySa string = "sa"
	// IPServiceIPCountrySb captures enum value "sb"
	IPServiceIPCountrySb string = "sb"
	// IPServiceIPCountrySc captures enum value "sc"
	IPServiceIPCountrySc string = "sc"
	// IPServiceIPCountrySd captures enum value "sd"
	IPServiceIPCountrySd string = "sd"
	// IPServiceIPCountrySe captures enum value "se"
	IPServiceIPCountrySe string = "se"
	// IPServiceIPCountrySg captures enum value "sg"
	IPServiceIPCountrySg string = "sg"
	// IPServiceIPCountrySh captures enum value "sh"
	IPServiceIPCountrySh string = "sh"
	// IPServiceIPCountrySi captures enum value "si"
	IPServiceIPCountrySi string = "si"
	// IPServiceIPCountrySj captures enum value "sj"
	IPServiceIPCountrySj string = "sj"
	// IPServiceIPCountrySk captures enum value "sk"
	IPServiceIPCountrySk string = "sk"
	// IPServiceIPCountrySl captures enum value "sl"
	IPServiceIPCountrySl string = "sl"
	// IPServiceIPCountrySm captures enum value "sm"
	IPServiceIPCountrySm string = "sm"
	// IPServiceIPCountrySn captures enum value "sn"
	IPServiceIPCountrySn string = "sn"
	// IPServiceIPCountrySo captures enum value "so"
	IPServiceIPCountrySo string = "so"
	// IPServiceIPCountrySr captures enum value "sr"
	IPServiceIPCountrySr string = "sr"
	// IPServiceIPCountrySs captures enum value "ss"
	IPServiceIPCountrySs string = "ss"
	// IPServiceIPCountrySt captures enum value "st"
	IPServiceIPCountrySt string = "st"
	// IPServiceIPCountrySv captures enum value "sv"
	IPServiceIPCountrySv string = "sv"
	// IPServiceIPCountrySx captures enum value "sx"
	IPServiceIPCountrySx string = "sx"
	// IPServiceIPCountrySy captures enum value "sy"
	IPServiceIPCountrySy string = "sy"
	// IPServiceIPCountrySz captures enum value "sz"
	IPServiceIPCountrySz string = "sz"
	// IPServiceIPCountryTc captures enum value "tc"
	IPServiceIPCountryTc string = "tc"
	// IPServiceIPCountryTd captures enum value "td"
	IPServiceIPCountryTd string = "td"
	// IPServiceIPCountryTf captures enum value "tf"
	IPServiceIPCountryTf string = "tf"
	// IPServiceIPCountryTg captures enum value "tg"
	IPServiceIPCountryTg string = "tg"
	// IPServiceIPCountryTh captures enum value "th"
	IPServiceIPCountryTh string = "th"
	// IPServiceIPCountryTj captures enum value "tj"
	IPServiceIPCountryTj string = "tj"
	// IPServiceIPCountryTk captures enum value "tk"
	IPServiceIPCountryTk string = "tk"
	// IPServiceIPCountryTl captures enum value "tl"
	IPServiceIPCountryTl string = "tl"
	// IPServiceIPCountryTm captures enum value "tm"
	IPServiceIPCountryTm string = "tm"
	// IPServiceIPCountryTn captures enum value "tn"
	IPServiceIPCountryTn string = "tn"
	// IPServiceIPCountryTo captures enum value "to"
	IPServiceIPCountryTo string = "to"
	// IPServiceIPCountryTp captures enum value "tp"
	IPServiceIPCountryTp string = "tp"
	// IPServiceIPCountryTr captures enum value "tr"
	IPServiceIPCountryTr string = "tr"
	// IPServiceIPCountryTt captures enum value "tt"
	IPServiceIPCountryTt string = "tt"
	// IPServiceIPCountryTv captures enum value "tv"
	IPServiceIPCountryTv string = "tv"
	// IPServiceIPCountryTw captures enum value "tw"
	IPServiceIPCountryTw string = "tw"
	// IPServiceIPCountryTz captures enum value "tz"
	IPServiceIPCountryTz string = "tz"
	// IPServiceIPCountryUa captures enum value "ua"
	IPServiceIPCountryUa string = "ua"
	// IPServiceIPCountryUg captures enum value "ug"
	IPServiceIPCountryUg string = "ug"
	// IPServiceIPCountryUk captures enum value "uk"
	IPServiceIPCountryUk string = "uk"
	// IPServiceIPCountryUm captures enum value "um"
	IPServiceIPCountryUm string = "um"
	// IPServiceIPCountryUs captures enum value "us"
	IPServiceIPCountryUs string = "us"
	// IPServiceIPCountryUy captures enum value "uy"
	IPServiceIPCountryUy string = "uy"
	// IPServiceIPCountryUz captures enum value "uz"
	IPServiceIPCountryUz string = "uz"
	// IPServiceIPCountryVa captures enum value "va"
	IPServiceIPCountryVa string = "va"
	// IPServiceIPCountryVc captures enum value "vc"
	IPServiceIPCountryVc string = "vc"
	// IPServiceIPCountryVe captures enum value "ve"
	IPServiceIPCountryVe string = "ve"
	// IPServiceIPCountryVg captures enum value "vg"
	IPServiceIPCountryVg string = "vg"
	// IPServiceIPCountryVi captures enum value "vi"
	IPServiceIPCountryVi string = "vi"
	// IPServiceIPCountryVn captures enum value "vn"
	IPServiceIPCountryVn string = "vn"
	// IPServiceIPCountryVu captures enum value "vu"
	IPServiceIPCountryVu string = "vu"
	// IPServiceIPCountryWe captures enum value "we"
	IPServiceIPCountryWe string = "we"
	// IPServiceIPCountryWf captures enum value "wf"
	IPServiceIPCountryWf string = "wf"
	// IPServiceIPCountryWs captures enum value "ws"
	IPServiceIPCountryWs string = "ws"
	// IPServiceIPCountryYe captures enum value "ye"
	IPServiceIPCountryYe string = "ye"
	// IPServiceIPCountryYt captures enum value "yt"
	IPServiceIPCountryYt string = "yt"
	// IPServiceIPCountryYu captures enum value "yu"
	IPServiceIPCountryYu string = "yu"
	// IPServiceIPCountryZa captures enum value "za"
	IPServiceIPCountryZa string = "za"
	// IPServiceIPCountryZm captures enum value "zm"
	IPServiceIPCountryZm string = "zm"
	// IPServiceIPCountryZw captures enum value "zw"
	IPServiceIPCountryZw string = "zw"
)

// prop value enum
func (m *IPServiceIP) validateCountryEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ipServiceIpTypeCountryPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IPServiceIP) validateCountry(formats strfmt.Registry) error {

	if swag.IsZero(m.Country) { // not required
		return nil
	}

	// value enum
	if err := m.validateCountryEnum("country", "body", m.Country); err != nil {
		return err
	}

	return nil
}

func (m *IPServiceIP) validateIP(formats strfmt.Registry) error {

	if err := validate.RequiredString("ip", "body", string(m.IP)); err != nil {
		return err
	}

	return nil
}

func (m *IPServiceIP) validateRoutedTo(formats strfmt.Registry) error {

	if swag.IsZero(m.RoutedTo) { // not required
		return nil
	}

	if m.RoutedTo != nil {

		if err := m.RoutedTo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("routedTo")
			}
			return err
		}
	}

	return nil
}

var ipServiceIpTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cdn","cloud","dedicated","failover","hosted_ssl","housing","loadBalancing","mail","overthebox","pcc","pci","private","vpn","vps","vrack","xdsl"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipServiceIpTypeTypePropEnum = append(ipServiceIpTypeTypePropEnum, v)
	}
}

const (
	// IPServiceIPTypeCdn captures enum value "cdn"
	IPServiceIPTypeCdn string = "cdn"
	// IPServiceIPTypeCloud captures enum value "cloud"
	IPServiceIPTypeCloud string = "cloud"
	// IPServiceIPTypeDedicated captures enum value "dedicated"
	IPServiceIPTypeDedicated string = "dedicated"
	// IPServiceIPTypeFailover captures enum value "failover"
	IPServiceIPTypeFailover string = "failover"
	// IPServiceIPTypeHostedSsl captures enum value "hosted_ssl"
	IPServiceIPTypeHostedSsl string = "hosted_ssl"
	// IPServiceIPTypeHousing captures enum value "housing"
	IPServiceIPTypeHousing string = "housing"
	// IPServiceIPTypeLoadBalancing captures enum value "loadBalancing"
	IPServiceIPTypeLoadBalancing string = "loadBalancing"
	// IPServiceIPTypeMail captures enum value "mail"
	IPServiceIPTypeMail string = "mail"
	// IPServiceIPTypeOverthebox captures enum value "overthebox"
	IPServiceIPTypeOverthebox string = "overthebox"
	// IPServiceIPTypePcc captures enum value "pcc"
	IPServiceIPTypePcc string = "pcc"
	// IPServiceIPTypePci captures enum value "pci"
	IPServiceIPTypePci string = "pci"
	// IPServiceIPTypePrivate captures enum value "private"
	IPServiceIPTypePrivate string = "private"
	// IPServiceIPTypeVpn captures enum value "vpn"
	IPServiceIPTypeVpn string = "vpn"
	// IPServiceIPTypeVps captures enum value "vps"
	IPServiceIPTypeVps string = "vps"
	// IPServiceIPTypeVrack captures enum value "vrack"
	IPServiceIPTypeVrack string = "vrack"
	// IPServiceIPTypeXdsl captures enum value "xdsl"
	IPServiceIPTypeXdsl string = "xdsl"
)

// prop value enum
func (m *IPServiceIP) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ipServiceIpTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IPServiceIP) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", string(m.Type)); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPServiceIP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPServiceIP) UnmarshalBinary(b []byte) error {
	var res IPServiceIP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}