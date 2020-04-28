// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { Selector } from "testcafe";
import { waitForReact } from 'testcafe-react-selectors'
import { axeCheck, createReport } from 'axe-testcafe'

const getBaseUrl = require('../getBaseUrl')
const baseUrl = getBaseUrl();

fixture `Not Found (404) page`
  .beforeEach(async (t) => {
    await t.navigateTo(`${baseUrl}/page-that-does-not-exist'`)
    await waitForReact()
  })

test('Automated accessibility testing', async t => {
  const { violations } = await axeCheck(t)
  await t.expect(violations.length === 0).ok(createReport(violations));
})

test('Page not found message is present', async t => {
  await t
    .expect(Selector('h1').innerText).eql('Pagina niet gevonden');
});
