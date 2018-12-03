import styled from 'styled-components'

const TableHead = styled.thead`
    display: table-header-group;

    th {
        font-size: ${p => p.theme.font.size.small};
        line-height: ${p => p.theme.font.lineHeight.small};

        color: ${p => p.theme.color.grey[60]};
        padding-bottom: .75rem;

        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
    }
`

export default TableHead