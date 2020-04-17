// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { node, oneOf } from 'prop-types'
import styled, { css } from 'styled-components'

const BodyCell = styled.td`
  background-color: #ffffff;
  display: table-cell;
  padding: ${(p) => (p.padding === 'none' ? 0 : '10px 16px 10px 16px')};

  font-size: 14px;
  line-height: 22px;
  font-weight: 400;

  text-align: ${(p) => p.align};

  ${(p) =>
    p.border === 'left' &&
    css`
      border-left: 1px solid #f0f2f7;
    `}

  ${(p) =>
    p.border === 'right' &&
    css`
      border-right: 1px solid #f0f2f7;
    `}
`

BodyCell.propTypes = {
  children: node,
  align: oneOf(['left', 'center', 'right']),
  border: oneOf(['left', 'right']),
  padding: oneOf(['none']),
}

BodyCell.defaultProps = {
  align: 'left',
}

export default BodyCell
