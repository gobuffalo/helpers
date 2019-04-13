package helpers

import "html/template"

// HTMLer generates HTML source
type HTMLer interface {
	HTML() template.HTML
}

// Helpers.Add("json", toJSONHelper)
// Helpers.Add("jsEscape", template.JSEscapeString)
// Helpers.Add("htmlEscape", htmlEscape)
// Helpers.Add("upcase", strings.ToUpper)
// Helpers.Add("downcase", strings.ToLower)
// Helpers.Add("contentFor", contentForHelper)
// Helpers.Add("contentOf", contentOfHelper)
// Helpers.Add("markdown", MarkdownHelper)
// Helpers.Add("len", lenHelper)
// Helpers.Add("debug", debugHelper)
// Helpers.Add("inspect", inspectHelper)
// Helpers.Add("range", rangeHelper)
// Helpers.Add("between", betweenHelper)
// Helpers.Add("until", untilHelper)
// Helpers.Add("groupBy", groupByHelper)
// Helpers.Add("form", BootstrapFormHelper)
// Helpers.Add("form_for", BootstrapFormForHelper)
// Helpers.Add("truncate", truncateHelper)
// Helpers.Add("env", envy.MustGet)
// Helpers.Add("envOr", envy.Get)
// Helpers.Add("partial", partialHelper)
// Helpers.Add("raw", func(s string) template.HTML {
// 	return template.HTML(s)
// })
// Helpers.AddMany(inflect.Helpers)
