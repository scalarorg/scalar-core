package cosmos

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
)

// ParseEvent maps rawData into the target struct or pointer to struct.
func ParseEvent(rawData map[string][]string, target interface{}) error {
	targetVal := reflect.ValueOf(target)

	// Handle both struct and pointer cases
	if targetVal.Kind() != reflect.Struct && targetVal.Kind() != reflect.Ptr && targetVal.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("target must be a struct or pointer to struct, got: %T", target)
	}

	if err := parseEventInternal(rawData, target); err != nil {
		return err
	}

	return nil
}

func parseEventInternal(rawData map[string][]string, target interface{}) error {
	targetVal := reflect.ValueOf(target)
	targetVal = targetVal.Elem()
	targetType := targetVal.Type()
	data := make(map[string]string)

	for key, value := range rawData {
		keys := strings.Split(key, ".")
		fieldName := keys[len(keys)-1]
		if len(value) > 0 {
			data[fieldName] = value[0]
		}
	}

	for i := 0; i < targetType.NumField(); i++ {
		field := targetType.Field(i)
		fieldValue := targetVal.Field(i)

		tag := getJSONTag(field)
		if value, exists := data[tag]; exists {
			if err := setFieldValue(fieldValue, value, tag); err != nil {
				return err
			}
		}
	}

	return nil
}

// getJSONTag extracts the JSON tag name or defaults to the field name.
func getJSONTag(field reflect.StructField) string {
	tag := field.Tag.Get("json")
	if tag == "" || tag == "-" {
		return field.Name
	}
	return strings.Split(tag, ",")[0]
}

// setFieldValue sets the appropriate value for a struct field.
func setFieldValue(fieldValue reflect.Value, value string, tag string) error {
	switch fieldValue.Kind() {
	case reflect.String:
		if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") {
			value = strings.Trim(value, "\"")
		}
		if strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'") {
			value = strings.Trim(value, "'")
		}
		fieldValue.SetString(value)
	case reflect.Slice:
		if fieldValue.Type().Elem().Kind() == reflect.Uint8 {
			return setByteArrayValue(fieldValue, value, tag)
		}
		return fmt.Errorf("unsupported slice type for field '%s'", tag)
	case reflect.Array:
		if fieldValue.Type().Elem().Kind() == reflect.Uint8 {
			return setArrayValue(fieldValue, value, tag)
		}
		return fmt.Errorf("unsupported array type for field '%s'", tag)
	default:
		return fmt.Errorf("unsupported field type %s for field '%s'", fieldValue.Kind(), tag)
	}
	return nil
}

// setByteArrayValue handles slices and decodes from either JSON array notation or hex string.
func setByteArrayValue(fieldValue reflect.Value, value string, tag string) error {
	var byteArray []byte
	var err error

	if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {
		byteArray, err = parseByteArray(value)
	} else {
		byteArray, err = hex.DecodeString(value)
	}

	if err != nil {
		return fmt.Errorf("failed to parse byte array for field '%s': %w", tag, err)
	}

	fieldValue.SetBytes(byteArray)
	return nil
}

// setArrayValue handles fixed-size arrays and decodes from JSON array notation or hex string.
func setArrayValue(fieldValue reflect.Value, value string, tag string) error {
	var byteArray []byte
	var err error

	if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {
		byteArray, err = parseByteArray(value)
	} else {
		byteArray, err = hex.DecodeString(value)
	}

	if err != nil {
		return fmt.Errorf("failed to parse array for field '%s': %w", tag, err)
	}

	reflect.Copy(fieldValue, reflect.ValueOf(byteArray))
	return nil
}

// parseByteArray converts a JSON array-like string into a byte slice.
func parseByteArray(value string) ([]byte, error) {
	value = strings.Trim(value, "[]")
	parts := strings.Split(value, ",")
	byteArray := make([]byte, len(parts))

	for i, part := range parts {
		num, err := parseByte(strings.TrimSpace(part))
		if err != nil {
			return nil, err
		}
		byteArray[i] = num
	}
	return byteArray, nil
}

// parseByte parses a single byte value from a string.
func parseByte(value string) (byte, error) {
	var num int
	_, err := fmt.Sscanf(value, "%d", &num)
	if err != nil {
		return 0, err
	}
	if num < 0 || num > 255 {
		return 0, fmt.Errorf("value out of range: %d", num)
	}
	return byte(num), nil
}
