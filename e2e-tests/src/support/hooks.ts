/**
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

import { CustomWorld } from "./custom-world";
import { logout } from "../utils/authenticate";
import { organizations } from "../utils/organizations";
import { default as logger } from "../debug";
import { Before, After, setDefaultTimeout } from "@cucumber/cucumber";
import { ITestCaseHookParameter } from "@cucumber/cucumber/lib/support_code_library_builder/types";
import webdriver from "selenium-webdriver";
import { ulid } from "ulid";
import dayjs from "dayjs";
const debug = logger("e2e-tests:hooks");

const config_file =
  "../../conf/" + (process.env.E2E_CONFIG_FILE || "default") + ".conf.js";
// eslint-disable-next-line @typescript-eslint/no-var-requires
const config = require(config_file).config;

const username = process.env.E2E_SELENIUM_USERNAME || config.username;
const accessKey = process.env.E2E_SELENIUM_ACCESS_KEY || config.accessKey;

// eslint-disable-next-line @typescript-eslint/no-explicit-any
function createSession(config: any, caps: any): webdriver.ThenableWebDriver {
  if (process.env.E2E_SELENIUM_URL) {
    config.server = process.env.E2E_SELENIUM_URL;
  }

  return new webdriver.Builder()
    .usingServer(config.server)
    .withCapabilities(caps)
    .build();
}

setDefaultTimeout(process.env.PWDEBUG ? -1 : 120 * 1000);

Before({ tags: "@debug" }, async function (this: CustomWorld) {
  this.debug = true;
});

Before({ tags: "@unauthenticated" }, async function (this: CustomWorld) {
  debug(`@unauthenticated tag detected, logging out all organizations`);

  const logoutActions = Object.keys(organizations).map((orgName) =>
    logout(this, orgName)
  );

  await Promise.all(logoutActions);
});

Before(async function (
  this: CustomWorld,
  { gherkinDocument, pickle }: ITestCaseHookParameter
) {
  this.id = `${ulid()}-e2e`;

  for (const [name] of Object.entries(organizations)) {
    organizations[name].createdItems[this.id] = {
      services: [],
    };
  }

  const time = new Date().toISOString().split(".")[0];
  this.testName =
    pickle.name.replace(/\W/g, "-") + "-" + time.replace(/:|T/g, "-");

  debug(`starting ${this.testName}`);

  const task_id = parseInt(process.env.TASK_ID || "0");
  const caps = config.capabilities[task_id];
  caps["browserstack.user"] = username;
  caps["browserstack.key"] = accessKey;

  // Set browserstack idle timeout to max allowed setting (300 seconds).
  // This is the max timeout allowed to not sending commands to browserstack,
  // we need this because some tests setup things outside of browserstack
  // and thus not sending commands to browserstack
  caps["browserstack.idleTimeout"] = "300";

  caps.name = this.testName;

  caps["bstack:options"] = caps["bstack:options"] || {};
  caps["bstack:options"].buildName = process.env.E2E_BUILD_NAME || "local";

  this.driver = createSession(config, caps);

  // set session name
  if (process.env.CI === "true") {
    await this.driver.executeScript(
      `browserstack_executor: {"action": "setSessionName", "arguments": {"name": "${gherkinDocument.feature?.name} - ${pickle.name}"}}`
    );
  }

  await this.driver?.manage().setTimeouts({
    implicit: 10000,
  });

  // maximize browser window
  await this.driver.manage().window().maximize();

  this.feature = pickle;

  // In order to ensure we use the same locale as in the browser,
  // we retrieve the browser its locale via Selenium and load the
  // corresponding locale file for Dayjs.
  const locale = (await this.driver.executeScript(
    "return window.navigator.userLanguage || window.navigator.language;"
  )) as string;

  const dayjsLocale = mapBrowserLocaleToDayjs(locale);
  require(`dayjs/locale/${dayjsLocale}`);
  dayjs.locale(dayjsLocale);
});

After(async function (this: CustomWorld, { result }: ITestCaseHookParameter) {
  if (result) {
    await this.attach(
      `Status: ${result?.status}. Duration:${result.duration?.seconds}s`
    );

    if (result.status === "FAILED") {
      debug(`creating screenshot of failed test`);
      await this.snapshot();

      const reason = process.env.E2E_LOG_URL
        ? `Test failed, see logs for more info: ${process.env.E2E_LOG_URL}`
        : "Test failed, see local logs for more info";

      if (process.env.CI === "true") {
        await this.driver.executeScript(
          `browserstack_executor: {"action": "setSessionStatus", "arguments": {"status":"failed","reason": "${reason}"}}`
        );
      }
    } else {
      if (process.env.CI === "true") {
        await this.driver.executeScript(
          `browserstack_executor: {"action": "setSessionStatus", "arguments": {"status":"passed","reason": "Test passed"}}`
        );
      }
    }
  }

  // quit driver here because we don't need it anymore for the rest of the cleanup
  await this.driver?.quit();

  const requests = [];
  for (const [name, org] of Object.entries(organizations)) {
    debug(`cleaning created items of ${name}`);

    for (const service of org.createdItems[this.id].services) {
      debug(`deleting service: ${service}`);
      requests.push(
        org.apiClients.management?.managementServiceDeleteService({
          name: service,
        })
      );
    }
  }

  try {
    await Promise.all(requests);
  } catch (e) {
    debug(`error cleaning created items: ${e}`);
  }

  debug(`done with ${this.testName}`);
});

const mapBrowserLocaleToDayjs = (input: string): string => {
  switch (input) {
    case "en-US":
      return "en";
    case "nl-NL":
      return "nl";

    default:
      return input.toLowerCase();
  }
};
