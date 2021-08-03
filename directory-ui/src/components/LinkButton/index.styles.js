// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import styled from 'styled-components'
import { Icon as BaseIcon } from '@commonground/design-system'
import { IconExternalLink as ExternalLink } from '../../icons'

export const Icon = styled(BaseIcon)`
  margin-left: ${(p) => p.theme.tokens.spacing03};
  margin-right: 0;
`

export const IconExternalLink = styled(ExternalLink)`
  width: ${(p) => p.theme.tokens.iconSizeSmall};
  height: ${(p) => p.theme.tokens.iconSizeSmall};
  fill: ${(p) => p.theme.tokens.colorPaletteGray700};
`
