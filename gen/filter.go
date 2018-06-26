package gen

import "github.com/dave/jennifer/jen"

func Filter(t string) *jen.File {
	f := jen.NewFile("f" + t)

	f.Type().Id("FilterIter").Struct(
		jen.Id("iter").Id("Iter"),
		jen.Id("filter").Func().Params(jen.Id(t)).Bool(),
	)
	f.Line()

	f.Func().Id("NewFilter").Params(jen.Id("iter").Id("Iter"), jen.Id("filter").Func().Params(jen.Id(t)).Bool()).Id("FilterIter").Block(
		jen.Return(jen.Id("FilterIter").Values(jen.Dict{
			jen.Id("iter"):   jen.Id("iter"),
			jen.Id("filter"): jen.Id("filter"),
		})),
	)
	f.Line()

	f.Func().Params(jen.Id("iter").Id("FilterIter")).Id("Next").Params().Id("Option").Block(
		jen.For().Block(
			jen.If(jen.Id("option").Op(":=").Id("iter").Dot("iter").Dot("Next").Call().Op(";").Op("!").Id("option").Dot("Present").Call().Op("||").Id("iter").Dot("filter").Call(jen.Id("option").Dot("Value"))).Block(
				jen.Return(jen.Id("option")),
			),
		),
	)

	return f
}
