import styled from 'styled-components'
import TableHeadCell from './TableHeadCell'

export const StyledSortableTableHeadCell = styled(TableHeadCell)`
    cursor: pointer;
`

export const StyledArrow = styled.svg`
  width: 8px;
  height: 10px;
  position: relative;
  left: 5px;
  top: 1px;
`