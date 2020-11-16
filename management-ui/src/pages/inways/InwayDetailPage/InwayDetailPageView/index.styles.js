// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

import SpecList from '../../../../components/SpecList'
import { IconInway } from '../../../../icons'

export const SubHeader = styled.div`
  display: flex;
  align-items: center;
  margin-top: -${(p) => p.theme.tokens.spacing07};
  margin-bottom: ${(p) => p.theme.tokens.spacing06};
`

export const StyledIconInway = styled(IconInway)`
  fill: ${(p) => p.theme.tokens.colorPaletteGray50};
`

export const StyledSpecList = styled(SpecList)`
  margin-bottom: ${(p) => p.theme.tokens.spacing05};
`
