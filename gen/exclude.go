package gen

import "github.com/dave/jennifer/jen"

func Exclude(t string) *jen.File {
	f := jen.NewFile("f" + t)

	f.Type().Id("ExcludeIter").Struct(
		jen.Id("iter").Id("Iter"),
		jen.Id("exclude").Func().Params(jen.Id(t)).Bool(),
	)
	f.Line()

	f.Func().Id("NewExclude").Params(
		jen.Id("iter").Id("Iter"),
		jen.Id("exclude").Func().Params(jen.Id(t)).Bool(),
	).Id("ExcludeIter").Block(
		jen.Return(jen.Id("ExcludeIter").Values(jen.Dict{
			jen.Id("iter"):    jen.Id("iter"),
			jen.Id("exclude"): jen.Id("exclude"),
		})),
	)
	f.Line()

	f.Func().Params(jen.Id("iter").Id("ExcludeIter")).Id("Next").Params().Id("Option").Block(
		jen.For().Block(
			jen.If(jen.Id("option").Op(":=").Id("iter").Dot("iter").Dot("Next").Call().Op(";").Op("!").Id("option").Dot("Present").Call().Op("||").Op("!").Id("iter").Dot("exclude").Call(jen.Id("option").Dot("Value"))).Block(
				jen.Return(jen.Id("option")),
			),
		),
	)

	return f
}
