import { CustomWorld } from './custom-world';
import { browserOptions } from './config';
import { Before, After, BeforeAll, AfterAll, Status, setDefaultTimeout } from '@cucumber/cucumber';
import {
  chromium,
  ChromiumBrowser,
  firefox,
  FirefoxBrowser,
  webkit,
  WebKitBrowser,
} from 'playwright';
import { ITestCaseHookParameter } from '@cucumber/cucumber/lib/support_code_library_builder/types';
import { ensureDir } from 'fs-extra';

let browser: ChromiumBrowser | FirefoxBrowser | WebKitBrowser;
const tracesDir = 'traces';

declare global {
  // eslint-disable-next-line no-var
  var browser: ChromiumBrowser | FirefoxBrowser | WebKitBrowser;
}

setDefaultTimeout(process.env.PWDEBUG ? -1 : 60 * 1000);

BeforeAll(async function () {
  switch (process.env.BROWSER) {
    case 'firefox':
      browser = await firefox.launch(browserOptions);
      break;
    case 'webkit':
      browser = await webkit.launch(browserOptions);
      break;
    default:
      browser = await chromium.launch(browserOptions);
  }
  await ensureDir(tracesDir);
});

Before({ tags: '@debug' }, async function (this: CustomWorld) {
  this.debug = true;
});

Before(async function (this: CustomWorld, { pickle }: ITestCaseHookParameter) {
  const time = new Date().toISOString().split('.')[0];
  this.testName = pickle.name.replace(/\W/g, '-') + '-' + time.replace(/:|T/g, '-');
  // customize the [browser context](https://playwright.dev/docs/next/api/class-browser#browsernewcontextoptions)
  this.context = await browser.newContext({
    acceptDownloads: true,
    recordVideo: process.env.PWVIDEO ? { dir: 'screenshots' } : undefined,
  });

  await this.context.tracing.start({ screenshots: true, snapshots: true });
  this.page = await this.context.newPage();
  this.page.on('console', async (msg) => {
    if (msg.type() === 'log') {
      await this.attach(msg.text());
    }
  });
  this.feature = pickle;
});

After(async function (this: CustomWorld, { result }: ITestCaseHookParameter) {
  if (result) {
    await this.attach(`Status: ${result?.status}. Duration:${result.duration?.seconds}s`);

    if (result.status !== Status.PASSED) {
      const image = await this.page?.screenshot();
      image && (await this.attach(image, 'image/png'));
      await this.context?.tracing.stop({ path: `${tracesDir}/${this.testName}-trace.zip` });
    }
  }
  await this.page?.close();
  await this.context?.close();
});

AfterAll(async function () {
  await browser.close();
});