// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { Button } from '@commonground/design-system'

export default styled(Button)`
  svg {
    width: ${(p) => p.theme.tokens.fontSizeMedium};
    height: ${(p) => p.theme.tokens.fontSizeMedium};
    margin-right: ${(p) => p.theme.tokens.spacing03};
  }
`
