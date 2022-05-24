import { CustomWorld } from "../../support/custom-world";
import { getOrgByName } from "../../utils/organizations";
import { default as logger } from "../../debug";
import { When } from "@cucumber/cucumber";
import { By } from "selenium-webdriver";
import assert from "assert";
const debug = logger("e2e-tests:inway");

When(
  "{string} unsets its organization inway",
  async function (this: CustomWorld, orgName: string) {
    const { driver } = this;

    const org = getOrgByName(orgName);

    await driver.get(`${org.management.url}/settings/general`);

    await driver.findElement(By.className("ReactSelect__control")).click();

    const options = await driver.findElements(
      By.className("ReactSelect__option")
    );

    let theOption;
    for (const o of options) {
      const text = await o.getText();
      if (text === "Geen") {
        theOption = o;
        await o.click();
        break;
      }
    }

    assert.notEqual(theOption, undefined, `option not found 'Geen'`);

    // Close options
    await driver.findElement(By.className("ReactSelect__control")).click();

    await driver.findElement(By.xpath("//button[@type='submit']")).click();
    await driver.findElement(By.xpath("//button[text()='Opslaan']")).click();

    await driver.findElement(
      By.xpath("//p[text()='De instellingen zijn bijgewerkt']")
    );
  }
);

When(
  "{string} removes its default inway",
  async function (this: CustomWorld, orgName: string) {
    const { driver } = this;

    const org = getOrgByName(orgName);

    await driver.get(
      `${org.management.url}/inways-and-outways/inways/${org.defaultInway.name}`
    );

    await driver
      .findElement(By.xpath("//button[@title='Inway verwijderen']"))
      .click();

    await driver
      .findElement(
        By.xpath(
          "//div[contains(@class, 'modal-content-enter-done')]//div[@role='dialog']//button[text()='Verwijderen']"
        )
      )
      .click();

    await driver.findElement(By.xpath("//p[text()='De inway is verwijderd']"));
  }
);
