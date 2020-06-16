// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { IconPencil } from '../../icons'

export const StyledPencil = styled(IconPencil)`
  fill: ${(p) => p.theme.colorTextButtonSecondary};
  width: ${(p) => p.theme.tokens.spacing05};
  height: ${(p) => p.theme.tokens.spacing05};
  margin-right: ${(p) => p.theme.tokens.spacing03};
`
