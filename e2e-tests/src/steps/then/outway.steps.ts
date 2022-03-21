import { CustomWorld } from "../../support/custom-world";
import { Then } from "@cucumber/cucumber";
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

    assert.equal(httpResponse?.status, 500);

    const responseText = await httpResponse?.text();
    const containsRevokedText = responseText.includes("order is revoked");

    if (!containsRevokedText) {
      throw new Error(
        `the response of the HTTP request seems not to be about a revoked order: ${responseText}`
      );
    }

    assert.equal(containsRevokedText, true);
  }
);

Then(
  "{string} receives an unauthorized response",
  async function (this: CustomWorld, orgName: string) {
    const httpResponse = await this.scenarioContext.organizations[orgName]
      .httpResponse;

    assert.equal(httpResponse?.status, 403);

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
