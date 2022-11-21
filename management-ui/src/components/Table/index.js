// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React, { createContext, useContext } from 'react'
import { bool, node, string } from 'prop-types'
import { useNavigate } from 'react-router-dom'
import {
  StyledChevron,
  StyledTable,
  Td,
  Th,
  StyledTr,
  TrAsLink,
  TrAsLinkLink,
  TrAsLinkTd,
} from './index.styles'

const TableContext = createContext(false)

const Tr = ({ children, to, name, ...props }) => {
  const navigate = useNavigate()
  const withLinks = useContext(TableContext)

  if (withLinks) {
    const handlePress = () => (to ? navigate(to) : null)

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
  return <StyledTr {...props}>{children}</StyledTr>
}

Tr.propTypes = {
  children: node,
  to: string,
  name: string,
  selected: bool,
}

Tr.defaultProps = {
  selected: false,
}

const TrHead = ({ children, ...props }) => {
  const withLinks = useContext(TableContext)

  if (withLinks) {
    return (
      <tr {...props}>
        {children}
        <TrAsLinkTd as="td" />
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
