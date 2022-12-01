package pkg

import(
	"time"
	"context"

  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
  "github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

type dogDataSource struct{}

func NewDogDataSource() datasource.DataSource{
	return &dogDataSource{}
}

	// Metadata should return the full name of the data source, such as
	// examplecloud_thing.
func (d *dogDataSource)	Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse){
	resp.TypeName = "tzzpg_dog"
}

	// GetSchema returns the schema for this data source.
func (d *dogDataSource)	GetSchema(context.Context) (tfsdk.Schema, diag.Diagnostics){
	return tfsdk.Schema{},nil

}

	// Read is called when the provider must read data source values in
	// order to update state. Config values should be read from the
	// ReadRequest and new state values set on the ReadResponse.
func(d *dogDataSource)	Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse){
	tflog.Info(ctx,"TODO")
	for{
		time.Sleep(time.Second)
		tflog.Info(ctx,"Creating TODO")
		tflog.Debug(ctx,"Test Debug")
		select{
		case <-ctx.Done():
			panic("context done")
		default:
		}
		tflog.Info(ctx,"ctx not done")
	}
	}
