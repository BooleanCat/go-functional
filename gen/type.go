package gen

import "github.com/dave/jennifer/jen"

func TypeFileContent(typeName string) *jen.File {
	f := jen.NewFile(packageName(typeName))

	f.Type().Defs(
		jen.Id("T").Id(typeName),
		jen.Id("tSlice").Index().Id(typeName),
		jen.Id("mapFunc").Func().Params(jen.Id(typeName)).Id(typeName),
		jen.Id("mapErrFunc").Func().Params(jen.Id(typeName)).Params(jen.Id(typeName), jen.Error()),
		jen.Id("foldFunc").Func().Params(jen.Id(typeName), jen.Id(typeName)).Id(typeName),
		jen.Id("filterFunc").Func().Params(jen.Id(typeName)).Bool(),
		jen.Id("filterErrFunc").Func().Params(jen.Id(typeName)).Params(jen.Bool(), jen.Error()),
	)

	f.Func().Id("fromT").Params(jen.Id("value").Id("T")).Id(typeName).Block(
		jen.Return(jen.Id(typeName).Call(jen.Id("value"))),
	)

	f.Func().Id("Collect").Params(jen.Id("iter").Id("Iter")).Params(jen.Index().Id(typeName), jen.Error()).Block(
		jen.Return(jen.Id("collect").Call(jen.Id("iter"))),
	)

	f.Func().Params(jen.Id("f").Op("*").Id("Functor")).Id("Collect").Params().Params(jen.Index().Id(typeName), jen.Error()).Block(
		jen.Return(jen.Id("Collect").Call(jen.Id("f").Dot("iter"))),
	)

	f.Func().Id("Collapse").Params(jen.Id("iter").Id("Iter")).Index().Id(typeName).Block(
		jen.Return(jen.Id("collapse").Call(jen.Id("iter"))),
	)

	f.Func().Params(jen.Id("f").Op("*").Id("Functor")).Id("Collapse").Params().Index().Id(typeName).Block(
		jen.Return(jen.Id("collapse").Call(jen.Id("f").Dot("iter"))),
	)

	f.Func().Id("Fold").Params(jen.Id("iter").Id("Iter"), jen.Id("initial").Id(typeName), jen.Id("op").Id("foldFunc")).Id(typeName).Block(
		jen.Return(jen.Id("fromT").Call(jen.Id("fold").Call(jen.Id("iter"), jen.Id("T").Call(jen.Id("initial")), jen.Id("op")))),
	)

	f.Func().Params(jen.Id("f").Op("*").Id("Functor")).Id("Fold").Params(jen.Id("initial").Id(typeName), jen.Id("op").Id("foldFunc")).Id(typeName).Block(
		jen.Return(jen.Id("Fold").Call(jen.Id("f").Dot("iter"), jen.Id("initial"), jen.Id("op"))),
	)

	f.Func().Id("asMapErrFunc").Params(jen.Id("f").Id("mapFunc")).Id("mapErrFunc").Block(
		jen.Return(jen.Func().Params(jen.Id("v").Id(typeName)).Params(jen.Id(typeName), jen.Error()).Block(
			jen.Return(jen.Id("f").Call(jen.Id("v")), jen.Nil()),
		)),
	)

	f.Func().Id("asFilterErrFunc").Params(jen.Id("f").Id("filterFunc")).Id("filterErrFunc").Block(
		jen.Return(jen.Func().Params(jen.Id("v").Id(typeName)).Params(jen.Bool(), jen.Error())).Block(
			jen.Return(jen.Id("f").Call(jen.Id("v")), jen.Nil()),
		),
	)

	return f
}

func packageName(typeName string) string {
	if typeName == "interface{}" {
		return "finterface"
	}
	return "f" + typeName
}
