package gen

import "github.com/dave/jennifer/jen"

func Map(t string) *jen.File {
	f := jen.NewFile("f" + t)

	f.Type().Id("MapIter").Struct(
		jen.Id("iter").Id("Iter"),
		jen.Id("op").Func().Params(jen.Id(t)).Id(t),
	)
	f.Line()

	f.Func().Id("NewMap").Params(jen.Id("iter").Id("Iter"), jen.Id("op").Func().Params(jen.Id(t)).Id(t)).Id("MapIter").Block(
		jen.Return(jen.Id("MapIter").Values(jen.Dict{
			jen.Id("iter"): jen.Id("iter"),
			jen.Id("op"):   jen.Id("op"),
		})),
	)
	f.Line()

	f.Func().Params(jen.Id("iter").Id("MapIter")).Id("Next").Params().Id("Option").Block(
		jen.Id("next").Op(":=").Id("iter").Dot("iter").Dot("Next").Call(),
		jen.If(jen.Op("!").Id("next").Dot("Present").Call()).Block(
			jen.Return(jen.Id("next")),
		),
		jen.Line(),
		jen.Return(jen.Id("Some").Call(jen.Id("iter").Dot("op").Call(jen.Id("next").Dot("Value")))),
	)

	return f
}
