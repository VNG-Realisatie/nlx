// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { Table } from '@commonground/design-system'
import ButtonWithIcon from '../../../../../../components/ButtonWithIcon'

export const TdActions = styled(Table.Td)`
  text-align: right;
  vertical-align: top;
  width: 140px;
`

export const StyledButtonWithIcon = styled(ButtonWithIcon)`
  min-width: 44px;
  min-height: 44px;
  justify-content: center;

  &:not(:first-child) {
    margin-left: 1rem;
  }

  svg {
    width: 20px;
    height: 20px;
    margin-right: 0;
  }
`
