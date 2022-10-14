/**
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

import { getOrgByName } from "./organizations";
import { CustomWorld } from "../support/custom-world";
import {
  Configuration,
  DirectoryServiceApi,
  ManagementServiceApi,
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
    await authenticateUsingBasicAuth(
      world,
      orgName,
      org.management.username,
      org.management.password
    );

    const credentialsBuffer = Buffer.from(
      `${org.management.username}:${org.management.password}`,
      "utf-8"
    );
    const credentialsBase64 = credentialsBuffer.toString("base64");

    org.apiClients.management = new ManagementServiceApi(
      new Configuration({
        basePath: org.management.url,
        fetchApi: fetch,
        headers: {
          Authorization: `Basic ${credentialsBase64}`,
        },
      })
    );

    org.apiClients.directory = new DirectoryServiceApi(
      new Configuration({
        basePath: org.management.url,
        fetchApi: fetch,
        headers: {
          Authorization: `Basic ${credentialsBase64}`,
        },
      })
    );
  } else {
    await authenticateUsingOIDC(
      world,
      orgName,
      org.management.username,
      org.management.password,
      true
    );

    const cookie = await driver.manage().getCookie("nlx_management_session");

    org.apiClients.management = new ManagementServiceApi(
      new Configuration({
        basePath: org.management.url,
        fetchApi: fetch,
        headers: {
          cookie: `nlx_management_session=${cookie.value}`,
        },
      })
    );

    org.apiClients.directory = new DirectoryServiceApi(
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

export const authenticateUsingOIDC = async (
  world: CustomWorld,
  orgName: string,
  username: string,
  password: string,
  credentialsAreCorrect: boolean
) => {
  debug(`authenticating '${orgName}' using oidc`);

  const { driver } = world;

  await driver
    .findElement(By.linkText("Inloggen met organisatieaccount"))
    .click();

  await driver.findElement(By.id("login")).sendKeys(username);
  await driver.findElement(By.id("password")).sendKeys(password);
  await driver.findElement(By.id("submit-login")).click();

  if (!credentialsAreCorrect) {
    return;
  }

  await driver
    .findElement(By.css(".theme-btn--success > .dex-btn-text"))
    .click();
};

export const authenticateUsingBasicAuth = async (
  world: CustomWorld,
  orgName: string,
  username: string,
  password: string
) => {
  debug(`authenticating '${orgName}' using basic auth`);

  const { driver } = world;

  await driver.findElement(By.id("email")).sendKeys(username);
  await driver.findElement(By.id("current-password")).sendKeys(password);
  await driver.findElement(By.xpath("//button[@type='submit']")).click();
};

export const logout = async (world: CustomWorld, orgName: string) => {
  const orgIsLoggedIn = world.scenarioContext.organizations[orgName].isLoggedIn;

  if (!orgIsLoggedIn) {
    debug(`organization '${orgName}' is logged out already`);
    return;
  }

  const { driver } = world;

  const org = getOrgByName(orgName);

  await driver.get(org.management.url);

  await driver.findElement(By.xpath("//[@aria-label='Account menu']")).click();
  await driver.findElement(By.xpath("//button[text()='Uitloggen']")).click();

  org.apiClients.management = undefined;
  org.apiClients.directory = undefined;

  world.scenarioContext.organizations[orgName].isLoggedIn = false;

  debug(`logout successful for '${orgName}'`);
};
