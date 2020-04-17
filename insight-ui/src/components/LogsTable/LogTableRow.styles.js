// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import styled, { css } from 'styled-components'
import Table from '../Table'

export const StyledLogTableRow = styled(Table.Row)`
  td {
    background-color: #ffffff;
    color: #2d3240;
  }

  ${(p) =>
    p.active &&
    css`
      cursor: default;

      td {
        background-color: #f7f9fc;
      }

      td:first-child {
        border-left: 2px solid #517fff;
        padding-left: 14px;
      }
    `}

  ${(p) =>
    !p.active &&
    css`
      cursor: pointer;

      &:hover td {
        background-color: #f7f9fc;
      }

      &:active td {
        background-color: #f0f2f7;
      }
    `}
`

export const StyledSubjectLabel = styled.span`
  display: inline-flex;
  font-size: 12px;
  line-height: 20px;
  height: 24px;
  align-items: center;
  padding: 0 8px 0 8px;
  border-radius: 3px;
  background-color: white;
  border: 1px solid #e6eaf5;
  white-space: nowrap;

  &:not(:last-child) {
    margin-right: 4px;
  }
`
