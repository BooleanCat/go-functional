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
	f.Line()

	f.Add(lambdaFunc(typeName))
	f.Line()

	f.Add(lambdaShortFunc(typeName))
	f.Line()

	f.Add(foldFunc(typeName))
	f.Line()

	f.Add(foldShortFunc(typeName))

	return f
}

func lambdaFunc(typeName string) *jen.Statement {
	return jen.Func().Id("Lambda").Params(jen.Id("f").Add(typedLambdaFunc(typeName))).Add(genericLambdaFunc(typeName)).Block(
		jen.Return(jen.Func().Params(jen.Id("a").Id("T")).Id("T").Block(
			jen.Return(jen.Id("T").Call(jen.Id("f").Call(jen.Id(typeName).Call(jen.Id("a"))))),
		)),
	)
}

func lambdaShortFunc(typeName string) *jen.Statement {
	return jen.Func().Id("Λ").Params(jen.Id("f").Add(typedLambdaFunc(typeName))).Add(genericLambdaFunc(typeName)).Block(
		jen.Return(jen.Id("Lambda").Call(jen.Id("f"))),
	)
}

func typedLambdaFunc(typeName string) *jen.Statement {
	return jen.Func().Params(jen.Id(typeName)).Id(typeName)
}

func genericLambdaFunc(typeName string) *jen.Statement {
	return jen.Func().Params(jen.Id("T")).Id("T")
}

func foldFunc(typeName string) *jen.Statement {
	return jen.Func().Id("TFold").Params(jen.Id("f").Add(typedFoldFunc(typeName))).Add(genericFoldFunc(typeName)).Block(
		jen.Return(jen.Func().Params(jen.Id("a").Id("T"), jen.Id("b").Id("T")).Id("T").Block(
			jen.Return(jen.Id("T").Call(jen.Id("f").Call(jen.Id(typeName).Call(jen.Id("a")), jen.Id(typeName).Call(jen.Id("b"))))),
		)),
	)
}

func foldShortFunc(typeName string) *jen.Statement {
	return jen.Func().Id("Π").Params(jen.Id("f").Add(typedFoldFunc(typeName))).Add(genericFoldFunc(typeName)).Block(
		jen.Return(jen.Id("TFold").Call(jen.Id("f"))),
	)
}

func typedFoldFunc(typeName string) *jen.Statement {
	return jen.Func().Params(jen.Id(typeName), jen.Id(typeName)).Id(typeName)
}

func genericFoldFunc(typeName string) *jen.Statement {
	return jen.Func().Params(jen.Id("T"), jen.Id("T")).Id("T")
}
