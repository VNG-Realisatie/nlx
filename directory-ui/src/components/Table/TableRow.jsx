import styled from 'styled-components'

const TableRow = styled.tr`
    display: table-row;
    
    td {
      box-shadow: 0 -1px 0 #F0F2F7;
    
      &:first-child {
        box-shadow: 0 -1px 0 #F0F2F7, -1px 0 #F0F2F7;
      }
      
      &:last-child {
        box-shadow: 0 -1px 0 #F0F2F7, 1px 0 #F0F2F7;
      }
    }
`

export default TableRow
