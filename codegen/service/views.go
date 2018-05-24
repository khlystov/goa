package service

import (
	"path/filepath"

	"goa.design/goa/codegen"
	"goa.design/goa/design"
)

// ViewsFile returns the views file for the given service containing types
// to render result types with more than one view appropriately.
func ViewsFile(genpkg string, service *design.ServiceExpr) *codegen.File {
	svc := Services.Get(service.Name)
	if len(svc.ProjectedTypes) == 0 {
		return nil
	}
	path := filepath.Join(codegen.Gendir, codegen.SnakeCase(service.Name), "views", "view.go")
	var (
		sections []*codegen.SectionTemplate
	)
	{
		header := codegen.Header(service.Name+" views", "views",
			[]*codegen.ImportSpec{
				{Path: "goa.design/goa", Name: "goa"},
				{Path: "unicode/utf8"},
			})
		sections = []*codegen.SectionTemplate{header}

		// type definitions
		for _, t := range svc.ProjectedTypes {
			sections = append(sections, &codegen.SectionTemplate{
				Name:   "projected-type",
				Source: userTypeT,
				Data:   t.UserTypeData,
			})
		}

		// validations
		for _, t := range svc.ProjectedTypes {
			if t.Validate != nil {
				sections = append(sections, &codegen.SectionTemplate{
					Name:   "validate-type",
					Source: validateT,
					Data:   t.Validate,
				})
			}
			if t.Views != nil {
				for _, v := range t.Views {
					if v.Validate != nil {
						sections = append(sections, &codegen.SectionTemplate{
							Name:   "validate-type",
							Source: validateT,
							Data:   v.Validate,
						})
					}
				}
			}
		}
	}

	return &codegen.File{Path: path, SectionTemplates: sections}
}

// input: ValidateData
const validateT = `{{ comment .Description }}
func (result {{ .Ref }}) {{ .Name }}() (err error) {
	{{ .Validate }}
  return
}
`
