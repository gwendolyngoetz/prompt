package prompt

import (
	"gwendolyngoetz/prompt/pkg/color"
)

type Part struct {
	FgColor   string
	BgColor   string
	Text      string
	Icon      string
	Separator string
	nextPart  *Part
}

type PromptBuilder struct {
	parts []Part
}

func (builder *PromptBuilder) AddPart(part Part) *PromptBuilder {
	if len(builder.parts) > 0 {
		builder.parts[len(builder.parts)-1].nextPart = &part
	}

	builder.parts = append(builder.parts, part)
	return builder
}

func (builder *PromptBuilder) Build() string {
	result := ""

	for _, part := range builder.parts {
		cb := color.ColorBuilder{}

		cb.
			SetFgColor(part.FgColor).
			SetBgColor(part.BgColor).
			SetIcon(part.Icon).
			SetText(part.Text)

		if part.nextPart != nil {
			cb.SetSeparator(part.Separator, part.BgColor, part.nextPart.BgColor)
		} else {
			cb.SetSeparator(part.Separator, part.BgColor, "")
		}

		result += cb.Build()
	}

	return result
}
