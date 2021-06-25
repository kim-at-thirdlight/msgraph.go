package main

import (
	"fmt"
	"strings"
)

const (
	colPrefix = "Collection("
)

var nsPrefixes = map[string]struct{}{
	"microsoft.graph.": {},
	"graph.":           {},
}

var reservedTypeTable = map[string]string{
	"Edm.Boolean":          "bool",
	"Edm.Byte":             "byte",
	"Edm.Int16":            "int",
	"Edm.Int32":            "int",
	"Edm.Int64":            "int",
	"Edm.Decimal":          "int",
	"Edm.Single":           "float64",
	"Edm.Double":           "float64",
	"Edm.Binary":           "Binary",
	"Edm.Stream":           "Stream",
	"Edm.Guid":             "UUID",
	"Edm.String":           "string",
	"Edm.DateTimeOffset":   "time.Time",
	"Edm.Duration":         "Duration",
	"Edm.TimeOfDay":        "TimeOfDay",
	"Edm.Date":             "Date",
	"microsoft.graph.Json": "json.RawMessage",
}

func ptrType(t string) string {
	switch t {
	case "json.RawMessage":
		return t
	}
	return "*" + t
}

func formatNS(t string) (string, bool) {
	for ns := range nsPrefixes {
		if strings.HasPrefix(t, ns) {
			return strings.Replace(
				t[len(ns):],
				".",
				"",
				-1,
			), true
		}
	}
	return t, false
}

func exported(n string) string {
	n = lintName(strings.Title(n))
	if strings.HasSuffix(n, "Request") {
		// NOTE: this is a really hacky place to put this
		//       logic, but it ensures it occurs to all exported
		//       names (which was previously an issue)
		n += "Object"
	}
	return n
}

func isCollectionType(t string) bool {
	return strings.HasPrefix(t, colPrefix)
}

func (g *Generator) SymBaseType(t string) string {
	if x, ok := g.SymTypeMap[t]; ok {
		return x
	}
	if x, ok := formatNS(t); ok {
		return exported(x)
	}
	if strings.HasPrefix(t, colPrefix) {
		return g.SymBaseType(t[len(colPrefix) : len(t)-1])
	}
	panic(fmt.Errorf("Unknown type %s", t))
}

func (g *Generator) SymFromType(t string) string {
	if x, ok := g.SymTypeMap[t]; ok {
		return x
	}
	if x, ok := formatNS(t); ok {
		return exported(x)
	}
	if strings.HasPrefix(t, colPrefix) {
		return g.SymBaseType(t[len(colPrefix):len(t)-1]) + "Collection"
	}
	panic(fmt.Errorf("Unknown type %s", t))
}

func (g *Generator) TypeFromType(t string) string {
	if x, ok := reservedTypeTable[t]; ok {
		return ptrType(x)
	}
	if x, ok := g.SymTypeMap[t]; ok {
		return ptrType(x)
	}
	if x, ok := formatNS(t); ok {
		return ptrType(exported(x))
	}
	if strings.HasPrefix(t, colPrefix) {
		return "[]" + g.TypeFromType(t[len(colPrefix) : len(t)-1])[1:]
	}
	panic(fmt.Errorf("Unknown type %s", t))
}
