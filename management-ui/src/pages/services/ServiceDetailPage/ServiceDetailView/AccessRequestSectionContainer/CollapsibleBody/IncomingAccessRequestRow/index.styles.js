// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { Table, Button } from '@commonground/design-system'

export const TdActions = styled(Table.Td)`
  text-align: right;
  vertical-align: top;
  padding: 0;
  width: 140px;
`

export const StyledButton = styled(Button)`
  min-width: 44px;
  min-height: 44px;
  justify-content: center;

  &:not(:first-child) {
    margin-left: 1rem;
  }
`
