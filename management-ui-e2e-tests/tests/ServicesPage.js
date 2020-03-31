// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { Selector } from 'testcafe'
import { axeCheck, createReport } from 'axe-testcafe'

const makeUrl = require('../utils/makeUrl')

fixture `Services page`
  .page(makeUrl('/services'))

test('Page is properly loaded', async t => {
  await t
    .expect(Selector('h1').innerText).eql('Services')
    .expect(Selector('[data-testid="services-list"]').exists).eql(true)

})

test('Automated accessibility testing', async t => {
  const { violations } = await axeCheck(t)
  await t.expect(violations.length === 0).ok(createReport(violations));
})
