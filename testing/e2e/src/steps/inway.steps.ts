import { CustomWorld } from '../support/custom-world';
import { organizations } from '../utils/organizations';
import { Given } from '@cucumber/cucumber';

Given('{string} has an inway with the {string} running', async function (this: CustomWorld, orgName: string, inwayName: string) {
  const { page } = this;

  const org = (organizations as any)[orgName];
  if (org === undefined) {
    return;
  }

  await page?.goto(`${org.management.url}/inways-and-outways/inways`);

  await page?.waitForSelector(`text="${inwayName}"`)
});