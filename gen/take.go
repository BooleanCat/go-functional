package gen

import "github.com/dave/jennifer/jen"

func Take(t string) *jen.File {
	f := jen.NewFile("f" + t)

	f.Type().Id("TakeIter").Struct(
		jen.Id("iter").Id("Iter"),
		jen.Id("n").Int(),
	)
	f.Line()

	f.Func().Id("NewTake").Params(jen.Id("iter").Id("Iter"), jen.Id("n").Int()).Op("*").Id("TakeIter").Block(
		jen.Return(jen.Op("&").Id("TakeIter").Values(jen.Dict{
			jen.Id("iter"): jen.Id("iter"),
			jen.Id("n"):    jen.Id("n"),
		})),
	)
	f.Line()

	f.Func().Params(jen.Id("iter").Id("*").Id("TakeIter")).Id("Next").Params().Id("Option").Block(
		jen.If(jen.Id("iter").Dot("n").Op("<=").Lit(0)).Block(
			jen.Return(jen.Id("None").Call()),
		),
		jen.Line(),
		jen.Id("iter").Dot("n").Op("--"),
		jen.Return(jen.Id("iter").Dot("iter").Dot("Next").Call()),
	)

	return f
}
