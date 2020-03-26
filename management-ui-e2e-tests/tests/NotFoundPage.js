// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { Selector } from 'testcafe'

const makeUrl = require('../utils/makeUrl')

fixture `NotFound (404) page`
  .page(makeUrl('/notfoundd'))

test('404 page is properly loaded', async t => {
  await t
    .expect(Selector('h1').innerText).eql('Page not found')
})
