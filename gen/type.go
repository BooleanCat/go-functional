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

	f.Func().Id("fromT").Params(jen.Id("value").Id("T")).Id(typeName).Block(
		jen.Return(jen.Id(typeName).Call(jen.Id("value"))),
	)

	// Generates...
	// func TFold(f func(<T>, <T>) <T>) func(T, T) T {
	//   return func(a, b T) T {
	//     return T(f(<T>(a), <T>(b)))
	//   }
	// }
	f.Add(foldFunc(typeName))
	f.Line()

	// Generates...
	// func Π(f func(<T>, <T>) <T>) func(T, T) T {
	//   return TFold(f)
	// }
	f.Add(foldShortFunc(typeName))

	return f
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
