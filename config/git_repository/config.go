package git_repository

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_git_repository", func(r *config.Resource) {
		r.ShortGroup = "git_repository"
		r.References["project_id"] = config.Reference{
			Type: "github.com/r3drun3/provider-azuredevops/apis/project/v1alpha1.Project",
		}
	})
}
