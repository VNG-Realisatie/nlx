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
      "nlx-inway: no access. delegator does not have access to the service for the public key in the claim"
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
    const containsRevokedText = responseText.includes(
      "nlx-inway: permission denied"
    );

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
  return Promise.resolve(result.status === 540);
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
    await pWaitFor.default(async () => await isUnauthorizedRequest(url), {
      interval: 1000,
      timeout: 1000 * 35,
    });

    const httpResponse = await fetch(url);
    assert.equal(httpResponse?.status, 540);

    const responseText = await httpResponse?.text();
    const containsRevokedText = responseText.includes(
      "nlx-inway: permission denied"
    );

    if (!containsRevokedText) {
      throw new Error(
        `the response of the HTTP request seems not to be about not having permission: ${responseText}`
      );
    }

    assert.equal(containsRevokedText, true);
  }
);
