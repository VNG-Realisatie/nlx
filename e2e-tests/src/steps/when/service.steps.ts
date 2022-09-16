/**
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

import { CustomWorld } from "../../support/custom-world";
import { getOrgByName } from "../../utils/organizations";
import { When } from "@cucumber/cucumber";
import { By } from "selenium-webdriver";
import assert from "assert";

When(
  "{string} create a service named {string} and exposed via the default Inway",
  async function (this: CustomWorld, orgName: string, serviceName: string) {
    serviceName = `${serviceName}-${this.id}`;

    const { driver } = this;

    const org = getOrgByName(orgName);

    await driver.get(`${org.management.url}/services`);

    await driver.findElement(By.linkText("Service toevoegen")).click();
    await driver.findElement(By.name("name")).sendKeys(serviceName);
    await driver
      .findElement(By.name("endpointURL"))
      .sendKeys("http://example.com");
    await driver
      .findElement(By.xpath(`//input[@value = '${org.defaultInway.name}']`))
      .click();
    await driver.findElement(By.xpath("//button[@type='submit']")).click();
    await driver.findElement(By.linkText("Service toevoegen"));

    org.createdItems[this.id].services.push(serviceName);
  }
);

When(
  "{string} removes the service {string}",
  async function (this: CustomWorld, orgName: string, serviceName: string) {
    serviceName = `${serviceName}-${this.id}`;

    const { driver } = this;

    const org = getOrgByName(orgName);

    await driver.get(`${org.management.url}/services/${serviceName}`);

    await driver
      .findElement(By.xpath("//button[@title='Service verwijderen']"))
      .click();

    // confirmation model
    await driver
      .findElement(
        By.xpath("//div[@role='dialog']//button[text()='Verwijderen']")
      )
      .click();

    const confirmationToast = await driver.findElement(
      By.xpath("//p[text()='De service is verwijderd']")
    );
    assert.notEqual(confirmationToast, undefined);
  }
);

When(
  "{string} revokes access of {string} to {string}",
  async function (
    this: CustomWorld,
    orgProviderName: string,
    orgConsumerName: string,
    serviceName: string
  ) {
    serviceName = `${serviceName}-${this.id}`;

    const { driver } = this;

    const org = getOrgByName(orgProviderName);

    await driver.get(`${org.management.url}/services/${serviceName}`);
    await driver
      .findElement(By.xpath("//*[text()='Organisaties met toegang']"))
      .click();

    await driver.findElement(By.xpath("//button[text()='Intrekken']")).click();

    const buttonRevoke = await driver.findElement(
      By.xpath(
        "//div[contains(@class, 'modal-content-enter-done')]//div[@role='dialog']//button[text()='Intrekken']"
      )
    );
    await buttonRevoke.click();

    const element = await driver.findElement(
      By.xpath("//p[text()='Toegang ingetrokken']")
    );
    assert.notEqual(element, undefined);
  }
);
