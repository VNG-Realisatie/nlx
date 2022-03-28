import { CustomWorld } from "../../support/custom-world";
import { getOrgByName } from "../../utils/organizations";
import { Then } from "@cucumber/cucumber";
import { By, until } from "selenium-webdriver";

Then(
  "the authentication page for {string} should display an error",
  async function (this: CustomWorld, orgName: string) {
    const { driver } = this;

    const org = getOrgByName(orgName);

    if (org.management.basicAuth) {
      await driver.findElement(
        By.xpath("//*[contains(text(), 'Ongeldig logingegevens.')]")
      );
    } else {
      await driver.findElement(
        By.xpath("//*[contains(text(), 'Invalid Email Address and password.')]")
      );
    }
  }
);

Then(
  "the Inways or Terms Of Service page of {string} should be visible after authenticating",
  async function (this: CustomWorld, orgName: string) {
    const { driver } = this;

    const inwaysOrTermsOfServicePageRegex = new RegExp(
      /(\/inways-and-outways\/inways|\/terms-of-service)$/
    );
    await driver.wait(until.urlMatches(inwaysOrTermsOfServicePageRegex));
  }
);
