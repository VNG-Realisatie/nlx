// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { bool, arrayOf, node } from 'prop-types'
import styled from 'styled-components'
import { Table } from '@commonground/design-system'

/**
 * Displays key/value pairs with the key in a deviating style.
 * Intended to be used with two td's per row only.
 */
const SpecTable = styled(Table)`
  & td:nth-child(1) {
    font-size: ${(p) => p.theme.tokens.fontSizeSmall};
    color: ${(p) => p.theme.tokens.colorPaletteGray500};
  }

  & td:nth-child(n + 2) {
    ${({ valueAlignRight }) => (valueAlignRight ? 'text-align: right' : '')}
  }
`

SpecTable.propTypes = {
  valueAlignRight: bool,
}

SpecTable.defaultProps = {
  valueAlignRight: false,
}

if (process.env.NODE_ENV !== 'production') {
  // Augment Tr to give developer a notice when using more then two cells in table row
  // Note: the production build won't contain this code
  const Tr = ({ children }) => {
    if (children.length > 2) {
      console.warn('Each row in `SpecTable` should only have two cells.')
    }

    return <Table.Tr>{children}</Table.Tr>
  }

  Tr.propTypes = {
    children: arrayOf(node),
  }

  SpecTable.Tr = Tr
}

export default SpecTable
