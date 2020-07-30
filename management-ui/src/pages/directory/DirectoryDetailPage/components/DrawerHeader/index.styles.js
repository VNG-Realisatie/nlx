// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

import StateIndicator from '../../../../../components/StateIndicator'

export const SubTitle = styled.p`
  margin-bottom: ${(p) => p.theme.tokens.spacing04};
  font-size: ${(p) => p.theme.tokens.fontSizeLarge};
`

export const Summary = styled.div`
  display: flex;

  & > * {
    font-size: ${(p) => p.theme.tokens.fontSizeSmall};
    margin-right: ${(p) => p.theme.tokens.spacing06};
    color: ${(p) => p.theme.colorTextLabel};
  }
`

export const StyledStateIndicator = styled(StateIndicator)`
  flex-grow: 1;
`
