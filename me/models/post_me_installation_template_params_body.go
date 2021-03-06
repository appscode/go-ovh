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

// PostMeInstallationTemplateParamsBody post me installation template params body
// swagger:model postMeInstallationTemplateParamsBody
type PostMeInstallationTemplateParamsBody struct {

	// base template name
	// Required: true
	BaseTemplateName *string `json:"baseTemplateName"`

	// default language
	// Required: true
	DefaultLanguage *string `json:"defaultLanguage"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this post me installation template params body
func (m *PostMeInstallationTemplateParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseTemplateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDefaultLanguage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostMeInstallationTemplateParamsBody) validateBaseTemplateName(formats strfmt.Registry) error {

	if err := validate.Required("baseTemplateName", "body", m.BaseTemplateName); err != nil {
		return err
	}

	return nil
}

var postMeInstallationTemplateParamsBodyTypeDefaultLanguagePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ar","bg","cs","da","de","el","en","es","et","fi","fr","he","hr","hu","it","ja","ko","lt","lv","nb","nl","no","pl","pt","ro","ru","sk","sl","sr","sv","th","tr","tu","uk","zh-Hans-CN","zh-Hans-HK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postMeInstallationTemplateParamsBodyTypeDefaultLanguagePropEnum = append(postMeInstallationTemplateParamsBodyTypeDefaultLanguagePropEnum, v)
	}
}

const (
	// PostMeInstallationTemplateParamsBodyDefaultLanguageAr captures enum value "ar"
	PostMeInstallationTemplateParamsBodyDefaultLanguageAr string = "ar"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageBg captures enum value "bg"
	PostMeInstallationTemplateParamsBodyDefaultLanguageBg string = "bg"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageCs captures enum value "cs"
	PostMeInstallationTemplateParamsBodyDefaultLanguageCs string = "cs"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageDa captures enum value "da"
	PostMeInstallationTemplateParamsBodyDefaultLanguageDa string = "da"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageDe captures enum value "de"
	PostMeInstallationTemplateParamsBodyDefaultLanguageDe string = "de"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageEl captures enum value "el"
	PostMeInstallationTemplateParamsBodyDefaultLanguageEl string = "el"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageEn captures enum value "en"
	PostMeInstallationTemplateParamsBodyDefaultLanguageEn string = "en"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageEs captures enum value "es"
	PostMeInstallationTemplateParamsBodyDefaultLanguageEs string = "es"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageEt captures enum value "et"
	PostMeInstallationTemplateParamsBodyDefaultLanguageEt string = "et"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageFi captures enum value "fi"
	PostMeInstallationTemplateParamsBodyDefaultLanguageFi string = "fi"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageFr captures enum value "fr"
	PostMeInstallationTemplateParamsBodyDefaultLanguageFr string = "fr"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageHe captures enum value "he"
	PostMeInstallationTemplateParamsBodyDefaultLanguageHe string = "he"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageHr captures enum value "hr"
	PostMeInstallationTemplateParamsBodyDefaultLanguageHr string = "hr"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageHu captures enum value "hu"
	PostMeInstallationTemplateParamsBodyDefaultLanguageHu string = "hu"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageIt captures enum value "it"
	PostMeInstallationTemplateParamsBodyDefaultLanguageIt string = "it"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageJa captures enum value "ja"
	PostMeInstallationTemplateParamsBodyDefaultLanguageJa string = "ja"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageKo captures enum value "ko"
	PostMeInstallationTemplateParamsBodyDefaultLanguageKo string = "ko"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageLt captures enum value "lt"
	PostMeInstallationTemplateParamsBodyDefaultLanguageLt string = "lt"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageLv captures enum value "lv"
	PostMeInstallationTemplateParamsBodyDefaultLanguageLv string = "lv"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageNb captures enum value "nb"
	PostMeInstallationTemplateParamsBodyDefaultLanguageNb string = "nb"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageNl captures enum value "nl"
	PostMeInstallationTemplateParamsBodyDefaultLanguageNl string = "nl"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageNo captures enum value "no"
	PostMeInstallationTemplateParamsBodyDefaultLanguageNo string = "no"
	// PostMeInstallationTemplateParamsBodyDefaultLanguagePl captures enum value "pl"
	PostMeInstallationTemplateParamsBodyDefaultLanguagePl string = "pl"
	// PostMeInstallationTemplateParamsBodyDefaultLanguagePt captures enum value "pt"
	PostMeInstallationTemplateParamsBodyDefaultLanguagePt string = "pt"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageRo captures enum value "ro"
	PostMeInstallationTemplateParamsBodyDefaultLanguageRo string = "ro"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageRu captures enum value "ru"
	PostMeInstallationTemplateParamsBodyDefaultLanguageRu string = "ru"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageSk captures enum value "sk"
	PostMeInstallationTemplateParamsBodyDefaultLanguageSk string = "sk"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageSl captures enum value "sl"
	PostMeInstallationTemplateParamsBodyDefaultLanguageSl string = "sl"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageSr captures enum value "sr"
	PostMeInstallationTemplateParamsBodyDefaultLanguageSr string = "sr"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageSv captures enum value "sv"
	PostMeInstallationTemplateParamsBodyDefaultLanguageSv string = "sv"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageTh captures enum value "th"
	PostMeInstallationTemplateParamsBodyDefaultLanguageTh string = "th"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageTr captures enum value "tr"
	PostMeInstallationTemplateParamsBodyDefaultLanguageTr string = "tr"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageTu captures enum value "tu"
	PostMeInstallationTemplateParamsBodyDefaultLanguageTu string = "tu"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageUk captures enum value "uk"
	PostMeInstallationTemplateParamsBodyDefaultLanguageUk string = "uk"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageZhHansCN captures enum value "zh-Hans-CN"
	PostMeInstallationTemplateParamsBodyDefaultLanguageZhHansCN string = "zh-Hans-CN"
	// PostMeInstallationTemplateParamsBodyDefaultLanguageZhHansHK captures enum value "zh-Hans-HK"
	PostMeInstallationTemplateParamsBodyDefaultLanguageZhHansHK string = "zh-Hans-HK"
)

// prop value enum
func (m *PostMeInstallationTemplateParamsBody) validateDefaultLanguageEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postMeInstallationTemplateParamsBodyTypeDefaultLanguagePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PostMeInstallationTemplateParamsBody) validateDefaultLanguage(formats strfmt.Registry) error {

	if err := validate.Required("defaultLanguage", "body", m.DefaultLanguage); err != nil {
		return err
	}

	// value enum
	if err := m.validateDefaultLanguageEnum("defaultLanguage", "body", *m.DefaultLanguage); err != nil {
		return err
	}

	return nil
}

func (m *PostMeInstallationTemplateParamsBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostMeInstallationTemplateParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostMeInstallationTemplateParamsBody) UnmarshalBinary(b []byte) error {
	var res PostMeInstallationTemplateParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
