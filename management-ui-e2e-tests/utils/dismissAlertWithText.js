// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { Selector } from 'testcafe'

export default async (t, text) => {
  const alert = Selector('div[role="alert"]').withText(text)

  // Toasters disappear automatically, so let's not throw an error if it's already gone
  if (alert.exists) {
    const dismissButton = alert.parent().find('[role="button"]')
    await t.click(dismissButton)
  }
}
