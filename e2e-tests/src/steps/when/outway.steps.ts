import { CustomWorld } from "../../support/custom-world";
import { getOrgByName, Organization } from "../../utils/organizations";
import { When } from "@cucumber/cucumber";
import fetch from "cross-fetch";
import pWaitFor from "p-wait-for";

const isNotBadRequest = async (url: string): Promise<boolean> => {
  const result = await fetch(url);
  return Promise.resolve(result.status !== 400);
};

When(
  "the default Outway of {string} calls the service {string} from {string}",
  async function (
    this: CustomWorld,
    orgNameConsumer: string,
    serviceName: string,
    orgNameProvider: string
  ) {
    serviceName = `${serviceName}-${this.id}`;

    const { scenarioContext } = this;

    const orgConsumer = getOrgByName(orgNameConsumer);
    const orgProvider = getOrgByName(orgNameProvider);

    const url = `${orgConsumer.defaultOutway.address}/${orgProvider.serialNumber}/${serviceName}/get`;

    // wait until the Outway has had the time to update its internal services list
    await pWaitFor.default(async () => await isNotBadRequest(url), {
      interval: 1000,
      timeout: 1000 * 35,
    });

    scenarioContext.organizations[orgNameConsumer].httpResponse = await fetch(
      url
    );
  }
);
