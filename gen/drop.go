package gen

import "github.com/dave/jennifer/jen"

func Drop(t string) *jen.File {
	f := jen.NewFile("f" + t)

	f.Type().Id("DropIter").Struct(
		jen.Id("iter").Id("Iter"),
		jen.Id("n").Int(),
	)
	f.Line()

	f.Func().Id("NewDrop").Params(jen.Id("iter").Id("Iter"), jen.Id("n").Int()).Op("*").Id("DropIter").Block(
		jen.Return(jen.Op("&").Id("DropIter").Values(jen.Dict{
			jen.Id("iter"): jen.Id("iter"),
			jen.Id("n"):    jen.Id("n"),
		})),
	)
	f.Line()

	f.Func().Params(jen.Id("iter").Op("*").Id("DropIter")).Id("Next").Params().Id("Option").Block(
		jen.For(jen.Id("iter").Dot("n").Op(">").Lit(0)).Block(
			jen.Id("iter").Dot("n").Op("--"),
			jen.Id("iter").Dot("iter").Dot("Next").Call(),
		),
		jen.Line(),
		jen.Return(jen.Id("iter").Dot("iter").Dot("Next").Call()),
	)

	return f
}
