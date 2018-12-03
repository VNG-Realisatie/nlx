import styled from 'styled-components'

const TableRow = styled.tr`
    display: table-row;

    &:last-child {
        td, th {
            border-bottom: none;
        }
    }
`

export default TableRow