import { CustomWorld } from "../../support/custom-world";
import { getOrgByName } from "../../utils/organizations";
import { default as logger } from "../../debug";
import { getOutwayByName } from "../../utils/outway";
import { When } from "@cucumber/cucumber";
import fetch from "cross-fetch";
import pWaitFor from "p-wait-for";
const debug = logger("e2e-tests:outway");

const isNotBadRequest = async (
  input: RequestInfo,
  init?: RequestInit
): Promise<boolean> => {
  const result = await fetch(input, init);
  return Promise.resolve(result.status !== 400);
};

When(
  "the Outway {string} of {string} calls the service {string} from {string}",
  async function (
    this: CustomWorld,
    outwayName: string,
    orgNameConsumer: string,
    serviceName: string,
    orgNameProvider: string
  ) {
    serviceName = `${serviceName}-${this.id}`;

    const { scenarioContext } = this;

    const orgProvider = getOrgByName(orgNameProvider);

    const outway = await getOutwayByName(orgNameConsumer, outwayName);

    const url = `${outway.selfAddress}/${orgProvider.serialNumber}/${serviceName}/get`;

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

When(
  "the Outway {string} of {string} calls the service {string} of {string} via the order of {string} with reference {string}",
  async function (
    this: CustomWorld,
    outwayName: string,
    orgNameConsumer: string,
    serviceName: string,
    orgNameServiceProvider: string,
    orgNameDelegator: string,
    orderReference: string
  ) {
    serviceName = `${serviceName}-${this.id}`;
    orderReference = `${orderReference}-${this.id}`;

    const { scenarioContext } = this;

    const orgProvider = getOrgByName(orgNameServiceProvider);
    const orgDelegator = getOrgByName(orgNameDelegator);

    const outway = await getOutwayByName(orgNameConsumer, outwayName);

    const url = `${outway.selfAddress}/${orgProvider.serialNumber}/${serviceName}/get`;

    const headers = {
      "X-NLX-Request-Delegator": orgDelegator.serialNumber,
      "X-NLX-Request-Order-Reference": orderReference,
    };

    debug(
      `using order to request a service using the following headers: `,
      headers
    );

    // wait until the Outway has had the time to update its internal services list
    await pWaitFor.default(
      async () =>
        await isNotBadRequest(url, {
          headers,
        }),
      {
        interval: 1000,
        timeout: 1000 * 35,
      }
    );

    scenarioContext.organizations[orgNameConsumer].httpResponse = await fetch(
      url,
      {
        headers,
      }
    );
  }
);
