package gen

import "github.com/dave/jennifer/jen"

func Iter(t string) *jen.File {
	f := jen.NewFile("f" + t)

	f.Type().Id("Iter").Interface(
		jen.Id("Next").Params().Id("Option"),
	)
	f.Line()

	f.Func().Id("Collect").Params(jen.Id("iter").Id("Iter")).Index().Id(t).Block(
		jen.Id("slice").Op(":=").Index().Id(t).Values(),
		jen.For().Block(
			jen.Id("option").Op(":=").Id("iter").Dot("Next").Call(),
			jen.If(jen.Op("!").Id("option").Dot("Present").Call()).Block(
				jen.Return(jen.Id("slice")),
			),
			jen.Line(),
			jen.Id("slice").Op("=").Append(jen.Id("slice"), jen.Id("option").Dot("Value")),
		),
	)
	f.Line()

	f.Func().Id("Fold").Params(jen.Id("iter").Id("Iter"), jen.Id("initial").Id(t), jen.Id("op").Id("foldOp")).Id(t).Block(
		jen.Id("result").Op(":=").Id("initial"),
		jen.For().Block(
			jen.Id("next").Op(":=").Id("iter").Dot("Next").Call(),
			jen.If(jen.Op("!").Id("next").Dot("Present").Call()).Block(
				jen.Return(jen.Id("result")),
			),
			jen.Line(),
			jen.Id("result").Op("=").Id("op").Call(jen.Id("result"), jen.Id("next").Dot("Value")),
		),
	)
	f.Line()

	f.Type().Id("foldOp").Func().Params(jen.Id(t), jen.Id(t)).Id(t)

	return f
}
