package gen

import "github.com/dave/jennifer/jen"

func Option(t string) *jen.File {
	f := jen.NewFile("f" + t)

	f.Type().Id("Option").Struct(
		jen.Id("Value").Id(t),
		jen.Id("present").Bool(),
	)

	f.Func().Id("Some").Params(jen.Id("value").Id(t)).Id("Option").Block(
		jen.Return(jen.Id("Option").Values(jen.Dict{
			jen.Id("Value"):   jen.Id("value"),
			jen.Id("present"): jen.True(),
		})),
	)
	f.Line()

	f.Func().Id("None").Params().Id("Option").Block(
		jen.Return(jen.Id("Option").Values()),
	)
	f.Line()

	f.Func().Params(jen.Id("o").Id("Option")).Id("Present").Params().Bool().Block(
		jen.Return(jen.Id("o").Dot("present")),
	)

	return f
}
