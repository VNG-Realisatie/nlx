import { Organization, getOrgByName } from "./organizations";
import { env } from "./env";
import { ManagementInway } from "../../../management-ui/src/api/models";
import { CustomWorld } from "../support/custom-world";
import { default as logger } from "../debug";
import pWaitFor from "p-wait-for";
import fetch from "cross-fetch";
import { strict as assert } from "assert";
const debug = logger("e2e-tests:inway");

export const isManagementAPIProxyAddressForDirectoryEqualTo = async (
  org: Organization,
  address: string
): Promise<boolean> => {
  const url = `${env.directoryUrl}/api/directory/organizations/${org.serialNumber}/inway/management-api-proxy-address`;
  const res = await fetch(url);

  assert.equal(res.status >= 400, false);
  const response = await res.json();
  if (response.address === undefined || response.address !== address) {
    return Promise.resolve(false);
  }

  return Promise.resolve(true);
};

export const hasDefaultInwayRunning = async (
  world: CustomWorld,
  orgName: string
) => {
  const org = getOrgByName(orgName);

  const response = await org.apiClients.management?.managementListInways();

  assert.equal(
    response?.inways?.some(
      (inway: ManagementInway) => inway.name === org.defaultInway.name
    ),
    true
  );
};

export const setDefaultInwayAsOrganizationInway = async (
  world: CustomWorld,
  organizationName: string
) => {
  debug(`setting default inway as organization inway for ${organizationName}`);
  const org = getOrgByName(organizationName);
  const response = await org.apiClients.management?.managementUpdateSettings({
    body: {
      organizationInway: org.defaultInway.name,
    },
  });

  if (!response) {
    throw new Error(
      `unable to set organization Inway for ${organizationName}, did you for get to authenticate this organization?`
    );
  }

  // wait until the inway is set as organization inway in the directory
  await pWaitFor.default(
    async () =>
      await isManagementAPIProxyAddressForDirectoryEqualTo(
        org,
        org.defaultInway.managementAPIProxyAddress
      ),
    {
      interval: 200,
      timeout: 1000 * 21,
    }
  );
};
