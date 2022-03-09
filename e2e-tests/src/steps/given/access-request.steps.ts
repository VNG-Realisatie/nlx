import { CustomWorld } from "../../support/custom-world";
import { getAccessToService } from "../../utils/service";
import { Given } from "@cucumber/cucumber";

Given(
  "the default Outway of {string} has access to {string} of {string}",
  async function (
    this: CustomWorld,
    serviceConsumerOrgName: string,
    serviceName: string,
    serviceProviderOrgName: string
  ) {
    await getAccessToService(
      this,
      serviceConsumerOrgName,
      serviceName,
      serviceProviderOrgName
    );
  }
);
