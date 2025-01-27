package models

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// Effect - struct for Effect
type Effect struct {
	BlurEffect        *BlurEffect
	DropShadowEffect  *DropShadowEffect
	InnerShadowEffect *InnerShadowEffect
}



// Unmarshal JSON data into one of the pointers in the struct
func (dst *Effect) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BlurEffect
	err = newStrictDecoder(data).Decode(&dst.BlurEffect)
	if err == nil {
		jsonBlurEffect, _ := json.Marshal(dst.BlurEffect)
		if string(jsonBlurEffect) == "{}" { // empty struct
			dst.BlurEffect = nil
		} else {
			if err = validator.Validate(dst.BlurEffect); err != nil {
				dst.BlurEffect = nil
			} else {
				match++
			}
		}
	} else {
		dst.BlurEffect = nil
	}

	// try to unmarshal data into DropShadowEffect
	err = newStrictDecoder(data).Decode(&dst.DropShadowEffect)
	if err == nil {
		jsonDropShadowEffect, _ := json.Marshal(dst.DropShadowEffect)
		if string(jsonDropShadowEffect) == "{}" { // empty struct
			dst.DropShadowEffect = nil
		} else {
			if err = validator.Validate(dst.DropShadowEffect); err != nil {
				dst.DropShadowEffect = nil
			} else {
				match++
			}
		}
	} else {
		dst.DropShadowEffect = nil
	}

	// try to unmarshal data into InnerShadowEffect
	err = newStrictDecoder(data).Decode(&dst.InnerShadowEffect)
	if err == nil {
		jsonInnerShadowEffect, _ := json.Marshal(dst.InnerShadowEffect)
		if string(jsonInnerShadowEffect) == "{}" { // empty struct
			dst.InnerShadowEffect = nil
		} else {
			if err = validator.Validate(dst.InnerShadowEffect); err != nil {
				dst.InnerShadowEffect = nil
			} else {
				match++
			}
		}
	} else {
		dst.InnerShadowEffect = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BlurEffect = nil
		dst.DropShadowEffect = nil
		dst.InnerShadowEffect = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Effect)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Effect)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Effect) MarshalJSON() ([]byte, error) {
	if src.BlurEffect != nil {
		return json.Marshal(&src.BlurEffect)
	}

	if src.DropShadowEffect != nil {
		return json.Marshal(&src.DropShadowEffect)
	}

	if src.InnerShadowEffect != nil {
		return json.Marshal(&src.InnerShadowEffect)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Effect) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BlurEffect != nil {
		return obj.BlurEffect
	}

	if obj.DropShadowEffect != nil {
		return obj.DropShadowEffect
	}

	if obj.InnerShadowEffect != nil {
		return obj.InnerShadowEffect
	}

	// all schemas are nil
	return nil
}
