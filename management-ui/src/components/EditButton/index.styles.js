// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { ReactComponent as Pencil } from './pencil.svg'

export const StyledPencil = styled(Pencil)`
  fill: ${(p) => p.theme.colorTextButtonSecondary};
  width: ${(p) => p.theme.tokens.spacing05};
  height: ${(p) => p.theme.tokens.spacing05};
  margin-right: ${(p) => p.theme.tokens.spacing03};
`
