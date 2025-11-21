package helper

import (
	"terraform-provider-objectscale/internal/clientgen"
	"terraform-provider-objectscale/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func UpdateReplicationGroupState(replicationGroups []clientgen.DataServiceVpoolServiceGetDataServiceVpoolsResponseDataServiceVpoolInner) []models.ReplicationGroupEntity {
	return SliceTransform(replicationGroups, func(v clientgen.DataServiceVpoolServiceGetDataServiceVpoolsResponseDataServiceVpoolInner) models.ReplicationGroupEntity {
		var vdcDetails models.VdcDetails
		if v.Vdc != nil {
			vdcDetails.Id = TfString(v.Vdc.Id)
			if v.Vdc.Link != nil {
				vdcDetails.Link.Rel = TfString(v.Vdc.Link.Rel)
				vdcDetails.Link.Href = TfString(v.Vdc.Link.Href)
			}
		}

		return models.ReplicationGroupEntity{
			Description:          TfString(v.Description),
			Name:                 TfString(v.Name),
			EnableRebalancing:    TfBool(v.EnableRebalancing),
			IsAllowAllNamespaces: TfBool(v.IsAllowAllNamespaces),
			IsFullRep:            TfBool(v.IsFullRep),
			UseReplicationTarget: TfBool(v.UseReplicationTarget),
			Id:                   TfString(v.Id),
			CreationTime:         types.Int64PointerValue(v.CreationTime),
			Inactive:             TfBool(v.Inactive),
			Global:               TfBool(v.Global),
			Remote:               TfBool(v.Remote),
			Internal:             TfBool(v.Internal),
			Vdc:                  vdcDetails,
			VarrayMappings: SliceTransform(v.VarrayMappings, func(vm clientgen.DataServiceVpoolServiceGetDataServiceVpoolsResponseDataServiceVpoolInnerVarrayMappingsInner) models.VarrayMapping {
				return models.VarrayMapping{
					Name:                TfString(vm.Name),
					Value:               TfString(vm.Value),
					IsReplicationTarget: TfBool(vm.IsReplicationTarget),
				}
			}),
		}
	})
}
