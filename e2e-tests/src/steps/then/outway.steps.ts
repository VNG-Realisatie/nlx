/**
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

import { CustomWorld } from "../../support/custom-world";
import { getOrgByName } from "../../utils/organizations";
import { getOutwayByName } from "../../utils/outway";
import { Then } from "@cucumber/cucumber";
import pWaitFor from "p-wait-for";
import fetch from "cross-fetch";
import { strict as assert } from "assert";

interface PostmanEchoGetResponse {
  url: string;
}

Then(
  "{string} receives a successful response",
  async function (this: CustomWorld, orgName: string) {
    const httpResponse = await this.scenarioContext.organizations[orgName]
      .httpResponse;

    if (httpResponse?.status === 200) {
      const json = (await httpResponse?.json()) as PostmanEchoGetResponse;

      assert.equal(httpResponse?.status, 200);
      assert.equal(json.url, "https://postman-echo.com/get");
      return;
    }

    const responseText = await httpResponse?.text();
    throw new Error(
      `unsuccessful response (status ${httpResponse?.status}) from ${httpResponse?.url}: ${responseText}`
    );
  }
);

Then(
  "{string} receives an order revoked response",
  async function (this: CustomWorld, orgName: string) {
    const httpResponse = await this.scenarioContext.organizations[orgName]
      .httpResponse;

    assert.equal(httpResponse?.status, 540);

    const responseText = await httpResponse?.text();
    const containsRevokedText = responseText.includes("order is revoked");

    if (!containsRevokedText) {
      throw new Error(
        `the response of the HTTP request seems not to be about an revoked order: ${responseText}`
      );
    }

    assert.equal(containsRevokedText, true);
  }
);

Then(
  "{string} receives an order expired response",
  async function (this: CustomWorld, orgName: string) {
    const httpResponse = await this.scenarioContext.organizations[orgName]
      .httpResponse;

    assert.equal(httpResponse?.status, 540);

    const responseText = await httpResponse?.text();
    const containsRevokedText = responseText.includes("order has expired");

    if (!containsRevokedText) {
      throw new Error(
        `the response of the HTTP request seems not to be about an expired order: ${responseText}`
      );
    }

    assert.equal(containsRevokedText, true);
  }
);

Then(
  "{string} receives a delegator does not have access response",
  async function (this: CustomWorld, orgName: string) {
    const httpResponse = await this.scenarioContext.organizations[orgName]
      .httpResponse;

    assert.equal(httpResponse?.status, 540);

    const responseText = await httpResponse?.text();
    const containsDelegatorNoAccessText = responseText.includes(
      "DELEGATOR_DOES_NOT_HAVE_ACCESS_TO_SERVICE"
    );

    if (!containsDelegatorNoAccessText) {
      throw new Error(
        `the response of the HTTP request seems not to be about a delegator without access: ${responseText}`
      );
    }

    assert.equal(containsDelegatorNoAccessText, true);
  }
);

Then(
  "{string} receives an unauthorized response",
  async function (this: CustomWorld, orgName: string) {
    const httpResponse = await this.scenarioContext.organizations[orgName]
      .httpResponse;

    assert.equal(httpResponse?.status, 540);

    const responseText = await httpResponse?.text();
    const containsRevokedText = responseText.includes("ACCESS_DENIED");

    if (!containsRevokedText) {
      throw new Error(
        `the response of the HTTP request seems not to be about not having permission: ${responseText}`
      );
    }

    assert.equal(containsRevokedText, true);
  }
);

const isUnauthorizedRequest = async (
  input: RequestInfo,
  init?: RequestInit
): Promise<boolean> => {
  const result = await fetch(input, init);
  const responseText = await result?.text();

  if (result.status === 540 && responseText.includes("ACCESS_DENIED")) {
    return Promise.resolve(true);
  }

  return Promise.resolve(false);
};

Then(
  "the Outway {string} of {string} no longer has access to the service {string} from {string}",
  async function (
    this: CustomWorld,
    outwayName: string,
    orgNameConsumer: string,
    serviceName: string,
    orgNameProvider: string
  ) {
    const orgProvider = getOrgByName(orgNameProvider);

    const outway = await getOutwayByName(orgNameConsumer, outwayName);

    serviceName = `${serviceName}-${this.id}`;

    const url = `${outway.selfAddress}/${orgProvider.serialNumber}/${serviceName}/get`;

    // wait until the Outway has had the time to update its internal services list
    try {
      await pWaitFor.default(async () => await isUnauthorizedRequest(url), {
        interval: 1000,
        timeout: 1000 * 35,
      });
    } catch (e) {
      const result = await fetch(url);
      const responseText = await result?.text();

      throw new Error(
        `the response of the HTTP request seems not to be about not having permission: ${responseText}`
      );
    }
  }
);

Then(
  "the outway {string} of {string} is removed",
  async function (this: CustomWorld, outwayName: string, orgName: string) {
    const org = getOrgByName(orgName);

    try {
      const res =
        await org.apiClients.management?.managementServiceListOutways();

      res?.outways?.forEach((o) => {
        if (o.name == outwayName) {
          throw new Error(
            "this code should not be triggered, since we expect the outway to be removed"
          );
        }
      });
    } catch (error: any) {
      if (!error.response) {
        throw error;
      }

      throw new Error(
        `unexpected status code '${error.response.status}' while getting outways: ${error}`
      );
    }
  }
);
