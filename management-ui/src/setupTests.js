// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom/extend-expect'

// the MutationObserver shim is added because CRA does not yet support Jest v25.
// open issue: https://github.com/facebook/create-react-app/pull/8362
import MutationObserver from '@sheerun/mutationobserver-shim'
window.MutationObserver = MutationObserver
