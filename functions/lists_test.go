package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestList provides unit test coverage for List()
func TestList(t *testing.T) {
	t.Parallel()
	type Args struct {
		s []interface{}
	}

	tests := []struct {
		name          string
		args          Args
		wantInterface interface{}
		wantError     bool
	}{
		{
			name: "nil",
			args: Args{
				s: nil,
			},
			wantInterface: []interface{}(nil),
			wantError:     false,
		},
		{
			name: "empty",
			args: Args{
				s: []interface{}{},
			},
			wantInterface: []interface{}{},
			wantError:     false,
		},
		{
			name: "one",
			args: Args{
				s: []interface{}{"one"},
			},
			wantInterface: []interface{}{"one"},
			wantError:     false,
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotInterface, gotError := List(tt.args.s...)
			if tt.wantError {
				require.Error(t, gotError)
			} else {
				require.NoError(t, gotError)
				assert.Equal(t, tt.wantInterface, gotInterface)
			}
		})
	}
}

// TestFirst provides unit test coverage for First()
func TestFirst(t *testing.T) {
	type Args struct {
		list interface{}
	}

	tests := []struct {
		name          string
		args          Args
		wantInterface interface{}
		wantError     bool
	}{
		{
			name: "nil",
			args: Args{
				list: nil,
			},
			wantError: true,
		},
		{
			name: "not a list",
			args: Args{
				list: "actually a string",
			},
			wantError: true,
		},
		{
			name: "from zero",
			args: Args{
				list: []int{},
			},
			wantInterface: nil,
			wantError:     false,
		},
		{
			name: "from two",
			args: Args{
				list: []string{"one", "two"},
			},
			wantInterface: "one",
			wantError:     false,
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotInterface, gotError := First(tt.args.list)
			if tt.wantError {
				require.Error(t, gotError)
			} else {
				require.NoError(t, gotError)
			}
			assert.Equal(t, tt.wantInterface, gotInterface)
		})
	}
}

// TestRest provides unit test coverage for Rest()
func TestRest(t *testing.T) {
	type Args struct {
		list interface{}
	}

	tests := []struct {
		name          string
		args          Args
		wantInterface interface{}
		wantError     bool
	}{
		{
			name: "nil",
			args: Args{
				list: nil,
			},
			wantError: true,
		},
		{
			name: "not a list",
			args: Args{
				list: "actually a string",
			},
			wantError: true,
		},
		{
			name: "from zero",
			args: Args{
				list: []int{},
			},
			wantInterface: nil,
			wantError:     false,
		},
		{
			name: "from two",
			args: Args{
				list: []string{"one", "two"},
			},
			wantInterface: []interface{}{"two"},
			wantError:     false,
		},
		{
			name: "with nils",
			args: Args{
				list: []interface{}{"one", "two", nil, "four"},
			},
			wantInterface: []interface{}{"two", nil, "four"},
			wantError:     false,
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotInterface, gotError := Rest(tt.args.list)
			if tt.wantError {
				require.Error(t, gotError)
			} else {
				require.NoError(t, gotError)
			}
			assert.Equal(t, tt.wantInterface, gotInterface)
		})
	}
}

// TestRest provides unit test coverage for Rest()
func TestPop(t *testing.T) {
	type Args struct {
		list interface{}
	}

	tests := []struct {
		name          string
		args          Args
		wantInterface interface{}
		wantError     bool
	}{
		{
			name: "nil",
			args: Args{
				list: nil,
			},
			wantError: true,
		},
		{
			name: "not a list",
			args: Args{
				list: "actually a string",
			},
			wantError: true,
		},
		{
			name: "from zero",
			args: Args{
				list: []int{},
			},
			wantInterface: nil,
			wantError:     false,
		},
		{
			name: "from two",
			args: Args{
				list: []string{"one", "two"},
			},
			wantInterface: []interface{}{"one"},
			wantError:     false,
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotInterface, gotError := Pop(tt.args.list)
			if tt.wantError {
				require.Error(t, gotError)
			} else {
				require.NoError(t, gotError)
			}
			assert.Equal(t, tt.wantInterface, gotInterface)
		})
	}
}

// TestLast provides unit test coverage for Last()
func TestLast(t *testing.T) {
	type Args struct {
		list interface{}
	}

	tests := []struct {
		name          string
		args          Args
		wantInterface interface{}
		wantError     bool
	}{
		{
			name: "nil",
			args: Args{
				list: nil,
			},
			wantError: true,
		},
		{
			name: "not a list",
			args: Args{
				list: "actually a string",
			},
			wantError: true,
		},
		{
			name: "from zero",
			args: Args{
				list: []int{},
			},
			wantInterface: nil,
			wantError:     false,
		},
		{
			name: "from two",
			args: Args{
				list: []string{"one", "two"},
			},
			wantInterface: "two",
			wantError:     false,
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotInterface, gotError := Last(tt.args.list)
			if tt.wantError {
				require.Error(t, gotError)
			} else {
				require.NoError(t, gotError)
			}
			assert.Equal(t, tt.wantInterface, gotInterface)
		})
	}
}

// TestContains provides unit test coverage for Contains()
func TestContains(t *testing.T) {
	t.Parallel()
	type Args struct {
		item interface{}
		list interface{}
	}

	tests := []struct {
		name      string
		args      Args
		wantBool  bool
		wantError bool
	}{
		{
			name: "nil",
			args: Args{
				list: nil,
			},
			wantError: true,
		},
		{
			name: "test against empty",
			args: Args{
				list: []int{},
				item: "2",
			},
			wantBool:  false,
			wantError: false,
		},
		{
			name: "not a list",
			args: Args{
				list: "actually a string",
			},
			wantError: true,
		},
		{
			name: "exists",
			args: Args{
				list: []string{"one", "two"},
				item: "two",
			},
			wantBool:  true,
			wantError: false,
		},
		{
			name: "doesn't exist",
			args: Args{
				list: []string{"one", "two"},
				item: "three",
			},
			wantBool:  false,
			wantError: false,
		},
		{
			name: "item of a different type",
			args: Args{
				list: []string{"one", "two"},
				item: 3,
			},
			wantBool:  false,
			wantError: false,
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotBool, gotError := Contains(tt.args.list, tt.args.item)
			if tt.wantError {
				require.Error(t, gotError)
			} else {
				require.NoError(t, gotError)
			}
			assert.Equal(t, tt.wantBool, gotBool)
		})
	}
}

// TestFilter provides unit test coverage for Filter()
func TestFilter(t *testing.T) {
	t.Parallel()
	type Args struct {
		item interface{}
		list interface{}
	}

	tests := []struct {
		name          string
		args          Args
		wantInterface interface{}
		wantError     bool
	}{
		{
			name: "filter from nil",
			args: Args{
				list: nil,
				item: nil,
			},
			wantError: true,
		},
		{
			name: "filter from empty",
			args: Args{
				list: []int{},
				item: 2,
			},
			wantInterface: []int{},
			wantError:     false,
		},
		{
			name: "not a list",
			args: Args{
				list: "actually a string",
			},
			wantError: true,
		},
		{
			name: "exists",
			args: Args{
				list: []string{"one", "two"},
				item: "two",
			},
			wantInterface: []interface{}{"one"},
			wantError:     false,
		},
		{
			name: "doesn't exist",
			args: Args{
				list: []string{"one", "two"},
				item: "three",
			},
			wantInterface: []interface{}{"one", "two"},
			wantError:     false,
		},
		{
			name: "item of a different type",
			args: Args{
				list: []string{"one", "two"},
				item: 3,
			},
			wantInterface: []interface{}{"one", "two"},
			wantError:     false,
		},
		{
			name: "remove multiple",
			args: Args{
				list: []string{"one", "two", "two", "three"},
				item: "two",
			},
			wantInterface: []interface{}{"one", "three"},
			wantError:     false,
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotInterface, gotError := Filter(tt.args.list, tt.args.item)
			if tt.wantError {
				require.Error(t, gotError)
			} else {
				require.NoError(t, gotError)
				assert.Equal(t, tt.wantInterface, gotInterface)
			}
		})
	}
}

// TestPush provides unit test coverage for Push()
func TestPush(t *testing.T) {
	t.Parallel()
	type Args struct {
		item interface{}
		list interface{}
	}

	tests := []struct {
		name          string
		args          Args
		wantInterface interface{}
		wantError     bool
	}{
		{
			name: "nil",
			args: Args{
				item: 7,
				list: nil,
			},
			wantError: true,
		},
		{
			name: "push to empty",
			args: Args{
				list: []int{},
				item: 2,
			},
			wantInterface: []interface{}{2},
		},
		{
			name: "not a list",
			args: Args{
				list: "actually a string",
			},
			wantError: true,
		},
		{
			name: "new",
			args: Args{
				list: []string{"one", "two"},
				item: "three",
			},
			wantInterface: []interface{}{"one", "two", "three"},
			wantError:     false,
		},
		{
			name: "item of a different type",
			args: Args{
				list: []string{"one", "two"},
				item: 3,
			},
			wantInterface: []interface{}{"one", "two", 3},
		},
		{
			name: "push nil to non-nillable",
			args: Args{
				list: []int{3, 4},
				item: nil,
			},
			wantInterface: []interface{}{3, 4, nil},
		},
		{
			name: "push nil to nillable",
			args: Args{
				list: []*int{helperPtrToInt(2)},
				item: nil,
			},
			wantInterface: []interface{}{helperPtrToInt(2), nil},
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotInterface, gotError := Push(tt.args.list, tt.args.item)
			if tt.wantError {
				require.Error(t, gotError, "with result %v", gotInterface)
			} else {
				require.NoError(t, gotError)
				assert.Equal(t, tt.wantInterface, gotInterface)
			}
		})
	}
}

// TestUnshift provides unit test coverage for Unshift()
func TestUnshift(t *testing.T) {
	t.Parallel()
	type Args struct {
		item interface{}
		list interface{}
	}

	tests := []struct {
		name          string
		args          Args
		wantInterface interface{}
		wantError     bool
	}{
		{
			name: "unshift to nil",
			args: Args{
				list: nil,
				item: nil,
			},
			wantError: true,
		},
		{
			name: "unshift to empty",
			args: Args{
				list: []int{},
				item: 2,
			},
			wantInterface: []interface{}{2},
		},
		{
			name: "not a list",
			args: Args{
				list: "actually a string",
			},
			wantError: true,
		},
		{
			name: "new",
			args: Args{
				list: []string{"one", "two"},
				item: "three",
			},
			wantInterface: []interface{}{"three", "one", "two"},
			wantError:     false,
		},
		{
			name: "item of a different type",
			args: Args{
				list: []string{"one", "two"},
				item: 3,
			},
			wantInterface: []interface{}{3, "one", "two"},
		},
		{
			name: "unshift nil to non-nillable",
			args: Args{
				list: []int{1, 2},
				item: nil,
			},
			wantInterface: []interface{}{nil, 1, 2},
		},
		{
			name: "unshift nil to nillable",
			args: Args{
				list: []*int{helperPtrToInt(2)},
				item: nil,
			},
			wantInterface: []interface{}{nil, helperPtrToInt(2)},
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotInterface, gotError := Unshift(tt.args.list, tt.args.item)
			if tt.wantError {
				require.Error(t, gotError, "with result %v", gotInterface)
			} else {
				require.NoError(t, gotError)
				assert.Equal(t, tt.wantInterface, gotInterface)
			}
		})
	}
}

// TestSlice provides unit test coverage for TestSlice()
func TestSlice(t *testing.T) {
	t.Parallel()
	type Args struct {
		i    int
		j    int
		list interface{}
	}

	tests := []struct {
		name          string
		args          Args
		wantInterface interface{}
		wantError     bool
	}{
		{
			name: "slice on nil",
			args: Args{
				i:    0,
				j:    0,
				list: nil,
			},
			wantError: true,
		},
		{
			name: "slice on empty",
			args: Args{
				i:    0,
				j:    0,
				list: []int{},
			},
			wantInterface: []interface{}{},
		},
		{
			name: "not a list",
			args: Args{
				list: "actually a string",
			},
			wantError: true,
		},
		{
			name: "middle slice",
			args: Args{
				i:    1,
				j:    3,
				list: []string{"one", "two", "three", "four"},
			},
			wantInterface: []interface{}{"two", "three"},
		},
		{
			name: "out of of bounds - leading",
			args: Args{
				i:    -1,
				j:    2,
				list: []string{"one", "two", "three", "four"},
			},
			wantError: true,
		},
		{
			name: "out of of bounds - trailing",
			args: Args{
				i:    3,
				j:    5,
				list: []string{"one", "two", "three", "four"},
			},
			wantError: true,
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotInterface, gotError := Slice(tt.args.i, tt.args.j, tt.args.list)
			if tt.wantError {
				require.Error(t, gotError, "with result %v", gotInterface)
			} else {
				require.NoError(t, gotError)
				assert.Equal(t, tt.wantInterface, gotInterface)
			}
		})
	}
}
