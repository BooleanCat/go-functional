package gen

import (
	"strings"

	"github.com/dave/jennifer/jen"
)

type TypeFileGen struct {
	typeName   string
	pointer    bool
	importPath string
}

func NewTypeFileGen(typeName string) TypeFileGen {
	g := TypeFileGen{typeName: typeName}
	if strings.HasPrefix(g.typeName, "*") {
		g.typeName = g.typeName[1:]
		g.pointer = true
	}

	return g
}

func (g TypeFileGen) ImportedFrom(path string) TypeFileGen {
	g.importPath = path
	return g
}

func (g TypeFileGen) File() *jen.File {
	f := jen.NewFile(g.packageName())

	for _, statement := range []*jen.Statement{
		g.defs(),
		g.fromT(),
		g.collect(),
		g.functorCollect(),
		g.collapse(),
		g.functorCollapse(),
		g.fold(),
		g.functorFold(),
		g.roll(),
		g.functorRoll(),
		g.transmute(),
		g.asMapErrFunc(),
		g.asFilterErrFunc(),
		g.asFoldErrFunc(),
	} {
		f.Add(statement)
	}

	return f
}

func (g TypeFileGen) defs() *jen.Statement {
	return jen.Type().Defs(
		jen.Id("T").Add(g.typeQualifier()),
		jen.Id("tSlice").Index().Add(g.typeQualifier()),
		jen.Id("mapFunc").Func().Params(g.typeQualifier()).Add(g.typeQualifier()),
		jen.Id("mapErrFunc").Func().Params(g.typeQualifier()).Params(g.typeQualifier(), jen.Error()),
		jen.Id("foldFunc").Func().Params(g.typeQualifier(), g.typeQualifier()).Add(g.typeQualifier()),
		jen.Id("foldErrFunc").Func().Params(g.typeQualifier(), g.typeQualifier()).Params(g.typeQualifier(), jen.Error()),
		jen.Id("filterFunc").Func().Params(g.typeQualifier()).Bool(),
		jen.Id("filterErrFunc").Func().Params(g.typeQualifier()).Params(jen.Bool(), jen.Error()),
		jen.Id("transformFunc").Func().Params(jen.Interface()).Params(g.typeQualifier(), jen.Error()),
	)
}

func (g TypeFileGen) fromT() *jen.Statement {
	body := jen.Return(jen.Add(g.typeQualifier()).Call(jen.Id("t")))

	if g.pointer {
		body = jen.Return(jen.Id("t"))
	}

	return jen.Func().Id("fromT").Params(jen.Id("t").Id("T")).Add(g.typeQualifier()).Block(body)
}

func (g TypeFileGen) collect() *jen.Statement {
	return jen.Func().Id("Collect").Params(jen.Id("iter").Id("Iter")).Params(jen.Index().Add(g.typeQualifier()), jen.Error()).Block(
		jen.Return(jen.Id("collect").Call(jen.Id("iter"))),
	)
}

func (g TypeFileGen) functorCollect() *jen.Statement {
	return jen.Func().Params(jen.Id("f").Op("*").Id("Functor")).Id("Collect").Params().Params(jen.Index().Add(g.typeQualifier()), jen.Error()).Block(
		jen.Return(jen.Id("collect").Call(jen.Id("f").Dot("iter"))),
	)
}

func (g TypeFileGen) collapse() *jen.Statement {
	return jen.Func().Id("Collapse").Params(jen.Id("iter").Id("Iter")).Index().Add(g.typeQualifier()).Block(
		jen.Return(jen.Id("collapse").Call(jen.Id("iter"))),
	)
}

func (g TypeFileGen) functorCollapse() *jen.Statement {
	return jen.Func().Params(jen.Id("f").Op("*").Id("Functor")).Id("Collapse").Params().Index().Add(g.typeQualifier()).Block(
		jen.Return(jen.Id("collapse").Call(jen.Id("f").Dot("iter"))),
	)
}

func (g TypeFileGen) fold() *jen.Statement {
	return jen.Func().Id("Fold").Params(jen.Id("iter").Id("Iter"), jen.Id("initial").Add(g.typeQualifier()), jen.Id("op").Id("foldErrFunc")).Params(g.typeQualifier(), jen.Error()).Block(
		jen.List(jen.Id("result"), jen.Id("err")).Op(":=").Id("fold").Call(jen.Id("iter"), jen.Id("T").Call(jen.Id("initial")), jen.Id("op")),
		jen.Return(jen.Id("fromT").Call(jen.Id("result")), jen.Id("err")),
	)
}

func (g TypeFileGen) functorFold() *jen.Statement {
	return jen.Func().Params(jen.Id("f").Op("*").Id("Functor")).Id("Fold").Params(jen.Id("initial").Add(g.typeQualifier()), jen.Id("op").Id("foldErrFunc")).Params(g.typeQualifier(), jen.Error()).Block(
		jen.Return(jen.Id("Fold").Call(jen.Id("f").Dot("iter"), jen.Id("initial"), jen.Id("op"))),
	)
}

func (g TypeFileGen) roll() *jen.Statement {
	return jen.Func().Id("Roll").Params(jen.Id("iter").Id("Iter"), jen.Id("initial").Add(g.typeQualifier()), jen.Id("op").Id("foldFunc")).Add(g.typeQualifier()).Block(
		jen.Return(jen.Id("fromT").Call(jen.Id("roll").Call(jen.Id("iter"), jen.Id("T").Call(jen.Id("initial")), jen.Id("op")))),
	)
}

func (g TypeFileGen) functorRoll() *jen.Statement {
	return jen.Func().Params(jen.Id("f").Op("*").Id("Functor")).Id("Roll").Params(jen.Id("initial").Add(g.typeQualifier()), jen.Id("op").Id("foldFunc")).Add(g.typeQualifier()).Block(
		jen.Return(jen.Id("Roll").Call(jen.Id("f").Dot("iter"), jen.Id("initial"), jen.Id("op"))),
	)
}

func (g TypeFileGen) transmute() *jen.Statement {
	return jen.Func().Id("Transmute").Params(jen.Id("v").Interface()).Add(g.typeQualifier()).Block(
		jen.List(jen.Id("result"), jen.Id("ok")).Op(":=").Id("v").Assert(g.typeQualifier()),
		jen.If(jen.Op("!").Id("ok")).Block(
			jen.Panic(jen.Qual("fmt", "Sprintf").Call(jen.Lit("could not transmute: %v"), jen.Id("v"))),
		),
		jen.Return(jen.Id("result")),
	)
}

func (g TypeFileGen) asMapErrFunc() *jen.Statement {
	return jen.Func().Id("asMapErrFunc").Params(jen.Id("f").Id("mapFunc")).Id("mapErrFunc").Block(
		jen.Return(jen.Func().Params(jen.Id("v").Add(g.typeQualifier())).Params(g.typeQualifier(), jen.Error()).Block(
			jen.Return(jen.Id("f").Call(jen.Id("v")), jen.Nil()),
		)),
	)
}

func (g TypeFileGen) asFilterErrFunc() *jen.Statement {
	return jen.Func().Id("asFilterErrFunc").Params(jen.Id("f").Id("filterFunc")).Id("filterErrFunc").Block(
		jen.Return(jen.Func().Params(jen.Id("v").Add(g.typeQualifier())).Params(jen.Bool(), jen.Error())).Block(
			jen.Return(jen.Id("f").Call(jen.Id("v")), jen.Nil()),
		),
	)
}

func (g TypeFileGen) asFoldErrFunc() *jen.Statement {
	return jen.Func().Id("asFoldErrFunc").Params(jen.Id("f").Id("foldFunc")).Id("foldErrFunc").Block(
		jen.Return(jen.Func().Params(jen.Id("v").Add(g.typeQualifier()), jen.Id("w").Add(g.typeQualifier())).Params(g.typeQualifier(), jen.Error())).Block(
			jen.Return(jen.Id("f").Call(jen.Id("v"), jen.Id("w")), jen.Nil()),
		),
	)
}

func (g TypeFileGen) typeQualifier() *jen.Statement {
	s := jen.Id(g.typeName)
	if g.importPath != "" {
		s = jen.Qual(g.importPath, g.typeName)
	}

	if g.pointer {
		return jen.Op("*").Add(s)
	}

	return s
}

func (g TypeFileGen) packageName() string {
	if g.typeName == "interface{}" {
		return "finterface"
	}
	if g.pointer {
		return "fp" + g.typeName
	}
	return "f" + g.typeName
}
