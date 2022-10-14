/**
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

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

  assert.equal(res.ok, true);

  const response = await res.json();

  if (response && response.address === address) {
    return Promise.resolve(true);
  }

  return Promise.resolve(false);
};

const isInwayRunning = async (
  org: Organization,
  inwayName: string
): Promise<boolean> => {
  try {
    const response =
      await org.apiClients.management?.managementServiceListInways();

    const isPresent = response?.inways?.some(
      (inway) => inway.name === inwayName
    );

    return Promise.resolve(!!isPresent);
  } catch (e) {
    return Promise.resolve(false);
  }
};

export const hasDefaultInwayRunning = async (
  world: CustomWorld,
  orgName: string
) => {
  const org = getOrgByName(orgName);

  await pWaitFor.default(
    async () => await isInwayRunning(org, org.defaultInway.name),
    {
      interval: 200,
      timeout: 1000 * 11,
    }
  );
};

export const setDefaultInwayAsOrganizationInway = async (
  world: CustomWorld,
  organizationName: string
) => {
  debug(`setting default inway as organization inway for ${organizationName}`);
  const org = getOrgByName(organizationName);
  const response =
    await org.apiClients.management?.managementServiceUpdateSettings({
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
