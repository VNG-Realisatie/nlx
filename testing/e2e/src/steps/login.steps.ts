import { CustomWorld } from '../support/custom-world';
import { organizations } from '../utils/organizations';
import { Given } from '@cucumber/cucumber';

Given('{string} is logged into NLX management', async function (this: CustomWorld, orgName: string) {
  const { page } = this;

  const org = (organizations as any)[orgName];
  if (org === undefined) {
    return;
  }

  await page?.goto(org.management.url);

  if (org.basicAuth) {
    return
  }

  await Promise.all([
    page?.click('[data-testid="login"]'),
    page?.waitForNavigation()
  ]);

  await page?.fill('[placeholder="email address"]', 'admin@nlx.local');
  await page?.fill('[placeholder="password"]', 'development');
  await page?.click('text=Login');

  await Promise.all([
    page?.click('button:has-text("Grant Access")'),
    page?.waitForNavigation()
  ]);
});