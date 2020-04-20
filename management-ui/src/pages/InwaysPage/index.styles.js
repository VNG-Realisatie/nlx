// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import styled, { css } from 'styled-components'
import Table from '../../components/Table'
import { ReactComponent as InwayIcon } from './inway.svg'

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

export const StyledIconTd = styled(Table.Td)`
  width: 3rem;
`

export const StyledInwayIcon = styled(InwayIcon)`
  display: block;
  width: 1rem;
  height: 1rem;
`
