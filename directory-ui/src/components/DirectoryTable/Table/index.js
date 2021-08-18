// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { createContext, useContext } from 'react'
import { bool, node, string } from 'prop-types'
import { Icon } from '@commonground/design-system'
import { useHistory } from 'react-router-dom'
import { IconChevronRight } from '../../../icons'
import {
  StyledTable,
  Td,
  MobileTd,
  MobileTdContent,
  Th,
  StyledTr,
  TrAsLink,
  TrAsLinkLink,
  TrAsLinkTd,
  Thead,
} from './index.styles'

const TableContext = createContext(false)

const Tr = ({ children, to, name, ...props }) => {
  const history = useHistory()
  const withLinks = useContext(TableContext)

  if (withLinks) {
    const handlePress = () => to && history.push(to)

    return (
      <TrAsLink
        onClick={handlePress}
        tabIndex="0"
        onKeyPress={handlePress}
        {...props}
      >
        {children}
        <TrAsLinkTd>
          {to && (
            <TrAsLinkLink to={to} aria-label={name} tabIndex="-1">
              <Icon as={IconChevronRight} />
            </TrAsLinkLink>
          )}
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
Table.MobileTd = MobileTd
Table.MobileTdContent = MobileTdContent
Table.Tr = Tr
Table.Thead = Thead
Table.TrHead = TrHead

export default Table
