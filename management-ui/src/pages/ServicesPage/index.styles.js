// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import styled, { css } from 'styled-components'
import { ReactComponent as IconPlus } from './plus.svg'

const verticallyAlignedContent = css`
  height: calc(100vh - 20rem);
  line-height: calc(100vh - 20rem);
`

export const StyledLoadingMessage = styled.p`
  align-items: center;
  justify-content: center;
  display: flex;
  margin-bottom: 0;
  ${verticallyAlignedContent}
`

export const StyledNoServicesMessage = styled.p`
  margin-bottom: 0;
  text-align: center;
  ${verticallyAlignedContent}
`

export const StyledActionsBar = styled.div`
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
`

export const StyledIconPlus = styled(IconPlus)`
  margin-right: ${(p) => p.theme.tokens.spacing03};
`
