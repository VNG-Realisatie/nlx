// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { render } from '@testing-library/react'
import AllProviders from './all-providers'

// based on https://testing-library.com/docs/react-testing-library/setup#custom-render
const renderWithProviders = (ui, options) => {
  const reactRoot = document.createElement('div')
  reactRoot.setAttribute('id', 'root')
  document.body.appendChild(reactRoot)
  return render(ui, { wrapper: AllProviders, ...options })
}

// re-export everything
export * from '@testing-library/react'

// override render method
export { renderWithProviders }
