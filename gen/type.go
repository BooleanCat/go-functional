package gen

import "github.com/dave/jennifer/jen"

func TypeFileContent(typeName string) *jen.File {
	f := jen.NewFile("f" + typeName)

	f.Type().Id("T").Id(typeName)
	f.Line()

	f.Type().Id("tSlice").Index().Id(typeName)
	f.Line()

	f.Type().Id("mapFunc").Func().Params(jen.Id(typeName)).Id(typeName)
	f.Line()

	f.Type().Id("foldFunc").Func().Params(jen.Id(typeName), jen.Id(typeName)).Id(typeName)
	f.Line()

	f.Func().Id("fromT").Params(jen.Id("value").Id("T")).Id(typeName).Block(
		jen.Return(jen.Id(typeName).Call(jen.Id("value"))),
	)

	f.Func().Id("Collect").Params(jen.Id("iter").Id("Iter")).Index().Id(typeName).Block(
		jen.Return(jen.Id("collect").Call(jen.Id("iter"))),
	)

	f.Func().Params(jen.Id("f").Op("*").Id("Functor")).Id("Collect").Params().Index().Id(typeName).Block(
		jen.Return(jen.Id("Collect").Call(jen.Id("f").Dot("iter"))),
	)

	return f
}
