// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { Selector } from 'testcafe'

export default async (t, text) => {
  const alert = Selector('div[role="alert"]').withText(text)
  const dismissButton = alert.parent().find('[role="button"]')
  await t.click(dismissButton)
}
