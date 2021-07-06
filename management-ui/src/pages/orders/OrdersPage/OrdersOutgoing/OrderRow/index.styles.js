// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import styled from 'styled-components'
import Table from '../../../../../components/Table'

export const Cell = styled(Table.Td)`
  vertical-align: top;
`

export const List = styled.ul`
  padding: 0;
  margin: 0;
  list-style: none;
`

export const Item = styled.li`
  font-size: ${(p) => p.theme.tokens.fontSizeSmall};
  color: ${(p) => p.theme.colorTextLabel};
  text-overflow: ellipsis;
  white-space: nowrap;
  overflow: hidden;
`

export const OrganizationName = styled.span`
  font-size: ${(p) => p.theme.tokens.fontSizeMedium};
  color: ${(p) => p.theme.colorText};
`

export const Separator = styled.span`
  margin: 0 0.25rem;
`
