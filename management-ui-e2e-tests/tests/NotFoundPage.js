// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { Selector } from 'testcafe'
import { axeCheck, createReport } from 'axe-testcafe'

const makeUrl = require('../utils/makeUrl')

fixture `NotFound (404) page`
  .page(makeUrl('/page-that-does-not-exist'))

test('Automated accessibility testing', async t => {
  const { violations } = await axeCheck(t)
  await t.expect(violations.length === 0).ok(createReport(violations));
})

test('404 page is properly loaded', async t => {
  await t
    .expect(Selector('h1').innerText).eql('Page not found')
})
