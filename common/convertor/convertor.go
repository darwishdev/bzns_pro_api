package convertor

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	rmsv1 "github.com/meloneg/mln_rms_core/common/pb/bznspro/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToPgType(str string) pgtype.Text {
	return pgtype.Text{String: str, Valid: str != ""}
}
func ToPgTypeBool(value bool) pgtype.Bool {
	return pgtype.Bool{Bool: value, Valid: true}
}
func ToPgTypeInt(value int32) pgtype.Int4 {
	return pgtype.Int4{Int32: value, Valid: true}
}
func ToPgTypeID(value int32) pgtype.Int4 {
	return pgtype.Int4{Int32: value, Valid: value > 0}
}
func ToPgTypeUInt(value int32) pgtype.Int4 {
	return pgtype.Int4{Int32: value, Valid: value > -1}
}
func ToTimeStamp(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}
func ToSelectInput(value int32, label string) *rmsv1.SelectInputOption {
	return &rmsv1.SelectInputOption{
		Value: value,
		Label: label,
	}
}

func SnakeToPascal(input string) string {
	words := strings.Split(input, "_")
	for i := range words {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}

func SetField(obj interface{}, name string, value interface{}) error {
	// Get the struct value from the interface
	structValue := reflect.ValueOf(obj).Elem()

	// Get the field value by name
	fieldValue := structValue.FieldByName(name)

	// Check if the field exists
	if !fieldValue.IsValid() {
		return fmt.Errorf("no such field: %s in obj", name)
	}

	// Get the type of the field
	fieldType := fieldValue.Type()

	// Get the value to be set as a reflect.Value
	newValue := reflect.ValueOf(value)

	// Check if the type of the value to be set matches the type of the field
	if !newValue.Type().AssignableTo(fieldType) {
		return fmt.Errorf("value type %v is not assignable to field type %v", newValue.Type(), fieldType)
	}

	// Set the field value to the new value
	fieldValue.Set(newValue)

	return nil
}
