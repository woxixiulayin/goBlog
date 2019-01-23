package template

import (
    "html/template"
    "io"
    "github.com/labstack/echo"
)

// echo需要自定义的html/template render
type TemplateRender struct {
    templates *template.Template
}

// 返回一个模板文档
func (t *TemplateRender) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    if viewContext, isMap := data.(map[string]interface{}); isMap {
        viewContext["reverse"] = c.Echo().Reverse
    }

    return t.templates.ExecuteTemplate(w, name, data)
}

func Renderer() *TemplateRender {
    renderer := &TemplateRender {
            templates: template.Must(template.ParseGlob("template/html/*.html")),
        }
    return renderer
}




