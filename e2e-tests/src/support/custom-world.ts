import { ManagementApi } from "../../../management-ui/src/api";
import { setWorldConstructor, World, IWorldOptions } from "@cucumber/cucumber";
import * as messages from "@cucumber/messages";
import webdriver from "selenium-webdriver";
import dayjs from "dayjs";
import fs from "fs/promises";
import { Buffer } from "buffer";

export interface CustomWorld extends World {
  id: string;
  testName: string;
  feature?: messages.Pickle;
  debug: boolean;
  driver: webdriver.ThenableWebDriver;
  managementApi: ManagementApi;
  snapshot(): Promise<void>;
}

export class CustomWorld extends World implements CustomWorld {
  constructor(options: IWorldOptions) {
    super(options);
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
