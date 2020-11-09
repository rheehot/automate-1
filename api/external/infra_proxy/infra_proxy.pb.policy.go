// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: external/infra_proxy/infra_proxy.proto

package infra_proxy

import (
	policy "github.com/chef/automate/api/external/iam/v2/policy"
	request "github.com/chef/automate/api/external/infra_proxy/request"
)

func init() {
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetServers", "infra:infraServers", "infra:infraServers:list", "GET", "/api/v0/infra/servers", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetServer", "infra:infraServers:{id}", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetServer); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/CreateServer", "infra:infraServers", "infra:infraServers:create", "POST", "/api/v0/infra/servers", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateServer); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "fqdn":
					return m.Fqdn
				case "ip_address":
					return m.IpAddress
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/UpdateServer", "infra:infraServers:{id}", "infra:infraServers:update", "PUT", "/api/v0/infra/servers/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateServer); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "fqdn":
					return m.Fqdn
				case "ip_address":
					return m.IpAddress
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/DeleteServer", "infra:infraServers:{id}", "infra:infraServers:delete", "DELETE", "/api/v0/infra/servers/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DeleteServer); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetOrgs", "infra:infraServers:{server_id}:orgs", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetOrgs); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetOrg", "infra:infraServers:{server_id}:orgs:{id}", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetOrg); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/CreateOrg", "infra:infraServers:{server_id}:orgs", "infra:infraServers:update", "POST", "/api/v0/infra/servers/{server_id}/orgs", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateOrg); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "admin_user":
					return m.AdminUser
				case "admin_key":
					return m.AdminKey
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/UpdateOrg", "infra:infraServers:{server_id}:orgs:{id}", "infra:infraServers:update", "PUT", "/api/v0/infra/servers/{server_id}/orgs/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateOrg); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "admin_user":
					return m.AdminUser
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/DeleteOrg", "infra:infraServers:{server_id}:orgs:{id}", "infra:infraServers:update", "DELETE", "/api/v0/infra/servers/{server_id}/orgs/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DeleteOrg); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/ResetOrgAdminKey", "infra:infraServers:{server_id}:orgs:{id}", "infra:infraServers:update", "PUT", "/api/v0/infra/servers/{server_id}/orgs/{id}/reset-key", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.ResetOrgAdminKey); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "server_id":
					return m.ServerId
				case "admin_key":
					return m.AdminKey
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetCookbooks", "infra:infraServers:{server_id}:orgs:{org_id}:cookbooks", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/cookbooks", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Cookbooks); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetCookbookVersions", "infra:infraServers:{server_id}:orgs:{org_id}:cookbooks", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/cookbooks/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CookbookVersions); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetCookbook", "infra:infraServers:{server_id}:orgs:{org_id}:cookbooks", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/cookbooks/{name}/{version}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Cookbook); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "version":
					return m.Version
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetCookbookFileContent", "infra:infraServers:{server_id}:orgs:{org_id}:cookbooks", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/cookbooks/{name}/{version}/file-content", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CookbookFileContent); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "version":
					return m.Version
				case "url":
					return m.Url
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetRoles", "infra:infraServers:{server_id}:orgs:{org_id}:roles", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/roles", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Roles); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetRole", "infra:infraServers:{server_id}:orgs:{org_id}:roles", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/roles/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Role); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/CreateRole", "infra:infraServers:{server_id}:orgs:{org_id}:roles", "infra:infraServers:update", "POST", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/roles", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateRole); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "description":
					return m.Description
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/DeleteRole", "infra:infraServers:{server_id}:orgs:{org_id}:roles", "infra:infraServers:update", "DELETE", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/roles/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Role); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/UpdateRole", "infra:infraServers:{server_id}:orgs:{org_id}:roles", "infra:infraServers:update", "PUT", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/roles/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateRole); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "description":
					return m.Description
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetClients", "infra:infraServers:{server_id}:orgs:{org_id}:clients", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/clients", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Clients); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetClient", "infra:servers:{server_id}:orgs:{org_id}:clients", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/clients/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Client); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetDataBags", "infra:infraServers:{server_id}:orgs:{org_id}:data_bags", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/data_bags", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DataBags); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetDataBagItem", "infra:infraServers:{server_id}:orgs:{org_id}:data_bags", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/data_bags/{name}/{item}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DataBag); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "item":
					return m.Item
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/CreateDataBag", "infra:infraServers:{server_id}:orgs:{org_id}:data_bags", "infra:infraServers:update", "POST", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/data_bags", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateDataBag); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/CreateDataBagItem", "infra:infraServers:{server_id}:orgs:{org_id}:data_bags", "infra:infraServers:update", "POST", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/data_bags/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateDataBagItem); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/DeleteDataBag", "infra:infraServers:{server_id}:orgs:{org_id}:data_bags", "infra:infraServers:update", "DELETE", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/data_bags/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DataBag); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "item":
					return m.Item
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/UpdateDataBagItem", "infra:infraServers:{server_id}:orgs:{org_id}:data_bags", "infra:infraServers:update", "PUT", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/data_bags/{name}/{item_id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateDataBagItem); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "item_id":
					return m.ItemId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetEnvironments", "infra:infraServers:{server_id}:orgs:{org_id}:environments", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/environments", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Environments); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetEnvironment", "infra:infraServers:{server_id}:orgs:{org_id}:environments", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/environments/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Environment); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/CreateEnvironment", "infra:infraServers:{server_id}:orgs:{org_id}:environments", "infra:infraServers:update", "POST", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/environments", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateEnvironment); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "description":
					return m.Description
				case "json_class":
					return m.JsonClass
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/DeleteEnvironment", "infra:infraServers:{server_id}:orgs:{org_id}:environments", "infra:infraServers:update", "DELETE", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/environments/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Environment); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/UpdateEnvironment", "infra:infraServers:{server_id}:orgs:{org_id}:environments", "infra:infraServers:update", "PUT", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/environments/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateEnvironment); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "description":
					return m.Description
				case "json_class":
					return m.JsonClass
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetAffectedNodes", "infra:infraServers:{server_id}:orgs:{org_id}:nodes", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/affected-nodes/{chef_type}/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.AffectedNodes); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "chef_type":
					return m.ChefType
				case "name":
					return m.Name
				case "version":
					return m.Version
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/DeleteNode", "infra:infraServers:{server_id}:orgs:{org_id}:nodes", "infra:infraServers:update", "DELETE", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/nodes/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DeleteNode); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/UpdateNode", "infra:infraServers:{server_id}:orgs:{org_id}:nodes", "infra:infraServers:update", "PUT", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/nodes/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateNode); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "environment":
					return m.Environment
				case "policy_name":
					return m.PolicyName
				case "policy_group":
					return m.PolicyGroup
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetPolicyfiles", "infra:infraServers:{server_id}:orgs:{org_id}:policyfiles", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/policyfiles", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Policyfiles); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetPolicyfile", "infra:infraServers:{server_id}:orgs:{org_id}:policyfiles", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/policyfiles/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Policyfile); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "revision_id":
					return m.RevisionId
				default:
					return ""
				}
			})
		}
		return ""
	})
}
