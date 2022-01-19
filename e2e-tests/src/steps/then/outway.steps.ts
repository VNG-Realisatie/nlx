import { CustomWorld } from "../../support/custom-world";
import { Then } from "@cucumber/cucumber";
import { strict as assert } from "assert";

Then(
  "{string} receives a successful response",
  async function (this: CustomWorld, orgName: string) {
    const body = await this.scenarioContext.organizations[
      orgName
    ].httpResponse?.json();

    assert.equal(
      this.scenarioContext.organizations[orgName].httpResponse?.status,
      200
    );
    assert.equal(body.url, "https://postman-echo.com/get");
  }
);
