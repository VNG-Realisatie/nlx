// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { axeCheck, createReport } from 'axe-testcafe'

export default async (t, axeContext, axeOptions) => {
  const { violations } = await axeCheck(t, axeContext, axeOptions)
  await t.expect(violations.length === 0).ok(createReport(violations))
}
