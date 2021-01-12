// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { IconServices } from '../../../icons'

export const StyledIconServices = styled(IconServices)`
  margin-right: ${(p) => p.theme.tokens.spacing05};
  fill: ${(p) => p.theme.tokens.colorPaletteGray500};
`

export const SectionHeader = styled.p`
  margin-bottom: ${(p) => p.theme.tokens.spacing03};
  color: ${(p) => p.theme.tokens.colorPaletteGray400};
`

export const ServiceField = styled.div`
  display: flex;
`

export const ServiceData = styled.div`
  display: flex;
  flex-direction: column;
`
