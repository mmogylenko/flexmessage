package flexmessage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

var colors = map[string]string{
	"default": "\033[1;36m%s\033[0m",
	"black":   "\033[1;30m%s\033[0m",
	"red":     "\033[1;31m%s\033[0m",
	"green":   "\033[1;32m%s\033[0m",
	"yellow":  "\033[1;33m%s\033[0m",
	"purple":  "\033[1;34m%s\033[0m",
	"magenta": "\033[1;35m%s\033[0m",
	"teal":    "\033[1;36m%s\033[0m",
	"white":   "\033[1;37m%s\033[0m",
}

// ColoringSchema defines our coloring schema :-)
// Your Cap.
type ColoringSchema struct {
	KeyColor        string `default:"purple"`
	StringColor     string `default:"yellow"`
	BoolColor       string `default:"purple"`
	NumberColor     string `default:"green"`
	NullColor       string `default:"magenta"`
	StringMaxLength int    `default:"0"`
	Indent          int    `default:"0"`
	DisableColors   bool   `default:"false"`
	RawStrings      bool   `default:"false"`
}

// ColoringSchemaOptions options for coloring and formatting
// TODO Implement this as an options in a separate file
type ColoringSchemaOptions struct {
	InitialDepth   int
	ValueSeparator string
	NullValue      string
	MapStart       string
	MapEnd         string
	ArrayStart     string
	ArrayEnd       string
	EmptyMap       string
	EmptyArray     string
}

const (
	initialDepth = 0
	valueSep     = ","
	null         = "null"
	startMap     = "{"
	endMap       = "}"
	startArray   = "["
	endArray     = "]"
	emptyMap     = startMap + endMap
	emptyArray   = startArray + endArray
)

// New is New. New is better than old. Or not
func (cs *ColoringSchema) New() *ColoringSchema {
	v := reflect.ValueOf(cs).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		if defaultVal := t.Field(i).Tag.Get("default"); defaultVal != "-" {
			field := v.Field(i)
			if field.IsZero() {
				switch field.Kind() {
				case reflect.Int:
					if val, err := strconv.ParseInt(defaultVal, 0, strconv.IntSize); err == nil {
						field.Set(reflect.ValueOf(int(val)).Convert(field.Type()))
					}
				case reflect.String:
					field.Set(reflect.ValueOf(defaultVal).Convert(field.Type()))

				case reflect.Bool:
					if val, err := strconv.ParseBool(defaultVal); err == nil {
						field.Set(reflect.ValueOf(val).Convert(field.Type()))
					}
				}
			}
		}

	}

	return cs
}

// indent
func (cs *ColoringSchema) indent(b *bytes.Buffer, depth int) {
	b.WriteString(strings.Repeat(" ", cs.Indent*depth))
}

// separator
func (cs *ColoringSchema) separator(b *bytes.Buffer) {
	if cs.Indent != 0 {
		b.WriteByte('\n')
	} else {
		b.WriteByte(' ')
	}
}

func (cs *ColoringSchema) colorize(colorString string, format string, args ...interface{}) string {
	color, found := colors[strings.ToLower(colorString)]
	if !found {
		color = colors[strings.ToLower("default")]
	}

	if cs.DisableColors == true {
		return fmt.Sprintf(format, args...)
	}
	return fmt.Sprintf(color,
		fmt.Sprintf(format, args...))
}

func (cs *ColoringSchema) colorizeValue(val interface{}, b *bytes.Buffer, depth int) {
	switch v := val.(type) {
	case map[string]interface{}:
		cs.colorizeMap(v, b, depth)
	case []interface{}:
		cs.colorizeArray(v, b, depth)
	case string:
		cs.colorizeString(v, b)
	case []string:
		b.WriteString("[")
		for i, s := range v {
			cs.colorizeString(s, b)
			if i != len(v)-1 {
				b.WriteString(", ")
			}
		}
		b.WriteString("]")
	case int:
		b.WriteString(cs.colorize(cs.NumberColor, strconv.FormatInt(int64(v), 10)))
	case int8:
		b.WriteString(cs.colorize(cs.NumberColor, strconv.FormatInt(int64(v), 10)))
	case int16:
		b.WriteString(cs.colorize(cs.NumberColor, strconv.FormatInt(int64(v), 10)))
	case int32:
		b.WriteString(cs.colorize(cs.NumberColor, strconv.FormatInt(int64(v), 10)))
	case int64:
		b.WriteString(cs.colorize(cs.NumberColor, strconv.FormatInt(v, 10)))
	case uint:
		b.WriteString(cs.colorize(cs.NumberColor, strconv.FormatInt(int64(v), 10)))
	case uint8:
		b.WriteString(cs.colorize(cs.NumberColor, strconv.FormatInt(int64(v), 10)))
	case uint16:
		b.WriteString(cs.colorize(cs.NumberColor, strconv.FormatInt(int64(v), 10)))
	case uint32:
		b.WriteString(cs.colorize(cs.NumberColor, strconv.FormatInt(int64(v), 10)))
	case uint64:
		b.WriteString(cs.colorize(cs.NumberColor, strconv.FormatInt(int64(v), 10)))
	case float64:
		b.WriteString(cs.colorize(cs.NumberColor, strconv.FormatFloat(v, 'f', -1, 64)))
	case bool:
		b.WriteString(cs.colorize(cs.BoolColor, (strconv.FormatBool(v))))
	case nil:
		b.WriteString(cs.colorize(cs.NullColor, null))
	case json.Number:
		b.WriteString(cs.colorize(cs.NumberColor, v.String()))
	default:
		fmt.Println(
			fmt.Errorf("Unsupported type %T for %v", val, val),
		)
	}

}

func (cs *ColoringSchema) colorizeMap(m map[string]interface{}, b *bytes.Buffer, depth int) {
	remaining := len(m)
	if remaining == 0 {
		b.WriteString(emptyMap)
		return
	}

	keys := make([]string, 0)
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	b.WriteString(startMap)
	cs.separator(b)

	for _, key := range keys {
		cs.indent(b, depth+1)
		b.WriteString(cs.colorize(cs.KeyColor, "\"%s\": ", key))

		cs.colorizeValue(m[key], b, depth+1)
		remaining--
		if remaining != 0 {
			b.WriteString(valueSep)
		}
		cs.separator(b)
	}
	cs.indent(b, depth)
	b.WriteString(endMap)
}

func (cs *ColoringSchema) colorizeString(str string, b *bytes.Buffer) {
	if !cs.RawStrings {
		strBytes, _ := json.Marshal(str)
		str = string(strBytes)
	}

	if cs.StringMaxLength != 0 && len(str) >= cs.StringMaxLength {
		str = fmt.Sprintf("%s...", str[0:cs.StringMaxLength])
	}

	b.WriteString(cs.colorize(cs.StringColor, str))
}

func (cs *ColoringSchema) colorizeArray(a []interface{}, b *bytes.Buffer, depth int) {
	if len(a) == 0 {
		b.WriteString(emptyArray)
		return
	}

	b.WriteString(startArray)
	cs.separator(b)

	for i, v := range a {
		cs.indent(b, depth+1)
		cs.colorizeValue(v, b, depth+1)
		if i < len(a)-1 {
			b.WriteString(valueSep)
		}
		cs.separator(b)
	}
	cs.indent(b, depth)
	b.WriteString(endArray)
}
