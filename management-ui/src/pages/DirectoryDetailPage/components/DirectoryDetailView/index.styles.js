// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

export const AccessSection = styled.section`
  display: flex;
  align-items: center;
  justify-content: space-between;
`

export const IconItem = styled.div`
  width: ${(p) => p.theme.tokens.spacing07};
  color: ${(p) => p.theme.colorTextLabel};
`

export const StatusItem = styled.div`
  flex: 1 1 auto;
`

export const AccessMessage = styled.p`
  font-size: ${(p) => p.theme.tokens.fontSizeSmall};
`
