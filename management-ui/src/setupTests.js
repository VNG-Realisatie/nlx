// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import '@testing-library/jest-dom/extend-expect'

// the MutationObserver shim is added because CRA does not yet support Jest v25.
// open issue: https://github.com/facebook/create-react-app/pull/8362
import MutationObserver from '@sheerun/mutationobserver-shim'
window.MutationObserver = MutationObserver

// Prevent fetch from going out to the network during test
global.fetch = require('jest-fetch-mock')
