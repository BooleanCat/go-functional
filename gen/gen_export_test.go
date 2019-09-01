package gen

import "github.com/dave/jennifer/jen"

func (g TypeFileGen) TypeQualifierImpl() *jen.Statement {
	return g.typeQualifier()
}
