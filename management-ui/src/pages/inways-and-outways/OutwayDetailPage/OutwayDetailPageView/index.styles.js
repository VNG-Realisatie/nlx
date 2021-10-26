// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

import SpecList from '../../../../components/SpecList'
import { IconOutway } from '../../../../icons'

export const SubHeader = styled.div`
  display: flex;
  align-items: center;
  margin-top: -${(p) => p.theme.tokens.spacing07};
  margin-bottom: ${(p) => p.theme.tokens.spacing06};
`

export const StyledIconOutway = styled(IconOutway)`
  fill: ${(p) => p.theme.colorFocus};
`

export const StyledSpecList = styled(SpecList)`
  margin-bottom: ${(p) => p.theme.tokens.spacing05};
`

export const StyledCode = styled.code`
  overflow: auto;
  font-size: 0.8rem;
  display: block;
  white-space: pre;
`
