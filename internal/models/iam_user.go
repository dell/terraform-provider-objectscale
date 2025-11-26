package models

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type IAMUserDatasourceModel struct {
	ID        types.String `tfsdk:"id"`
	Namespace types.String `tfsdk:"namespace"`
	Username  types.String `tfsdk:"username"`
	Groupname types.String `tfsdk:"groupname"`
	Users     []IAMUser    `tfsdk:"users"`
}

type IAMUser struct {
	ID                  types.String       `tfsdk:"id"`
	UserName            types.String       `tfsdk:"username"`
	Arn                 types.String       `tfsdk:"arn"`
	Path                types.String       `tfsdk:"path"`
	CreateDate          types.String       `tfsdk:"create_date"`
	PermissionsBoundary types.String       `tfsdk:"permissions_boundary"` // can be empty
	Tags                []IAMUserTag       `tfsdk:"tags"`
	AccessKeys          []IAMUserAccessKey `tfsdk:"access_keys"`
}

type IAMUserTag struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

type IAMUserAccessKey struct {
	AccessKeyId types.String `tfsdk:"access_key_id"`
	CreateDate  types.String `tfsdk:"create_date"`
	Status      types.String `tfsdk:"status"`
}
