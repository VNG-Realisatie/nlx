// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import styled, { css } from 'styled-components'
import { Table } from '@commonground/design-system'

export const StyledLogTableRow = styled(Table.Row)`
  td {
    background-color: #FFFFFF;
    color: #2D3240;
  }

  ${p => p.active && css`
    cursor: default;

    td {
      background-color: #F7F9FC;
    }

    td:first-child {
      border-left: 2px solid #517FFF;
      padding-left: 14px;
    }
  `}

  ${p => !p.active && css`
    cursor: pointer;

    &:hover td {
      background-color: #F7F9FC;
    }

    &:active td {
      background-color: #F0F2F7;
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
  border: 1px solid #E6EAF5;
  white-space: nowrap;

  &:not(:last-child) {
    margin-right: 4px;
  }
`
