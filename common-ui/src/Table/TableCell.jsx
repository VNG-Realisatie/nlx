import PropTypes from 'prop-types'
import styled, {css} from 'styled-components'

const TableCell = styled.td`
    display: table-cell;
    border-bottom: 1px solid ${p => p.theme.color.grey[30]};
    padding: .9rem 3rem 1rem 1.5rem;

    font-size: ${p => p.theme.font.size.normal};
    line-height: ${p => p.theme.font.lineHeight.normal};
    font-weight: ${p => p.theme.font.weight.normal};
    color: ${p => p.theme.color.black};

    ${p => p.align === 'left' && css`
        text-align: left;
    `}

    ${p => p.align === 'center' && css`
        text-align: center;
    `}

    ${p => p.align === 'right' && css`
        text-align: right;
    `}

    &:last-child {
        padding-right: 1.5rem;
    }
`

TableCell.propTypes = {
    children: PropTypes.node,
    align: PropTypes.oneOf(['left', 'center', 'right']).isRequired,
}

TableCell.defaultProps = {
    align: 'left',
}

export default TableCell