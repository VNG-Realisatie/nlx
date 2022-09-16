/**
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

import { organizations } from "../utils/organizations";
import { setWorldConstructor, World, IWorldOptions } from "@cucumber/cucumber";
import * as messages from "@cucumber/messages";
import webdriver from "selenium-webdriver";
import dayjs from "dayjs";
import fs from "fs/promises";
import { Buffer } from "buffer";

interface OrganizationContext {
  httpResponse: Response | undefined;
  isLoggedIn: boolean;
}

interface OrganizationsContext {
  [key: string]: OrganizationContext;
}

interface ScenarioContext {
  organizations: OrganizationsContext;
}

export interface CustomWorld extends World {
  id: string;
  testName: string;
  feature?: messages.Pickle;
  debug: boolean;
  driver: webdriver.ThenableWebDriver;
  snapshot(): Promise<void>;
  scenarioContext: ScenarioContext;
}

export class CustomWorld extends World implements CustomWorld {
  constructor(options: IWorldOptions) {
    super(options);

    this.scenarioContext = {
      organizations: {},
    };

    Object.keys(organizations).forEach((organizationName) => {
      this.scenarioContext.organizations[organizationName] = {
        httpResponse: undefined,
        isLoggedIn: false,
      };
    });
  }
  debug = false;

  async snapshot(): Promise<void> {
    const image = Buffer.from(await this.driver.takeScreenshot(), "base64");
    return fs.writeFile(
      `reports/screenshots/${this.testName}-${dayjs().format(
        "YYYY-MM-DD-HH:mm:ss"
      )}.png`,
      image,
      "base64"
    );
  }
}

setWorldConstructor(CustomWorld);
