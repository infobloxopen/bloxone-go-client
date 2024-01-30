package internal

import (
    "reflect"
    "testing"
)

func TestConfiguration_AddDefaultTags(t *testing.T) {
    type fields struct {
        DefaultTags map[string]string
    }
    type args struct {
        m map[string]string
    }
    tests := []struct {
        name     string
        fields   fields
        args     args
        expected map[string]string
    }{
        {
            name: "empty fields",
            fields: fields{
                DefaultTags: map[string]string{},
            },
            args: args{
                m: map[string]string{},
            },
            expected: map[string]string{},
        }, {
            name: "existing fields",
            fields: fields{
                DefaultTags: map[string]string{
                    "key1": "value1",
                },
            },
            args: args{
                m: map[string]string{
                    "key2": "value2",
                },
            },
            expected: map[string]string{
                "key1": "value1",
                "key2": "value2",
            },
        },
        {
            name: "overriding fields",
            fields: fields{
                DefaultTags: map[string]string{
                    "key1": "value1",
                },
            },
            args: args{
                m: map[string]string{
                    "key1": "value2",
                },
            },
            expected: map[string]string{
                "key1": "value2",
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            c := &Configuration{
                DefaultTags: tt.fields.DefaultTags,
            }
            c.AddDefaultTags(tt.args.m)
            if !reflect.DeepEqual(c.GetDefaultTags(), tt.expected) {
                t.Errorf("internal() want = %v, got = %v", tt.expected, c.GetDefaultTags())
                return
            }
        })
    }
}
