package gen

import (
	"strings"

	"github.com/dave/jennifer/jen"
)

type TypeFileGen struct {
	typeName string
}

func NewTypeFileGen(typeName string) TypeFileGen {
	return TypeFileGen{typeName}
}

func (g TypeFileGen) File() *jen.File {
	f := jen.NewFile(packageName(g.typeName))

	f.Add(g.defs())
	f.Add(g.fromT())
	f.Add(g.collect())
	f.Add(g.functorCollect())

	f.Func().Id("Collapse").Params(jen.Id("iter").Id("Iter")).Index().Id(g.typeName).Block(
		jen.Return(jen.Id("collapse").Call(jen.Id("iter"))),
	)

	f.Func().Params(jen.Id("f").Op("*").Id("Functor")).Id("Collapse").Params().Index().Id(g.typeName).Block(
		jen.Return(jen.Id("collapse").Call(jen.Id("f").Dot("iter"))),
	)

	f.Func().Id("Fold").Params(jen.Id("iter").Id("Iter"), jen.Id("initial").Id(g.typeName), jen.Id("op").Id("foldErrFunc")).Params(jen.Id(g.typeName), jen.Error()).Block(
		jen.List(jen.Id("result"), jen.Id("err")).Op(":=").Id("fold").Call(jen.Id("iter"), jen.Id("T").Call(jen.Id("initial")), jen.Id("op")),
		jen.Return(jen.Id("fromT").Call(jen.Id("result")), jen.Id("err")),
	)

	f.Func().Params(jen.Id("f").Op("*").Id("Functor")).Id("Fold").Params(jen.Id("initial").Id(g.typeName), jen.Id("op").Id("foldErrFunc")).Params(jen.Id(g.typeName), jen.Error()).Block(
		jen.Return(jen.Id("Fold").Call(jen.Id("f").Dot("iter"), jen.Id("initial"), jen.Id("op"))),
	)

	f.Func().Id("Roll").Params(jen.Id("iter").Id("Iter"), jen.Id("initial").Id(g.typeName), jen.Id("op").Id("foldFunc")).Id(g.typeName).Block(
		jen.Return(jen.Id("fromT").Call(jen.Id("roll").Call(jen.Id("iter"), jen.Id("T").Call(jen.Id("initial")), jen.Id("op")))),
	)

	f.Func().Params(jen.Id("f").Op("*").Id("Functor")).Id("Roll").Params(jen.Id("initial").Id(g.typeName), jen.Id("op").Id("foldFunc")).Id(g.typeName).Block(
		jen.Return(jen.Id("Roll").Call(jen.Id("f").Dot("iter"), jen.Id("initial"), jen.Id("op"))),
	)

	f.Func().Id("Transmute").Params(jen.Id("v").Interface()).Id(g.typeName).Block(
		jen.List(jen.Id("result"), jen.Id("ok")).Op(":=").Id("v").Assert(jen.Id(g.typeName)),
		jen.If(jen.Op("!").Id("ok")).Block(
			jen.Panic(jen.Qual("fmt", "Sprintf").Call(jen.Lit("could not transmute: %v"), jen.Id("v"))),
		),
		jen.Return(jen.Id("result")),
	)

	f.Func().Id("asMapErrFunc").Params(jen.Id("f").Id("mapFunc")).Id("mapErrFunc").Block(
		jen.Return(jen.Func().Params(jen.Id("v").Id(g.typeName)).Params(jen.Id(g.typeName), jen.Error()).Block(
			jen.Return(jen.Id("f").Call(jen.Id("v")), jen.Nil()),
		)),
	)

	f.Func().Id("asFilterErrFunc").Params(jen.Id("f").Id("filterFunc")).Id("filterErrFunc").Block(
		jen.Return(jen.Func().Params(jen.Id("v").Id(g.typeName)).Params(jen.Bool(), jen.Error())).Block(
			jen.Return(jen.Id("f").Call(jen.Id("v")), jen.Nil()),
		),
	)

	f.Func().Id("asFoldErrFunc").Params(jen.Id("f").Id("foldFunc")).Id("foldErrFunc").Block(
		jen.Return(jen.Func().Params(jen.Id("v").Id(g.typeName), jen.Id("w").Id(g.typeName)).Params(jen.Id(g.typeName), jen.Error())).Block(
			jen.Return(jen.Id("f").Call(jen.Id("v"), jen.Id("w")), jen.Nil()),
		),
	)

	return f
}

func (g TypeFileGen) defs() *jen.Statement {
	return jen.Type().Defs(
		jen.Id("T").Id(g.typeName),
		jen.Id("tSlice").Index().Id(g.typeName),
		jen.Id("mapFunc").Func().Params(jen.Id(g.typeName)).Id(g.typeName),
		jen.Id("mapErrFunc").Func().Params(jen.Id(g.typeName)).Params(jen.Id(g.typeName), jen.Error()),
		jen.Id("foldFunc").Func().Params(jen.Id(g.typeName), jen.Id(g.typeName)).Id(g.typeName),
		jen.Id("foldErrFunc").Func().Params(jen.Id(g.typeName), jen.Id(g.typeName)).Params(jen.Id(g.typeName), jen.Error()),
		jen.Id("filterFunc").Func().Params(jen.Id(g.typeName)).Bool(),
		jen.Id("filterErrFunc").Func().Params(jen.Id(g.typeName)).Params(jen.Bool(), jen.Error()),
		jen.Id("transformFunc").Func().Params(jen.Interface()).Params(jen.Id(g.typeName), jen.Error()),
	)
}

func (g TypeFileGen) fromT() *jen.Statement {
	body := jen.Return(jen.Id(g.typeName).Call(jen.Id("t")))

	if strings.HasPrefix(g.typeName, "*") {
		body = jen.Return(jen.Id("t"))
	}

	return jen.Func().Id("fromT").Params(jen.Id("t").Id("T")).Id(g.typeName).Block(body)
}

func (g TypeFileGen) collect() *jen.Statement {
	return jen.Func().Id("Collect").Params(jen.Id("iter").Id("Iter")).Params(jen.Index().Id(g.typeName), jen.Error()).Block(
		jen.Return(jen.Id("collect").Call(jen.Id("iter"))),
	)
}

func (g TypeFileGen) functorCollect() *jen.Statement {
	return jen.Func().Params(jen.Id("f").Op("*").Id("Functor")).Id("Collect").Params().Params(jen.Index().Id(g.typeName), jen.Error()).Block(
		jen.Return(jen.Id("collect").Call(jen.Id("f").Dot("iter"))),
	)
}

func packageName(typeName string) string {
	if typeName == "interface{}" {
		return "finterface"
	}
	if strings.HasPrefix(typeName, "*") {
		return "fp" + typeName[1:len(typeName)]
	}
	return "f" + typeName
}
