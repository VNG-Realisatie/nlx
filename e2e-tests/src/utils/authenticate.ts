import { getOrgByName } from "./organizations";
import { CustomWorld } from "../support/custom-world";
import {
  Configuration,
  DirectoryApi,
  ManagementApi,
} from "../../../management-ui/src/api";
import { default as logger } from "../debug";
import { By } from "selenium-webdriver";
import fetch from "cross-fetch";

const debug = logger("e2e-tests:authentication");

export const authenticate = async (world: CustomWorld, orgName: string) => {
  const orgIsLoggedIn = world.scenarioContext.organizations[orgName].isLoggedIn;

  if (orgIsLoggedIn) {
    debug(`organization '${orgName}' is logged in`);
    return;
  }

  const { driver } = world;

  const org = getOrgByName(orgName);

  await driver.get(org.management.url);

  if (org.management.basicAuth) {
    debug(`authenticating '${orgName}' using basic auth`);
    await driver.findElement(By.id("email")).sendKeys(org.management.username);
    await driver
      .findElement(By.id("current-password"))
      .sendKeys(org.management.password);
    await driver.findElement(By.xpath("//button[@type='submit']")).click();

    const credentialsBuffer = Buffer.from(
      `${org.management.username}:${org.management.password}`,
      "utf-8"
    );
    const credentialsBase64 = credentialsBuffer.toString("base64");

    org.apiClients.management = new ManagementApi(
      new Configuration({
        basePath: org.management.url,
        fetchApi: fetch,
        headers: {
          Authorization: `Basic ${credentialsBase64}`,
        },
      })
    );

    org.apiClients.directory = new DirectoryApi(
      new Configuration({
        basePath: org.management.url,
        fetchApi: fetch,
        headers: {
          Authorization: `Basic ${credentialsBase64}`,
        },
      })
    );
  } else {
    debug(`authenticating '${orgName}' using oidc`);

    await driver
      .findElement(By.linkText("Inloggen met organisatieaccount"))
      .click();

    await driver.findElement(By.id("login")).sendKeys(org.management.username);
    await driver
      .findElement(By.id("password"))
      .sendKeys(org.management.password);
    await driver.findElement(By.id("submit-login")).click();

    await driver
      .findElement(By.css(".theme-btn--success > .dex-btn-text"))
      .click();

    const cookie = await driver.manage().getCookie("nlx_management_session");

    org.apiClients.management = new ManagementApi(
      new Configuration({
        basePath: org.management.url,
        fetchApi: fetch,
        headers: {
          cookie: `nlx_management_session=${cookie.value}`,
        },
      })
    );

    org.apiClients.directory = new DirectoryApi(
      new Configuration({
        basePath: org.management.url,
        fetchApi: fetch,
        headers: {
          cookie: `nlx_management_session=${cookie.value}`,
        },
      })
    );
  }

  world.scenarioContext.organizations[orgName].isLoggedIn = true;
  debug(`authentication successful for '${orgName}'`);
};
