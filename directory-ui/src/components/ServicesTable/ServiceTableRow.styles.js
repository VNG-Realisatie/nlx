// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import styled from 'styled-components'
import Table from '../Table'

export const StyledServiceTableRow = styled(Table.Row)`
  cursor: pointer;

  td {
    background-color: #ffffff;
    color: ${(p) => (p.status === 'offline' ? '#A3AABF' : '#2D3240')};
  }

  &:hover td {
    background-color: #f7f9fc;
  }
`

export const StyledApiTypeLabel = styled.span`
  display: inline-flex;
  font-size: 12px;
  line-height: 20px;
  padding: 1px 8px 2px;
  border-radius: 3px;
  border: 1px solid #cad0e0;
  white-space: nowrap;
  opacity: ${(p) => (p.status === 'offline' ? 0.4 : 1)};
`
