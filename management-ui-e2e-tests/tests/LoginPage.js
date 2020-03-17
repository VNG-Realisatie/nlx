// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { Selector } from 'testcafe'

const getBaseUrl = require('../getBaseUrl')
const baseUrl = getBaseUrl();

fixture `Login page`
  .page `${baseUrl}`

test('Welcome message is present', async t => {
  await t
    .expect(Selector('h1').innerText).eql('Welkom');
});
