package response_test

import (
	"encoding/json"
	"github.com/go-lean/response"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func Test_JSONStream_BeginObject(t *testing.T) {
	strBuilder := &strings.Builder{}
	str := response.NewJSONStream(strBuilder)

	require.Nil(t, str.BeginObject())

	require.Equal(t, "{", strBuilder.String())
}

func Test_JSONStream_EndObject(t *testing.T) {
	strBuilder := &strings.Builder{}
	str := response.NewJSONStream(strBuilder)

	require.Nil(t, str.EndObject())

	require.Equal(t, "}", strBuilder.String())
}

func Test_JSONStream_BeginArray(t *testing.T) {
	strBuilder := &strings.Builder{}
	str := response.NewJSONStream(strBuilder)

	require.Nil(t, str.BeginArray())

	require.Equal(t, "[", strBuilder.String())
}

func Test_JSONStream_EndArray(t *testing.T) {
	strBuilder := &strings.Builder{}
	str := response.NewJSONStream(strBuilder)

	require.Nil(t, str.EndArray())

	require.Equal(t, "]", strBuilder.String())
}

func Test_JSONStream_WriteKey(t *testing.T) {
	strBuilder := &strings.Builder{}
	str := response.NewJSONStream(strBuilder)

	require.Nil(t, str.WriteKey("baba"))

	require.Equal(t, "\"baba\":", strBuilder.String())
}

type serializable struct {
	t          *testing.T
	TextData   string
	IntData    int
	Int8Data   int8
	Int16Data  int16
	Int32Data  int32
	Int64Data  int64
	TextArray  []string
	PairsArray []struct {
		Key   string
		Value int
	}
}

func (s *serializable) Encode(stream *response.JSONStream) error {
	require.Nil(s.t, stream.BeginObject())
	require.Nil(s.t, stream.WriteKey("TextData"))
	require.Nil(s.t, stream.WriteString(s.TextData))
	require.Nil(s.t, stream.WriteKey("IntData"))
	require.Nil(s.t, stream.WriteInt(s.IntData))
	require.Nil(s.t, stream.WriteKey("Int8Data"))
	require.Nil(s.t, stream.WriteInt8(s.Int8Data))
	require.Nil(s.t, stream.WriteKey("Int16Data"))
	require.Nil(s.t, stream.WriteInt16(s.Int16Data))
	require.Nil(s.t, stream.WriteKey("Int32Data"))
	require.Nil(s.t, stream.WriteInt32(s.Int32Data))
	require.Nil(s.t, stream.WriteKey("Int64Data"))
	require.Nil(s.t, stream.WriteInt64(s.Int64Data))

	require.Nil(s.t, stream.WriteKey("TextArray"))
	require.Nil(s.t, stream.BeginArray())
	for _, param := range s.TextArray {
		require.Nil(s.t, stream.WriteString(param))
	}
	require.Nil(s.t, stream.EndArray())

	require.Nil(s.t, stream.WriteKey("PairsArray"))
	require.Nil(s.t, stream.BeginArray())
	for _, param := range s.PairsArray {
		require.Nil(s.t, stream.BeginObject())
		require.Nil(s.t, stream.WriteKey("Key"))
		require.Nil(s.t, stream.WriteString(param.Key))
		require.Nil(s.t, stream.WriteKey("Value"))
		require.Nil(s.t, stream.WriteInt(param.Value))
		require.Nil(s.t, stream.EndObject())
	}
	require.Nil(s.t, stream.EndArray())

	require.Nil(s.t, stream.EndObject())

	return nil
}

func Test_JSONStream_WriteObject(t *testing.T) {
	strBuilder := &strings.Builder{}
	str := response.NewJSONStream(strBuilder)
	obj := &serializable{
		t:         t,
		TextData:  "baba",
		IntData:   32,
		Int8Data:  8,
		Int16Data: 16,
		Int32Data: 32,
		Int64Data: 64,
		TextArray: []string{"baba", "is", "you"},
		PairsArray: []struct {
			Key   string
			Value int
		}{
			{"first", 1},
			{"second", 2},
		},
	}

	require.Nil(t, str.WriteObject(obj))
	var decObj serializable
	require.Nil(t, json.NewDecoder(strings.NewReader(strBuilder.String())).Decode(&decObj), strBuilder.String())

	require.Equal(t, "baba", decObj.TextData)
	require.Equal(t, 32, decObj.IntData)
	require.Equal(t, int8(8), decObj.Int8Data)
	require.Equal(t, int16(16), decObj.Int16Data)
	require.Equal(t, int32(32), decObj.Int32Data)
	require.Equal(t, int64(64), decObj.Int64Data)
	require.Len(t, decObj.TextArray, 3)
	require.Equal(t, "baba", decObj.TextArray[0])
	require.Equal(t, "is", decObj.TextArray[1])
	require.Equal(t, "you", decObj.TextArray[2])
	require.Len(t, decObj.PairsArray, 2)
	require.Equal(t, "first", decObj.PairsArray[0].Key)
	require.Equal(t, 1, decObj.PairsArray[0].Value)
	require.Equal(t, "second", decObj.PairsArray[1].Key)
	require.Equal(t, 2, decObj.PairsArray[1].Value)
}

func Test_WriteUint(t *testing.T) {
	builder := strings.Builder{}
	str := response.NewJSONStream(&builder)
	require.Nil(t, str.WriteUint(32))
	require.Nil(t, str.WriteUint(32))
	require.Equal(t, "32,32", builder.String())

	builder.Reset()
	str = response.NewJSONStream(&builder)
	require.Nil(t, str.WriteUint8(8))
	require.Nil(t, str.WriteUint8(8))
	require.Equal(t, "8,8", builder.String())

	builder.Reset()
	str = response.NewJSONStream(&builder)
	require.Nil(t, str.WriteUint16(16))
	require.Nil(t, str.WriteUint16(16))
	require.Equal(t, "16,16", builder.String())

	builder.Reset()
	str = response.NewJSONStream(&builder)
	require.Nil(t, str.WriteUint32(32))
	require.Nil(t, str.WriteUint32(32))
	require.Equal(t, "32,32", builder.String())

	builder.Reset()
	str = response.NewJSONStream(&builder)
	require.Nil(t, str.WriteUint64(64))
	require.Nil(t, str.WriteUint64(64))
	require.Equal(t, "64,64", builder.String())
}

func Test_WriteFloat(t *testing.T) {
	builder := strings.Builder{}
	str := response.NewJSONStream(&builder)
	require.Nil(t, str.WriteFloat32(32))
	require.Nil(t, str.WriteFloat32(32.0))
	require.Nil(t, str.WriteFloat32(32.5))
	require.Nil(t, str.WriteFloat32(32.55))
	require.Equal(t, "32,32,32.5,32.55", builder.String())

	builder.Reset()
	str = response.NewJSONStream(&builder)
	require.Nil(t, str.WriteFloat64(64))
	require.Nil(t, str.WriteFloat64(64.0))
	require.Nil(t, str.WriteFloat64(64.5))
	require.Nil(t, str.WriteFloat64(64.55))
	require.Equal(t, "64,64,64.5,64.55", builder.String())
}
