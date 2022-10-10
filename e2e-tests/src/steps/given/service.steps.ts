/**
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

import { CustomWorld } from "../../support/custom-world";
import { createService } from "../../utils/service";
import { getOrgByName } from "../../utils/organizations";
import { getOutwayByName } from "../../utils/outway";
import { default as logger } from "../../debug";
import pWaitFor from "p-wait-for";
import fetch from "cross-fetch";
import { Given } from "@cucumber/cucumber";
import assert from "assert";
const debug = logger("e2e-tests:service");
Given(
  "{string} has a service named {string}",
  async function (
    this: CustomWorld,
    serviceProviderOrgName: string,
    serviceName: string
  ) {
    await createService(this, serviceName, serviceProviderOrgName);
  }
);

Given(
  "{string} revokes access of Outway {string} from {string} to {string}",
  async function (
    this: CustomWorld,
    orgProviderName: string,
    outwayName: string,
    orgConsumerName: string,
    serviceName: string
  ) {
    serviceName = `${serviceName}-${this.id}`;

    const org = getOrgByName(orgProviderName);
    const orgConsumer = getOrgByName(orgConsumerName);

    const outway = await getOutwayByName(orgConsumerName, outwayName);

    const response =
      await org.apiClients.management?.managementListAccessGrantsForService({
        serviceName: serviceName,
      });

    const accessGrant = response?.accessGrants?.find((accessGrant) => {
      return (
        accessGrant.organization?.serialNumber === orgConsumer.serialNumber &&
        accessGrant.publicKeyFingerprint === outway.publicKeyFingerprint
      );
    });

    assert.notEqual(accessGrant, undefined);

    await org.apiClients.management?.managementRevokeAccessGrant({
      accessGrantId: `${accessGrant?.id}`,
    });

    debug(`revoked access for service ${serviceName}`);

    const url = `${outway.selfAddress}/${org.serialNumber}/${serviceName}/get`;

    await pWaitFor.default(async () => await isNotAllowedAccess(url), {
      interval: 1000,
      timeout: 1000 * 35,
    });
  }
);

const isNotAllowedAccess = async (
  input: RequestInfo,
  init?: RequestInit
): Promise<boolean> => {
  const result = await fetch(input, init);
  const responseText = await result.text();
  const responseContainsInvalidService = responseText.includes(
    "is not allowed access"
  );
  return Promise.resolve(responseContainsInvalidService);
};
