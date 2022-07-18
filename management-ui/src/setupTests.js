// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import '@testing-library/jest-dom/extend-expect'

global.scrollTo = jest.fn()

// Prevent fetch from going out to the network during test
global.fetch = require('jest-fetch-mock')
