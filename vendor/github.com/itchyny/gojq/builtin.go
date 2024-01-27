// Code generated by _tools/gen_builtin.go; DO NOT EDIT.

package gojq

func init() {
	builtinFuncDefs = map[string][]*FuncDef{
		"IN": []*FuncDef{&FuncDef{Name: "IN", Args: []string{"s"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "any", Args: []*Query{&Query{Left: &Query{Func: "s"}, Op: OpEq, Right: &Query{Func: "."}}, &Query{Func: "."}}}}}}, &FuncDef{Name: "IN", Args: []string{"src", "s"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "any", Args: []*Query{&Query{Left: &Query{Func: "src"}, Op: OpEq, Right: &Query{Func: "s"}}, &Query{Func: "."}}}}}}},
		"INDEX": []*FuncDef{&FuncDef{Name: "INDEX", Args: []string{"stream", "idx_expr"}, Body: &Query{Term: &Term{Type: TermTypeReduce, Reduce: &Reduce{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "stream"}}, Pattern: &Pattern{Name: "$row"}, Start: &Query{Term: &Term{Type: TermTypeObject, Object: &Object{}}}, Update: &Query{Left: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Start: &Query{Left: &Query{Func: "$row"}, Op: OpPipe, Right: &Query{Left: &Query{Func: "idx_expr"}, Op: OpPipe, Right: &Query{Func: "tostring"}}}}}}, Op: OpAssign, Right: &Query{Func: "$row"}}}}}}, &FuncDef{Name: "INDEX", Args: []string{"idx_expr"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "INDEX", Args: []*Query{&Query{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{&Suffix{Iter: true}}}}, &Query{Func: "idx_expr"}}}}}}},
		"JOIN": []*FuncDef{&FuncDef{Name: "JOIN", Args: []string{"$idx", "idx_expr"}, Body: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{&Suffix{Iter: true}}}}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Func: "."}, Op: OpComma, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "$idx"}, SuffixList: []*Suffix{&Suffix{Index: &Index{Start: &Query{Func: "idx_expr"}}}}}}}}}}}}}}}, &FuncDef{Name: "JOIN", Args: []string{"$idx", "stream", "idx_expr"}, Body: &Query{Left: &Query{Func: "stream"}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Func: "."}, Op: OpComma, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "$idx"}, SuffixList: []*Suffix{&Suffix{Index: &Index{Start: &Query{Func: "idx_expr"}}}}}}}}}}}}, &FuncDef{Name: "JOIN", Args: []string{"$idx", "stream", "idx_expr", "join_expr"}, Body: &Query{Left: &Query{Func: "stream"}, Op: OpPipe, Right: &Query{Left: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Func: "."}, Op: OpComma, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "$idx"}, SuffixList: []*Suffix{&Suffix{Index: &Index{Start: &Query{Func: "idx_expr"}}}}}}}}}}, Op: OpPipe, Right: &Query{Func: "join_expr"}}}}},
		"_assign": []*FuncDef{&FuncDef{Name: "_assign", Args: []string{"ps", "$v"}, Body: &Query{Term: &Term{Type: TermTypeReduce, Reduce: &Reduce{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "path", Args: []*Query{&Query{Func: "ps"}}}}, Pattern: &Pattern{Name: "$p"}, Start: &Query{Func: "."}, Update: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "setpath", Args: []*Query{&Query{Func: "$p"}, &Query{Func: "$v"}}}}}}}}}},
		"_modify": []*FuncDef{&FuncDef{Name: "_modify", Args: []string{"ps", "f"}, Body: &Query{Left: &Query{Term: &Term{Type: TermTypeReduce, Reduce: &Reduce{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "path", Args: []*Query{&Query{Func: "ps"}}}}, Pattern: &Pattern{Name: "$p"}, Start: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Func: "."}, Op: OpComma, Right: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{}}}}}}}, Update: &Query{Term: &Term{Type: TermTypeLabel, Label: &Label{Ident: "$out", Body: &Query{Left: &Query{Term: &Term{Type: TermTypeQuery, Query: &Query{Term: &Term{Type: TermTypeQuery, Query: &Query{Left: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}}}, Op: OpAdd, Right: &Query{Func: "$p"}}, SuffixList: []*Suffix{&Suffix{Bind: &Bind{Patterns: []*Pattern{&Pattern{Name: "$q"}}, Body: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "setpath", Args: []*Query{&Query{Func: "$q"}, &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "getpath", Args: []*Query{&Query{Func: "$q"}}}}}, Op: OpPipe, Right: &Query{Func: "f"}}}}}}, Op: OpPipe, Right: &Query{Left: &Query{Func: "."}, Op: OpComma, Right: &Query{Term: &Term{Type: TermTypeBreak, Break: "$out"}}}}}}}}}}}, Op: OpComma, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "setpath", Args: []*Query{&Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Term: &Term{Type: TermTypeNumber, Number: "1"}}}}}, &Query{Left: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Start: &Query{Term: &Term{Type: TermTypeNumber, Number: "1"}}}}}, Op: OpAdd, Right: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Func: "$p"}}}}}}}}}}}}}}}}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{&Suffix{Bind: &Bind{Patterns: []*Pattern{&Pattern{Name: "$x"}}, Body: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "$x"}, SuffixList: []*Suffix{&Suffix{Index: &Index{Start: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}}}}}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "delpaths", Args: []*Query{&Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "$x"}, SuffixList: []*Suffix{&Suffix{Index: &Index{Start: &Query{Term: &Term{Type: TermTypeNumber, Number: "1"}}}}}}}}}}}}}}}}}}}},
		"all": []*FuncDef{&FuncDef{Name: "all", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "all", Args: []*Query{&Query{Func: "."}}}}}}, &FuncDef{Name: "all", Args: []string{"y"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "all", Args: []*Query{&Query{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{&Suffix{Iter: true}}}}, &Query{Func: "y"}}}}}}, &FuncDef{Name: "all", Args: []string{"g", "y"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "isempty", Args: []*Query{&Query{Left: &Query{Func: "g"}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{&Query{Left: &Query{Func: "y"}, Op: OpPipe, Right: &Query{Func: "not"}}}}}}}}}}}}},
		"any": []*FuncDef{&FuncDef{Name: "any", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "any", Args: []*Query{&Query{Func: "."}}}}}}, &FuncDef{Name: "any", Args: []string{"y"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "any", Args: []*Query{&Query{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{&Suffix{Iter: true}}}}, &Query{Func: "y"}}}}}}, &FuncDef{Name: "any", Args: []string{"g", "y"}, Body: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "isempty", Args: []*Query{&Query{Left: &Query{Func: "g"}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{&Query{Func: "y"}}}}}}}}}}, Op: OpPipe, Right: &Query{Func: "not"}}}},
		"arrays": []*FuncDef{&FuncDef{Name: "arrays", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{&Query{Left: &Query{Func: "type"}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "array"}}}}}}}}}},
		"ascii_downcase": []*FuncDef{&FuncDef{Name: "ascii_downcase", Body: &Query{Left: &Query{Func: "explode"}, Op: OpPipe, Right: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "map", Args: []*Query{&Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Left: &Query{Left: &Query{Term: &Term{Type: TermTypeNumber, Number: "65"}}, Op: OpLe, Right: &Query{Func: "."}}, Op: OpAnd, Right: &Query{Left: &Query{Func: "."}, Op: OpLe, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "90"}}}}, Then: &Query{Left: &Query{Func: "."}, Op: OpAdd, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "32"}}}}}}}}}}, Op: OpPipe, Right: &Query{Func: "implode"}}}}},
		"ascii_upcase": []*FuncDef{&FuncDef{Name: "ascii_upcase", Body: &Query{Left: &Query{Func: "explode"}, Op: OpPipe, Right: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "map", Args: []*Query{&Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Left: &Query{Left: &Query{Term: &Term{Type: TermTypeNumber, Number: "97"}}, Op: OpLe, Right: &Query{Func: "."}}, Op: OpAnd, Right: &Query{Left: &Query{Func: "."}, Op: OpLe, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "122"}}}}, Then: &Query{Left: &Query{Func: "."}, Op: OpSub, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "32"}}}}}}}}}}, Op: OpPipe, Right: &Query{Func: "implode"}}}}},
		"booleans": []*FuncDef{&FuncDef{Name: "booleans", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{&Query{Left: &Query{Func: "type"}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "boolean"}}}}}}}}}},
		"capture": []*FuncDef{&FuncDef{Name: "capture", Args: []string{"$re"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "capture", Args: []*Query{&Query{Func: "$re"}, &Query{Func: "null"}}}}}}, &FuncDef{Name: "capture", Args: []string{"$re", "$flags"}, Body: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "match", Args: []*Query{&Query{Func: "$re"}, &Query{Func: "$flags"}}}}}, Op: OpPipe, Right: &Query{Func: "_capture"}}}},
		"combinations": []*FuncDef{&FuncDef{Name: "combinations", Body: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Left: &Query{Func: "length"}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}, Then: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{}}}, Else: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Start: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}, SuffixList: []*Suffix{&Suffix{Iter: true}, &Suffix{Bind: &Bind{Patterns: []*Pattern{&Pattern{Name: "$x"}}, Body: &Query{Left: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Func: "$x"}}}}, Op: OpAdd, Right: &Query{Term: &Term{Type: TermTypeQuery, Query: &Query{Left: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Start: &Query{Term: &Term{Type: TermTypeNumber, Number: "1"}}, IsSlice: true}}}, Op: OpPipe, Right: &Query{Func: "combinations"}}}}}}}}}}}}}}, &FuncDef{Name: "combinations", Args: []string{"n"}, Body: &Query{Left: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "limit", Args: []*Query{&Query{Func: "n"}, &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "repeat", Args: []*Query{&Query{Func: "."}}}}}}}}}}}}, Op: OpPipe, Right: &Query{Func: "combinations"}}}},
		"del": []*FuncDef{&FuncDef{Name: "del", Args: []string{"f"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "delpaths", Args: []*Query{&Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "path", Args: []*Query{&Query{Func: "f"}}}}}}}}}}}}}},
		"finites": []*FuncDef{&FuncDef{Name: "finites", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{&Query{Func: "isfinite"}}}}}}},
		"first": []*FuncDef{&FuncDef{Name: "first", Body: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Start: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}}}}, &FuncDef{Name: "first", Args: []string{"g"}, Body: &Query{Term: &Term{Type: TermTypeLabel, Label: &Label{Ident: "$out", Body: &Query{Left: &Query{Func: "g"}, Op: OpPipe, Right: &Query{Left: &Query{Func: "."}, Op: OpComma, Right: &Query{Term: &Term{Type: TermTypeBreak, Break: "$out"}}}}}}}}},
		"fromdate": []*FuncDef{&FuncDef{Name: "fromdate", Body: &Query{Func: "fromdateiso8601"}}},
		"fromdateiso8601": []*FuncDef{&FuncDef{Name: "fromdateiso8601", Body: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "strptime", Args: []*Query{&Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "%Y-%m-%dT%H:%M:%S%z"}}}}}}}, Op: OpPipe, Right: &Query{Func: "mktime"}}}},
		"fromstream": []*FuncDef{&FuncDef{Name: "fromstream", Args: []string{"f"}, Body: &Query{Term: &Term{Type: TermTypeObject, Object: &Object{KeyVals: []*ObjectKeyVal{&ObjectKeyVal{Key: "x", Val: &ObjectVal{Queries: []*Query{&Query{Func: "null"}}}}, &ObjectKeyVal{Key: "e", Val: &ObjectVal{Queries: []*Query{&Query{Func: "false"}}}}}}, SuffixList: []*Suffix{&Suffix{Bind: &Bind{Patterns: []*Pattern{&Pattern{Name: "$init"}}, Body: &Query{Term: &Term{Type: TermTypeForeach, Foreach: &Foreach{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "f"}}, Pattern: &Pattern{Name: "$i"}, Start: &Query{Func: "$init"}, Update: &Query{Left: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Name: "e"}}}, Then: &Query{Func: "$init"}, Else: &Query{Func: "."}}}}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Left: &Query{Func: "$i"}, Op: OpPipe, Right: &Query{Left: &Query{Func: "length"}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "2"}}}}, Then: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "setpath", Args: []*Query{&Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "e"}}}}}}, &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "$i"}, SuffixList: []*Suffix{&Suffix{Index: &Index{Start: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}}}}}, Op: OpPipe, Right: &Query{Left: &Query{Func: "length"}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}}}}}}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "setpath", Args: []*Query{&Query{Left: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "x"}}}}}}, Op: OpAdd, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "$i"}, SuffixList: []*Suffix{&Suffix{Index: &Index{Start: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}}}}}}, &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "$i"}, SuffixList: []*Suffix{&Suffix{Index: &Index{Start: &Query{Term: &Term{Type: TermTypeNumber, Number: "1"}}}}}}}}}}}}, Else: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "setpath", Args: []*Query{&Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "e"}}}}}}, &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "$i"}, SuffixList: []*Suffix{&Suffix{Index: &Index{Start: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}}}}}, Op: OpPipe, Right: &Query{Left: &Query{Func: "length"}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "1"}}}}}}}}}}}}, Extract: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Name: "e"}}}, Then: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Name: "x"}}}, Else: &Query{Func: "empty"}}}}}}}}}}}}}},
		"group_by": []*FuncDef{&FuncDef{Name: "group_by", Args: []string{"f"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "_group_by", Args: []*Query{&Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "map", Args: []*Query{&Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Func: "f"}}}}}}}}}}}}}},
		"gsub": []*FuncDef{&FuncDef{Name: "gsub", Args: []string{"$re", "str"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "sub", Args: []*Query{&Query{Func: "$re"}, &Query{Func: "str"}, &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "g"}}}}}}}}, &FuncDef{Name: "gsub", Args: []string{"$re", "str", "$flags"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "sub", Args: []*Query{&Query{Func: "$re"}, &Query{Func: "str"}, &Query{Left: &Query{Func: "$flags"}, Op: OpAdd, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "g"}}}}}}}}}},
		"in": []*FuncDef{&FuncDef{Name: "in", Args: []string{"xs"}, Body: &Query{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{&Suffix{Bind: &Bind{Patterns: []*Pattern{&Pattern{Name: "$x"}}, Body: &Query{Left: &Query{Func: "xs"}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "has", Args: []*Query{&Query{Func: "$x"}}}}}}}}}}}}},
		"inputs": []*FuncDef{&FuncDef{Name: "inputs", Body: &Query{Term: &Term{Type: TermTypeTry, Try: &Try{Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "repeat", Args: []*Query{&Query{Func: "input"}}}}}, Catch: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Left: &Query{Func: "."}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "break"}}}}, Then: &Query{Func: "empty"}, Else: &Query{Func: "error"}}}}}}}}},
		"inside": []*FuncDef{&FuncDef{Name: "inside", Args: []string{"xs"}, Body: &Query{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{&Suffix{Bind: &Bind{Patterns: []*Pattern{&Pattern{Name: "$x"}}, Body: &Query{Left: &Query{Func: "xs"}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "contains", Args: []*Query{&Query{Func: "$x"}}}}}}}}}}}}},
		"isempty": []*FuncDef{&FuncDef{Name: "isempty", Args: []string{"g"}, Body: &Query{Term: &Term{Type: TermTypeLabel, Label: &Label{Ident: "$out", Body: &Query{Left: &Query{Term: &Term{Type: TermTypeQuery, Query: &Query{Left: &Query{Func: "g"}, Op: OpPipe, Right: &Query{Left: &Query{Func: "false"}, Op: OpComma, Right: &Query{Term: &Term{Type: TermTypeBreak, Break: "$out"}}}}}}, Op: OpComma, Right: &Query{Func: "true"}}}}}}},
		"iterables": []*FuncDef{&FuncDef{Name: "iterables", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{&Query{Left: &Query{Func: "type"}, Op: OpPipe, Right: &Query{Left: &Query{Left: &Query{Func: "."}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "array"}}}}, Op: OpOr, Right: &Query{Left: &Query{Func: "."}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "object"}}}}}}}}}}}},
		"last": []*FuncDef{&FuncDef{Name: "last", Body: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Start: &Query{Term: &Term{Type: TermTypeUnary, Unary: &Unary{Op: OpSub, Term: &Term{Type: TermTypeNumber, Number: "1"}}}}}}}}, &FuncDef{Name: "last", Args: []string{"g"}, Body: &Query{Term: &Term{Type: TermTypeReduce, Reduce: &Reduce{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "g"}}, Pattern: &Pattern{Name: "$item"}, Start: &Query{Func: "null"}, Update: &Query{Func: "$item"}}}}}},
		"leaf_paths": []*FuncDef{&FuncDef{Name: "leaf_paths", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "paths", Args: []*Query{&Query{Func: "scalars"}}}}}}},
		"limit": []*FuncDef{&FuncDef{Name: "limit", Args: []string{"$n", "g"}, Body: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Left: &Query{Func: "$n"}, Op: OpGt, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}, Then: &Query{Term: &Term{Type: TermTypeLabel, Label: &Label{Ident: "$out", Body: &Query{Term: &Term{Type: TermTypeForeach, Foreach: &Foreach{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "g"}}, Pattern: &Pattern{Name: "$item"}, Start: &Query{Func: "$n"}, Update: &Query{Left: &Query{Func: "."}, Op: OpSub, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "1"}}}, Extract: &Query{Left: &Query{Func: "$item"}, Op: OpComma, Right: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Left: &Query{Func: "."}, Op: OpLe, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}, Then: &Query{Term: &Term{Type: TermTypeBreak, Break: "$out"}}, Else: &Query{Func: "empty"}}}}}}}}}}}, Elif: []*IfElif{&IfElif{Cond: &Query{Left: &Query{Func: "$n"}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}, Then: &Query{Func: "empty"}}}, Else: &Query{Func: "g"}}}}}},
		"map": []*FuncDef{&FuncDef{Name: "map", Args: []string{"f"}, Body: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{&Suffix{Iter: true}}}}, Op: OpPipe, Right: &Query{Func: "f"}}}}}}},
		"map_values": []*FuncDef{&FuncDef{Name: "map_values", Args: []string{"f"}, Body: &Query{Left: &Query{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{&Suffix{Iter: true}}}}, Op: OpModify, Right: &Query{Func: "f"}}}},
		"match": []*FuncDef{&FuncDef{Name: "match", Args: []string{"$re"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "match", Args: []*Query{&Query{Func: "$re"}, &Query{Func: "null"}}}}}}, &FuncDef{Name: "match", Args: []string{"$re", "$flags"}, Body: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "_match", Args: []*Query{&Query{Func: "$re"}, &Query{Func: "$flags"}, &Query{Func: "false"}}}}}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{&Suffix{Iter: true}}}}}}},
		"max_by": []*FuncDef{&FuncDef{Name: "max_by", Args: []string{"f"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "_max_by", Args: []*Query{&Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "map", Args: []*Query{&Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Func: "f"}}}}}}}}}}}}}},
		"min_by": []*FuncDef{&FuncDef{Name: "min_by", Args: []string{"f"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "_min_by", Args: []*Query{&Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "map", Args: []*Query{&Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Func: "f"}}}}}}}}}}}}}},
		"normals": []*FuncDef{&FuncDef{Name: "normals", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{&Query{Func: "isnormal"}}}}}}},
		"not": []*FuncDef{&FuncDef{Name: "not", Body: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Func: "."}, Then: &Query{Func: "false"}, Else: &Query{Func: "true"}}}}}},
		"nth": []*FuncDef{&FuncDef{Name: "nth", Args: []string{"$n"}, Body: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Start: &Query{Func: "$n"}}}}}, &FuncDef{Name: "nth", Args: []string{"$n", "g"}, Body: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Left: &Query{Func: "$n"}, Op: OpLt, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}, Then: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "error", Args: []*Query{&Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "nth doesn't support negative indices"}}}}}}}, Else: &Query{Term: &Term{Type: TermTypeLabel, Label: &Label{Ident: "$out", Body: &Query{Term: &Term{Type: TermTypeForeach, Foreach: &Foreach{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "g"}}, Pattern: &Pattern{Name: "$item"}, Start: &Query{Left: &Query{Func: "$n"}, Op: OpAdd, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "1"}}}, Update: &Query{Left: &Query{Func: "."}, Op: OpSub, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "1"}}}, Extract: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Left: &Query{Func: "."}, Op: OpLe, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}, Then: &Query{Left: &Query{Func: "$item"}, Op: OpComma, Right: &Query{Term: &Term{Type: TermTypeBreak, Break: "$out"}}}, Else: &Query{Func: "empty"}}}}}}}}}}}}}}},
		"nulls": []*FuncDef{&FuncDef{Name: "nulls", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{&Query{Left: &Query{Func: "."}, Op: OpEq, Right: &Query{Func: "null"}}}}}}}},
		"numbers": []*FuncDef{&FuncDef{Name: "numbers", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{&Query{Left: &Query{Func: "type"}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "number"}}}}}}}}}},
		"objects": []*FuncDef{&FuncDef{Name: "objects", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{&Query{Left: &Query{Func: "type"}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "object"}}}}}}}}}},
		"paths": []*FuncDef{&FuncDef{Name: "paths", Body: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "path", Args: []*Query{&Query{Func: ".."}}}}}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{&Query{Left: &Query{Func: "."}, Op: OpNe, Right: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{}}}}}}}}}}, &FuncDef{Name: "paths", Args: []string{"f"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "paths"}, SuffixList: []*Suffix{&Suffix{Bind: &Bind{Patterns: []*Pattern{&Pattern{Name: "$p"}}, Body: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{&Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "getpath", Args: []*Query{&Query{Func: "$p"}}}}}, Op: OpPipe, Right: &Query{Func: "f"}}}}}}, Op: OpPipe, Right: &Query{Func: "$p"}}}}}}}}},
		"range": []*FuncDef{&FuncDef{Name: "range", Args: []string{"$end"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "_range", Args: []*Query{&Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}, &Query{Func: "$end"}, &Query{Term: &Term{Type: TermTypeNumber, Number: "1"}}}}}}}, &FuncDef{Name: "range", Args: []string{"$start", "$end"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "_range", Args: []*Query{&Query{Func: "$start"}, &Query{Func: "$end"}, &Query{Term: &Term{Type: TermTypeNumber, Number: "1"}}}}}}}, &FuncDef{Name: "range", Args: []string{"$start", "$end", "$step"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "_range", Args: []*Query{&Query{Func: "$start"}, &Query{Func: "$end"}, &Query{Func: "$step"}}}}}}},
		"recurse": []*FuncDef{&FuncDef{Name: "recurse", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "recurse", Args: []*Query{&Query{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{&Suffix{Iter: true}, &Suffix{Optional: true}}}}}}}}}, &FuncDef{Name: "recurse", Args: []string{"f"}, Body: &Query{FuncDefs: []*FuncDef{&FuncDef{Name: "r", Body: &Query{Left: &Query{Func: "."}, Op: OpComma, Right: &Query{Term: &Term{Type: TermTypeQuery, Query: &Query{Left: &Query{Func: "f"}, Op: OpPipe, Right: &Query{Func: "r"}}}}}}}, Func: "r"}}, &FuncDef{Name: "recurse", Args: []string{"f", "cond"}, Body: &Query{FuncDefs: []*FuncDef{&FuncDef{Name: "r", Body: &Query{Left: &Query{Func: "."}, Op: OpComma, Right: &Query{Term: &Term{Type: TermTypeQuery, Query: &Query{Left: &Query{Func: "f"}, Op: OpPipe, Right: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{&Query{Func: "cond"}}}}}, Op: OpPipe, Right: &Query{Func: "r"}}}}}}}}, Func: "r"}}},
		"repeat": []*FuncDef{&FuncDef{Name: "repeat", Args: []string{"f"}, Body: &Query{FuncDefs: []*FuncDef{&FuncDef{Name: "_repeat", Body: &Query{Left: &Query{Func: "f"}, Op: OpComma, Right: &Query{Func: "_repeat"}}}}, Func: "_repeat"}}},
		"scalars": []*FuncDef{&FuncDef{Name: "scalars", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{&Query{Left: &Query{Func: "type"}, Op: OpPipe, Right: &Query{Left: &Query{Left: &Query{Func: "."}, Op: OpNe, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "array"}}}}, Op: OpAnd, Right: &Query{Left: &Query{Func: "."}, Op: OpNe, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "object"}}}}}}}}}}}},
		"scan": []*FuncDef{&FuncDef{Name: "scan", Args: []string{"$re"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "scan", Args: []*Query{&Query{Func: "$re"}, &Query{Func: "null"}}}}}}, &FuncDef{Name: "scan", Args: []string{"$re", "$flags"}, Body: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "match", Args: []*Query{&Query{Func: "$re"}, &Query{Left: &Query{Func: "$flags"}, Op: OpAdd, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "g"}}}}}}}}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Left: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Name: "captures"}}}, Op: OpPipe, Right: &Query{Left: &Query{Func: "length"}, Op: OpGt, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}}, Then: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Name: "captures"}, SuffixList: []*Suffix{&Suffix{Iter: true}, &Suffix{Index: &Index{Name: "string"}}}}}}}}, Else: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Name: "string"}}}}}}}}},
		"select": []*FuncDef{&FuncDef{Name: "select", Args: []string{"f"}, Body: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Func: "f"}, Then: &Query{Func: "."}, Else: &Query{Func: "empty"}}}}}},
		"sort_by": []*FuncDef{&FuncDef{Name: "sort_by", Args: []string{"f"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "_sort_by", Args: []*Query{&Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "map", Args: []*Query{&Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Func: "f"}}}}}}}}}}}}}},
		"splits": []*FuncDef{&FuncDef{Name: "splits", Args: []string{"$re"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "splits", Args: []*Query{&Query{Func: "$re"}, &Query{Func: "null"}}}}}}, &FuncDef{Name: "splits", Args: []string{"$re", "$flags"}, Body: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "split", Args: []*Query{&Query{Func: "$re"}, &Query{Func: "$flags"}}}}}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{&Suffix{Iter: true}}}}}}},
		"strings": []*FuncDef{&FuncDef{Name: "strings", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{&Query{Left: &Query{Func: "type"}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "string"}}}}}}}}}},
		"sub": []*FuncDef{&FuncDef{Name: "sub", Args: []string{"$re", "str"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "sub", Args: []*Query{&Query{Func: "$re"}, &Query{Func: "str"}, &Query{Func: "null"}}}}}}, &FuncDef{Name: "sub", Args: []string{"$re", "str", "$flags"}, Body: &Query{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{&Suffix{Bind: &Bind{Patterns: []*Pattern{&Pattern{Name: "$str"}}, Body: &Query{FuncDefs: []*FuncDef{&FuncDef{Name: "_sub", Body: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Left: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Name: "matches"}}}, Op: OpPipe, Right: &Query{Left: &Query{Func: "length"}, Op: OpGt, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}}, Then: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Name: "matches"}, SuffixList: []*Suffix{&Suffix{Index: &Index{Start: &Query{Term: &Term{Type: TermTypeUnary, Unary: &Unary{Op: OpSub, Term: &Term{Type: TermTypeNumber, Number: "1"}}}}}}, &Suffix{Bind: &Bind{Patterns: []*Pattern{&Pattern{Name: "$r"}}, Body: &Query{Left: &Query{Term: &Term{Type: TermTypeObject, Object: &Object{KeyVals: []*ObjectKeyVal{&ObjectKeyVal{Key: "string", Val: &ObjectVal{Queries: []*Query{&Query{Term: &Term{Type: TermTypeQuery, Query: &Query{Left: &Query{Left: &Query{Term: &Term{Type: TermTypeQuery, Query: &Query{Left: &Query{Func: "$r"}, Op: OpPipe, Right: &Query{Left: &Query{Func: "_capture"}, Op: OpPipe, Right: &Query{Func: "str"}}}}}, Op: OpAdd, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "$str"}, SuffixList: []*Suffix{&Suffix{Index: &Index{Start: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "$r"}, SuffixList: []*Suffix{&Suffix{Index: &Index{Name: "offset"}}}}}, Op: OpAdd, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "$r"}, SuffixList: []*Suffix{&Suffix{Index: &Index{Name: "length"}}}}}}, End: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Name: "offset"}}}, IsSlice: true}}}}}}, Op: OpAdd, Right: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Name: "string"}}}}}}}}}, &ObjectKeyVal{Key: "offset", Val: &ObjectVal{Queries: []*Query{&Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "$r"}, SuffixList: []*Suffix{&Suffix{Index: &Index{Name: "offset"}}}}}}}}, &ObjectKeyVal{Key: "matches", Val: &ObjectVal{Queries: []*Query{&Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Name: "matches"}, SuffixList: []*Suffix{&Suffix{Index: &Index{End: &Query{Term: &Term{Type: TermTypeUnary, Unary: &Unary{Op: OpSub, Term: &Term{Type: TermTypeNumber, Number: "1"}}}}, IsSlice: true}}}}}}}}}}}}, Op: OpPipe, Right: &Query{Func: "_sub"}}}}}}}, Else: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "$str"}, SuffixList: []*Suffix{&Suffix{Index: &Index{End: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Name: "offset"}}}, IsSlice: true}}}}}, Op: OpAdd, Right: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Name: "string"}}}}}}}}}, Left: &Query{Term: &Term{Type: TermTypeObject, Object: &Object{KeyVals: []*ObjectKeyVal{&ObjectKeyVal{Key: "string", Val: &ObjectVal{Queries: []*Query{&Query{Term: &Term{Type: TermTypeString, Str: &String{}}}}}}, &ObjectKeyVal{Key: "matches", Val: &ObjectVal{Queries: []*Query{&Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "match", Args: []*Query{&Query{Func: "$re"}, &Query{Func: "$flags"}}}}}}}}}}}}}}}, Op: OpPipe, Right: &Query{Func: "_sub"}}}}}}}}},
		"test": []*FuncDef{&FuncDef{Name: "test", Args: []string{"$re"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "test", Args: []*Query{&Query{Func: "$re"}, &Query{Func: "null"}}}}}}, &FuncDef{Name: "test", Args: []string{"$re", "$flags"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "_match", Args: []*Query{&Query{Func: "$re"}, &Query{Func: "$flags"}, &Query{Func: "true"}}}}}}},
		"todate": []*FuncDef{&FuncDef{Name: "todate", Body: &Query{Func: "todateiso8601"}}},
		"todateiso8601": []*FuncDef{&FuncDef{Name: "todateiso8601", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "strftime", Args: []*Query{&Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "%Y-%m-%dT%H:%M:%SZ"}}}}}}}}},
		"tostream": []*FuncDef{&FuncDef{Name: "tostream", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "path", Args: []*Query{&Query{FuncDefs: []*FuncDef{&FuncDef{Name: "r", Body: &Query{Left: &Query{Term: &Term{Type: TermTypeQuery, Query: &Query{Left: &Query{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{&Suffix{Iter: true}, &Suffix{Optional: true}}}}, Op: OpPipe, Right: &Query{Func: "r"}}}}, Op: OpComma, Right: &Query{Func: "."}}}}, Func: "r"}}}, SuffixList: []*Suffix{&Suffix{Bind: &Bind{Patterns: []*Pattern{&Pattern{Name: "$p"}}, Body: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "getpath", Args: []*Query{&Query{Func: "$p"}}}}}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeReduce, Reduce: &Reduce{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "path", Args: []*Query{&Query{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{&Suffix{Iter: true}, &Suffix{Optional: true}}}}}}}, Pattern: &Pattern{Name: "$q"}, Start: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Func: "$p"}, Op: OpComma, Right: &Query{Func: "."}}}}}, Update: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Func: "$p"}, Op: OpAdd, Right: &Query{Func: "$q"}}}}}}}}}}}}}}}},
		"truncate_stream": []*FuncDef{&FuncDef{Name: "truncate_stream", Args: []string{"f"}, Body: &Query{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{&Suffix{Bind: &Bind{Patterns: []*Pattern{&Pattern{Name: "$n"}}, Body: &Query{Left: &Query{Func: "null"}, Op: OpPipe, Right: &Query{Left: &Query{Func: "f"}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Left: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Start: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}}}, Op: OpPipe, Right: &Query{Left: &Query{Func: "length"}, Op: OpGt, Right: &Query{Func: "$n"}}}, Then: &Query{Left: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Start: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}}}, Op: OpModify, Right: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Start: &Query{Func: "$n"}, IsSlice: true}}}}, Else: &Query{Func: "empty"}}}}}}}}}}}}},
		"unique_by": []*FuncDef{&FuncDef{Name: "unique_by", Args: []string{"f"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "_unique_by", Args: []*Query{&Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "map", Args: []*Query{&Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Func: "f"}}}}}}}}}}}}}},
		"until": []*FuncDef{&FuncDef{Name: "until", Args: []string{"cond", "next"}, Body: &Query{FuncDefs: []*FuncDef{&FuncDef{Name: "_until", Body: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Func: "cond"}, Then: &Query{Func: "."}, Else: &Query{Left: &Query{Func: "next"}, Op: OpPipe, Right: &Query{Func: "_until"}}}}}}}, Func: "_until"}}},
		"values": []*FuncDef{&FuncDef{Name: "values", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{&Query{Left: &Query{Func: "."}, Op: OpNe, Right: &Query{Func: "null"}}}}}}}},
		"walk": []*FuncDef{&FuncDef{Name: "walk", Args: []string{"f"}, Body: &Query{FuncDefs: []*FuncDef{&FuncDef{Name: "_walk", Body: &Query{Left: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Left: &Query{Func: "type"}, Op: OpPipe, Right: &Query{Left: &Query{Left: &Query{Func: "."}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "array"}}}}, Op: OpOr, Right: &Query{Left: &Query{Func: "."}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "object"}}}}}}, Then: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "map_values", Args: []*Query{&Query{Func: "_walk"}}}}}}}}, Op: OpPipe, Right: &Query{Func: "f"}}}}, Func: "_walk"}}},
		"while": []*FuncDef{&FuncDef{Name: "while", Args: []string{"cond", "update"}, Body: &Query{FuncDefs: []*FuncDef{&FuncDef{Name: "_while", Body: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Func: "cond"}, Then: &Query{Left: &Query{Func: "."}, Op: OpComma, Right: &Query{Term: &Term{Type: TermTypeQuery, Query: &Query{Left: &Query{Func: "update"}, Op: OpPipe, Right: &Query{Func: "_while"}}}}}, Else: &Query{Func: "empty"}}}}}}, Func: "_while"}}},
		"with_entries": []*FuncDef{&FuncDef{Name: "with_entries", Args: []string{"f"}, Body: &Query{Left: &Query{Func: "to_entries"}, Op: OpPipe, Right: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "map", Args: []*Query{&Query{Func: "f"}}}}}, Op: OpPipe, Right: &Query{Func: "from_entries"}}}}},
	}
}
