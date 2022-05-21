package color

import "strings"

var (
	resetColor = "\033[0m"
)

type ColorBuilder struct {
	fgColor          string
	bgColor          string
	text             string
	icon             string
	separator        string
	separatorFgColor string
	separatorBgColor string
}

func (builder *ColorBuilder) SetFgColor(color string) *ColorBuilder {
	builder.fgColor = color
	return builder
}

func (builder *ColorBuilder) SetBgColor(color string) *ColorBuilder {
	builder.bgColor = color
	return builder
}

func (builder *ColorBuilder) SetText(text string) *ColorBuilder {
	builder.text = text
	return builder
}

func (builder *ColorBuilder) SetIcon(icon string) *ColorBuilder {
	builder.icon = icon
	return builder
}

func (builder *ColorBuilder) SetSeparator(separator string, fgColor string, bgColor string) *ColorBuilder {
	builder.separator = separator
	builder.separatorFgColor = fgColor
	builder.separatorBgColor = bgColor
	return builder
}

func (builder *ColorBuilder) Build() string {
	c := ""

	if len(builder.bgColor) > 0 {
		c += formatBgColor(builder.bgColor)
	}

	if len(builder.fgColor) > 0 {
		c += formatFgColor(builder.fgColor)
	}

	if len(builder.icon) > 0 {
		c += " " + builder.icon + " "
	}

	if len(builder.text) > 0 {
		c += builder.text
	}

	if len(builder.fgColor) > 0 {
		c += resetColor
	}

	if len(builder.bgColor) > 0 {
		c += resetColor
	}

	if len(builder.separator) > 0 {
		if len(builder.separatorBgColor) > 0 {
			c += formatBgColor(builder.separatorBgColor)
		}

		if len(builder.separatorFgColor) > 0 {
			c += formatFgColorBold(builder.separatorFgColor)
		}

		c += builder.separator

		if len(builder.separatorFgColor) > 0 {
			c += resetColor
		}

		if len(builder.separatorBgColor) > 0 {
			c += resetColor
		}
	}

	return c
}

func formatFgColor(color string) string {
	return "\033[38;2;" + formatRgb(color) + "m"
}

func formatFgColorBold(color string) string {
	return "\033[1;38;2;" + formatRgb(color) + "m"
}

func formatBgColor(color string) string {
	return "\033[48;2;" + formatRgb(color) + "m"
}

func formatRgb(color string) string {
	return strings.Replace(color, ",", ";", -1)
}
