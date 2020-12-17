// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { Button } from '@commonground/design-system'

export const StyledButton = styled(Button)`
  display: none;
  padding-left: ${(p) => p.theme.tokens.spacing05};

  tr:hover & {
    display: block;
  }
`
