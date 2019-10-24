package azurerm

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	//"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/features"
)

func TestAccAzureRMKubernetesClusterAgentPoolCreateUpdate(t *testing.T) {
	AKSResourceName := "azurerm_kubernetes_cluster.test"
	agentPoolResourceName2 := "azurerm_kubernetes_cluster_agent_pool.linux2"
	agentPoolResourceName3 := "azurerm_kubernetes_cluster_agent_pool.linux3"
	ri := tf.AccRandTimeInt()
	clientId := os.Getenv("ARM_CLIENT_ID")
	clientSecret := os.Getenv("ARM_CLIENT_SECRET")
	config := testAccAzureRMKubernetesClusterAgentPool_basic(ri, clientId, clientSecret, testLocation())
	config2 := testAccAzureRMKubernetesClusterAgentPool_basic2(ri, clientId, clientSecret, testLocation())
	config3 := testAccAzureRMKubernetesClusterAgentPool_basic3(ri, clientId, clientSecret, testLocation())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMKubernetesClusterAgentPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMKubernetesClusterExists(AKSResourceName),
					testCheckAzureRMKubernetesAgentPoolExists(agentPoolResourceName2),
					resource.TestCheckResourceAttr(agentPoolResourceName2, "node_count", "1"),
					resource.TestCheckResourceAttr(AKSResourceName, "role_based_access_control.#", "1"),
					resource.TestCheckResourceAttr(AKSResourceName, "role_based_access_control.0.enabled", "false"),
					resource.TestCheckResourceAttr(AKSResourceName, "role_based_access_control.0.azure_active_directory.#", "0"),
					resource.TestCheckResourceAttrSet(AKSResourceName, "kube_config.0.client_key"),
					resource.TestCheckResourceAttrSet(AKSResourceName, "kube_config.0.client_certificate"),
					resource.TestCheckResourceAttrSet(AKSResourceName, "kube_config.0.cluster_ca_certificate"),
					resource.TestCheckResourceAttrSet(AKSResourceName, "kube_config.0.host"),
					resource.TestCheckResourceAttrSet(AKSResourceName, "kube_config.0.username"),
					resource.TestCheckResourceAttrSet(AKSResourceName, "kube_config.0.password"),
					resource.TestCheckResourceAttr(AKSResourceName, "kube_admin_config.#", "0"),
					resource.TestCheckResourceAttr(AKSResourceName, "kube_admin_config_raw", ""),
					resource.TestCheckResourceAttrSet(AKSResourceName, "agent_pool_profile.0.max_pods"),
					resource.TestCheckResourceAttr(AKSResourceName, "network_profile.0.load_balancer_sku", "Basic"),
				),
			},
			{
				Config: config2,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMKubernetesClusterExists(AKSResourceName),
					testCheckAzureRMKubernetesAgentPoolExists(agentPoolResourceName2),
					resource.TestCheckResourceAttr(agentPoolResourceName2, "node_count", "2"),
					resource.TestCheckResourceAttr(AKSResourceName, "role_based_access_control.#", "1"),
					resource.TestCheckResourceAttr(AKSResourceName, "role_based_access_control.0.enabled", "false"),
					resource.TestCheckResourceAttr(AKSResourceName, "role_based_access_control.0.azure_active_directory.#", "0"),
					resource.TestCheckResourceAttrSet(AKSResourceName, "kube_config.0.client_key"),
					resource.TestCheckResourceAttrSet(AKSResourceName, "kube_config.0.client_certificate"),
					resource.TestCheckResourceAttrSet(AKSResourceName, "kube_config.0.cluster_ca_certificate"),
					resource.TestCheckResourceAttrSet(AKSResourceName, "kube_config.0.host"),
					resource.TestCheckResourceAttrSet(AKSResourceName, "kube_config.0.username"),
					resource.TestCheckResourceAttrSet(AKSResourceName, "kube_config.0.password"),
					resource.TestCheckResourceAttr(AKSResourceName, "kube_admin_config.#", "0"),
					resource.TestCheckResourceAttr(AKSResourceName, "kube_admin_config_raw", ""),
					resource.TestCheckResourceAttrSet(AKSResourceName, "agent_pool_profile.0.max_pods"),
					resource.TestCheckResourceAttr(AKSResourceName, "network_profile.0.load_balancer_sku", "Basic"),
				),
			},
			{
				Config: config3,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMKubernetesClusterExists(AKSResourceName),
					testCheckAzureRMKubernetesAgentPoolExists(agentPoolResourceName2),
					testCheckAzureRMKubernetesAgentPoolExists(agentPoolResourceName3),
					resource.TestCheckResourceAttr(agentPoolResourceName2, "node_count", "2"),
					resource.TestCheckResourceAttr(agentPoolResourceName3, "node_count", "1"),
					resource.TestCheckResourceAttr(AKSResourceName, "role_based_access_control.#", "1"),
					resource.TestCheckResourceAttr(AKSResourceName, "role_based_access_control.0.enabled", "false"),
					resource.TestCheckResourceAttr(AKSResourceName, "role_based_access_control.0.azure_active_directory.#", "0"),
					resource.TestCheckResourceAttrSet(AKSResourceName, "kube_config.0.client_key"),
					resource.TestCheckResourceAttrSet(AKSResourceName, "kube_config.0.client_certificate"),
					resource.TestCheckResourceAttrSet(AKSResourceName, "kube_config.0.cluster_ca_certificate"),
					resource.TestCheckResourceAttrSet(AKSResourceName, "kube_config.0.host"),
					resource.TestCheckResourceAttrSet(AKSResourceName, "kube_config.0.username"),
					resource.TestCheckResourceAttrSet(AKSResourceName, "kube_config.0.password"),
					resource.TestCheckResourceAttr(AKSResourceName, "kube_admin_config.#", "0"),
					resource.TestCheckResourceAttr(AKSResourceName, "kube_admin_config_raw", ""),
					resource.TestCheckResourceAttrSet(AKSResourceName, "agent_pool_profile.0.max_pods"),
					resource.TestCheckResourceAttr(AKSResourceName, "network_profile.0.load_balancer_sku", "Basic"),
				),
			},
			{
				ResourceName:            AKSResourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"service_principal.0.client_secret"},
			},
		},
	})
}

func testCheckAzureRMKubernetesClusterAgentPoolDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*ArmClient).Containers.KubernetesClustersClient

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azurerm_kubernetes_cluster" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		resourceGroup := rs.Primary.Attributes["resource_group_name"]

		ctx := testAccProvider.Meta().(*ArmClient).StopContext
		resp, err := conn.Get(ctx, resourceGroup, name)

		if err != nil {
			return nil
		}

		if resp.StatusCode != http.StatusNotFound {
			return fmt.Errorf("Managed Kubernetes Cluster still exists:\n%#v", resp)
		}
	}

	return nil
}

func testAccAzureRMKubernetesClusterAgentPool_basic(rInt int, clientId string, clientSecret string, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_kubernetes_cluster" "test" {
  name                = "acctestaks%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
  dns_prefix          = "acctestaks%d"

  agent_pool_profile {
    name    = "default"
    count   = "1"
    vm_size = "Standard_DS2_v2"
    type = "VirtualMachineScaleSets"
  }

  service_principal {
    client_id     = "%s"
    client_secret = "%s"
  }
}

resource "azurerm_kubernetes_cluster_agent_pool" "linux2" {
    agent_pool_name = "linux2"
	name = azurerm_kubernetes_cluster.test.name
    resource_group_name = azurerm_resource_group.test.name

    os_type = "Linux"
    vm_size = "Standard_DS2_v2"
    node_count = 1
}
`, rInt, location, rInt, rInt, clientId, clientSecret)
}

func testAccAzureRMKubernetesClusterAgentPool_basic2(rInt int, clientId string, clientSecret string, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_kubernetes_cluster" "test" {
  name                = "acctestaks%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
  dns_prefix          = "acctestaks%d"

  agent_pool_profile {
    name    = "default"
    count   = "1"
    vm_size = "Standard_DS2_v2"
    type = "VirtualMachineScaleSets"
  }

  service_principal {
    client_id     = "%s"
    client_secret = "%s"
  }
}

resource "azurerm_kubernetes_cluster_agent_pool" "linux2" {
    agent_pool_name = "linux2"
	name = azurerm_kubernetes_cluster.test.name
    resource_group_name = azurerm_resource_group.test.name

    agent_pool_type = "VirtualMachineScaleSets"
    os_type = "Linux"
    vm_size = "Standard_DS2_v2"
    node_count = 2
}
`, rInt, location, rInt, rInt, clientId, clientSecret)
}

func testAccAzureRMKubernetesClusterAgentPool_basic3(rInt int, clientId string, clientSecret string, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_kubernetes_cluster" "test" {
  name                = "acctestaks%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
  dns_prefix          = "acctestaks%d"

  agent_pool_profile {
    name    = "default"
    count   = "1"
    vm_size = "Standard_DS2_v2"
    type = "VirtualMachineScaleSets"
  }

  service_principal {
    client_id     = "%s"
    client_secret = "%s"
  }
}

resource "azurerm_kubernetes_cluster_agent_pool" "linux2" {
    agent_pool_name = "linux2"
	name = azurerm_kubernetes_cluster.test.name
    resource_group_name = azurerm_resource_group.test.name

    agent_pool_type = "VirtualMachineScaleSets"
    os_type = "Linux"
    vm_size = "Standard_DS2_v2"
    node_count = 2
}

resource "azurerm_kubernetes_cluster_agent_pool" "linux3" {
    agent_pool_name = "linux3"
	name = azurerm_kubernetes_cluster.test.name
    resource_group_name = azurerm_resource_group.test.name

    agent_pool_type = "VirtualMachineScaleSets"
    os_type = "Linux"
    vm_size = "Standard_DS2_v2"
    node_count = 1
}
`, rInt, location, rInt, rInt, clientId, clientSecret)
}

func testCheckAzureRMKubernetesAgentPoolExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		name := rs.Primary.Attributes["name"]
		resourceGroup, hasResourceGroup := rs.Primary.Attributes["resource_group_name"]
		if !hasResourceGroup {
			return fmt.Errorf("Bad: no resource group found in state for Managed Kubernetes Cluster: %s", name)
		}
		agentPoolName := rs.Primary.Attributes["agent_pool_name"]

		client := testAccProvider.Meta().(*ArmClient).Containers.AgentPoolsClient
		//meta.(*ArmClient).Containers.AgentPoolsClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext

		agent_pool, err := client.Get(ctx, resourceGroup, name, agentPoolName)
		if err != nil {
			return fmt.Errorf("Bad: Get on kubernetesClustersClient: %+v", err)
		}

		if agent_pool.StatusCode == http.StatusNotFound {
			return fmt.Errorf("Bad: Managed Kubernetes Cluster %q (Resource Group: %q) does not exist", name, resourceGroup)
		}

		return nil
	}
}
