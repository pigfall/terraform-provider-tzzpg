package pkg

import(
	"context"

  tfprovider "github.com/hashicorp/terraform-plugin-framework/provider"
  "github.com/hashicorp/terraform-plugin-framework/tfsdk"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/resource"
)

type Provider struct{

}

func NewProvider() tfprovider.Provider{
	return &Provider{}
}

func(*Provider)	GetSchema(context.Context) (tfsdk.Schema, diag.Diagnostics){
	return tfsdk.Schema{},nil
}

func(*Provider)	Configure(context.Context, tfprovider.ConfigureRequest, *tfprovider.ConfigureResponse){
	return
}

func(*Provider)	DataSources(context.Context) []func() datasource.DataSource{
	return []func() datasource.DataSource{
		NewDogDataSource,
	}
}

func(*Provider)	Resources(context.Context) []func() resource.Resource{
	return nil
}
