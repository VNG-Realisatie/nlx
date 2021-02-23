// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { IconExternalLink } from '../../icons'

export const Centered = styled.section`
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 10rem auto;
  max-width: 32rem;
  text-align: center;
`

export const StyledIconExternalLink = styled(IconExternalLink)`
  margin-left: ${(p) => p.theme.tokens.spacing03};
  fill: ${(p) => p.theme.tokens.colorPaletteGray400};
`
