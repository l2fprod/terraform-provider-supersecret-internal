package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = &SuperSecretProvider{}

type SuperSecretProvider struct{}

func New() func() provider.Provider {
	return func() provider.Provider {
		return &SuperSecretProvider{}
	}
}

func (p *SuperSecretProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "supersecret"
}

func (p *SuperSecretProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (p *SuperSecretProvider) Configure(_ context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {
}

func (p *SuperSecretProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewHelloResource,
	}
}

func (p *SuperSecretProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}
