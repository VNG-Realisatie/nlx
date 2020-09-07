// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { Selector, t } from 'testcafe'
const label = Selector('label')

export const AUTHORIZATION_TYPE_NONE = 'none'
export const AUTHORIZATION_TYPE_WHITELIST = 'whitelist'

class Checkbox {
  constructor(text) {
    this.label = label.withText(text)
    this.checkbox = this.label.find('input[type=checkbox]')
  }

  async enable() {
    const checked = await this.checkbox.checked
    if (!checked) {
      await t.click(this.checkbox)
    }
  }

  async disable() {
    const checked = await this.checkbox.checked
    if (checked) {
      await t.click(this.checkbox)
    }
  }
}

class AuthorizationType {
  constructor(text) {
    this.label = label.withText(text)
    this.radioButton = this.label.find('input[type=radio]')
  }
}

class Page {
  constructor() {
    this.title = Selector('h1')

    this.nameInput = Selector('[data-testid="name"]')
    this.endpointUrlInput = Selector('[data-testid="endpointURL"]')
    this.documentationUrlInput = Selector('[data-testid="documentationURL"]')
    this.apiSpecificationUrlInput = Selector(
      '[data-testid="apiSpecificationURL"]',
    )
    this.publishToCentralDirectory = new Checkbox(
      'Publiceren in de centrale directory',
    )
    this.publishedInDirectoryWarning = Selector(
      '[data-testid="publishedInDirectory-warning"]',
    )
    this.techSupportContactInput = Selector(
      '[data-testid="techSupportContact"]',
    )
    this.publicSupportContactInput = Selector(
      '[data-testid="publicSupportContact"]',
    )
    this.authorizationModes = {
      whitelist: new AuthorizationType(
        'Whitelist voor geauthorizeerde organisaties',
      ),
      none: new AuthorizationType('Alle organisaties toestaan'),
    }
    this.inway = (name) => Selector(`[name="inways"][value=${name}]`)

    this.submitButton = Selector('button[type="submit"]')

    this.nameFieldError = Selector('[data-testid="error-name"]')
  }

  async fillAndSubmitForm({
    name,
    endpointUrl,
    documentationUrl,
    apiSpecificationUrl,
    publishToCentralDirectory,
    techSupportContact,
    publicSupportContact,
    authorizationType,
    inways,
    performSubmit,
  }) {
    if (typeof name !== 'undefined') {
      await t.typeText(this.nameInput, name, { replace: true })
    }

    if (typeof endpointUrl !== 'undefined') {
      await t.typeText(this.endpointUrlInput, endpointUrl, { replace: true })
    }
    if (typeof documentationUrl !== 'undefined') {
      await t.typeText(this.documentationUrlInput, documentationUrl, {
        replace: true,
      })
    }

    if (typeof apiSpecificationUrl !== 'undefined') {
      await t.typeText(this.apiSpecificationUrlInput, apiSpecificationUrl, {
        replace: true,
      })
    }

    if (typeof publishToCentralDirectory !== 'undefined') {
      if (publishToCentralDirectory) {
        await this.publishToCentralDirectory.enable()
      } else {
        await this.publishToCentralDirectory.disable()
      }
    }

    if (typeof techSupportContact !== 'undefined') {
      await t.typeText(this.techSupportContactInput, techSupportContact, {
        replace: true,
      })
    }

    if (typeof publicSupportContact !== 'undefined') {
      await t.typeText(this.publicSupportContactInput, publicSupportContact, {
        replace: true,
      })
    }

    if (typeof inways !== 'undefined') {
      for (const inway in inways) {
        await t.click(this.inway(inways[inway]))
      }
    }

    if (typeof authorizationType !== 'undefined') {
      switch (authorizationType) {
        case AUTHORIZATION_TYPE_NONE:
          await t.click(this.authorizationModes.none.radioButton)
          break

        case AUTHORIZATION_TYPE_WHITELIST:
          await t.click(this.authorizationModes.whitelist.radioButton)
          break

        default:
          throw new Error(`invalid authorization type '${authorizationType}'`)
      }
    }

    if (typeof performSubmit === 'undefined' || performSubmit === true) {
      await t.click(this.submitButton)
      await t.wait(3000)
    }
  }
}

export default new Page()
