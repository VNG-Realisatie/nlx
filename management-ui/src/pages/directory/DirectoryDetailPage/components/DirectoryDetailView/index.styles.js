// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { Alert } from '@commonground/design-system'

import { IconExternalLink } from '../../../../../icons'

export const StyledAlert = styled(Alert)`
  margin-bottom: ${(p) => p.theme.tokens.spacing05};
`

export const AccessMessage = styled.p`
  font-size: ${(p) => p.theme.tokens.fontSizeSmall};
`

export const ExternalLinkSection = styled.section`
  display: flex;
  justify-content: space-between;
  margin-bottom: ${(p) => p.theme.tokens.spacing05};

  & > * {
    flex: 1 1 50%;

    &:first-child {
      margin-right: ${(p) => p.theme.tokens.spacing05};
    }
  }
`

export const StyledIconExternalLink = styled(IconExternalLink)`
  margin-left: ${(p) => p.theme.tokens.spacing03};
  fill: ${(p) => p.theme.tokens.colorPaletteGray400};
  /* TODO: Remove */
  width: 1.25rem;
  height: 1.25rem;
`
