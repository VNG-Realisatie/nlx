/**
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

import { CustomWorld } from "../../support/custom-world";
import { getAccessToService } from "../../utils/service";
import { Given } from "@cucumber/cucumber";

Given(
  "the Outway {string} of {string} has access to {string} of {string}",
  async function (
    this: CustomWorld,
    outwayName: string,
    serviceConsumerOrgName: string,
    serviceName: string,
    serviceProviderOrgName: string
  ) {
    await getAccessToService(
      this,
      serviceConsumerOrgName,
      serviceName,
      serviceProviderOrgName,
      outwayName
    );
  }
);
