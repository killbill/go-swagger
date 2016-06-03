package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	"github.com/willf/bitset"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

var nullableItemBitmap bitset.BitSet

func init() {
	nullableItemBitmap = *bitset.New(3)
	nullableItemBitmap.Set(1)
}

// NewItem creates a new item with the bitsets initialized
func NewItem() *Item {
	return &Item{
		setValues: *bitset.New(3),
		nulls:     *bitset.New(3),
	}
}

/*Item item

swagger:model item
*/
type Item struct {
	setValues bitset.BitSet
	nulls     bitset.BitSet

	/* completed
	 */
	Completed bool `json:"completed,omitempty"`

	/* description

	Required: true
	Min Length: 1
	Nullable: true
	*/
	Description string `json:"description"`

	/* id

	Read Only: true
	*/
	ID int64 `json:"id,omitempty"`
}

// FlagCompletedSet flags the completed field as set
func (m *Item) FlagCompletedSet() {
	m.setValues.Set(0)
}

// FlagCompletedUnset flag completed field as not set
func (m *Item) FlagCompletedUnset() {
	m.setValues.Clear(0)
}

// IsCompletedSet returns true if the value for completed was set
func (m *Item) IsCompletedSet() bool {
	return m.setValues.Test(0) || m.Completed
}

// SetCompleted set the value of the completed field
func (m *Item) SetCompleted(value bool) {
	m.FlagCompletedSet()
	m.Completed = value
}

// ClearCompleted clears the value for the completed field
func (m *Item) ClearCompleted() {
	m.FlagCompletedUnset()
	m.nulls.Clear(0)
	m.Completed = false
}

// GetCompleted returns the value for the completed field but also if it was set in the wire format.
func (m *Item) GetCompleted() (value bool, haskey bool) {
	return m.Completed, m.IsCompletedSet()
}

// GetCompletedPtr returns a pointer for the value of the completed field.
// if completed was not set it will return nil
func (m *Item) GetCompletedPtr() *bool {
	if !m.IsCompletedSet() {
		return nil
	}
	return &m.Completed
}

// FlagDescriptionSet flag the description field as set
func (m *Item) FlagDescriptionSet() {
	m.setValues.Set(1)
}

// FlagDescriptionUnset flags the description field as unset
func (m *Item) FlagDescriptionUnset() {
	m.setValues.Clear(1)
}

// FlagDescriptionNil flags the description field as nil
func (m *Item) FlagDescriptionNil() {
	if nullableItemBitmap.Test(1) {
		m.nulls.Set(1)
	}
}

// FlagDescriptionZero flags the description field as zero
func (m *Item) FlagDescriptionZero() {
	m.nulls.Clear(1)
}

// IsDescriptionNil returns true if the description field is nil
func (m *Item) IsDescriptionNil() bool {
	return nullableItemBitmap.Test(1) && m.nulls.Test(1)
}

// IsDescriptionSet returns true if the description field is set
func (m *Item) IsDescriptionSet() bool {
	return m.setValues.Test(1) || m.Completed
}

// HasDescriptionValue returns true if the description field has a value
func (m *Item) HasDescriptionValue() bool {
	return m.IsDescriptionSet() && !m.IsDescriptionNil()
}

// SetDescription sets the description of the field to the specified value
func (m *Item) SetDescription(value *string) {
	m.FlagDescriptionSet()
	if value == nil {
		m.FlagDescriptionNil()
		m.Description = ""
		return
	}
	m.Description = *value
}

// ClearDescription clears the value of the description field
func (m *Item) ClearDescription() {
	m.FlagDescriptionUnset()
	m.Description = ""
	m.FlagDescriptionNil()
}

// GetDescription value with indication if it should be nil or unset
func (m *Item) GetDescription() (value string, null bool, haskey bool) {
	return m.Description, nullableItemBitmap.Test(1) && m.nulls.Test(1), m.setValues.Test(1) || len(m.Description) > 0
}

// GetDescriptionPtr gets the description field returns nil when unset
func (m *Item) GetDescriptionPtr() *string {
	if !m.HasDescriptionValue() {
		return nil
	}
	return &m.Description
}

// FlagIDSet flags id field as set
func (m *Item) FlagIDSet() {
	m.setValues.Set(2)
}

// IsIDSet returns true if the id is set
func (m *Item) IsIDSet() bool {
	return m.setValues.Test(2) || m.ID > 0
}

// SetID sets the id field
func (m *Item) SetID(value int64) {
	m.FlagDescriptionSet()
	m.ID = value
}

// Validate validates this item
func (m *Item) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Item) validateDescription(formats strfmt.Registry) error {

	if !m.HasDescriptionValue() {
		return errors.Required("description", "body")
	}

	if err := validate.MinLength("description", "body", m.Description, 1); err != nil {
		return err
	}

	return nil
}

// MarshalEasyJSON for this item
func (m Item) MarshalEasyJSON(out *jwriter.Writer) {
	out.RawByte('{')
	first := true
	_ = first
	if m.Completed || m.setValues.Test(0) {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"completed\":")
		if nullableItemBitmap.Test(0) && m.nulls.Test(0) {
			out.RawString("null")
		} else {
			out.Bool(m.Completed)
		}
	}
	if m.Description != "" || m.setValues.Test(1) {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"description\":")
		if nullableItemBitmap.Test(1) && m.nulls.Test(1) {
			out.RawString("null")
		} else {
			out.String(m.Description)
		}
	}
	if m.ID != 0 || m.setValues.Test(2) {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"id\":")
		if nullableItemBitmap.Test(1) && m.nulls.Test(1) {
			out.RawString("null")
		} else {
			out.Int64(m.ID)
		}
	}
	out.RawByte('}')
}

// MarshalJSON for this item
func (m Item) MarshalJSON() ([]byte, error) {
	out := jwriter.Writer{}
	m.MarshalEasyJSON(&out)
	return out.BuildBytes()
}

// UnmarshalEasyJSON set the correct bitmap fields etc when deserializing from JSON
func (m *Item) UnmarshalEasyJSON(in *jlexer.Lexer) {
	m.nulls = *bitset.New(3)
	m.setValues = *bitset.New(3)
	if in.IsNull() {
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			switch key {
			case "id":
				m.ID = 0
				m.setValues.Set(2)
				m.nulls.Set(2)
			case "description":
				m.Description = ""
				m.setValues.Set(1)
				m.nulls.Set(1)
			case "completed":
				m.Completed = false
				m.setValues.Set(0)
				m.nulls.Set(0)
			default:
				in.SkipRecursive()
			}
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			m.ID = in.Int64()
			m.setValues.Set(2)
			m.nulls.Clear(2)
		case "description":
			m.Description = in.String()
			m.setValues.Set(1)
			m.nulls.Clear(1)
		case "completed":
			m.Completed = in.Bool()
			m.setValues.Set(0)
			m.nulls.Clear(0)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}

// UnmarshalJSON read object from json
func (m *Item) UnmarshalJSON(data []byte) error {
	in := jlexer.Lexer{Data: data}
	m.UnmarshalEasyJSON(&in)
	return in.Error()
}

// PatchWith other object, only use the properties that were actually set
func (m *Item) PatchWith(other *Item) error {
	if other.IsCompletedSet() {
		m.SetCompleted(other.Completed)
	}

	if other.IsDescriptionSet() {
		if other.IsDescriptionNil() {
			other.SetDescription(nil)
		} else {
			other.SetDescription(&other.Description)
		}
	}

	if other.IsIDSet() {
		if other.ID == 0 {
			m.setValues.Clear(2)
			m.ID = 0
		} else {
			m.SetID(m.ID)
		}
	}
	return nil
}

// Clone creates a deep clone of the item
func (m *Item) Clone() *Item {
	n := NewItem()
	n.PatchWith(m)
	return n
}
