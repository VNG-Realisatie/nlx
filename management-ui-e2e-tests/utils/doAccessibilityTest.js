// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { axeCheck, createReport } from 'axe-testcafe'

export default async (t) => {
  const { violations } = await axeCheck(t)
  await t.expect(violations.length === 0).ok(createReport(violations))
}
