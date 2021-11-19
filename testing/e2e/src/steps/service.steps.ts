import { CustomWorld } from '../support/custom-world';
import { organizations } from '../utils/organizations';
import { When, Then } from '@cucumber/cucumber';

When('{string} create a service with valid properties named {string} and as inway {string}', async function (this: CustomWorld, orgName: string, serviceName: string, inwayName: string) {
    const { page } = this;

    const org = (organizations as any)[orgName];
    if (org === undefined) {
        return;
    }

    await page?.goto(`${org.management.url}/services`);

    await page?.click('text=Service toevoegen');

    await page?.fill('[data-testid="name"]', serviceName);
    await page?.fill('[data-testid="endpointURL"]', 'https://google.com');
    await page?.click(`text="${inwayName}"`);

    await page?.click('[data-testid="form"] >> text=Service toevoegen')
});

Then('the service {string} shows up under My services of the management interface of {string}', async function (this: CustomWorld, serviceName: string, orgName: string) {
    const { page } = this;

    const org = (organizations as any)[orgName];
    if (org === undefined) {
        return;
    }
    
    await page?.goto(`${org.management.url}/services`),

    await page?.waitForSelector(`text="${serviceName}"`)
});