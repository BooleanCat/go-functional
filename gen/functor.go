package gen

import "github.com/dave/jennifer/jen"

func Functor(t string) *jen.File {
	f := jen.NewFile("f" + t)

	f.Type().Id("Functor").Struct(jen.Id("iter").Id("Iter"))
	f.Line()

	f.Func().Id("New").Params(jen.Id("iter").Id("Iter")).Op("*").Id("Functor").Block(
		jen.Return(jen.Op("&").Id("Functor").Values(jen.Dict{jen.Id("iter"): jen.Id("iter")})),
	)

	f.Type().Id("Lifted").Struct(
		jen.Id("slice").Index().Id(t),
		jen.Id("index").Int(),
	)
	f.Line()

	f.Func().Id("Lift").Params(jen.Id("slice").Index().Id(t)).Op("*").Id("Functor").Block(
		jen.Return(jen.Op("&").Id("Functor").Values(jen.Dict{
			jen.Id("iter"): jen.Op("&").Id("Lifted").Values(jen.Dict{
				jen.Id("slice"): jen.Id("slice"),
			}),
		})),
	)
	f.Line()

	f.Func().Params(jen.Id("f").Op("*").Id("Lifted")).Id("Next").Params().Id("Option").Block(
		jen.If(jen.Id("f").Dot("index").Op(">=").Len(jen.Id("f").Dot("slice"))).Block(
			jen.Return(jen.Id("None").Call()),
		),
		jen.Line(),
		jen.Id("f").Dot("index").Op("++"),
		jen.Return(jen.Id("Some").Call(jen.Id("f").Dot("slice").Index(jen.Id("f").Dot("index").Op("-").Lit(1)))),
	)
	f.Line()

	f.Add(functorMethod("Filter")).Params(jen.Id("filter").Func().Params(jen.Id("value").Id(t)).Bool()).Op("*").Id("Functor").Block(
		jen.Id("f").Dot("iter").Op("=").Id("NewFilter").Call(jen.Id("f").Dot("iter"), jen.Id("filter")),
		jen.Return(jen.Id("f")),
	)
	f.Line()

	f.Add(functorMethod("Exclude")).Params(jen.Id("exclude").Func().Params(jen.Id("value").Id(t)).Bool()).Op("*").Id("Functor").Block(
		jen.Id("f").Dot("iter").Op("=").Id("NewExclude").Call(jen.Id("f").Dot("iter"), jen.Id("exclude")),
		jen.Return(jen.Id("f")),
	)
	f.Line()

	f.Add(functorMethod("Drop")).Params(jen.Id("n").Int()).Op("*").Id("Functor").Block(
		jen.Id("f").Dot("iter").Op("=").Id("NewDrop").Call(jen.Id("f").Dot("iter"), jen.Id("n")),
		jen.Return(jen.Id("f")),
	)
	f.Line()

	f.Add(functorMethod("Take")).Params(jen.Id("n").Int()).Op("*").Id("Functor").Block(
		jen.Id("f").Dot("iter").Op("=").Id("NewTake").Call(jen.Id("f").Dot("iter"), jen.Id("n")),
		jen.Return(jen.Id("f")),
	)
	f.Line()

	f.Add(functorMethod("Map")).Params(jen.Id("op").Func().Params(jen.Id(t)).Id(t)).Op("*").Id("Functor").Block(
		jen.Id("f").Dot("iter").Op("=").Id("NewMap").Call(jen.Id("f").Dot("iter"), jen.Id("op")),
		jen.Return(jen.Id("f")),
	)
	f.Line()

	f.Add(functorMethod("Fold")).Params(jen.Id("initial").Id(t), jen.Id("op").Id("foldOp")).Id(t).Block(
		jen.Return(jen.Id("Fold").Call(jen.Id("f").Dot("iter"), jen.Id("initial"), jen.Id("op"))),
	)
	f.Line()

	f.Add(functorMethod("Collect")).Params().Index().Id(t).Block(
		jen.Return(jen.Id("Collect").Call(jen.Id("f").Dot("iter"))),
	)

	return f
}

func functorMethod(name string) *jen.Statement {
	return jen.Func().Params(jen.Id("f").Op("*").Id("Functor")).Id(name)
}
