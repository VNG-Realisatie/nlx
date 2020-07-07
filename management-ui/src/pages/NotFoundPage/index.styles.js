// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

import { IconWarningCircle } from '../../icons'

export const NotFoundContainer = styled.section`
  position: relative;
  margin-top: 100px;
  margin-left: 242px;
  max-width: 420px;
`

export const StyledIconErrorCircle = styled(IconWarningCircle)`
  position: absolute;
  top: ${(p) => p.theme.tokens.spacing04};
  left: -${(p) => p.theme.tokens.spacing11};
  width: 51px;
  height: 51px;
  fill: ${(p) => p.theme.tokens.colorError};
`
