// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { Selector } from 'testcafe'

const getBaseUrl = require('../getBaseUrl')
const baseUrl = getBaseUrl();

fixture `NotFound (404) page`
  .page `${baseUrl}/asdf`

test('Page not found message is present', async t => {
  await t
    .expect(Selector('h1').innerText).eql('Page not found');
});
