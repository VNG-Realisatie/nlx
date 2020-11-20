// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

import { IconExternalLink } from '../../../../../../icons'

export const Section = styled.section`
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
  fill: ${(p) =>
    p.$disabled
      ? p.theme.colorTextButtonSecondaryDisabled
      : p.theme.tokens.colorPaletteGray400};
`
