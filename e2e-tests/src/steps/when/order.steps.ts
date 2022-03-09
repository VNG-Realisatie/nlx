import { CustomWorld } from "../../support/custom-world";
import { getOrgByName } from "../../utils/organizations";
import { getOutwayByName } from "../../utils/outway";
import { When } from "@cucumber/cucumber";
import { By } from "selenium-webdriver";
import dayjs from "dayjs";
import localizedFormat from "dayjs/plugin/localizedFormat";
import { strict as assert } from "assert";

dayjs.extend(localizedFormat);

const randomPublicKeyPEM = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArN5xGkM73tJsCpKny59e
5lXNRY+eT0sbWyEGsR1qIPRKmLSiRHl3xMsovn5mo6jN3eeK/Q4wKd6Ae5XGzP63
pTG6U5KVVB74eQxSFfV3UEOrDaJ78X5mBZO+Ku21V2QFr44tvMh5IZDX3RbMB/4K
ad6sapmSF00HWrqTVMkrEsZ98DTb5nwGLh3kISnct4tLyVSpsl9s1rtkSgGUcs1T
IvWxS2D2mOsSL1HRdUNcFQmzchbfG87kXPvicoOISAZDJKDqWp3iuH0gJpQ+XMBf
mcD90I7Z/cRQjWP3P93B3V06cJkd00cEIRcIQqF8N+lE01H88Fi+wePhZRy92NP5
4wIDAQAB
-----END PUBLIC KEY-----`;

When(
  "{string} creates an order with reference {string} for {string} including the service {string} of {string} via Outway {string}",
  async function (
    this: CustomWorld,
    delegatorOrgName: string,
    orderReference: string,
    delegateeOrgName: string,
    serviceName: string,
    orgName: string,
    delegatorOutwayName: string
  ) {
    serviceName = `${serviceName}-${this.id}`;
    orderReference = `${orderReference}-${this.id}`;

    const { driver } = this;

    const delegator = getOrgByName(delegatorOrgName);
    const delegatee = getOrgByName(delegateeOrgName);
    const org = getOrgByName(orgName);

    await driver.get(`${delegator.management.url}/orders`);

    await driver.findElement(By.linkText("Opdracht toevoegen")).click();

    await driver
      .findElement(By.xpath("//input[@name='description']"))
      .sendKeys("description");
    await driver.findElement(By.name("reference")).sendKeys(orderReference);
    await driver
      .findElement(By.name("delegatee"))
      .sendKeys(delegatee.serialNumber);
    await driver
      .findElement(By.name("publicKeyPEM"))
      .sendKeys(randomPublicKeyPEM);

    const today = dayjs();
    const validFrom = today.format("L");
    const validUntil = today.add(1, "month").format("L");

    await driver.findElement(By.name("validFrom")).sendKeys(validFrom);
    await driver.findElement(By.name("validUntil")).sendKeys(validUntil);

    await driver.findElement(By.className("ReactSelect__control")).click();

    const options = await driver.findElements(
      By.className("ReactSelect__option")
    );

    const delegatorOutway = await getOutwayByName(
      delegatorOrgName,
      delegatorOutwayName
    );

    const serviceOptionText = `${serviceName} - ${orgName} (${org.serialNumber}) - via ${delegatorOutway.name} (${delegatorOutway.publicKeyFingerprint})`;

    let theOption;
    for (const o of options) {
      const text = await o.getText();
      if (text === serviceOptionText) {
        theOption = o;
        await o.click();
        break;
      }
    }

    assert.notEqual(
      theOption,
      undefined,
      `option not found '${serviceOptionText}'`
    );

    await driver.findElement(By.xpath("//button[@type='submit']")).click();

    await driver.findElement(
      By.xpath("//button[text()='Overzicht bijwerken']")
    );
  }
);
