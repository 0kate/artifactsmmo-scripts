package shared

type ContentType = string
type ContentCode = string

const (
	ContentTypeMonster ContentType = "monster"
	ContentTypeUnknown ContentType = "unknown"
)

type Content struct {
	type_ ContentType
	code  ContentCode
}

func NewContent(type_ ContentType, code ContentCode) *Content {
	return &Content{
		type_: type_,
		code:  code,
	}
}

func (c *Content) Type() ContentType {
	return c.type_
}

func (c *Content) Code() ContentCode {
	return c.code
}
