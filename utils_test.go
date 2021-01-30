package gbfs

import (
	"reflect"
	"testing"
)

func TestID_UnmarshalJSON(t *testing.T) {
	type fields struct {
		id  string
		raw []byte
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// success cases
		{"numeric id", fields{"9000", []byte("9000")}, args{[]byte("9000")}, false},
		{"string id", fields{`"9000"`, []byte(`"9000"`)}, args{[]byte(`"9000"`)}, false},
		// error cases
		{"null", fields{"null", []byte("null")}, args{[]byte("null")}, true},
		{"empty string", fields{"", []byte("")}, args{[]byte("")}, true},
		{"object", fields{"{}", []byte("{}")}, args{[]byte("{}")}, true},
		{"array", fields{"[]", []byte("[]")}, args{[]byte("[]")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := &ID{
				id:  tt.fields.id,
				raw: tt.fields.raw,
			}
			if err := id.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("ID.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestID_MarshalJSON(t *testing.T) {
	type fields struct {
		id  string
		raw []byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// success cases
		{"numeric id", fields{"9000", []byte("9000")}, []byte("9000"), false},
		{"string id", fields{`"9000"`, []byte(`"9000"`)}, []byte(`"9000"`), false},
		// error cases
		{"empty id", fields{"", nil}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := ID{
				id:  tt.fields.id,
				raw: tt.fields.raw,
			}
			got, err := id.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("ID.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ID.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
