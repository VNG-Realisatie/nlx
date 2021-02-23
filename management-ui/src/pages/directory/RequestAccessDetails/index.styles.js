// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { IconServices } from '../../../icons'

export const SectionHeader = styled.p`
  margin-bottom: 0;
  color: ${(p) => p.theme.tokens.colorPaletteGray400};
`

export const SectionContent = styled.div`
  display: flex;
  margin-bottom: ${(p) => p.theme.tokens.spacing07};
`

export const SectionContentService = styled(SectionContent)`
  margin-top: ${(p) => p.theme.tokens.spacing03};
`

export const StyledIconServices = styled(IconServices)`
  margin-right: ${(p) => p.theme.tokens.spacing05};
  fill: ${(p) => p.theme.tokens.colorPaletteGray500};
`

export const ServiceData = styled.div`
  display: flex;
  flex-direction: column;
`
