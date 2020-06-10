// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

import SpecList from '../SpecList'
import { ReactComponent as IconInway } from './inway.svg'

export const SubHeader = styled.div`
  display: flex;
  align-items: center;
  margin-top: -${(p) => p.theme.tokens.spacing07};
  margin-bottom: ${(p) => p.theme.tokens.spacing06};
`

export const StyledIconInway = styled(IconInway)`
  width: ${(p) => p.theme.tokens.spacing05};
  height: ${(p) => p.theme.tokens.spacing05};
  margin-right: ${(p) => p.theme.tokens.spacing03};
`

export const StyledSpecList = styled(SpecList)`
  margin-bottom: ${(p) => p.theme.tokens.spacing05};
`
