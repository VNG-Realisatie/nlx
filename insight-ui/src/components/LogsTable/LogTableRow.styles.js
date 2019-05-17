import styled from 'styled-components'
import { Table } from '@commonground/design-system'

export const StyledLogTableRow = styled(Table.Row)`
  td {
    background-color: #FFFFFF;
    color: #2D3240;
  }
`

export const StyledSubjectLabel = styled.span`
  display: inline-flex;
  font-size: 12px;
  height: 24px;
  align-items: center;
  padding: 0 8px 0 8px;
  border-radius: 3px;
  border: 1px solid #CAD0E0;
  white-space: nowrap;
  
  &:not(:last-child) {
    margin-right: 4px;
  }
`
