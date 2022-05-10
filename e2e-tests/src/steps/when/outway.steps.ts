import { CustomWorld } from "../../support/custom-world";
import { getOrgByName } from "../../utils/organizations";
import { default as logger } from "../../debug";
import { getOutwayByName } from "../../utils/outway";
import { When } from "@cucumber/cucumber";
import fetch from "cross-fetch";
import pWaitFor from "p-wait-for";
const debug = logger("e2e-tests:outway");

const isServiceKnownInServiceListOfOutway = async (
  input: RequestInfo,
  init?: RequestInit
): Promise<boolean> => {
  const result = await fetch(input, init);
  const responseText = await result.text()
  const responseContainsInvalidService = responseText.includes('nlx-outway: invalid serialNumber/service path: valid services')
  return Promise.resolve(!responseContainsInvalidService)
};

When(
  "the Outway {string} of {string} calls the service {string} from {string} with valid authorization",
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
    await pWaitFor.default(async () => await isServiceKnownInServiceListOfOutway(url), {
      interval: 1000,
      timeout: 1000 * 35,
    });

    const headers =  {"Proxy-Authorization": "Bearer 8bb0cf6eb9b17d0f7d22b456f121257dc1254e1f01665370476383ea776df414"}

    scenarioContext.organizations[orgNameConsumer].httpResponse = await fetch(
      url,{
        headers
      }
    );
  }
);

When(
  "the Outway {string} of {string} calls the service {string} of {string} with valid authorization details via the order of {string} with reference {string}",
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

    const validAuthorizationDetails = {
      "X-Nlx-Authorization": "Bearer 8bb0cf6eb9b17d0f7d22b456f121257dc1254e1f01665370476383ea776df414",
    }

    const headers = {
      "X-Nlx-Request-Delegator": orgDelegator.serialNumber,
      "X-Nlx-Request-Order-Reference": orderReference,
      ...validAuthorizationDetails
    };

    debug(
      `using order to request a service using the following headers: `,
      headers
    );

    // wait until the Outway has had the time to update its internal services list
    await pWaitFor.default(
      async () =>
        await isServiceKnownInServiceListOfOutway(url),
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
