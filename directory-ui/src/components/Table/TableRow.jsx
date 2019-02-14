import styled from 'styled-components'

const TableRow = styled.tr`
    display: table-row;
    
    &:first-child {
      td {
        border-top: 1px solid #EDEFF1;
      }
    
      td {
        &:first-child {
          border-top-left-radius: 3px;
          border-left: 1px solid #EDEFF1;
        }
        
        &:last-child {
          border-top-right-radius: 3px;
          border-right: 1px solid #EDEFF1;
        }
      }
    }
    
    &:not(:last-child) {
      td {
        border-bottom: 1px solid #EDEFF1;
      }
    }
    
    &:not(:first-child):not(:last-child) {
      td {
        &:first-child {
          border-left: 1px solid #EDEFF1;
        }
        
        &:last-child {
          border-right: 1px solid #EDEFF1;
        }
      }        
    }
    
    &:last-child {
      td {
        border-bottom: 1px solid #EDEFF1;
      
        &:first-child {
          border-bottom-left-radius: 3px;
          border-left: 1px solid #EDEFF1;
        }
        
        &:last-child {
          border-bottom-right-radius: 3px;
          border-right: 1px solid #EDEFF1;
        }
      }
    }
`

export default TableRow
