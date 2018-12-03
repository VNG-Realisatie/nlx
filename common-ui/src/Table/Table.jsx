import styled from 'styled-components'

const Table = styled.table`
    display: table;
    width: 100%;

    border-spacing: 0;
    border-collapse: collapse;
`

export default Table
export { default as TableBody } from './TableBody'
export { default as TableCell } from './TableCell'
export { default as TableHead } from './TableHead'
export { default as TableRow } from './TableRow'