// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { ClientFunction } from 'testcafe'

const getLocation = ClientFunction(() => document.location.href)

module.exports = getLocation
