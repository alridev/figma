package models

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// Paint - struct for Paint
type Paint struct {
	GradientPaint *GradientPaint
	ImagePaint    *ImagePaint
	SolidPaint    *SolidPaint
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Paint) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GradientPaint
	err = newStrictDecoder(data).Decode(&dst.GradientPaint)
	if err == nil {
		jsonGradientPaint, _ := json.Marshal(dst.GradientPaint)
		if string(jsonGradientPaint) == "{}" { // empty struct
			dst.GradientPaint = nil
		} else {
			if err = validator.Validate(dst.GradientPaint); err != nil {
				dst.GradientPaint = nil
			} else {
				match++
			}
		}
	} else {
		dst.GradientPaint = nil
	}

	// try to unmarshal data into ImagePaint
	err = newStrictDecoder(data).Decode(&dst.ImagePaint)
	if err == nil {
		jsonImagePaint, _ := json.Marshal(dst.ImagePaint)
		if string(jsonImagePaint) == "{}" { // empty struct
			dst.ImagePaint = nil
		} else {
			if err = validator.Validate(dst.ImagePaint); err != nil {
				dst.ImagePaint = nil
			} else {
				match++
			}
		}
	} else {
		dst.ImagePaint = nil
	}

	// try to unmarshal data into SolidPaint
	err = newStrictDecoder(data).Decode(&dst.SolidPaint)
	if err == nil {
		jsonSolidPaint, _ := json.Marshal(dst.SolidPaint)
		if string(jsonSolidPaint) == "{}" { // empty struct
			dst.SolidPaint = nil
		} else {
			if err = validator.Validate(dst.SolidPaint); err != nil {
				dst.SolidPaint = nil
			} else {
				match++
			}
		}
	} else {
		dst.SolidPaint = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GradientPaint = nil
		dst.ImagePaint = nil
		dst.SolidPaint = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Paint)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Paint)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Paint) MarshalJSON() ([]byte, error) {
	if src.GradientPaint != nil {
		return json.Marshal(&src.GradientPaint)
	}

	if src.ImagePaint != nil {
		return json.Marshal(&src.ImagePaint)
	}

	if src.SolidPaint != nil {
		return json.Marshal(&src.SolidPaint)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Paint) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.GradientPaint != nil {
		return obj.GradientPaint
	}

	if obj.ImagePaint != nil {
		return obj.ImagePaint
	}

	if obj.SolidPaint != nil {
		return obj.SolidPaint
	}

	// all schemas are nil
	return nil
}
