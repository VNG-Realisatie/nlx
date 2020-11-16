// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { Button } from '@commonground/design-system'

export const StyledButton = styled(Button)`
  float: right;
  visibility: hidden;

  tr:hover & {
    visibility: visible;
  }
`
