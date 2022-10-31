/**
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

import { CustomWorld } from "../../support/custom-world";
import { getOrgByName } from "../../utils/organizations";
import { getOutwayByName } from "../../utils/outway";
import { Then } from "@cucumber/cucumber";
import { strict as assert } from "assert";

Then(
  "the service {string} of {string} is created",
  async function (this: CustomWorld, serviceName: string, orgName: string) {
    serviceName = `${serviceName}-${this.id}`;

    const org = getOrgByName(orgName);

    const response =
      await org.apiClients.management?.managementServiceGetService({
        name: serviceName,
      });

    assert.equal(response?.name, serviceName);
  }
);

Then(
  "the service {string} of {string} is no longer available",
  async function (this: CustomWorld, serviceName: string, orgName: string) {
    serviceName = `${serviceName}-${this.id}`;

    const org = getOrgByName(orgName);

    try {
      await org.apiClients.management?.managementServiceGetService({
        name: serviceName,
      });
      throw new Error(
        "this code should not be triggered, since we expect the service to be removed"
      );
    } catch (error: any) {
      if (error.response.status !== 404) {
        throw new Error(
          `unexpected status code '${error.response.status}' while getting a service, expected 404: ${error}`
        );
      }
    }
  }
);

Then(
  "{string} no longer has an access request to {string} from {string} for the Outway {string}",
  async function (
    this: CustomWorld,
    orgNameConsumer: string,
    serviceName: string,
    orgNameProvider: string,
    outwayName: string
  ) {
    serviceName = `${serviceName}-${this.id}`;

    const orgConsumer = getOrgByName(orgNameConsumer);
    const orgProvider = getOrgByName(orgNameProvider);

    const outway = await getOutwayByName(orgNameConsumer, outwayName);

    const response =
      await orgConsumer.apiClients.directory?.directoryServiceGetOrganizationService(
        {
          organizationSerialNumber: orgProvider.serialNumber,
          serviceName: serviceName,
        }
      );

    const filteredAccessRequests =
      response?.directoryService?.accessStates?.filter((accessState) => {
        return (
          accessState?.accessRequest?.publicKeyFingerprint ==
          outway?.publicKeyFingerprint
        );
      });
    assert.equal(filteredAccessRequests?.length, 1);
    const accessRequest = filteredAccessRequests[0];

    assert.equal(
      accessRequest.accessRequest?.state,
      "ACCESS_REQUEST_STATE_WITHDRAWN"
    );
  }
);
