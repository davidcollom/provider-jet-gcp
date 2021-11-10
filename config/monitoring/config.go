package monitoring

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_monitoring_group", func(r *config.Resource) {
		// Note(turkenh): We have to override the default kind here,
		// which is "Group", since it conflicts otherwise
		// with: Group redeclared in this block
		r.Kind = "MonitoringGroup"
	})
}
