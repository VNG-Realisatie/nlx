import { CustomWorld } from "../../support/custom-world";
import { createService } from "../../utils/service";
import { Given } from "@cucumber/cucumber";

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

