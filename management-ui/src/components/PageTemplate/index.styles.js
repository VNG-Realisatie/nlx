// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

export const Page = styled.div`
  display: flex;
  align-items: flex-start;
  height: 100%;
`

export const MainWrapper = styled.div`
  flex: 1;
  overflow: auto;
  height: 100%;
`

export const Main = styled.main`
  padding: ${(p) => p.theme.tokens.spacing09};
`
