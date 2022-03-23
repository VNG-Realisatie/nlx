import { CustomWorld } from "../../support/custom-world";
import { checkForAccessibilityIssues } from "../../utils/accessibility";
import { Then } from "@cucumber/cucumber";
import { strict as assert } from "assert";

Then("the page is accessible", async function (this: CustomWorld) {
  const { driver } = this;
  const violations = await checkForAccessibilityIssues(driver);
  assert.deepEqual(violations, []);
});
