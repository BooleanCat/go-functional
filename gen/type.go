package gen

import "github.com/dave/jennifer/jen"

func Type(typeName string) *jen.File {
	f := jen.NewFile("f" + typeName)

	f.Type().Id("T").Id(typeName)
	f.Line()

	f.Func().Id("TFrom").Params(jen.Id("s").Index().Id(typeName)).Index().Id("T").Block(
		jen.Id("slice").Op(":=").Make(jen.Index().Id("T"), jen.Len(jen.Id("s"))),
		jen.For().Id("i").Op(":=").Range().Id("s").Block(
			jen.Id("slice").Index(jen.Id("i")).Op("=").Id("T").Call(jen.Id("s").Index(jen.Id("i"))),
		),
		jen.Return(jen.Id("slice")),
	)

	return f
}
