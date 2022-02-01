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
