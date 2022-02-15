// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { Alert } from '@commonground/design-system'

export const StyledAlert = styled(Alert)`
  margin-bottom: ${(p) => p.theme.tokens.spacing05};
`

export const Section = styled.section`
  display: flex;
  align-items: center;
  justify-content: space-between;
`

export const IconItem = styled.div`
  margin-right: ${(p) => p.theme.tokens.spacing03};
  color: ${(p) => p.theme.colorTextLabel};
`

export const StateDetail = styled.div`
  flex: 1 1 auto;
  display: flex;
  flex-direction: column;
  justify-content: center;
  min-height: 50px;
`
