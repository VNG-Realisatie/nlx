// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import styled from 'styled-components'

const Table = styled.table`
  border-spacing: unset;
  width: 100%;
`

const Th = styled.th`
  text-transform: uppercase;
  text-align: left;
  padding: ${(p) => p.theme.tokens.spacing04} 0;
  border-bottom: 1px solid ${(p) => p.theme.colorBorderTable};
  font-size: ${(p) => p.theme.tokens.fontSizeSmall};
`

const Td = styled.td`
  padding: ${(p) => p.theme.tokens.spacing04} 0;
  border-bottom: 1px solid ${(p) => p.theme.colorBorderTable};
`

Table.Th = Th
Table.Td = Td

export default Table
