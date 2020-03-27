// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React, { createContext, useContext } from 'react'
import { bool, node, string } from 'prop-types'
import { useHistory } from 'react-router-dom'
import {
  StyledChevron,
  StyledTable,
  Td,
  Th,
  TrAsLink,
  TrAsLinkLink,
  TrAsLinkTd,
} from './index.styles'

const TableContext = createContext(false)

const Tr = ({ children, to, name, ...props }) => {
  const history = useHistory()
  const withLinks = useContext(TableContext)

  if (withLinks) {
    const handlePress = () => (to ? history.push(to) : null)

    return (
      <TrAsLink
        onClick={handlePress}
        tabIndex="0"
        onKeyPress={handlePress}
        {...props}
      >
        {children}
        <TrAsLinkTd>
          {to ? (
            <TrAsLinkLink to={to} aria-label={name} tabIndex="-1">
              <StyledChevron />
            </TrAsLinkLink>
          ) : null}
        </TrAsLinkTd>
      </TrAsLink>
    )
  }
  return <tr {...props}>{children}</tr>
}

Tr.propTypes = {
  children: node,
  to: string,
  name: string,
}

const TrHead = ({ children, ...props }) => {
  const withLinks = useContext(TableContext)

  if (withLinks) {
    return (
      <tr {...props}>
        {children}
        <TrAsLinkTd as="th" />
      </tr>
    )
  }
  return <tr {...props}>{children}</tr>
}
TrHead.propTypes = {
  children: node,
}

const Table = ({ withLinks, children, ...props }) => {
  return (
    <TableContext.Provider value={withLinks}>
      <StyledTable {...props}>{children}</StyledTable>
    </TableContext.Provider>
  )
}
Table.propTypes = {
  children: node,
  withLinks: bool,
}
Table.defaultProps = {
  withLinks: false,
}

Table.Th = Th
Table.Td = Td
Table.Tr = Tr
Table.TrHead = TrHead

export default Table
