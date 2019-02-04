import styled from 'styled-components'

const TableBody = styled.tbody`
    background-color: #ffffff;
    display: table-row-group;
    
    tr {    
      &:first-child {
        td {
          &:first-child {
            border-top-left-radius: 3px;
            box-shadow: 0 -1px 0 #F0F2F7,
                        -1px 0 #F0F2F7;
          }
          
          
          &:last-child {
            border-top-right-radius: 3px;
            box-shadow: 0 -1px 0 #F0F2F7,
                        1px 0 #F0F2F7;
          }
        }
      }
    }
    
    tr:last-child {
      td {
        box-shadow: 0 -1px 0 #F0F2F7,0 1px 0 #F0F2F7;
      
        &:first-child {
          border-bottom-left-radius: 3px;
          box-shadow: 0 -1px 0 #F0F2F7,
                      0 1px 0 #F0F2F7,
                      -1px 0 #F0F2F7;
        }
        
        &:last-child {
          border-bottom-right-radius: 3px;
          box-shadow: 0 -1px 0 #F0F2F7,
                      0 1px 0 #F0F2F7,
                      1px 0 #F0F2F7;
        }
      }
    }
`

export default TableBody
