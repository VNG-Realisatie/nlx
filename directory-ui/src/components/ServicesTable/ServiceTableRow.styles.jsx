import styled from 'styled-components'
import Table from '../Table'

export const StyledServiceTableRow = styled(Table.Row)`
  td {
    background-color: #FFFFFF;
    color: ${
    p => p.status === 'offline' ? '#A3AABF' : '#2D3240'};
  }
`

export const StyledApiTypeLabel = styled.span`
  display: inline-flex;
  font-size: 12px;
  line-height: 20px;
  padding: 1px 8px 2px;
  border-radius: 3px;
  border: 1px solid #CAD0E0;
  opacity: ${
  p => p.status === 'offline' ?
    .4 : 1
  }
`
