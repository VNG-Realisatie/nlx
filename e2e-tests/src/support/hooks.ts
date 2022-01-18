import { CustomWorld } from "./custom-world";
// import { snapshot } from "../utils/snapshot";
import { Before, After, setDefaultTimeout } from "@cucumber/cucumber";
import { ITestCaseHookParameter } from "@cucumber/cucumber/lib/support_code_library_builder/types";
import webdriver from "selenium-webdriver";
import { ulid } from "ulid";

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

setDefaultTimeout(process.env.PWDEBUG ? -1 : 60 * 1000);

Before({ tags: "@debug" }, async function (this: CustomWorld) {
  this.debug = true;
});

Before(async function (this: CustomWorld, { pickle }: ITestCaseHookParameter) {
  this.id = ulid();

  const time = new Date().toISOString().split(".")[0];
  this.testName =
    pickle.name.replace(/\W/g, "-") + "-" + time.replace(/:|T/g, "-");

  console.log(`starting ${this.testName}`);

  const task_id = parseInt(process.env.TASK_ID || "0");
  const caps = config.capabilities[task_id];
  caps["browserstack.user"] = username;
  caps["browserstack.key"] = accessKey;

  caps.name = this.testName;
  caps.build = process.env.E2E_BUILD_NAME || "local";

  this.driver = createSession(config, caps);
  await this.driver?.manage().setTimeouts({
    implicit: 10000,
  });

  this.feature = pickle;
});

After(async function (this: CustomWorld, { result }: ITestCaseHookParameter) {
  if (result) {
    await this.attach(
      `Status: ${result?.status}. Duration:${result.duration?.seconds}s`
    );

    if (result.status === "FAILED") {
      await this.snapshot();
    }
  }
  await this.driver?.quit();
  console.log(`done with ${this.testName}`);
});
